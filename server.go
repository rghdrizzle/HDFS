package main

import "rghdrizzle/hdfs/p2p"

type FileServerOpts struct {
	StorageRoot           string
	PathTransformFromFunc PathTransformFromFunc
	Transport             p2p.Transport
}
type FileServer struct {
	FileServerOpts FileServerOpts
	store       *Store
}

func NewFileServer(opts FileServerOpts) *FileServer {
	storageOpts := StoreOpts{
		Root:                  opts.StorageRoot,
		PathTransformFromFunc: opts.PathTransformFromFunc,
	}
	return &FileServer{
		FileServerOpts: opts,
		store:          NewStore(storageOpts),
	}
}

func (fs *FileServer) Start() error {
	if err:= fs.FileServerOpts.Transport.ListenAndAccept();err!=nil{
		return err
	}
	return nil
}