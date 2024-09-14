package p2p

import (
	"fmt"
	"net"
	"sync"
)
//TCPPeer represents the remote node established over the tcp connection.
type TCPPeer struct{
	// conn is the underlying connection of the peer
	conn net.Conn
	//if we dial and retrieve a conn then outbound => true
	// if we accept and retrieve a conn then outbound => false
	outbound bool
}
type TCPTransportOpts struct{
	ListenAddr string
	Decoder Decoder
	HandShakeFunc HandShakeFunc
}

type TCPTransport struct{
	TCPTransportOpts
	listener net.Listener

	mu sync.RWMutex // It protects the peer below
	peer  map[net.Addr]Peer

}

func NewTCPPeer(conn net.Conn , outbound bool) *TCPPeer{
	return &TCPPeer{
		conn: conn,
		outbound: outbound,
	}
}


func NewTCPTransport(opts TCPTransportOpts) *TCPTransport{
	return &TCPTransport{
		TCPTransportOpts: opts,
	}
}
func (t *TCPTransport) ListenAndAccept() error{
	var err error
	t.listener, err = net.Listen("tcp",t.ListenAddr)
	if err!=nil{
		return err
	}
	fmt.Println("Listening....")
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
type temp struct{}
func (t *TCPTransport) handleConn(conn net.Conn){
	peer := NewTCPPeer(conn , true)
	fmt.Printf("New incoming connection:%v\n",peer)
	if err:= t.HandShakeFunc(peer); err!=nil{
		conn.Close()
		fmt.Printf("TCP handshake error: %s\n",err)
		return
	}
	
	msg := &Message{}
	//read loop
	for{
		if err :=t.Decoder.Decode(conn,msg); err!=nil{
			fmt.Printf("TCP error: %s\n",err)
			continue
		}
		msg.From = conn.RemoteAddr()
		fmt.Println("Message:", msg)
	}
	
	
}