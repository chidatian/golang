package message

import (
	"encoding/json"
)

type Msg struct {}

func (this *Msg) GetConn(host string) net.Conn{
	conn, err := net.Dial("tcp", host)
	if err != nil {
		// handle error
		panic(err)
	}
	return conn
}