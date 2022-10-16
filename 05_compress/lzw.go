package main

import (
	"compress/lzw"
	"io"
	"os"
)

func main() {
	inFile, err := os.Open("./lines.txt")
	checkError(err)
	defer inFile.Close()

	outFile, err := os.Create("./lines.lzw")
	checkError(err)
	defer outFile.Close()

	lzww := lzw.NewWriter(outFile, lzw.LSB, 8)
	io.Copy(lzww, inFile)
	checkError(lzww.Close())

	// Read from compressed file
	lzwFile, err := os.Open("./lines.lzw")
	checkError(err)
	defer lzwFile.Close()

	lzwr := lzw.NewReader(lzwFile, lzw.LSB, 8)
	io.Copy(os.Stdout, lzwr)
	checkError(lzwr.Close())
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
