package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
)

//go:embed embed_example.go
var b []byte

//go:embed *
var fsys embed.FS

func main() {
	fmt.Printf("bytes: %v\n", b)

	entries, err := fs.ReadDir(fsys, ".")
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("name: %s %d\n", entry.Name(), info.Size())
	}
}
