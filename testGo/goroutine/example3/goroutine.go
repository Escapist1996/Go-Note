package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		go func(i int) {
			for {
				fmt.Println(i) //	IO操作，会交出控制权
			}
		}(i)
	}
	time.Sleep(time.Microsecond)
}

// 如果没有 time.Sleep，在for执行完 i 的自增后就会立刻结束程序，此时协程还没来得及处理就结束了，没有任何打印

// func main() {
// 	var arr [10]int
// 	for i := 0; i < 100; i++ {
// 		go func(i int) {
// 			for {
// 				arr[i]++		// 不会交出控制权，会死循环
// 			}
// 		}(i)
// 	}
// 	time.Sleep(time.Microsecond)
// }
