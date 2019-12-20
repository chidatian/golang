<<<<<<< HEAD:gochat/server/main.go
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
=======
package server

import (
	"fmt"
	"github.com/astaxie/goredis"
	"github.com/fatih/color"
	"gochat/common/message"
	"gochat/common/conf"
	"net"
)

type Server struct {
	Redis goredis.Client
	Msg message.Msg
	Config conf.Configure
	Message chan string
	User map[uint16]*User
	LinkPool *Pool
}

func NewServer() *Server {
	return &Server{
		Redis : goredis.Client{Addr: conf.ConfigureInfo.Redis.Host, Db: 0},
		Config : conf.ConfigureInfo,
		Message : make(chan string, 1),
		User : make(map[uint16]*User),
		LinkPool : NewPool(),
	}
}

func (this *Server) Response() {
	for {
		select {
		case info := <- this.Message :
			for _, item := range this.User {
				if item == nil {
					continue	
				}
				this.SendData(item.Index, []byte(info))
			}
		}
	}
}

func (this *Server) ResponseUsers(index uint16) {
	var users string = "Online username: \n"
	for _, item := range this.User {
		if item == nil || item.Name == ""{
			continue
		}
		users = users + " [ " + item.Name + " ] "
	}
	users = users + "\nPlease enter your username: "
	this.SendData(index, []byte(users))
}

func (this *Server) HandleConnection(index uint16) {
	conn := this.LinkPool.Get(index)
	defer conn.Close()
	this.ResponseUsers(index)
	var info string
	var err error
	var isUserName bool = true
	for {
		info, err = this.ReadData(index)
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
	this.LinkPool.Del(index)
}

func (this *Server) Run() {
	color.Green("Server on %s\n", this.Config.Server.Host)
	ln, err := net.Listen("tcp", this.Config.Server.Host)
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
		index := this.LinkPool.Set(conn)
		this.User[index] = &User{Index : index}
		color.Yellow("[ CLIENT %d ] %T\n", index, conn)
		go this.HandleConnection(index)
	}
>>>>>>> 65888be0467c07f28b57cd4073746ad38fcc0782:gochat/server/server.go
}