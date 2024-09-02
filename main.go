package main

import (
	"fmt"
	"log"
	"rghdrizzle/hdfs/p2p"
)

func main(){
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr: ":3333",
		Decoder: p2p.GOBDecoder{},
		HandShakeFunc: p2p.NOPHandshakeFunc,
	}
	tr := p2p.NewTCPTransport(tcpOpts)
	fmt.Println("Connection waiting in port 3333")
	if err:= tr.ListenAndAccept(); err!=nil{
		log.Fatal(err)
	}

	select {}
}