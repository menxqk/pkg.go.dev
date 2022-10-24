package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	data := []byte("This page is intentionally left blank.")
	fmt.Printf("%x\n", sha1.Sum(data))
}
