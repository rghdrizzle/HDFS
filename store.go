package main

import (
	//"fmt"
	"io"
	"log"
	"os"
)


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