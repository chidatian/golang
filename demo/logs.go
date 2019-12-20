package main

import (
	"os"
	"log"
)

func main() {
	_, err := os.Open("./a.txt")
	if err != nil {
		log.Fatal(err)
	}
}