package main

import (
	"crypto/aes"
	"fmt"
)

var key = []byte("abc123mno456abc123mno456mno456ab")

func main() {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	msg := []byte("A secret message")  // 16 characters
	cod := make([]byte, aes.BlockSize) // aes.BlockSize is 16
	cipher.Encrypt(cod, msg)
	fmt.Printf("cod: %v\n", cod)

	dec := make([]byte, aes.BlockSize)
	cipher.Decrypt(dec, cod)
	fmt.Printf("dec: %s\n", dec)
}
