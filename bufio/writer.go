package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fStrings, err := os.OpenFile("./strings.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	checkError(err)
	defer fStrings.Close()

	writer := bufio.NewWriter(fStrings)
	i := 5
	n := 1
	for i > 0 {
		writer.WriteString(fmt.Sprintf("String %d\n", n))
		n++
		i--
	}

	fmt.Printf("Number of bytes written to the current buffer: %d\n", writer.Buffered())
	fmt.Printf("Size of the underlying buffer in bytes: %d\n", writer.Size())

	err = writer.Flush()
	checkError(err)

	fmt.Printf("After Flush, number of bytes written to the current buffer: %d\n", writer.Buffered())
	fmt.Printf("After Flush, size of the underlying buffer in bytes: %d\n", writer.Size())
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
