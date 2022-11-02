package main

import (
	"fmt"
	"io"
	"log"
)

func main() {
	r := strings.NewReaer("some io.Reader stream to be read\n")
	s := io.NewSectionReader(r, 5, 17)

	buf := make([]byte, 9)
	if _, err := s.Read(buf); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", buf)
}
