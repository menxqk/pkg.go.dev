package main

import (
	"fmt"
	"hash/maphash"
)

func main() {
	// The zero Hash value is valid and ready to use; setting an
	// initial seed is not necessary
	var h maphash.Hash

	// Add a string to the hash, and print the current hash value.
	h.WriteString("hello, ")
	fmt.Printf("%#x\n", h.Sum64())

	// Append additional data (in the form of a byte array).
	h.Write([]byte{'w', 'o', 'r', 'l', 'd'})
	fmt.Printf("%#x\n", h.Sum64())

	// Reset discards all data previously added to the Hash, without
	// changing its seed.
	h.Reset()

	// Use SetSeed to create a new Hash h2 which will behave
	// identically to h.
	var h2 maphash.Hash
	h2.SetSeed(h.Seed())

	h.WriteString("name")
	h2.WriteString("name")
	fmt.Printf("%#x == %#x\n", h.Sum64(), h2.Sum64())
}
