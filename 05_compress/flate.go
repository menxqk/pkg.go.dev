package main

import (
	"compress/flate"
	"io"
	"os"
)

func main() {
	// Write compressed file
	inputFile, err := os.Open("./lines.txt")
	checkError(err)
	defer inputFile.Close()

	outputFile, err := os.Create("./lines.flate")
	checkError(err)
	defer outputFile.Close()

	flateWriter, err := flate.NewWriter(outputFile, flate.BestCompression)
	checkError(err)

	io.Copy(flateWriter, inputFile)
	flateWriter.Flush()
	flateWriter.Close()

	// Read compressed file
	compressedFile, err := os.Open("./lines.flate")
	checkError(err)
	defer compressedFile.Close()

	flateReader := flate.NewReader(compressedFile)
	defer flateReader.Close()
	io.Copy(os.Stdout, flateReader)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
