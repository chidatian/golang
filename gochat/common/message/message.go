package message

import (
	"encoding/binary"
	// "encoding/json"
)

type Msg struct {}

func (this *Msg) encode(array interface{}) string {
	return " "
}

func (this *Msg) decode(json string) interface{} {
	return nil
}

func (this *Msg) Uint32ToByte(info []byte) []byte{
    var bytes [4]byte
	number := uint32(len(info))
    // PutUint16([]byte, uint16)
	binary.BigEndian.PutUint32(bytes[:], number)
	return bytes[:]
}

func (this *Msg) ToUint32(info []byte) uint32{
	// Uint16([]byte) uint16
	i := binary.BigEndian.Uint32(info)
	return i
}