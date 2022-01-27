package main

import (
	"fmt"
)

type Phone interface {
	Call()
	SendMsg(s string) bool
}

type Android struct {
	Name string
}

func (a Android) Call() {
	fmt.Println("call")
}
func (a Android) SendMsg(s string) bool {
	fmt.Println(s)
	return true
}

type Ios struct{}

func (a Ios) Call() {
	fmt.Println("call")
}
func (a Ios) SendMsg(s string) bool {
	fmt.Println(s)
	return true
}

func main() {
	str := "hello"
	fmt.Println(len(str))
	for _, v := range str {
		fmt.Println(string(v))
	}

	/*
		k := []int{1, 2, 3, 4, 5}
		i := 4

		copy(k[i:], k[i+1:])
		k = k[:len(k)-1]
		fmt.Println(k)

		x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		for idx := 0; idx < len(x); idx++ {
			if x[idx]%3 == 0 {
				copy(x[idx:], x[idx+1:])
				x = x[:len(x)-1]
			}
			if idx+1 > len(x) {
				break
			}
		}
		fmt.Println(x)
	*/
}
