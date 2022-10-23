package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	bytes, err := ioutil.ReadFile("./priv_key.pem")
	if err != nil {
		panic(err)
	}

	block, _ := pem.Decode(bytes)
	if block == nil {
		panic("block is nil")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	ciphertext, _ := hex.DecodeString("950508d5c42fa77119799c062f6f5db4d67b75e9dd40604351e872c31566cda33abc555250b4e7763ec2d81ba3a631769b176851b2c1fc7f9ca3b5f240928e31928c9343be1f5f199f862334fe610111060b8d1bc641851927d188a19527d4467510cdc3c912e6a30e40ab1f1e1ba3d550c6070467d2dae9e74fc77503e7df1075f67e76f1c1dc5ee0dc8f7c18625cb19ed7649dd10596105df8b510e1dcd8c5c0f10edbe82678923793e9fc13800568ea674fe1081446482f2188f6029dd3c49cdc1b03765f7b8a0540321b7b579a01f194389595934d0c55ac244c79059092047b627d1b519ca49b6b7822456ae1e44adcf501f6487e828a0e0025b10206f3")
	label := []byte("orders")

	// crypto/rand.Reader is a good source of entrupy for blinding the RSA
	// operation.
	rng := rand.Reader

	plaintext, err := rsa.DecryptOAEP(sha256.New(), rng, priv, ciphertext, label)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from decrption: %s\n", err)
		return
	}

	fmt.Printf("Plaintext: %s\n", string(plaintext))

}
