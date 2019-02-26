package main

import (
	"fmt"
	"runtime"
	"sync"
		)

func main() {
	runtime.GOMAXPROCS(2)
	
	var wg sync.WaitGroup		//wg用来等待程序完成
	wg.Add(2)

	fmt.Println("Start Goroutines")

	go func() {
		defer wg.Done()			//函数退出调用Done来通知main函数工作完成

		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}
