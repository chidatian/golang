package controllers

import (
	"fmt"
	"tools/tcc"
)

type Goods struct {
	tcc.Controller
}

func (this Goods) Get() {
	fmt.Println(this.Request)
}