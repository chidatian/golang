package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a = 'A'
	fmt.Println(a)
	res := reflect.TypeOf(a)
	fmt.Println(res)
	re := reflect.ValueOf(a)
	fmt.Println(re)
}