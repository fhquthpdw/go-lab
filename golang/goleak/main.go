package main

import (
	"time"
)

func main() {
	leak()
}

var ch = make(chan int)

func leak() {
	var handleDone = make(chan struct{})

	go func() {
		do()
		handleDone <- struct{}{}
	}()

	select {
	case <-handleDone:
	case <-time.After(2 * time.Second):
	}
}

func do() {
	time.Sleep(5 * time.Second)
}
