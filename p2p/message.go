package p2p

import "net"

//RPC represents any data that is sent over each transport between
// two nodes in the network
type RPC struct{
	From net.Addr
	Payload []byte
}