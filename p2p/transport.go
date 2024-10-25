package p2p

import "net"

//Peer is an interface that is the representation the remote node
type Peer interface{
	RemoteAddr() net.Addr
	Close() error
}

// Transport is an object that handles the communication between multiple nodes in the net
type Transport interface{
	Dial(string) error
	ListenAndAccept() error
	Consume()<-chan RPC
	Close() error
}



