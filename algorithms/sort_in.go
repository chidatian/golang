package main

import "fmt"

func main() {
	a1 := []int{1, 2, 3, 4, 5}
	a2 := a1

	a1 = append(a1[:1], a1[2:]...)

	fmt.Println(a2)
	fmt.Println(a1)
	sortA(3)
}

func sortA(b int) {
	var a []int = []int{1,5,7,9}
	var n int = len(a) - 1
	i := n
	for ; i >= 0 && a[i] > b; i-- {
		if i == n {
			a = append(a, a[i])
		} else {
			a[i+1] = a[i]
		}
	}
	if i == n {
		a = append(a, b)
	} else {
		a[i+1] = b
	}
	fmt.Println(a)
}