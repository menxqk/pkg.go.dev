package main

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	secretHash := md5.New()
	secretHash.Write([]byte("secret"))
	key := secretHash.Sum(nil)
	fmt.Println("The secret key is ", hex.EncodeToString(key))

	msg := "A string message"

	sig := hmac.New(sha256.New, key)
	sig.Write([]byte(msg))
	enc := hex.EncodeToString(sig.Sum(nil))
	fmt.Printf("enc: %s\n", enc)
}
