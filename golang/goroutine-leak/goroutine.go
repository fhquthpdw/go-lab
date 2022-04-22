package main

import (
	"fmt"
	"sync"
)

func queryAll() {
	ch := make(chan int)
	go func() { ch <- 1 }()
	go func() { ch <- 2 }()
	go func() { ch <- 3 }()
	// <-ch
	// <-ch
	<-ch
}

func WaitingGoroutineLeak() {
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Add(-1)

			for k := range ch {
				fmt.Println(k)
			}
		}()
	}
	wg.Wait()
	close(ch)
}
