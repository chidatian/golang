package main

import (
	"fmt"
)

// 生产者 消费者模式
func main() {
	data := make(chan int)
	exit := make(chan bool)

	// 消费者
	go func () {
		for d := range data {
			fmt.Println("=====")
			fmt.Println(d)
		}
		exit <- true
	}()
	
	// 生产者
	for i := 1; i< 10; i++ {
		fmt.Println("-----")
		data <- i
	}
	close(data)
	<-exit
}