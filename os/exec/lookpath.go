package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	path, err := exec.LookPath("./lookpath")
	if err != nil {
		log.Fatal("lookpath executable not found")
	}
	fmt.Printf("lookpath is available at %s\n", path)
}
