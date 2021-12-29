package main

import "fmt"

type Number interface {
	int | float64
}

func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
	}
	x := Sum(m)
	fmt.Println(x)
}

func Sum[K comparable, V Number](m map[K]V) V {
	var r V
	for _, val := range m {
		r += val
	}
	return r
}
