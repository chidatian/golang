package main

import (
	"fmt"
	"time"
	"runtime"
)

func main() {
	data := make(chan int)
	// 协程
	go func(){
		for {
			select {
			case v := <- data :
				fmt.Println("---------", v)
			}
		}
	}()
	// 填充数据
	for i := 0; i < 10; i++ {
		data <- i
	}

	// 循环打印协程的个数
	for {
		fmt.Println("runtime.number:= ", runtime.NumGoroutine())
		time.Sleep(2 * time.Second);
	}
}