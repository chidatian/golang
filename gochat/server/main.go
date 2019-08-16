package main

import (
	"fmt"
	"github.com/astaxie/goredis"
	"gochat/conf"
	"gochat/server/user"
	"net"
)

type Server struct {
	Config conf.Configure
	Redis goredis.Client
	User map[int]*user.User
	Message chan string
	LinkNum int
}

func (this *Server) Response() {
	for {
		select {
		case info := <- this.Message :
			for _, item := range this.User {
				if item == nil {
					continue	
				}
				item.SendData([]byte(info))
			}
		}
	}
}

func (this *Server) ResponseUsers(index int) {
	var users string = "Online username: \n"
	for _, item := range this.User {
		if item == nil || item.Name == ""{
			continue	
		}
		users = users + " [ " + item.Name + " ] "
	}
	users = users + "\nPlease enter your username: "
	this.User[index].SendData([]byte(users))
}

func (this *Server) HandleConnection(index int) {
	defer this.User[index].Conn.Close()
	this.ResponseUsers(index)
	var info string
	var err error
	var isUserName bool = true
	for {
		info, err = this.User[index].ReadData()
		if err != nil {
			break
		}
		if isUserName {
			this.User[index].Name, isUserName = info, false
			continue
		}
		if info == "exit" {
			break
		}
		// fmt.Printf("CLIENT : %s \n", s)
		this.Message <- fmt.Sprintf("[ %s ] : %s", this.User[index].Name, info)
	}
	this.User[index] = nil
}

func (this *Server) Run(tcp string, host string) {
	fmt.Printf("Server on %s\n", host)
	ln, err := net.Listen(tcp, host)
	defer ln.Close()
	if err != nil {
		// handle error
		panic(err)
	}
	go this.Response()
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			continue
		}
		this.LinkNum++
		this.User[this.LinkNum] = &user.User{Conn: conn}
		fmt.Printf("[ CLIENT %d ] %T\n", this.LinkNum, conn)
		go this.HandleConnection(this.LinkNum)
	}
}

var server Server

func main() {
	server.Run("tcp", server.Config.Server.Host)
}

func init() {
	server.LinkNum = 0
	server.User = make(map[int]*user.User)
	server.Message = make(chan string, 1)
	server.Config = conf.ConfigureInfo
	server.Redis.Addr = server.Config.Redis.Host
	server.Redis.Db = 0
	// redis.Set("chat", []byte("test"))
}