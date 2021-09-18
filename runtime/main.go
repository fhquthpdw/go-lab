package main

import (
	"fmt"
	"runtime"
)

func main() {
	caller()
}

func caller() {
	fmt.Println(subFunc())
}

func subFunc() int {
	pc, file, line, ok := runtime.Caller(1)
	pcFunc := runtime.FuncForPC(pc)
	callerName := pcFunc.Name()
	fmt.Println(fmt.Sprintf("%p, %s, %d, %t, %s", caller, file, line, ok, callerName))
	return 1
}
