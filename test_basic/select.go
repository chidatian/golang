package main

import (
	"fmt"
)

func main() {
	// select_test()
	select_test2()
}

func select_test2() {
	arr := []int{1,2}
	select {
	default:
		fmt.Println("--default --")
	}
	fmt.Println(arr)
}


func select_test() {
	exit := make(chan bool)
	data1, data2 := make(chan int), make(chan int)
	
	go func () {
		v, ok:= 0, false
		for {
			select {
			case v,ok = <- data1 :
			case v,ok = <- data2 :
			}

			fmt.Println(ok)
			if ok {
				fmt.Println(v)
			} else {
				break
			}
		}
		exit <- true
	}()

	data1 <- 1
	data2 <- 11
	data1 <- 2
	data2 <- 12
	data1 <- 3
	data2 <- 13
	close(data1)
	close(data2)
	<- exit
	fmt.Println("--end --")
}