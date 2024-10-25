package main

import (
	"log"
	"rghdrizzle/hdfs/p2p"
	//"time"
)

func makeServer(listenAddr string,nodes ...string) *FileServer{ // ...string is used to recieve n number of arguments of type string ( n can be 0 as well )

	tcpTransportOpts := p2p.TCPTransportOpts{
		ListenAddr: listenAddr,
		Decoder: p2p.DefaultDecoder{},
		HandShakeFunc: p2p.	NOPHandshakeFunc,
		
	}
	tcpTransport := p2p.NewTCPTransport(tcpTransportOpts)
	fileServerOpts := FileServerOpts	{
		StorageRoot: listenAddr+"_network",
		PathTransformFromFunc: CASpathTransformFunc,
		Transport: tcpTransport,
		BootstrapNodes: nodes,
	}
	s := NewFileServer(fileServerOpts)
	return s
}

func main(){
	
	s1 := makeServer(":3000","")
	s2 := makeServer(":4000",":3000")
	go func(){
		log.Fatal(s1.Start())	
		
	}()
	s2.Start()
	// go func(){
	// 	time.Sleep(time.Second * 3)
	// 	s.Stop()

	// }()
	// if err := s.Start();err!=nil{
	// 	log.Fatal(err)
	// }
	
}