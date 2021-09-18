package main

import "fmt"

// 运行命令: go run -gcflags=-G=3 ./main.go
func main() {
	strs := []string{"Hello", "World", "Generics"}
	decs := []float64{3.14, 1.14, 1.618, 2.718}
	nums := []int{2, 4, 6, 8}
	p(strs)
	p(decs)
	p(nums)
}

func p[T any] (arr []T) {
	for _, v := range arr {
		fmt.Print(v)
		fmt.Print(" ")
	}
	fmt.Println("")
}
