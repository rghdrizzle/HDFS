package main

import (
	//"fmt"
	// "bytes"
	// "crypto/md5"
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"strings"
)
const RootDefaultFolderName = "drizzleNet"
//CASpathTransformFunc is a function which converts the key to a hash path so that it follows content-addressable storage system
// The function computes the sha1 for the key and then converts that to a hexadecimal string. The hexadecimal string is then broken down to 5 parts which is then joined together to create a hased path.
// example:  3c4df/ead25/99fb1/00c3c/76053/d04ca/81af6/4ab70
func CASpathTransformFunc( key string ) PathKey{
	hash := sha1.Sum([]byte(key)) // [20]bytes => to a slice([]byte) by doing => [:]
	hashString := hex.EncodeToString(hash[:])

	blockSize := 5
	sliceLength := len(hashString) / blockSize
	paths := make([]string, sliceLength)
	for i:=0;i<sliceLength;i++{
		from, to := i*blockSize, (i*blockSize)+blockSize
		paths[i]= hashString[from:to]
	}
	
	return PathKey{
		PathName: strings.Join(paths,"/"),
		Filename: hashString,
		}
}

type PathTransformFromFunc func(string) PathKey
type StoreOpts struct {
	//Root is the folder name of the root which is the root directory which contains files
	Root string
	PathTransformFromFunc PathTransformFromFunc
}

type PathKey struct{
	PathName string
	Filename string
}

func (p PathKey) FullPath() string{
	return fmt.Sprintf("%s/%s",p.PathName,p.Filename)
}
func (p PathKey) FirstFileName() string{
	path := strings.Split(p.PathName,"/")
	if len(path)==0{
		return ""
	}
	return path[0]
}
var DefaultTransformFunc = func(key string) PathKey {
	return PathKey{
		PathName: key,
		Filename: key,
	}
}


type Store struct {
	StoreOpts
}

func NewStore(opts StoreOpts) *Store{
	if opts.PathTransformFromFunc == nil{
		opts.PathTransformFromFunc = DefaultTransformFunc
	}
	if len(opts.Root)==0{
		opts.Root = RootDefaultFolderName
	}
	return &Store{
		StoreOpts: opts,
	}
}
func (s *Store) Has(key string) bool{
	PathKey := s.PathTransformFromFunc(key)
	fullPathWithRoot := fmt.Sprintf("%s/%s",s.Root,PathKey.FullPath())
	_,err := os.Stat(fullPathWithRoot)
	if err== fs.ErrNotExist{
		return false
	}
	return true
}

func (s *Store) Delete(key string) error{
	PathKey := s.PathTransformFromFunc(key)
	defer func(){
		log.Printf("File [%s] deleted from the disk",PathKey.Filename)
	}()

	firstFileNameWithRoot := fmt.Sprintf("%s/%s",s.Root,PathKey.FirstFileName())
	err := os.RemoveAll(firstFileNameWithRoot)
	return err
}

func (s *Store) Read(key string) (io.Reader,error){
	f,err:= s.ReadStream(key)
	if err!=nil{
		return nil,err
	}
	defer f.Close() 

	buf := new(bytes.Buffer)
	_,err = io.Copy(buf,f)
	return buf,err
}
func (s *Store) ReadStream(key string) (io.ReadCloser,error){ //io.ReadCloser is similar to a file returned by os pkg and using readCloser we will be able to close the opened file
	PathKey := s.PathTransformFromFunc(key)
	pathNameWithRoot := fmt.Sprintf("%s/%s",s.Root, PathKey.FullPath())
	return  os.Open(pathNameWithRoot)
}

func (s *Store) WriteStream(key string, r io.Reader) error{
	pathKey := s.PathTransformFromFunc(key)
	pathNameWithRoot := fmt.Sprintf("%s/%s",s.Root, pathKey.PathName)
	if err:= os.MkdirAll(pathNameWithRoot,os.ModePerm); err!=nil{
		return err
	}
	// buf := new(bytes.Buffer)
	// io.Copy(buf,r)
	// filenameBytes := md5.Sum(buf.Bytes())
	// filename:=hex.EncodeToString(filenameBytes[:])
	filename:= pathKey.FullPath()
	fullPathWithRoot := fmt.Sprintf("%s/%s",s.Root, filename)
	f,err := os.Create(fullPathWithRoot)
	if err!=nil{
		return err
	}
	defer f.Close()
	n , err := io.Copy(f,r) 
	log.Printf("Written %d bytes to disk %s:",n,fullPathWithRoot)
	return nil


}