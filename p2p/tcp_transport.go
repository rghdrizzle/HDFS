package p2p

import (
	"net"
	"sync"
)

type TCPTransport struct{
	listenAddress string
	listener net.Listener
	mu sync.RWMutex // It protects the peer below
	peer  map[net.Addr]Peer

}

func NewTCPTransport(listenAddr string) Transport{
	return &TCPTransport{
		listenAddress: listenAddr,
	}
}
