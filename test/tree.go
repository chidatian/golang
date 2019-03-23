package main

import(
	"fmt"
)
// 二叉查找树
type Tree struct {
	Root *Node
}

type Node struct {
	Value int
	Left *Node
	Right *Node
	Parent *Node
}
// 在树中插入节点
func (this *Tree) Insert(v int) {
	n := new(Node)
	n.Value = v
	if this.Root == nil {
		this.Root = n
		return
	}
	cur := this.Root
	for {
		n.Parent = cur
		if v < cur.Value {
			if cur.Left == nil {
				cur.Left = n
				break
			} else {
				cur = cur.Left
			}
		} else {
			if cur.Right == nil {
				cur.Right = n
				break
			} else {
				cur = cur.Right
			}
		}
	}
}
// 遍历所有节点
func (this *Tree) Select() {
	var mapSearch func(n *Node)
	mapSearch = func(n *Node){
		if n != nil {
			mapSearch(n.Left)
			fmt.Println(n.Value)
			mapSearch(n.Right)
		}
	}
	mapSearch(this.Root)
}
// 查找节点
func (this *Tree) Search(v int) *Node{
	cur := this.Root
	for {
		if v < cur.Value {
			cur = cur.Left
			if cur == nil {
				return nil
			}
		} else if v > cur.Value {
			cur = cur.Right
			if cur == nil {
				return nil
			}
		} else {
			return cur
		}
	}
}
// 最小值
func (this *Tree) Min() *Node{
	cur := this.Root
	for {
		if cur.Left != nil {
			cur = cur.Left
		} else {
			return cur
		}
	}
}
// 最大值
func (this *Tree) Max() *Node{
	cur := this.Root
	for {
		if cur.Right != nil {
			cur = cur.Right
		} else {
			return cur
		}
	}
}
// 删除一个节点
func (this *Tree) Delete(v int) {
	fmt.Printf("del: %d\n",v)
	var (
		mapDel func (n *Node)
		cur = this.Search(v)
	)
	mapDel = func (cur *Node) {
		if cur == nil {
			return
		}
		if cur.Left == nil && cur.Right == nil {
			if cur.Parent.Left == cur {
				cur.Parent.Left = nil
			} else {
				cur.Parent.Right = nil
			}
		} else if cur.Left == nil && cur.Right != nil {
			cur.Right.Parent = cur.Parent
			if cur.Parent.Left == cur {
				cur.Parent.Left = cur.Right
			} else {
				cur.Parent.Right = cur.Right
			}
		} else if cur.Left != nil && cur.Right == nil {
			cur.Left.Parent = cur.Parent
			if cur.Parent.Left == cur {
				cur.Parent.Left = cur.Left
			} else {
				cur.Parent.Right = cur.Left
			}
		} else {
			fmt.Println("both")
			curs := cur.Right
			for {
				if curs.Left != nil {
					curs = curs.Left
				} else {
					break
				}
			}
			cur.Value = curs.Value
			mapDel(curs)
		}
	}
	mapDel(cur)
}
// 修改一个节点
func (this *Tree) Edit(v int, v2 int) {
	this.Insert(v2)
	this.Delete(v)
}
func main() {
	var t Tree
	t.Insert(100)
	t.Insert(120)
	t.Insert(80)
	t.Insert(70)
	t.Insert(90)
	t.Insert(85)
	t.Insert(95)
	t.Insert(110)
	t.Select()
	t.Delete(80)
	t.Select()
	fmt.Println("---------------")
	val := t.Search(110)
	fmt.Println(val)
	fmt.Println(t.Min())
	fmt.Println(t.Root.Value)
	fmt.Println(t.Root.Left.Value)
	fmt.Println(t.Root.Right.Value)
	t.Edit(80,125)
	t.Select()
}