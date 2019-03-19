package controllers

import(
	"tii"
)

type Goods struct {
	tii.Controller
}

func (this *Goods) Get() {
	r := this.Request.URL
	
}