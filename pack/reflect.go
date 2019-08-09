package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id int `field:"id"`
	Name string `field:"name"`
}

func (this User) Test() {
	fmt.Println("--")
}
// 通过反射给struct 赋值
func controller() {
	i := 1
	j := "tian"
	u := new(User)
	e := reflect.ValueOf(u).Elem()
	e.FieldByName("Id").Set(reflect.ValueOf(i))
	e.FieldByName("Name").Set(reflect.ValueOf(j))
	res := e.MethodByName("Test").IsValid()
	fmt.Println(res)
}

// struct 标签
func tag() {
	var u User
	t := reflect.TypeOf(u)
	f,_ := t.FieldByName("Id")
	fmt.Println(f)
	fmt.Println(f.Tag.Get("field"))
}

func main() {
	tag()
}