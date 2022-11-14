package main

import (
	"fmt"
	"time"
)

func main() {
	u, _ := time.ParseDuration("1us")
	fmt.Printf("One microsecond is %d nanoseconds.\n", u.Nanoseconds())
}
