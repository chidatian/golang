package main

import (
	"fmt"
)

func main() {
	var arr []int = []int{6, 2, 7, 3, 8, 9}
	fmt.Println(arr)
	qsort_children(arr)
	fmt.Println(arr)
}

/* func qsort(arr []int){
	res, mid := qsort_children(arr)
	c_left, c_right := len(res[:mid]), len(res[mid+1:])
	for true {
		if c_left > 1 {

		}
		if c_right > 1 {

		}
		res, mid := qsort_children(res[:mid])
		c_left, c_right := len(res[:mid]), len(res[mid+1:])
	}
} */
// 快排
func qsort_children(arr []int){
	if len(arr) <= 1 {
		return
	}
	nums := len(arr) - 1
	mid, i := arr[0], 1
	left, right := 0, nums
	for left < right {
		if arr[i] > mid {
			arr[i], arr[right] = arr[right], arr[i]
			right--
		} else {
			arr[i], arr[left] = arr[left], arr[i]
			i++
			left++
		}
	}
	qsort_children(arr[:left])
	qsort_children(arr[left+1:])
	// return arr, left
}