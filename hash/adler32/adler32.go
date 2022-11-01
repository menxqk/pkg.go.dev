package main

import (
	"fmt"
	"hash/adler32"
)

func main() {
	s := []byte("A string")
	result := adler32.Checksum(s)
	fmt.Printf("%d\n", result)

	ad := adler32.New()
	ad.Write(s)
	fmt.Printf("%d\n", ad.Sum32())
}
