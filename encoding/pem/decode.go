package main

import (
	"crypto/dsa"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

func main() {
	files := []string{"./priv_key.pem", "./pub_key.pem", "cert.pem"}

	for _, f := range files {
		decodeFile(f)
	}
}

func decodeFile(name string) {
	pemBytes, err := ioutil.ReadFile(name)
	checkError(err)

	block, _ := pem.Decode(pemBytes)

	fmt.Printf("Type: %s\n", block.Type)

	switch block.Type {
	case "RSA PRIVATE KEY":
		key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		checkError(err)
		fmt.Printf("D: %d\nE: %d\nN: %d\n", key.D, key.E, key.N)
	case "PUBLIC KEY":
		key, err := x509.ParsePKIXPublicKey(block.Bytes)
		checkError(err)
		switch key := key.(type) {
		case *rsa.PublicKey:
			fmt.Printf("rsa.PublicKey\nE: %d\nN: %d\n", key.E, key.N)
		case *dsa.PublicKey:
			fmt.Printf("dsa.PublicKey\nG: %d\nP: %d\nQ: %d\nY: %d\n", key.G, key.P, key.Q, key.Y)
		case *ecdsa.PublicKey:
			fmt.Printf("ecdsa.PublicKey\nX: %d\nY: %d\n", key.X, key.Y)
		case ed25519.PublicKey:
			fmt.Println("ed25519.PublicKey\n", key)
		}
	case "CERTIFICATE":
		cert, err := x509.ParseCertificate(block.Bytes)
		checkError(err)
		fmt.Printf("Issuer: %s, Version: %d\n", cert.Issuer, cert.Version)

	default:
		fmt.Println("could not parse bytes")
	}

	fmt.Println()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
