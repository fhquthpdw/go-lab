package main

import (
	"context"
	"golang.org/x/sync/semaphore"
)

type pFunc func()

func doParallel(total, parallel int, f pFunc) {
	//maxWorkers := runtime.GOMAXPROCS(0)
	sem := semaphore.NewWeighted(int64(parallel))
	ctx := context.TODO()

	for i := 0; i < total; i++ {
		if err := sem.Acquire(ctx, 1); err != nil {
			panic(err)
		}

		go func(i int) {
			defer sem.Release(int64(1))
			f()
		}(i)
	}

	if err := sem.Acquire(ctx, int64(total)); err != nil {
		panic(err)
	}

	return
}

func main() {

}
