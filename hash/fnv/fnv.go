package main

import (
	"fmt"
	"hash/fnv"
)

func main() {
	s := []byte("A string")
	hash := fnv.New32()
	hash.Write(s)
	fmt.Printf("%d\n", hash.Sum32())
}
