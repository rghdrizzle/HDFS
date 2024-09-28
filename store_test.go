package main

import (
	"bytes"
	"fmt"
	//"fmt"
	"testing"
)
func TestPathTransformFunc(t *testing.T){
	key:= "testingtransfromfunc"
	pathKey := CASpathTransformFunc(key)
	expectedOriginal:="3c4dfead2599fb100c3c76053d04ca81af64ab70"
	expectedPath := "3c4df/ead25/99fb1/00c3c/76053/d04ca/81af6/4ab70"
	fmt.Printf(pathKey.PathName)
	fmt.Println(pathKey.Original)
	if pathKey.PathName != expectedPath{
		t.Errorf("Error: want: %s, got: %s",pathKey.PathName,expectedPath)
	}
	if pathKey.Original != expectedOriginal{
		t.Errorf("Error: want: %s, got: %s",pathKey.Original,expectedOriginal)
	}
}
func TestStore(t *testing.T){
	opts := StoreOpts{
		PathTransformFromFunc: CASpathTransformFunc,
	}
	s := NewStore(opts)
	data := bytes.NewReader([]byte("some file"))
	if err := s.WriteStream("hello_world",data); err!=nil{
		t.Error(err)
	}
}