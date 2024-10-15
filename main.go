package main

import (
	"log"
	"rghdrizzle/hdfs/p2p"
	"time"
)

func main(){
	tcpTransportOpts := p2p.TCPTransportOpts{
		ListenAddr: ":3000",
		Decoder: p2p.DefaultDecoder{},
		HandShakeFunc: p2p.	NOPHandshakeFunc,
	}
	tcpTransport := p2p.NewTCPTransport(tcpTransportOpts)
	fileServerOpts := FileServerOpts	{
		StorageRoot: "3000_network",
		Transport: tcpTransport,
	}
	s := NewFileServer(fileServerOpts)
	go func(){
		time.Sleep(time.Second * 3)
		s.Stop()

	}()
	if err := s.Start();err!=nil{
		log.Fatal(err)
	}
	
}