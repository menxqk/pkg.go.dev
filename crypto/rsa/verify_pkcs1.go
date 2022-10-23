package main

import (
	"crypto"
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

	message := []byte("message to be signed")

	// Only small messages can be signed directly; thus the hash of a
	// mesage, rather than the message itself, is signed. This requires
	// that the hash function be collision resistant. SHA-256 is the
	// least-strong hash function that should for this.
	hashed := sha256.Sum256(message)

	signature, _ := hex.DecodeString("411a0d9db6951cee2eac0e853568eda1f30704eb7d73632807007907bd0c7ce66bd1eae023040f6ec4488d330afbdd470e90d414249f3abd8fdf9713a4196d08b83ee43a669ce223771e142a0728a7babb8756367d319534a2bb0969e709c5a21d9c34c740a842cfe9b3d3fc613555213b0ed6d7c271e00ec0bd5ea5aa433bc7994d06ccdb5543797e4a8440b69b2a555bc812788fe08944a453191fc9ee734ad8f84eab8f44f472d5067e9b93364f4f88cbc611f7560c67ff238caaf560fb8d867cc36dc3b31ffb47244376184c0395b36708459da233ee0bf82c7a7bc11c7b2fe0a0e51a6f7bdedb0c45852f95379c05c77a78c90928040ff57d817b120204")

	err = rsa.VerifyPKCS1v15(&priv.PublicKey, crypto.SHA256, hashed[:], signature)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from verification: %s\n", err)
		return
	} else {
		fmt.Println("signature veryfied")
	}
}
