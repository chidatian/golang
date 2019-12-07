package main

import (
	"gochat/server"
)

var ser *server.Server

func main() {
	ser = server.NewServer()
	ser.Run()
}