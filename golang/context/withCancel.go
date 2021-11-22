package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()

	for i := 0; i < 5; i++ {
		go Speak(ctx, i)
	}

	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}

func Speak(ctx context.Context, idx int) {
	i := 0
	for range time.Tick(time.Second) {
		select {
		case <-ctx.Done():
			fmt.Println(idx, "tick")
			return
		default:
			i++
			fmt.Println(idx, "=>", i)
		}
	}
}
