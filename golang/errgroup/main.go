package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	var c = make(chan int, 4)
	<-c
	fmt.Println("gogogo")

	var g errgroup.Group

	g.Go(func() error {
		time.Sleep(3 * time.Second)
		return errors.New("ggo1 error")
	})

	g.Go(func() error {
		time.Sleep(15 * time.Second)
		return errors.New("ggo2 error")
	})

	if err := g.Wait(); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("all done")
}

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
				_ = resp.Body.Close()
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
				fmt.Println("canceled:", i)
				return nil
			default:
				if i > 90 {
					fmt.Println("error:", i)
					return fmt.Errorf("error: %d", i)
				}
				return nil
			}
		})
	}

	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}
}
