package main

import (
	"fmt"
	"net"
	"context"
)

func main() {
	nd := net.Dialer{}
	res, err := nd.DialContext(context.Background(), "tcp", "192.168.1.86:3308")
	if err != nil {
		panic(err)
	}

	fmt.Println("-------------------------")
	fmt.Println(res)

	// func (c *TCPConn) SetReadDeadline(t time.Time) error
	var query string = "select * from weather_tmp where id = 10"
	if _, err := res.Write([]byte(query)); err != nil {
		panic(err)
	}

	var b []byte
	fmt.Println(res.Read(b))
}