package main

import(
	"fmt"
)

func main() {
	var i int = 10
	var j *int
	var k int
	fmt.Println(i)
	j = &i
	fmt.Println(j)
	k = *j
	fmt.Println(k)
}