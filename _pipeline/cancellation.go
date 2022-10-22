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

func sq(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
			case <-done:
				return
			}

		}
	}()
	return out
}

func merge(done <-chan struct{}, chs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start and output goroutine for each channel in chs.
	// output copies values from c ou out until c is closed
	// or it receives a value from done, then outputs wg.Done
	output := func(ch <-chan int) {
		defer wg.Done()
		for n := range ch {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
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
	// Set up a done channel that's shared by whole pipeline,
	// and close that channel when this pipeline exists, as a signal
	// for all the goroutines we started to exit.
	done := make(chan struct{})
	defer close(done)

	in := gen(2, 3)

	// Distribute the sq work across two goroutines both read from in.
	c1 := sq(done, in)
	c2 := sq(done, in)

	// Consume the first value from output.
	out := merge(done, c1, c2)
	fmt.Println(<-out)
}
