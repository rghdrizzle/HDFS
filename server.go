package main

import (
	// "io"
	"fmt"
	"log"
	"rghdrizzle/hdfs/p2p"
	"sync"
)

type FileServerOpts struct {
	StorageRoot           string
	PathTransformFromFunc PathTransformFromFunc
	Transport             p2p.Transport
	BootstrapNodes		  []string
}
type FileServer struct {
	FileServerOpts 

	peerLock sync.Mutex
	peers map[string]p2p.Peer
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
		peers: make(map[string]p2p.Peer),
	}
}

func (fs *FileServer) BootstrapNetwork() error{

	for _, addr := range(fs.BootstrapNodes){
		if len(addr)==0 {continue}
		go func(addr string){
			if err:= fs.Transport.Dial(addr); err!=nil{
				log.Println("Dial error:",err)
			}
		}(addr)
	}
	return nil
}

func (fs *FileServer) Start() error {
	if err:= fs.FileServerOpts.Transport.ListenAndAccept();err!=nil{
		return err
	}
	fs.BootstrapNetwork()
	fs.loop()
	return nil
}
func (fs *FileServer) Stop(){
	close(fs.quitch)
}

func (fs *FileServer) OnPeer(){
		
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
