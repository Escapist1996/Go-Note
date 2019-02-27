package main

import (
	"fmt"
	"testing"
)

func work(c chan int) {
	for {
		r := <-c // 管道读出
		fmt.Println(r)
	}
}

func workFactory() chan int {
	c := make(chan int)
	go func(c chan int) {
		fmt.Println(<-c)
	}(c)
	return c
}

// func main() {
// 	// c := make(chan int)
// 	// go work(c)
// 	// c <- 1 // 管道写入
// 	// c <- 2
// 	// time.Sleep(time.Microsecond)

// 	var i int
// 	var arr [10]chan int
// 	for i = 0; i < 10; i++ {
// 		arr[i] = workFactory()
// 	}
// 	for i = 0; i < 10; i++ {
// 		arr[i] <- i
// 	}

// 	time.Sleep(time.Microsecond)
// }
func create(i int) chan int {
	c := make(chan int)
	go func(c chan int, i int) {
		for {
			c <- 'a' + i
		}
	}(c, i)
	return c
}
func main() {
	c1, c2 := create(1), create(2)
	for {
		select { // 由于没有 default,会一直阻塞读
		case n1 := <-c1:
			fmt.Printf("%c \n", n1)
		case n2 := <-c2:
			fmt.Printf("%c \n", n2)
		}
	}
}

func testAdd1(t *testing.T) { // 开头小写，会被跳过

}
