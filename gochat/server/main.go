package main

import (
	"fmt"
	"gochat/conf"
	"github.com/astaxie/goredis"
	"net"
)

type Server struct {
	Config conf.Configure
	Redis goredis.Client

}

func (this *Server) HandleConnection(con net.Conn) {
	defer con.Close()
	var resp [1024]byte
	for {
		// Read(b []byte) (n int, err error)
		n, err := con.Read(resp[:100])
		if err != nil {
			panic(err)
		}
		fmt.Println(string(resp[:n]))
	}
}

func (this *Server) Run(tcp string, host string) {
	fmt.Printf("Server on %s\n", host)
	ln, err := net.Listen(tcp, host)
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