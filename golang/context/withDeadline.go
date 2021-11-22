package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	later, _ := time.ParseDuration("11s")
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(later))
	//defer cancel()

	for i := 0; i < 5; i++ {
		go Monitor(ctx, i)
	}

	time.Sleep(13 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}

func Monitor(ctx context.Context, idx int) {
	select {
	case <-ctx.Done():
		fmt.Println(idx, ctx.Err())
	case <-time.After(10 * time.Second):
		fmt.Println("stop monitor")
	}
}
