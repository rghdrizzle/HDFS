package main

import (
	"bytes"
	"fmt"
	"io"

	//"fmt"
	"testing"
)
func TestPathTransformFunc(t *testing.T){
	key:= "testingtransfromfunc"
	pathKey := CASpathTransformFunc(key)
	expectedFilename:="3c4dfead2599fb100c3c76053d04ca81af64ab70"
	expectedPath := "3c4df/ead25/99fb1/00c3c/76053/d04ca/81af6/4ab70"
	fmt.Printf(pathKey.PathName)
	fmt.Println(pathKey.Filename)
	if pathKey.PathName != expectedPath{
		t.Errorf("Error: want: %s, got: %s",pathKey.PathName,expectedPath)
	}
	if pathKey.Filename != expectedFilename{
		t.Errorf("Error: want: %s, got: %s",pathKey.Filename,expectedFilename)
	}
}
func TestStore(t *testing.T){
	opts := StoreOpts{
		PathTransformFromFunc: CASpathTransformFunc,
	}
	s := NewStore(opts)
	key := "hello_world"
	data := []byte("hello world this is a storage system")
	if err := s.WriteStream(key,bytes.NewReader(data)); err!=nil{
		t.Error(err)
	}
	if ok:=s.Has(key); !ok{
		t.Errorf("Expected key %s missing:",key)
	}

	r , err := s.Read(key)
	if err!=nil{
		t.Error(err)
	}
	b,_ := io.ReadAll(r)
	if string(b) != string(data){
		t.Errorf("want %s, got %s",data,b)
	}
	fmt.Println(b)
	s.Delete(key)
}

func TestStoreDeleteKey(t *testing.T) {
	opts := StoreOpts{
		PathTransformFromFunc: CASpathTransformFunc,
	}
	s := NewStore(opts)
	key := "hello_world"
	data := []byte("hello world this is a storage system")
	if err := s.WriteStream(key,bytes.NewReader(data)); err!=nil{
		t.Error(err)
	}
	if err:=s.Delete(key); err!=nil{
		t.Error(err)
	}

}