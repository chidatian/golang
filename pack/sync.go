 package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	i int = 1
	l sync.Mutex
	w sync.WaitGroup
)

func decr(a int) {
	defer w.Done()
	l.Lock()
	defer l.Unlock()
	i = i + 1
	time.Sleep(time.Second)
	fmt.Printf("-%d--%d\n", a, i)
}

func main() {
	for a := 0; a < 100; a++ {
		go decr(a)
		w.Add(1)
	}
	w.Wait()
}