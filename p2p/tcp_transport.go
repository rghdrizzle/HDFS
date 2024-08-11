package p2p

import (
	"fmt"
	"net"
	"sync"
)
//TCPPeer represents the remote node established over the tcp connection.
type TCPPeer struct{
	conn net.Conn
	//if we dial and retrieve a conn then outbound => true
	// if we accept and retrieve a conn then outbound => false
	outbound bool
}

type TCPTransport struct{
	listenAddress string
	listener net.Listener
	mu sync.RWMutex // It protects the peer below
	peer  map[net.Addr]Peer

}

func NewTCPTransport(listenAddr string) *TCPTransport{
	return &TCPTransport{
		listenAddress: listenAddr,
	}
}
func (t *TCPTransport) ListenAndAccept() error{
	var err error
	t.listener, err = net.Listen("tcp",t.listenAddress)
	if err!=nil{
		return err
	}

	go t.acceptLoop()
	
	return nil
}

func (t *TCPTransport) acceptLoop(){
	for{
		conn , err := t.listener.Accept()
		if err!=nil{
			fmt.Printf("TCP Accept failure %s",err)
		}
		go t.handleConn(conn) // we use a goroutine here to prevent a delay in the connection because handling will consume time 
		//and when one client is handling the other has to wait for the loop to end so they can accept in another thread
		//if we use a go routine then handling will be done concurrently so the other clients can also connect

	}
}

func (t *TCPTransport) handleConn(conn net.Conn){
	fmt.Printf("New incoming connection:%v\n",conn)
}