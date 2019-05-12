package main

import (
	"fmt"
)

func main() {
	var c chan int = make(chan int)
	fmt.Println(c)
	i := make([]int, 2, 10)
	fmt.Println(i)
}

func test_go() {
	i := 0
	f := func() {
		i++
		println(i)
	}
	f()
	println(i)
}

func test_chan() {
	i := 10
	tag := make(chan bool)
	go func() {
		for i < 20 {
			i++
			fmt.Println(i)
		}
		tag <- true
	}()
	fmt.Printf("num: %d\n",i)
	<-tag
}