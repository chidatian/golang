package main

import (
	"fmt"
	"github.com/astaxie/goredis"
	"gochat/conf"
	"gochat/common/message"
	"gochat/server/user"
	"net"
)

type Server struct {
	Config conf.Configure
	Redis goredis.Client
	Msg message.Msg
	User []user.User
}

func (this *Server) HandleUser(info []byte) {
	for _, item := range this.User {
		item.Response(info)
	}
}

func (this *Server) HandleConnection(con net.Conn) {
	defer con.Close()
	for {
		resp := make([]byte, 10240)
		number := make([]byte, 5)
		// Read(b []byte) (n int, err error)
		_, err := con.Read(number[:4])
		right := this.Msg.ToUint32(number[:4])
		if err != nil {
			panic(err)
		}

		n, err := con.Read(resp[:right])
		if err != nil {
			panic(err)
		}

		if string(resp[:n]) == "exit" {
			break
		}
		// fmt.Printf("CLIENT : %s \n", s)
		this.HandleUser(resp[:n])
	}
}

func (this *Server) Run(tcp string, host string) {
	fmt.Printf("Server on %s\n", host)
	ln, err := net.Listen(tcp, host)
	defer ln.Close()
	if err != nil {
		// handle error
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			continue
		}
		this.User = append(this.User, user.User{Conn: conn})
		go this.HandleConnection(conn)
	}
}

var server Server

func main() {
	server.Run("tcp", server.Config.Server.Host)
}

func init() {
	server.Config = conf.ConfigureInfo
	server.Redis.Addr = server.Config.Redis.Host
	server.Redis.Db = 0
	// redis.Set("chat", []byte("test"))
}