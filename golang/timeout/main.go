package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

const (
	Done    = 1
	Timeout = 2
)

func main() {
	//timeout()
	timeout2()
}

func timeout() {
	fmt.Println("Starting...")
	doBusiness := make(chan struct{})
	timeoutOrBusinessDoneQuit := make(chan int)

	go func() {
		for {
			select {
			case <-doBusiness: // 这个里面处理业务逻辑
				fmt.Println("Doing Business...")
				time.Sleep(10 * time.Second)
				fmt.Println("BusinessDone")
				timeoutOrBusinessDoneQuit <- Done // 业务处理完成
			case <-time.After(5 * time.Second): // 上面的ch如果一直没数据会阻塞，那么select也会检测其他case条件，检测到后3秒超时
				fmt.Println("TimeOut")
				timeoutOrBusinessDoneQuit <- Timeout // 业务处理完成
			}
		}
	}()
	doBusiness <- struct{}{}    // 开始处理业务
	<-timeoutOrBusinessDoneQuit // 这里暂时阻塞，直到业务处理完成或者超时
	fmt.Println("Over")
}

func timeout2() {
	fmt.Println("Timeout 2 Starting...")
	businessDone := make(chan struct{})

	go func() {
		// do business logic
		for i := 0; i < 10; i++ {
			time.Sleep(5 * time.Second)
			appendFile("a.txt", i)
			fmt.Println("loop ", i)
		}
		businessDone <- struct{}{}
	}()
	for {
		select {
		case <-businessDone:
			fmt.Println("business done")
		case <-time.After(3 * time.Second):
			fmt.Println("timeout")
		}
		break
	}

	time.Sleep(10 * time.Second)
}

func appendFile(fileName string, i int) {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer func() {
		_ = f.Close()
	}()

	newLine := fmt.Sprintf("new line %d", i)
	_, err = fmt.Fprintln(f, newLine)
	if err != nil {
		log.Fatal(err.Error())
	}
}
