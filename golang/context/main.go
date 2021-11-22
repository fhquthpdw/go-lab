// https://studygolang.com/articles/30675

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go Hello(ctx, 2000*time.Millisecond)

	select {
	case <-ctx.Done():
		fmt.Println("Hello:", ctx.Err())
	}
}

func Hello(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("Hello Done:", ctx.Err())
	case <-time.After(duration):
		fmt.Println("Hello process request with", duration)
	}
}
