package main

import (
	"fmt"
	"sync"
)

var (
	once      sync.Once
	singleIns *single
)

type single struct {
}

func getInstance() *single {
	if singleIns == nil {
		once.Do(func() {
			fmt.Println("Creating single instance now")
			singleIns = &single{}
		})
	} else {
		fmt.Println("Single instance already created")
	}

	return singleIns
}
