package main

import (
	//"fmt"
	// "bytes"
	// "crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

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
		Original: hashString,
		}
}

type PathTransformFromFunc func(string) PathKey
type StoreOpts struct {
	PathTransformFromFunc PathTransformFromFunc
}

type PathKey struct{
	PathName string
	Original string
}

func (p PathKey) Filename() string{
	return fmt.Sprintf("%s/%s",p.PathName,p.Original)
}

var DefaultTransformFunc = func(key string) string {
	return key
}


type Store struct {
	StoreOpts
}

func NewStore(opts StoreOpts) *Store{
	return &Store{
		StoreOpts: opts,
	}
}

func (s *Store) WriteStream(key string, r io.Reader) error{
	pathKey := s.PathTransformFromFunc(key)
	if err:= os.MkdirAll(pathKey.PathName,os.ModePerm); err!=nil{
		return err
	}
	// buf := new(bytes.Buffer)
	// io.Copy(buf,r)
	// filenameBytes := md5.Sum(buf.Bytes())
	// filename:=hex.EncodeToString(filenameBytes[:])
	filename:= pathKey.Filename()
	pathAndFilename := filename
	f,err := os.Create(pathAndFilename)
	if err!=nil{
		return err
	}
	n , err := io.Copy(f,r) // The reason why we did not do io.Copy(f,r) is because the reader r has already been exhausted 
	//i.e if r completed reading already then if u use it again it will return something empty 
	//which is why we used the buffer which copied the contents in reader r
	log.Printf("Written %d bytes to disk %s:",n,pathAndFilename)
	return nil


}