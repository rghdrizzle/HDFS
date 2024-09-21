package main

import (
	"fmt"
	"log"
	"rghdrizzle/hdfs/p2p"
)
func OnPeer(peer p2p.Peer) error{
	//fmt.Println("Doing some logic outside of tcp transport")
	peer.Close()
	return nil
}
func main(){
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr: ":3333",
		Decoder: p2p.DefaultDecoder{},
		HandShakeFunc: p2p.NOPHandshakeFunc,
		OnPeer: OnPeer,
	}

	tr := p2p.NewTCPTransport(tcpOpts)
	go func(){
		for{
			msg := <-tr.Consume()
			fmt.Println(msg)

		}
	}()
	fmt.Println("Connection waiting in port 3333")
	if err:= tr.ListenAndAccept(); err!=nil{
		log.Fatal(err)
	}

	select {}
}