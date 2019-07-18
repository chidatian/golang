package main

import (
	"fmt"
)

func main() {
	m := make(map[string]interface{})

	m["id"] = 1
	m["name"] = "tian"
	fmt.Println(m)
	fmt.Println(m["name"])
	fmt.Println(m[""])
}