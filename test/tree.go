package main

import(
	"fmt"
)

type Tree struct {
	Root Node
}

type Node struct {
	Value int
	Left *Node
	Right *Node
}

func (this *Tree) Insert(v int) {
	this.Root.Insert(v)
}

func (this *Node) Insert(v int) {
	if v > this.Value {
		if this.Right == nil {
			this.Right = &Node{v,nil,nil}
		} else {
			this.Right.Insert(v)
		}
	} else {
		if this.Left == nil {
			this.Left = &Node{v,nil,nil}
		} else {
			this.Left.Insert(v)
		}
	}
}

func (this *Tree) List() []int{
	var res []int
	this.Root.List(&res)
	return res
}

func (this *Node) List(result *[]int) {
	if this.Left != nil {
		this.Left.List(result)
	}
	*result = append(*result, this.Value)
	if this.Right != nil {
		this.Right.List(result)
	}
}

func main() {
	var t Tree
	t.Insert(10)
	t.Insert(2)
	t.Insert(22)
	data := t.List()
	fmt.Println(data)
}