package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"
)

var gf singleflight.Group

func main() {
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done()

			result, _, _ := gf.Do("test_gf", func() (interface{}, error) {
				return gfGetData("test_gf-" + strconv.Itoa(i)), nil
			})
			resultStr := result.(string)
			fmt.Println(resultStr)
		}(i)
	}
	wg.Wait()
}

func gfGetData(key string) string {
	fmt.Printf("getting data by key: %s\n", key)
	time.Sleep(2 * time.Second)
	return "return data from gfGetData"
}
