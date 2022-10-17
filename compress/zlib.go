package main

import (
	"compress/zlib"
	"io"
	"os"
)

func main() {
	inFile, err := os.Open("./lines.txt")
	checkError(err)
	defer inFile.Close()

	outFile, err := os.Create("./lines.zip")
	checkError(err)

	zlw := zlib.NewWriter(outFile)
	io.Copy(zlw, inFile)
	zlw.Flush()
	zlw.Close()

	// Read contents from compressed file
	zFile, err := os.Open("./lines.zip")
	checkError(err)
	defer zFile.Close()

	zlr, err := zlib.NewReader(zFile)
	checkError(err)
	defer zlr.Close()

	io.Copy(os.Stdout, zlr)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
