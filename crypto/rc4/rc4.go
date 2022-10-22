package main

import (
	"crypto/rc4"
	"fmt"
)

func main() {
	key := []byte("secret key")
	cipher, err := rc4.NewCipher(key)
	if err != nil {
		panic(err)
	}

	msg := "A string message"
	enc := make([]byte, len(msg))
	cipher.XORKeyStream(enc, []byte(msg))
	fmt.Printf("enc: %x\n", enc)

	dec := make([]byte, len(msg))
	cipher.XORKeyStream(dec, enc)
	fmt.Printf("dec: %x\n", dec)
}
