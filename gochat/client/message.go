<<<<<<< HEAD:gochat/server/user/user.go
package user

import (
	"net"
	"gochat/common/message"
)

type User struct {
	Name string
	Conn net.Conn
	Msg message.Msg
}

// 向用户 发送数据
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

// 接收用户的数据
func (this *User) ReadData() (string, error){
	resp := make([]byte, 10240)
	number := make([]byte, 5)
	// Read(b []byte) (n int, err error)
	_, err := this.Conn.Read(number[:4])
	if err != nil {
		return "", err
	}

	right := this.Msg.ToUint32(number[:4])
	n, err := this.Conn.Read(resp[:right])
	if err != nil {
		return "", err
	}
	return string(resp[:n]), nil
=======
package client


func (this *Client) SendData(info []byte) {
	// Write(b []byte) (n int, err error)
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

func (this *Client) ReadData() string{
	resp := make([]byte, 10240)
	number := make([]byte, 5)
	// Read(b []byte) (n int, err error)
	_, err := this.Conn.Read(number[:4])
	right := this.Msg.ToUint32(number[:4])
	if err != nil {
		panic(err)
	}

	n, err := this.Conn.Read(resp[:right])
	if err != nil {
		panic(err)
	}
	return string(resp[:n])
>>>>>>> 65888be0467c07f28b57cd4073746ad38fcc0782:gochat/client/message.go
}