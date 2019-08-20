package client

import (
	"fmt"
	"gochat/common/conf"
	"gochat/common/message"
	"net"
	"bufio"
	"os"
)

type Client struct {
	Config conf.Configure
	Input *bufio.Reader
	Conn net.Conn
	Msg message.Msg
}

func NewClient() *Client{
	return &Client{
		Config : conf.ConfigureInfo,
		Input : bufio.NewReader(os.Stdin),
	}
}

func (this *Client) InitConn() {
	var err error
	this.Conn, err = net.Dial("tcp", this.Config.Server.Host)
	if err != nil {
		// handle error
		panic(err)
	}
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
			this.Msg.System(fmt.Sprintf("Welcome to my home  [ %s ]  You can say now\n", string(info)))
		}
	}
}