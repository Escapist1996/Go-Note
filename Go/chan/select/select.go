package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	// ch <- 1

	select {
	case msg := <-ch:
		fmt.Println(msg)
	default:
		fmt.Println("default")
	}
	fmt.Println("-----")
}
