package main

import (
	"bytes"
	"testing"
)
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