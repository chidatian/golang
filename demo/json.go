package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	m := make(map[string]string)
	m["a"] = "this is a"
	m["b"] = "this is b"
	m["c"] = "this is c"
	fmt.Println(m)
	// func Marshal(v interface{}) ([]byte, error)
	res, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
	for k, item := range m {
		fmt.Println(k, item)
	}
	delete(m, "b")
	res, err = json.Marshal(m)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
}