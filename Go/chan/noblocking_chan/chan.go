package main

import "fmt"

func send(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("send date: ", i)
	}
}

func main() {
	snch := make(chan int, 10)
	rvch := make(chan string, 10)
	go send(snch)

	rvch <- "ch"
	select {
	case a := <-snch:
		{
			fmt.Println("recv channel: ", a)
		}
	case a := <-rvch:
		{
			fmt.Println("send channel: ", a)
		}
	default:
		fmt.Println("no date in channel")
	}
}
