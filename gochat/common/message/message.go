<<<<<<< HEAD
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
=======
package message

import (
	"encoding/binary"
	// "encoding/json"
	"github.com/fatih/color"
)

type Msg struct {}

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

func (this *Msg) System(info string) {
	color.Yellow(info)
}

func (this *Msg) User(info string) {

>>>>>>> 65888be0467c07f28b57cd4073746ad38fcc0782
}