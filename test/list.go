package main

import (
	"fmt"
	"pack"
)
// 两个栈 实现队列
type Lister interface {
	Set(val string)
	Get() string
	move()
}

type List struct {
	In *pack.Stack
	Out *pack.Stack
}

func New() *List{
	obj := new(List)
	obj.In, obj.Out = pack.New(), pack.New()
	return obj
}

func (this *List) Set(val string) {
	this.In.Push(val)
}

func (this *List) Get() string {
	if this.Out.Count == 0 {
		this.move()
	}
	res := this.Out.Pop()
	return res
}

func (this *List) move() {
	tmp := *this.In
	this.In.Clear()
	for tmp.Count > 0 {
		this.Out.Push(tmp.Pop())
	}
}

func main() {
	l := New()
	l.Set("demo")
	l.Set("test")
	fmt.Println(l.Get())
	fmt.Println(l.Get())
}