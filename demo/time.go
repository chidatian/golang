package main

import (
	"time"
	"fmt"
	"strings"
)

func main() {
	/*res := time.Now().Unix()
	fmt.Println(res)
	str := "2019-09-03 16:40:34.0143498 +0800 CST"
	re, _ := time.Parse("Mon Jan 2 15:04:05 -0700 MST 2006", str)
	fmt.Println(re)*/

	// 2006-01-02
	for {

		date := time.Now().Format("15:04:05")
		// func Split(s, sep string) []string
		res := strings.Split(date, ":")
		fmt.Println(res)
		time.Sleep(time.Second)
	}
}