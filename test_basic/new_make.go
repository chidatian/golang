package main

import(
	"fmt"
)

func main() {
	// new：用来初始化一个对象，并且返回该对象的首地址．其自身是一个指针．可用于初始化任何类型
	// make:返回一个初始化的实例，返回的是一个实例，而不是指针，其只能用来初始化：slice,map和channel三种类型
	a := new(int)
	b := make([]int,2)
	fmt.Println(a)
	fmt.Println(b)
}