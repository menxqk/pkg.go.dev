package main

import (
	"fmt"
	"strconv"
)

func main() {
	b32 := []byte("float32:")
	b32 = strconv.AppendFloat(b32, 3.14155926535, 'E', -1, 32)
	fmt.Println(string(b32))

	b64 := []byte("float64:")
	b64 = strconv.AppendFloat(b64, 3.1415926535, 'E', -1, 64)
	fmt.Println(string(b64))

	b64 = []byte("float64:")
	b64 = strconv.AppendFloat(b64, 3.1415926535, 'f', 2, 64)
	fmt.Println(string(b64))
}
