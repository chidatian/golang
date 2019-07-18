package main

import (
	"tools/tcc"
	"tcc/controllers"
)

func main() {
	a := tcc.New()
	a.Routes["Goods"] = new(controllers.Goods)
	a.Routes["Student"] = new(controllers.Student)
	a.Run("127.0.0.1:9090")
}

