package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	u, err := url.Parse("http://example.com/#x/y%2fz")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Fragment:", u.Fragment)
	fmt.Println("RawFragment:", u.RawFragment)
	fmt.Println("EscapedFragment:", u.EscapedFragment())
}
