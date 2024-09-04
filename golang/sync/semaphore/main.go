package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"time"

	"golang.org/x/sync/semaphore"
)

func main() {
	ctx := context.TODO()
	var (
		maxWorkers = runtime.GOMAXPROCS(20)
		sem        = semaphore.NewWeighted(int64(maxWorkers)) // 设置最大并发数
		out        = make([]int, 132)
	)

	for i := range out {
		if err := sem.Acquire(ctx, 1); err != nil {
			log.Printf("Failed to acquire semaphore: %v", err)
		}

		go func(i int) {
			defer sem.Release(1)

			out[i] = returnInt(i)
		}(i)
	}

	// 这里算是回收 goroutine，全部回收之后流程才可以往下走
	// 注意这里的 Acquire 方法的第二个参数的值不是1哦
	if err := sem.Acquire(ctx, int64(maxWorkers)); err != nil {
		log.Printf("Failed to acquire semaphore: %v", err)
	}

	fmt.Println(out)
}

func returnInt(i int) int {
	fmt.Println(i)
	n := rand.Int31n(100)
	time.Sleep(time.Duration(n) * time.Millisecond)
	return i
}
