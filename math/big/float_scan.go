package main

import (
	"fmt"
	"log"
	"math/big"
)

func main() {
	// The Scan function is rarely used directly;
	// the fmt package recognizes it as an implmentation of fmt.Scanner.
	f := new(big.Float)
	_, err := fmt.Sscan("1.19282e99", f)
	if err != nil {
		log.Println("error scanning value:", err)
	} else {
		fmt.Println(f)
	}
}
