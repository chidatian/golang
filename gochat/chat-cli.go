package main

import (
	"gochat/client"
)

var cli *client.Client

func main() {
	cli = client.NewClient()
	cli.Run()
}