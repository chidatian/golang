package main

import (
	"fmt"
	"net/http"
	"reflect"
)

func main() {
	var url string = "http://www.baidu.com"
	res,_ := http.Get(url)
	r := reflect.TypeOf(res)
	fmt.Println(r)
}