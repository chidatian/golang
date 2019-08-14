package main

import (
	"fmt"
	"gochat/conf"
	"gochat/client/message"
	"net"
)

type Client struct {
	Config conf.Configure
	Conn net.Conn
	Msg message.Msg
}

func (this *Client) InitConn() {
	var err error
	this.Conn, err = net.Dial("tcp", this.Config.Server.Host)
	if err != nil {
		// handle error
		panic(err)
	}
}

func (this *Client) SendData(info string) {
	// Write(b []byte) (n int, err error)
	_, err := this.Conn.Write([]byte(info))
	if err != nil {
		panic(err)
	}
}

func (this *Client) Run() {
	this.InitConn()
	s := this.Msg.InitBuf()
	var info string
	for {
		fmt.Scanf("%s", &info)
		this.SendData(info)
		if info == "exit" {
			break
		}
	}
}

var client Client

func main() {
	client.Config = conf.ConfigureInfo
	client.Run()
}