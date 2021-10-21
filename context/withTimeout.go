package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 11*time.Second)
	//defer cancel()

	for i := 0; i < 5; i++ {
		go MonitorTimeout(ctx, i)
	}

	time.Sleep(13 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}

func MonitorTimeout(ctx context.Context, idx int) {
	select {
	case <-ctx.Done():
		fmt.Println(idx, ctx.Err())
	case <-time.After(10 * time.Second):
		fmt.Println("stop monitor")
	}
}
