package user

import (
	"net"
	"gochat/common/message"
)

type User struct {
	Conn net.Conn
	Msg message.Msg
}

func (this *User) Response(info []byte) {
	this.SendData(info)
}

func (this *User) SendData(info []byte) {
	number := this.Msg.Uint32ToByte(info)
	_, err := this.Conn.Write(number)
	if err != nil {
		panic(err)
	}

	_, err = this.Conn.Write(info)
	if err != nil {
		panic(err)
	}
}

func (this *User) ReadData() string{
	return ""
}