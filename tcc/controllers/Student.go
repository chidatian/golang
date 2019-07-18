package controllers

import (
	"fmt"
	"tools/tcc"
)

type Student struct {
	tcc.Controller
}

func (this Student) Index() {
	fmt.Println(this.Request)
}