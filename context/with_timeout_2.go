package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	gen := func(ctx context.Context, c chan<- int) {
		n := 1
		for {
			select {
			case <-ctx.Done():
				close(c)
				fmt.Println("Done...")
				return
			case c <- n:
				time.Sleep(1 * time.Second)
				n++
			}
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	c := make(chan int)

	go gen(ctx, c)

	for i := range c {
		fmt.Println(i)
	}
}
