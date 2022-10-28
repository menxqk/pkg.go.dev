package main

import (
	"debug/buildinfo"
	"fmt"
	"log"
	"os"
)

func main() {
	name := os.Args[0]

	info, err := buildinfo.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ReadFile info: %v\n", info)

	f, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	info, err = buildinfo.Read(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read info: %v\n", info)
}
