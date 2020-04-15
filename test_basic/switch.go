package main

import "fmt"

func main() {
	var a = 2
	switch a {
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("nothing")
	}

	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("----", err)
			}
		}()
		panic("123")
	}()
	select{}
}