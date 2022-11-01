package main

import (
	"fmt"
	"hash/crc64"
)

func main() {
	s := []byte("A string")
	t := crc64.Table{} // 256 positions
	result := crc64.Checksum(s, &t)
	fmt.Printf("result: %d\n", result)
}
