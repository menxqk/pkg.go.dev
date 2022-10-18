package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// gen generates integers in a separate goroutine and
	// send them to the returned channel.
	// The callers of gen need to cancel the context once
	// they are done consuming generate integers not to leak
	// the internal goroutine started by gen.
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return // returning not to leak the goroutine
				case dst <- n:
					time.Sleep(1 * time.Second)
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}
