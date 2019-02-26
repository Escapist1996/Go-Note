package main

import (
	"fmt"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func ch() {

	sigchan := make(chan int, 2)
	sigchan <- 9
	sigchan <- 10
	fmt.Println(<-sigchan)
	fmt.Println(<-sigchan)
}

func main() {
	ch()
	// s := []int{7, 2, 8, -9, 4, 0}

	// c := make(chan int)
	// go sum(s[:4], c)
	// go sum(s[4:], c)
	// x, y := <-c, <-c // receive from c
	// a := s[:4]
	// for _, v := range a {
	// 	fmt.Println(v)
	// }
	// fmt.Println("==========")
	// b := s[2:]
	// for _, v := range b {
	// 	fmt.Println(v)
	// }

	// fmt.Println(x, y, x+y)
}
