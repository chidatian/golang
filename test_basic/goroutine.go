package main

import (
	"fmt"
	"sync"
	"runtime"
)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i < 10; i++ {
			fmt.Println("-----------")
			fmt.Println(i)
		}
	}()
	wg.Wait()
	fmt.Println("========")
}