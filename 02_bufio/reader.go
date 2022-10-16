package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fLines, err := os.Open("./lines.txt")
	checkError(err)
	defer fLines.Close()

	reader := bufio.NewReader(fLines)

	var s string
	err = nil
	for err == nil {
		s, err = reader.ReadString('\n')
		fmt.Printf("%s\n", dropLF(s))
	}
	fmt.Println()

	fPipes, err := os.Open("./pipes.txt")
	checkError(err)
	defer fPipes.Close()

	reader = bufio.NewReader(fPipes)
	for err == nil {
		s, err = reader.ReadString('|')
		fmt.Printf("%s\n", dropPipe(s))
	}
}

func dropLF(s string) string {
	if len(s) > 0 && s[len(s)-1] == '\n' {
		return string(s[0 : len(s)-1])
	}
	return s
}

func dropPipe(s string) string {
	if len(s) > 0 && s[len(s)-1] == '|' {
		return string(s[0 : len(s)-1])
	}
	return s
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
