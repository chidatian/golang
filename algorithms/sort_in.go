package main

import "fmt"

func main() {
	// sortA(3)
	sortB()
}

func sortB() {
	var a []int = []int{3,2,5,1,8}
	var j int
	for i := 1; i < len(a); i ++ {
		b := a[i]
		for j = i-1; j >= 0 && a[j] > b; j-- {
			a[j+1] = a[j]
		}
		a[j+1] = b
	}
	fmt.Println(a)
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