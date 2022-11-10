package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%q\n", strings.SplitAfterN("a,b,c", ",", 2))
	fmt.Printf("%q\n", strings.SplitAfterN("a,b,c", ",", -1))
}
