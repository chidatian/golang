<<<<<<< HEAD
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
=======
package main

import (
	"github.com/fatih/color"
)

func main() {
	color.Red("Prints %s in blue.", "text")
	color.Yellow("Prints %s in blue.", "text")
	color.Green("Prints %s in blue.", "text")
	color.Blue("Prints %s in blue.", "text")
>>>>>>> 65888be0467c07f28b57cd4073746ad38fcc0782
}