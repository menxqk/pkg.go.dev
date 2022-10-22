package main

import (
	"crypto/dsa"
	"crypto/rand"
	"fmt"
)

func main() {
	params := dsa.Parameters{}
	err := dsa.GenerateParameters(&params, rand.Reader, dsa.L1024N160)
	if err != nil {
		panic(err)
	}

	publicKey := dsa.PublicKey{Parameters: params}
	privateKey := dsa.PrivateKey{PublicKey: publicKey}
	err = dsa.GenerateKey(&privateKey, rand.Reader)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", privateKey)
}
