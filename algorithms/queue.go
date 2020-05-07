package main

import "fmt"

type queue struct {
	head int
	tail int
	box []int
	length int
	size int
}

func (this *queue) Enqueue(i int) {
	if this.size == this.length{
		panic("overflow")
	}
	this.box[this.tail] = i
	this.tail++
	this.size++
	if this.tail == this.length {
		this.tail = 0
	}
}

func (this *queue) Dequeue() int {
	if this.size == 0 {
		panic("underflow")
	}
	k := this.box[this.head]
	this.head++
	this.size--
	if this.head == this.length {
		this.head = 0
	}
	return k
}

func New(length int) *queue {
	return &queue{
		head : 0,
		tail : 0,
		box : make([]int, length),
		length : length,
		size : 0,
	}
}

func main() {
	q := New(6)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	q.Enqueue(6)
	fmt.Println(q)
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	q.Enqueue(7)
	q.Enqueue(8)
	q.Enqueue(9)
	fmt.Println(q)
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q)
	fmt.Println(q.Dequeue())
	fmt.Println(q)
	q.Enqueue(91)
	q.Enqueue(92)
	q.Enqueue(93)
	q.Enqueue(94)
	q.Enqueue(95)
	q.Enqueue(96)
	fmt.Println(q)
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q)
	q.Enqueue(961)
	q.Enqueue(962)
	fmt.Println(q)

}