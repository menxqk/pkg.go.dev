package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []float64{1.0, 2.0, 3.3, 4.6, 6.1, 7.2, 8.0}

	x := 2.0
	i := sort.SearchFloat64s(a, x)
	fmt.Printf("foung %g at index %d in %v\n", x, i, a)

	x = 1.5
	i = sort.SearchFloat64s(a, x)
	fmt.Printf("%g not found, can be inserted at index %d in %v\n", x, i, a)

}
