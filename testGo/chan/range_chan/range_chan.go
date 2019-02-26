package main

import "fmt"

func main() {

	mychan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		mychan <- i
	}
	close(mychan) // 关闭管道
	fmt.Println("date length: ", len(mychan))
	for v := range mychan { // 循环管道
		fmt.Println(v)
	}
	fmt.Printf("date length: %d\n", len(mychan))
}
