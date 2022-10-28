package main

import (
	"encoding/asn1"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	pemBytes, err := ioutil.ReadFile("./root_cert.pem")
	if err != nil {
		log.Fatal(err)
	}

	derBytes, _ := pem.Decode(pemBytes)

	var raw asn1.RawValue
	_, err = asn1.Unmarshal(derBytes.Bytes, &raw)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Class: %v\n", raw.Class)
	fmt.Printf("Tag: %v\n", raw.Tag)
	fmt.Printf("IsCompund: %v\n", raw.IsCompound)
}
