package main

import (
	"fmt"
)
// 堆栈 test
type Stacker interface {
	Pop() string		// 出栈
	Push(val string)	// 入栈
}

type Stack struct {
	top *child
}

type child struct {
	value string
	next *child
}

func (this *Stack) Pop() string {
	res := this.top.value
	this.top = this.top.next
	return res
}

func (this *Stack) Push(val string) {
	c := child{}
	c.value = val
	if this.top == nil {
		this.top = &c
	} else {
		c.next = this.top
		this.top = &c
	}
}

func main() {
	var s Stacker
	s = new(Stack)
	s.Push("test")
	s.Push("demo")
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}