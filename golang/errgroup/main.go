package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"
)

func justError() {
	var g errgroup.Group
	var urls = []string{
		"https://www.baidu.com",
		"https://www.sina.com",
		"https://www.sohu.com",
	}

	for _, url := range urls {
		g.Go(func() error {
			resp, err := http.Get(url)
			if err == nil {
				resp.Body.Close()
			}
			return err
		})
	}

	if err := g.Wait(); err == nil {
		fmt.Println("all success")
	}
}

func errorWithContext() {
	var times [100][0]int
	eg, ctx := errgroup.WithContext(context.Background())
	i := 0
	for range times { // 长度为 0 的数组不占用内存空间，0 内存的100次迭代
		i = i + 1
		eg.Go(func() error {
			time.Sleep(2 * time.Second)
			select {
			case <-ctx.Done():
				fmt.Println("Canceled:", i)
				return nil
			default:
				if i > 90 {
					fmt.Println("Error:", i)
					return fmt.Errorf("Error: %d", i)
				}
				return nil
			}
		})
	}

	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}
}
