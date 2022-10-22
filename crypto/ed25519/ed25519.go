package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"fmt"
)

func main() {
	pb, pv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}

	msg := []byte("A string message")
	fmt.Printf("message: %s\n", msg)

	signed := ed25519.Sign(pv, msg)
	fmt.Printf("signed: %x\n", signed)

	verified := ed25519.Verify(pb, msg, signed)
	fmt.Printf("verified: %t\n", verified)
}
