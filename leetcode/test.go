package main

import (
	"fmt"
	// "reflect"
)

func main() {
	repeat()
}

func repeat() {
	var str string = "abcabcbb"
	var res map[int32]int = make(map[int32]int)
	for _,v := range(str) {
		res[v] = 1
	}
	n := len(res)
	fmt.Println(n)
}