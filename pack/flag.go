package main

import (
	"fmt"
	"flag"
)

func main() {
	name := flag.String("u", "", "input your url")
	file := flag.String("f", "", "input your filename")
	flag.Parse()
	fmt.Printf("%s, %s\n", *name, *file)
}