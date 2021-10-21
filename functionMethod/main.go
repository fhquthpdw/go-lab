package main

import "fmt"

type FType func() string

func (f FType) Print() {
	fmt.Println(f())
}

func main() {
	WayOne()
	WayTwo()
}

func WayOne() {
	var f FType
	f = func() string {
		return "this is from way one"
	}
	f.Print()
}

func WayTwo() {
	f := FType(func() string {
		return "this is from way two"
	})
	f.Print()
}
