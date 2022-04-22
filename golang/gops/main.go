package main

// usage:
// go build -o gops-example
// ./gops-example
// gops ...

import (
	"log"
	"runtime"
	"time"

	"github.com/google/gops/agent"
)

func main() {
	if err := agent.Listen(agent.Options{
		ShutdownCleanup: true, // automatically closes on os.Interrupt
	}); err != nil {
		log.Fatal(err)
	}

	var t = time.NewTicker(time.Second * 1)
	var ia []int
	var idx int

	for {
		if idx%5 == 0 {
			runtime.GC()
		}
		idx++
		select {
		case <-t.C:
			go func() {
				ia = append(ia, genIna()...)
				select {}
			}()
		}
	}
	//time.Sleep(time.Hour)
}

func genIna() (a []int) {
	for i := 0; i < 10000; i++ {
		a = append(a, i)
	}
	return a
}
