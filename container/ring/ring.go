package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// Create a new ring of size 5.
	r := ring.New(5)
	// Get the length of the ring.
	nr := r.Len()
	// Initialize the ring with some integer values.
	for i := 0; i < nr; i++ {
		r.Value = i
		r = r.Next()
	}
	// Iterate through the ring and print its contents.
	r.Do(func(p any) {
		fmt.Printf("%d ", p.(int))
	})
	fmt.Println()

	// Create another ring, size 10.
	s := ring.New(10)
	// Get the length of the ring.
	ns := s.Len()
	// Initialize the ring with some values.
	for i := 0; i < ns; i++ {
		s.Value = nr + i
		s = s.Next()
	}
	// Iterate through the ring and print its contents.
	s.Do(func(p any) {
		fmt.Printf("%d ", p.(int))
	})
	fmt.Println()

	// Link rings r and s
	rs := r.Link(s)
	// Iterate through the ring and print its contents.
	rs.Do(func(p any) {
		fmt.Printf("%d ", p.(int))
	})
	fmt.Println()
}
