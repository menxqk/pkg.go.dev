package main

import (
	"fmt"
	"math"
)

func main() {
	c := math.Ceil(1.49)
	fmt.Printf("%.1f\n", c)
	fmt.Printf("%.2f\n", c)
}
