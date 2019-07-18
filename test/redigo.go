package main

import (
	"fmt"
	"time"
	"github.com/gomodule/redigo/redis"
)

func main() {
	c, _ := redis.DialTimeout("tcp", "127.0.0.1:6379", 0, 1*time.Second, 1*time.Second)
	_, _ = c.Do("SELECT", "0")
	val := c.Send("GET", "a")
	fmt.Println(val)
}