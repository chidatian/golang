package main

import (
	"fmt"
	"reflect"
)

func test() {
	var a []int = []int{1,2,3}
	var b int = 123
	re := reflect.TypeOf(b)
	fmt.Println(reflect.PtrTo(re))
	for _,item := range a {
		// func PtrTo(t Type) Type
		res := reflect.TypeOf(item)
		fmt.Println(reflect.PtrTo(res))
	}
}

func main() {
	test()
}