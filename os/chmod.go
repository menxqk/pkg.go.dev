package main

import (
	"log"
	"os"
)

func main() {
	if err := os.Chmod("some-file.txt", 0644); err != nil {
		log.Fatal(err)
	}
}
