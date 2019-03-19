package main

import(
	"tii"
)

func main() {
	app := tii.New()
	app.Router("Get", &controllers.Goods{})
	app.Run(":9090")
}