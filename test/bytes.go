package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "asdf fdsa  qwer"
	ret := strings.Split(s, " ")
	for k,v := range ret {
		fmt.Println(k,v,len(v))
	}
}