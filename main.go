package main

import (
	"bytes"

	"log"
	"rghdrizzle/hdfs/p2p"
	"time"
	"strings"
)

func makeServer(listenAddr string,nodes ...string) *FileServer{ // ...string is used to recieve n number of arguments of type string ( n can be 0 as well )

	tcpTransportOpts := p2p.TCPTransportOpts{
		ListenAddr: listenAddr,
		Decoder: p2p.DefaultDecoder{},
		HandShakeFunc: p2p.	NOPHandshakeFunc,
		
	}
	tcpTransport := p2p.NewTCPTransport(tcpTransportOpts)
	fileServerOpts := FileServerOpts	{
		StorageRoot: strings.Replace(listenAddr, ":", "_", -1)+"_network", // colon will not be considered to be a valid directory name in windows so we replace the colon here, if we are storing the files in a linux based system then simply having ":3000_network" should be fine
		PathTransformFromFunc: CASpathTransformFunc,
		Transport: tcpTransport,
		BootstrapNodes: nodes,
	}
	s := NewFileServer(fileServerOpts)

	tcpTransport.OnPeer = s.OnPeer
	return s
}

func main(){
	
	s1 := makeServer(":3000","")
	s2 := makeServer(":4000",":3000")
	go func(){
		log.Fatal(s1.Start())	
		
	}()
	go s2.Start()

	time.Sleep(1*time.Second)

	data := bytes.NewReader([]byte("large file"))

	s2.StoreData("privatedata",data)
	select{}
	// go func(){
	// 	time.Sleep(time.Second * 3)
	// 	s.Stop()

	// }()
	// if err := s.Start();err!=nil{
	// 	log.Fatal(err)
	// }
	
}