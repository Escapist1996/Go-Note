package main

import (
	"fmt"
	"runtime"
	"sync"
	   )

var (
	counter int
	wg sync.WaitGroup	//等待程序结束

	mutex sync.Mutex	//定义临界区
	)

func main() {
	wg.Add(2)
	
	go incCounter(1)
	go incCounter(2)

	wg.Wait()
	fmt.Printf("Final Counter: %d\n", counter)
}
func incCounter(id int) {

	defer wg.Done()

	for count := 0; count < 2; count++ {
		
		mutex.Lock()
		{
			value := counter
			runtime.Gosched()

			value++
			counter = value
		}
		mutex.Unlock()
	}
}
