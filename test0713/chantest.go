package main

import (
	"fmt"
	   )


func Count(c chan int) {
	c <- 1
	fmt.Println("Counting")
}

func main() {
	ch := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		ch[i] = make(chan int)
		go Count(ch[i])
	}
	
	for _, c := range(ch) {
		<-c			//读取数据
	}
}
