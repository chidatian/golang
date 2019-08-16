package main

import (
	"fmt"
	"gochat/conf"
	"gochat/common/message"
	"net"
	"bufio"
	"os"
)

type Client struct {
	Config conf.Configure
	Conn net.Conn
	Input *bufio.Reader
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
}

func (this *Client) Response() {
	for {
		res := this.ReadData()
		fmt.Println(res)
	}
}

func (this *Client) Run() {
	this.InitConn()
	go this.Response()
	var isUserName bool = true
	// fmt.Println("Please enter your username: ")
	for {
		// func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)
		info, _, err := this.Input.ReadLine()
		if err != nil {
			continue
		}
		this.SendData(info)
		if string(info) == "exit" {
			break
		}
		if isUserName {
			isUserName = false
			fmt.Printf("Welcome to my home ------ [ %s ] ------ You can say now\n", string(info))
		}
	}
}

var client Client

func main() {
	client.Input = bufio.NewReader(os.Stdin)
	client.Config = conf.ConfigureInfo
	client.Run()
}