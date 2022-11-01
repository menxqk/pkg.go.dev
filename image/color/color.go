package main

import (
	"fmt"
	"image/color"
)

func main() {
	alpha := color.Alpha{1}
	r, g, b, a := alpha.RGBA()
	fmt.Printf("r: %d, g: %d, b: %d, a: %d\n", r, g, b, a)

}
