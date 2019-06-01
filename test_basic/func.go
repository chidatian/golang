package main

import(
	"fmt"
)

func main() {
	fmt.Println("test-----")
	s := sum(10, 10)
	fmt.Println(s)
}

func sum(a int, b int) (int,int) {
	return a+b,10
}