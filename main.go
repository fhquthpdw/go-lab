package main

import "fmt"

// 如何避免出现连续的 if err := f(); err != nil {} 的语句
type person struct {
	Age1 int
	Age2 int
	Age3 int
	Age4 int
	Age5 int
}

func f(p person) {
	var err error
	checkErr := func() {
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
