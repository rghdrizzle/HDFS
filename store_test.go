package main

import (
	"bytes"
	"fmt"
	"testing"
)
func TestPathTransformFunc(t *testing.T){
	key:= "testingtransfromfunc"
	pathname := CASpathTransformFunc(key)
	fmt.Println(pathname)
}
func TestStore(t *testing.T){
	opts := StoreOpts{
		PathTransformFromFunc: DefaultTransformFunc,
	}
	s := NewStore(opts)
	data := bytes.NewReader([]byte("some file"))
	if err := s.WriteStream("hello_world",data); err!=nil{
		t.Error(err)
	}
}