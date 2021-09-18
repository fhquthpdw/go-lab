package main

import (
	"sync/atomic"
	"time"
)

type config struct {
	url  string
	port int
}

func loadConfig() config {
	return config{
		url:  "url",
		port: 80,
	}
}

func requests() chan int {
	return make(chan int)
}

func main() {
	var conf atomic.Value
	conf.Store(loadConfig())

	go func() {
		for {
			time.Sleep(1 * time.Millisecond)
			conf.Store(loadConfig())
		}
	}()

	for i := 0; i < 1000; i++ {
		go func() {
			//for r := range requests() {
			c := conf.Load().(config)
			_ = c
			//}
		}()
	}
}
