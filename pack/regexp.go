package main

import (
	"fmt"
	"regexp"
)

func test() {
	var s string = "chidatian"
	var pattern string = "da"
	isB, err := regexp.MatchString(pattern, s)
	if err != nil {
		panic(err)
	}
	fmt.Println(isB)
}

func main() {
	test()
}