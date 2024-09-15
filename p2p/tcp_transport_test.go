package p2p

import(
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T){
	tcpOpts := TCPTransportOpts{
		ListenAddr: ":3333",
		Decoder: DefaultDecoder{},
		HandShakeFunc: NOPHandshakeFunc,
	}
	listenAddr := ":3333"
	tr := NewTCPTransport(tcpOpts)
	assert.Equal(t,tr.ListenAddr,listenAddr)

	assert.Nil(t,tr.ListenAndAccept())
	

}