package main

import (
	"fmt"
	"time"
)

func statusUpdate() string { return "" }

func main() {
	c := time.Tick(5 * time.Second)
	u := time.After(30 * time.Second)
	// for next := range c {
	// 	fmt.Printf("%v %s\n", next, statusUpdate())
	// }
	for {
		select {
		case next := <-c:
			fmt.Printf("%v %s\n", next, statusUpdate())
		case <-time.After(10 * time.Second):
			// This will never be printed
			fmt.Println("after 10 seconds...")
			return
		case <-u:
			fmt.Println("after 30 seconds...")
			return
		}
	}
}
