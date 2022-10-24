package main

import (
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
)

func main() {
	rootPEM, err := ioutil.ReadFile("./rootCA_cert.pem")
	if err != nil {
		panic(err)
	}

	serverPEM, err := ioutil.ReadFile("./server_cert.pem")
	if err != nil {
		panic(err)
	}

	// First, create the set of root certificates. For this example we only
	// have one. It's also possible to omit this in order to use the
	// default root set of the current operating system.
	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM(rootPEM)
	if !ok {
		panic("failed to parse root certificate")
	}

	block, _ := pem.Decode(serverPEM)
	if block == nil {
		panic("failed to parse certificate PEM")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		panic("failed to parse certificate: " + err.Error())
	}

	opts := x509.VerifyOptions{
		DNSName: "test.domain.com",
		Roots:   roots,
	}

	if _, err := cert.Verify(opts); err != nil {
		panic("failed to verify certificate: " + err.Error())
	}

}
