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
}