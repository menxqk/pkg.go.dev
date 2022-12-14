package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func main() {
	r, theta := cmplx.Polar(2i)
	fmt.Printf("r: %.1f, θ: %.1f*π\n", r, theta/math.Pi)
}
