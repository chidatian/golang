package main


import (
	"fmt"
	"time"
)

func main() {
	testChan := make(chan string, 3)
	exit := make(chan bool)

	go func() {
		for {
			select {
			case item := <-testChan :
				fmt.Println(item)
			}
		}
		exit <- true
	}()

	testChan <- "china"
	testChan <- "china"
	testChan <- "china"
	fmt.Println(len(testChan))
	testChan <- "china"
	time.Sleep(3* time.Second)
	fmt.Println(len(testChan))
	testChan <- "china"
	testChan <- "china"

	<- exit
}