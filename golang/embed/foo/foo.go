package foo

import "fmt"

type A struct {
}

func (a A) M() {
	fmt.Println("M")
}
