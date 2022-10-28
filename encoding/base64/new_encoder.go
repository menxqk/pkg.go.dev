package main

import (
	"encoding/base64"
	"os"
)

func main() {
	input := []byte("foo\nbar")
	encoder := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	encoder.Write(input)
	encoder.Close()
	os.Stdout.Write([]byte("\n"))
}
