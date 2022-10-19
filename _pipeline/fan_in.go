package main

import (
	"fmt"
	"sync"
)

func gen(nums ...int) <-chan int {
	c := make(chan int)
	go func() {
		for _, n := range nums {
			c <- n
		}
		close(c)
	}()
	return c
}

func sq(in <-chan int) <-chan int {
	c := make(chan int)

	go func() {
		for v := range in {
			c <- v
		}
		close(c)
	}()

	return c
}

func merge(chs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel in cs.
	// output copies values from c to out until c is closed, then
	// calls wg.Done().
	output := func(ch <-chan int) {
		for v := range ch {
			out <- v
		}
		wg.Done()
	}

	for _, ch := range chs {
		wg.Add(1)
		go output(ch)
	}

	// Start a goroutine to close out once all the output goroutines
	// are done. This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	ch := gen(1, 2, 3, 4, 5)

	// Distribute the sq work across two goroutines
	// that both read from in.
	out1 := sq(ch)
	out2 := sq(ch)

	// Consume the merged output from c1 and c2
	for n := range merge(out1, out2) {
		fmt.Println(n)
	}
}
