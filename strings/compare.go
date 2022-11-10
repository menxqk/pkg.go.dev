package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Compare("a", "b"))
	fmt.Println(strings.Compare("a", "a"))
	fmt.Println(strings.Compare("b", "a"))
	fmt.Println(strings.Compare("abc", "abc"))
	fmt.Println(strings.Compare("abc", "acb"))
}
