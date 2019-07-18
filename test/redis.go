package main

import (
	"fmt"
	"time"
	"github.com/astaxie/goredis"
)

const (
	ADDR = "127.0.0.1:6379"
	AUTH = "root"
	DB   = 0
)

var redis goredis.Client

func init() {
	redis.Addr = ADDR
	redis.Db = DB
}

func order(name string) {
	for {
		fmt.Println(name)
		time.Sleep(2 * time.Second)
	}
}

func main() {
	for {
		el,_ := redis.Rpop("shop")
		if len(el) == 0 {
			time.Sleep(2 * time.Second)
			continue
		}
		go order(string(el))
	}
}