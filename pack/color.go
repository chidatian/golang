package main

import (
	"github.com/fatih/color"
)

func main() {
	color.Red("Prints %s in blue.", "text")
	color.Yellow("Prints %s in blue.", "text")
	color.Green("Prints %s in blue.", "text")
	color.Blue("Prints %s in blue.", "text")
}