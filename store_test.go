package main

import (
	"bytes"
	//"fmt"
	"testing"
)
func TestPathTransformFunc(t *testing.T){
	key:= "testingtransfromfunc"
	pathname := CASpathTransformFunc(key)
	expectedPath := "3c4df/ead25/99fb1/00c3c/76053/d04ca/81af6/4ab70"
	if pathname != expectedPath{
		t.Errorf("Error: want: %s, got: %s",expectedPath,pathname)
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