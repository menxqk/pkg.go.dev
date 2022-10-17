package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// Load contents of files to byte slice
	b, err := ioutil.ReadFile("/etc/passwd")
	checkError(err)

	// Create bytes reader
	reader := bytes.NewReader(b)
	fmt.Printf("Unread bytes: %d\n", reader.Len())

	// Use buffered reader to read contents of file
	bufReader := bufio.NewReader(reader)
	var users = []string{}
	var s string
	err = nil
	for err == nil {
		s, err = bufReader.ReadString('\n')
		if s == "" {
			continue
		}
		fields := strings.Split(s, ":")
		if len(fields) > 0 {
			users = append(users, fields[0])
		}
	}

	// Create bytes Buffer and copy users to it
	buf := bytes.NewBuffer(nil)
	for _, u := range users {
		buf.WriteString(fmt.Sprintf("user: %q\n", u))
	}
	fmt.Printf("Bytes written: %d\n", buf.Len())
	// Print users to stdout
	buf.WriteTo(os.Stdout)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
