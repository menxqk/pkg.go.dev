package main

import (
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("read_file.go")
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(data)
}
