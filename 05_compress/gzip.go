package main

import (
	"compress/gzip"
	"io"
	"os"
	"time"
)

func main() {
	inFile, err := os.Open("./lines.txt")
	checkError(err)
	defer inFile.Close()

	outFile, err := os.Create("./lines.gzip")
	checkError(err)
	defer outFile.Close()

	gzw := gzip.NewWriter(outFile)
	gzw.Name = "lines"
	gzw.Comment = "Some lines"
	gzw.ModTime = time.Now()
	_, err = io.Copy(gzw, inFile)
	gzw.Flush()
	gzw.Close()

	gzFile, err := os.Open("./lines.gzip")
	checkError(err)
	defer gzFile.Close()

	gzr, err := gzip.NewReader(gzFile)
	checkError(err)
	defer gzr.Close()

	io.Copy(os.Stdout, gzr)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
