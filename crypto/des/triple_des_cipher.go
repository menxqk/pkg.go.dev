package main

import (
	"crypto/des"
)

func main() {
	// NewTripleDESChiper can also be used when EDE2 is required by
	// duplicating the first 8 bytes of the 16-byte key.
	ede2key := []byte("example key 1234")

	var tripleDESKey []byte
	tripleDESKey = append(tripleDESKey, ede2key[:16]...)
	tripleDESKey = append(tripleDESKey, ede2key[:8]...)

	_, err := des.NewTripleDESCipher(tripleDESKey)
	if err != nil {
		panic(err)
	}
}
