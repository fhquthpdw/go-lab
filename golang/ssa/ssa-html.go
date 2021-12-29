// https://blog.huoding.com/2021/10/11/960
// 通过SSA工具，可以展现出源代码到汇编过程中，编译器做了哪些工作，并且可以把结果生成 html 文件
// SSA 工具最方便的地方是它可以把源代码和汇编通过颜色对应起来
// 使用如下命令编译：
// $ GOSSAFUNC=main go build -gcflags="-N -l" ./main.go
// 会在当前目录下生成一个 ssa.html 文件，用浏览器打开就能看到内容了
// 汇编可以看这里： https://github.com/cch123/golang-notes/blob/master/assembly.md

package main

import "time"

func main() {
	running := true
	go func() {
		println("start thread1")
		count := 1
		for running {
			count++
		}
		println("end thread1: count =", count)
	}()
	go func() {
		println("start thread2")
		for {
			running = false
		}
	}()
	time.Sleep(time.Hour)
}

// 不过为什么「running = false」这行源代码没有对应的汇编呢？这是因为 SSA 的工作单位是函数，
// 上面的结果是 main 函数的结果，而「running = false」实际上属于 main 函数里第 2 个 goroutine，
// 相当于 main.func2，重新运行 SSA
// $ GOSSAFUNC=main.func2 go build -gcflags="-N -l" ./main.go
