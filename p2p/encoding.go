package p2p

import (
	"encoding/gob"
	"io"
)

type Decoder interface{
	Decode(io.Reader,*Message) error
}

type GOBDecoder struct{}

func (dec GOBDecoder) Decode(r io.Reader,v *Message) error{
	return gob.NewDecoder(r).Decode(v)
}

type DefaultDecoder struct{

}
func (dec DefaultDecoder) Decode(r io.Reader,msg *Message) error{
	buf:= make([]byte,1028)

	m ,err:= r.Read(buf)
	if err!=nil{
		return err
	}
	msg.Payload = buf[:m]

	return nil

}
