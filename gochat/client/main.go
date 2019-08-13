package main

import (
	"fmt"
	"gochat/conf"
	"net"
	"encoding/json"
)

type Client struct {
	Config conf.Configure
	Conn net.Conn
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
	_, err := this.Conn.Write([]byte(info))
	if err != nil {
		panic(err)
	}
}

func (this *Client) Run() {
	this.InitConn()
	// Write(b []byte) (n int, err error)
	var info string
	for {
		fmt.Scanf("%s\n", &info)
		fmt.Println(info)
		this.SendData(info)
	}
}

var client Client

func main() {
	client.Config = conf.ConfigureInfo
	client.Run()
}