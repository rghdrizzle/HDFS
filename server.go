package main

import (
	// "io"
	"fmt"
	"rghdrizzle/hdfs/p2p"
)

type FileServerOpts struct {
	StorageRoot           string
	PathTransformFromFunc PathTransformFromFunc
	Transport             p2p.Transport
}
type FileServer struct {
	FileServerOpts 
	store       *Store
	quitch chan struct{}
}

func NewFileServer(opts FileServerOpts) *FileServer {
	storageOpts := StoreOpts{
		Root:                  opts.StorageRoot,
		PathTransformFromFunc: opts.PathTransformFromFunc,
	}
	return &FileServer{
		FileServerOpts: opts,
		store:          NewStore(storageOpts),
		quitch: make(chan struct{}),
	}
}

func (fs *FileServer) Start() error {
	if err:= fs.FileServerOpts.Transport.ListenAndAccept();err!=nil{
		return err
	}
	fs.loop()
	return nil
}
func (fs *FileServer) Stop(){
	close(fs.quitch)
}
func (fs *FileServer) loop(){
	defer func(){
		fmt.Println("Server stopped")
		fs.Transport.Close()
		
	}()
	for{
		select{
		case msg:= <- fs.Transport.Consume():
			fmt.Println(msg)
		case <-fs.quitch:
			return 
		}
	}
	
}