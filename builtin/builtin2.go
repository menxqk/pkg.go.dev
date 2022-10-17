package main

import (
	"time"
)

func main() {
	s := "A string\n"
	ss := "Another string\n"
	print(s, ss)

	print(time.Now().String() + "\n")
	c := make(chan string)
	go sendAfter5(c)
	print(<-c)
	close(c)

	defer func() {
		e := recover()
		es, ok := e.(string)
		if ok {
			print("recovered from: ", string(es+"\n"))
		}
	}()
	panic("Intentional panic!")
}

func sendAfter5(c chan<- string) {
	time.Sleep(5 * time.Second)
	c <- time.Now().String() + "\n"
}
