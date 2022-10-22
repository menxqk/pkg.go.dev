package main

import (
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
)

func main() {
	curve := elliptic.P256()
	key, x, y, err := elliptic.GenerateKey(curve, rand.Reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("key: %x\nx: %d\ny: %d\n", key, x, y)
}
