package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Count("cheese", "e"))
	fmt.Println(strings.Count("five", "e"))
	fmt.Println(strings.Count("five", "")) // before and after each rune
}
