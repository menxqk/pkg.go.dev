package main

import (
	"compress/bzip2"
	"io"
	"os"
)

func main() {
	f, err := os.Open("./lines.txt.bz2")
	checkEror(err)

	reader := bzip2.NewReader(f)
	io.Copy(os.Stdout, reader)
}

func checkEror(err error) {
	if err != nil {
		panic(err)
	}
}
