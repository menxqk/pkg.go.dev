package main

import (
	"log"
	"os"
)

func main() {
	err := os.WriteFile("some-file.txt", []byte("Hello, Gophers!"), 0666)
	if err != nil {
		log.Fatal(err)
	}
}
