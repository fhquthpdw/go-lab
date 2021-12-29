// 一个 select 死锁的问题
// 总结一下就是如果遇到 ch1 <- <- ch0 的情况，那么，<- cho 会被先求职，ch1 就会一直阻塞，直到 ch0 发送数据，这样就会导致 select 死锁
// From:
//	- https://stackoverflow.com/questions/51167940/chained-channel-operations-in-a-single-select-case
//	- https://blog.huoding.com/2021/08/29/947
package main

import (
	"fmt"
	"time"
)

func talk(msg string, sleep int) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(sleep) * time.Millisecond)
		}
	}()
	return ch
}

func fanIn(input1, input2 <-chan string) <-chan string {
	ch := make(chan string)
	go func() {
		for {
			select {
			case ch <- <-input1:
			case ch <- <-input2:
			}

			// 下面这种 select 的写法是对的
			/*
				select {
				case t := <-input1:
					ch <- t
				case t := <-input2:
					ch <- t
				}
			*/
		}
	}()
	return ch
}

func main() {
	// 版本一

	// 版本二
	// 这个版本中，只会输出 talk 中循环次数的个数数据，因为任何一个 talk 执行完了， select 中的一个 case 就会永远阻塞
	// 为什么 talk 里的 i 是顺序的呢？
	// 答：因为在 fanIn 中每次 for 循环后，select 语句里面的 <-input1 和 <-input2 都会被重新求值
	//	  所有每一次循环 <-input1 和 <-input2 就都会执行一次
	//	  所以，input1 和 input2 对应的 talk 是每次 for 里面都会执行一次
	// 	  注意：如果多个case满足读写条件，select会随机选择一个语句执行，其他的会被丢弃掉
	ch := fanIn(talk("A", 10), talk("B", 1000))
	for i := 0; i < 10; i++ {
		fmt.Printf("%q\n", <-ch)
	}
}
