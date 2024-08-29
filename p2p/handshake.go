package p2p

//HandleShaekFunc is
type HandShakeFunc func(Peer) error

func NOPHandshakeFunc(Peer) error{return nil}