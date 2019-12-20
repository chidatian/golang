package main


import (
	"fmt"
	"strconv"
)

func main() {
	a := "a123a"
	res,_ := strconv.Atoi(a)
	fmt.Println(res)
}