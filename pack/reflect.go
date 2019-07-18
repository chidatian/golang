package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id int
	Name string
}

func (this User) Test() {
	fmt.Println("--")
}

func main() {
	i := 1
	j := "tian"
	u := new(User)
	e := reflect.ValueOf(u).Elem()
	e.FieldByName("Id").Set(reflect.ValueOf(i))
	e.FieldByName("Name").Set(reflect.ValueOf(j))
	res := e.MethodByName("Test").IsValid()
	fmt.Println(res)
}