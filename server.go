package main

import (
	// "io"
	//"bytes"
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
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

type PayLoad struct{
	Key string
	Data []byte

}
func (fs *FileServer) broadcast(p *PayLoad) error{
	peers:= []io.Writer{}

	for _,peer := range(fs.peers){
		peers= append(peers,peer)
	}
	mw := io.MultiWriter(peers...)
	return gob.NewEncoder(mw).Encode(p)
	
}
func (fs  *FileServer) StoreData(key string, r io.Reader) error{
	// Storing file to the disk
	// then we broadcast the file to other known peers
	if err:= fs.store.Write(key,r);err!=nil{
		return err
	}
	buf := new(bytes.Buffer)

	_, err:= io.Copy(buf,r)
	if err!=nil{
		return err
	}
	p := &PayLoad{
		Key: key,
		Data: buf.Bytes(),
	}

	fmt.Println(buf.Bytes())

	return fs.broadcast(p)
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

func (fs *FileServer) OnPeer(p p2p.Peer)error{
	fs.peerLock.Lock()
	defer fs.peerLock.Unlock()
	fs.peers[p.RemoteAddr().String()]= p
	log.Printf("Connected to remote %s",p.RemoteAddr().String())
	return nil
}


func (fs *FileServer) loop(){
	defer func(){
		fmt.Println("Server stopped")
		fs.Transport.Close()
		
	}()
	for{
		select{
		case msg:= <- fs.Transport.Consume():
			fmt.Printf("recev\n")
			var p PayLoad
			fmt.Println(msg.Payload)
			if err:=gob.NewDecoder(bytes.NewReader(msg.Payload)).Decode(&p);err!=nil{
				log.Fatal(err)
			}
			fmt.Printf("%+v\n",p)
		case <-fs.quitch:
			return 
		}
	}
	
}
