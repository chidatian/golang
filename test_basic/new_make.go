package main

import(
	"fmt"
)

func main() {
	a := new(int)
	b := make([]int,2)
	fmt.Println(*a)
	fmt.Println(b)
}