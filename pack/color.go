package main

import (
	"fmt"
	"image/color"
)

func main() {
	// func (c RGBA) RGBA() (r, g, b, a uint32)
	var c color.RGBA
	r, g, b, a := c.RGBA()
	fmt.Println(r, g, b, a)
}