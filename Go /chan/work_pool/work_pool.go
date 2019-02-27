package main

import "fmt"

func Task(taskch, resch chan int, exitch chan bool) {
	// 异常处理
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Task error: ", err)
			return
		}
	}()

	// 处理任务
	for t := range taskch {
		resch <- t
	}
	// 处理完任务发送信号
	exitch <- true
}
func main() {
	taskch := make(chan int, 20)
	resch := make(chan int, 20)
	exitch := make(chan bool, 5)

	go func() {
		for i := 0; i < 20; i++ {
			taskch <- i
		}
		close(taskch)
	}()

	// 启动5个goroutine
	for i := 0; i < 5; i++ {
		go Task(taskch, resch, exitch)
	}
	//等待5个goroutine结束
	go func() {
		for i := 0; i < 5; i++ {
			<-exitch
		}
		close(resch)
		close(exitch)
	}()

	// 打印结果
	for res := range resch {
		fmt.Println("Task res: ", res)
	}
}
