package main

import (
	//"fmt"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"log"
	"os"
	"strings"
)

func CASpathTransformFunc( key string ) string{
	hash := sha1.Sum([]byte(key)) // [20]bytes => to a slice([]byte) by doing => [:]
	hashString := hex.EncodeToString(hash[:])

	blockSize := 5
	sliceLength := len(hashString) / blockSize
	paths := make([]string, sliceLength)
	for i:=0;i<sliceLength;i++{
		from, to := i*blockSize, (i*blockSize)+blockSize
		paths[i]= hashString[from:to]
	}
	return strings.Join(paths,"/")
}

type PathTransformFromFunc func(string) string
type StoreOpts struct {
	PathTransformFromFunc PathTransformFromFunc
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
	pathName := s.StoreOpts.PathTransformFromFunc(key)
	if err:= os.MkdirAll(pathName,os.ModePerm); err!=nil{
		return err
	}
	filename := "test.txt"
	pathAndFilename := pathName + "/" + filename
	f,err := os.Create(pathAndFilename)
	if err!=nil{
		return err
	}
	n , err := io.Copy(f,r)
	log.Printf("Written %d bytes to disk %s:",n,pathAndFilename)
	return nil


}