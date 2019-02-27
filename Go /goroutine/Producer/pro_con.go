package main

import (
	"fmt"
	"sync"
)

func Productor(mychan chan int, date int, wait *sync.WaitGroup) {
	mychan <- date
	fmt.Println("productor date: ", date)
	wait.Done()
}

func Consumer(mychan chan int, date int, wait *sync.WaitGroup) {
	a := <-mychan
	fmt.Println("consumer date: ", a)
	wait.Done()
}

func main() {

	// var readOnlyChan <-chan int  // 只读chan
	// var writeOnlyChan chan<- int // 只写chan
	// var mychan chan int          // 可读可写chan
	// // 定义完成以后需要make分配内存，不然使用会deadlock
	// mychan = make(chan int, 10)

	// // Or
	// read_only := make(<-chan int, 10)
	// write_only := make(chan<- int, 10)
	// rw_chan := make(chan int, 10)

	datechan := make(chan int, 10)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		go Productor(datechan, i, &wg)
		wg.Add(1)
	}
	for j := 0; j < 10; j++ {
		go Consumer(datechan, j, &wg)
		wg.Add(1)
	}
	wg.Wait()
}
