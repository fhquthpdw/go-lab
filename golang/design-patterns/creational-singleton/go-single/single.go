package main

import (
	"fmt"
	"sync"
)

var (
	lock      = &sync.Mutex{}
	singleIns *single
)

type single struct {
}

func getInstance() *single {
	if singleIns == nil {
		lock.Lock()
		defer lock.Unlock()

		if singleIns == nil {
			fmt.Println("Creating single instance now")
			singleIns = &single{}
		} else {
			fmt.Println("Single instance already created")
		}
	} else {
		fmt.Println("Single instance already created")
	}

	return singleIns
}
