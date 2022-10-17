package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	fLines, err := os.Open("./lines.txt")
	checkError(err)
	defer fLines.Close()
	// ScanLines (default)
	scanner := bufio.NewScanner(fLines)
	fmt.Println("Iterate using lines ('\\n' is the default separator):")
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("line read: %s\n", line)
	}
	fmt.Println()

	fPipes, err := os.Open("./pipes.txt")
	checkError(err)
	defer fPipes.Close()
	// ScanPipes
	scanner = bufio.NewScanner(fPipes)
	scanner.Split(ScanPipes)
	fmt.Println("Iterate using pipes ('|' as the default separator):")
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), "\r\n")
		fmt.Printf("line read: %s\n", line)
	}
	fmt.Println()
}

func ScanPipes(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// If it is EOF or there is not data return 0
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	// Search for the position of the '|' character
	if i := bytes.IndexByte(data, '|'); i >= 0 {
		// dropCR
		return i + 1, dropPipe(data[0:i]), nil
	}
	// If we are at EOF, we have a final token. Return it.
	if atEOF {
		return len(data), dropPipe(data), nil
	}
	// Request more data;
	return 0, nil, nil
}

func dropPipe(data []byte) []byte {
	// If there is data and the last character is '|', remove last character
	if len(data) > 0 && data[len(data)-1] == '|' {
		return data[0 : len(data)-1]
	}
	return data
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
