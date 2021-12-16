package main

import (
	"fmt"
)

func main() {
	eerr := ee2()
	fmt.Println(eerr)

	err := e2()
	fmt.Println(err)
}

func e1() error {
	return fmt.Errorf("error 1")
}

func e2() error {
	return fmt.Errorf("error 2: %w", e1())
}

func ee2() error {
	err := e1()
	return fmt.Errorf("%v, ee 2: %v", err, e1())
}
