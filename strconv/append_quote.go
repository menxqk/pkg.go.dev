package main

import (
	"fmt"
	"strconv"
)

func main() {
	b := []byte("quote:")
	b = strconv.AppendQuote(b, `"Fran & Freddie's Diner"`)
	fmt.Println(string(b))

	c := []byte("quote:")
	c = strconv.AppendQuote(c, `Fran & Freddie's Diner`)
	fmt.Println(string(c))
}
