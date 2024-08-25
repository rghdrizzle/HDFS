package main

import (
	"fmt"
	"log"
	"rghdrizzle/hdfs/p2p"
)

func main(){
	tr := p2p.NewTCPTransport(":3333")
	fmt.Println("Connection waiting in port 3333")
	if err:= tr.ListenAndAccept(); err!=nil{
		log.Fatal(err)
	}

	select {}
}