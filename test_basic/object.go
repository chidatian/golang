package main

import (
	"fmt"
)

type Controller struct {

}

type Goods struct {
	Controller
}

func (this Controller) PP() {
	fmt.Println("pp")
}

func (this Goods) DD() {
	fmt.Println("dd")
}

func (this Goods) PP() {
	fmt.Println("goods pp")
}

func main() {
	var g Goods
	g.DD()
	g.PP()
}