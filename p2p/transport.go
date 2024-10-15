package p2p


//Peer is an interface that is the representation the remote node
type Peer interface{
	Close() error
}

// Transport is an object that handles the communication between multiple nodes in the net
type Transport interface{
	ListenAndAccept() error
	Consume()<-chan RPC
	Close() error
}



