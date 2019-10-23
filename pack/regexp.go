package main

import (
	"fmt"
	"regexp"
)

func test() {
	var s string = "chidatian"
	var pattern string = "da"
	isB, err := regexp.MatchString(pattern, s)
	if err != nil {
		panic(err)
	}
	fmt.Println(isB)
}

func test1() {
	// func Compile(expr string) (*Regexp, error)
	// func (re *Regexp) Split(s string, n int) []string
	obj, err := regexp.Compile(`f:/git/test/a/`)
	if err != nil {
		panic(err)
	}
	res := obj.Split("f:/git/test/a/b/qwer.txt", -1)
	fmt.Println(res)
	for _, item := range res {
		fmt.Println(item)
	}
}

func test2() {
	var dirs []string
	dirs = append(dirs, "a", "b")
	fmt.Println(dirs)
}

func main() {
	test2()
}