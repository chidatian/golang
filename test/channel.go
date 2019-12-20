package main

import (
	"fmt"
	"time"
)

type Customer struct {

}

type Producter struct {

}

type Factory struct {
	Cer *Customer
	Per *Producter
	Com chan string
	Exit chan bool
}

func New() *Factory {
	return &Factory{
		Cer : new(Customer),
		Per : new(Producter),
		Com : make(chan string),
		Exit : make(chan bool),
	}
}

func main() {
	f := New()
	defer close(f.Exit)
	defer close(f.Com)
	var arr [3]string = [3]string{"a", "b", "c"}
	go func() {
		A:
		for {
			select {
			case item := <- f.Com:
				fmt.Println(item)
			default:
				fmt.Println("wait ...")
				break A;
			}
			time.Sleep(2 * time.Second)
		}
		f.Exit <- true
	}()
	go func() {
		for _,v := range arr {
			fmt.Println("-----")
			f.Com <- v
		}
	}()
	<- f.Exit
}