package main

import "fmt"

const (
	i = 7
	j
	k
)

func main() {
	a := []string{"a", "b", "c"}
	a = nil
	fmt.Println(a, len(a), cap(a))
}
