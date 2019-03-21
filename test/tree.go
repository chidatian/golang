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
	fmt.Printf("del %d\n",v)
	cur := this.Root
	TAG:
	for {
		if v < cur.Value {
			if cur.Left == nil {
				return
			}
			cur = cur.Left
		} else if v > cur.Value {
			if cur.Right == nil {
				return
			}
			cur = cur.Right
		} else {
			if cur.Left == nil && cur.Right == nil {
				if cur.Parent.Left == cur {
					cur.Parent.Left = nil
				} else {
					cur.Parent.Right = nil
				}
				break
			} else if cur.Left == nil && cur.Right != nil {
				*cur = *(cur.Right)
				break
			} else if cur.Left != nil && cur.Right == nil {
				*cur = *(cur.Left)
				break
			} else {
				curMin := cur
				for {
					if curMin.Left != nil {
						curMin = curMin.Left
					} else {
						fmt.Println(cur)
						fmt.Println(cur.Left)
						fmt.Println(cur.Right)
						fmt.Println(curMin)
						fmt.Println(curMin.Left)
						fmt.Println(curMin.Right)
						*cur = *curMin
						fmt.Println(cur.Parent)
						fmt.Println(cur.Parent.Left)
						fmt.Println(cur.Parent.Right)
						break TAG
					}
				}
			}
		}
	}
}
// 修改一个节点
func (this *Tree) Edit(v int) {

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
	val := t.Search(95)
	fmt.Println(val)
	t.Delete(90)
	t.Select()
	fmt.Println(t.Min())
	fmt.Println(t.Root.Value)
	fmt.Println(t.Root.Left.Value)
	fmt.Println(t.Root.Right.Value)
}