package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

func main() {
	fi, err := os.Lstat("some-file.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("permissions: %#o\n", fi.Mode().Perm()) // 0400, 0700, etc.
	switch mode := fi.Mode(); {
	case mode.IsRegular():
		fmt.Println("regular file")
	case mode.IsDir():
		fmt.Println("directory")
	case mode&fs.ModeSymlink != 0:
		fmt.Println("symbolic link")
	case mode&fs.ModeNamedPipe != 0:
		fmt.Println("named pipe")
	}
}
