package main

import "fmt"

func main() {
	var f []func()
	array := []string{
		"1", "2", "3", "4",
	}
	for _, e := range array {
		// elem := e				// 每次循环后，每个匿名函数中保存的是局部变量elem的值，相当于定义了四个局部变量
		// f = append(f, func() {
		// 	fmt.Println(elem)		// 1 2 3 4
		// })
		f = append(f, func() { // 匿名函数中记录的是循环变量的内存地址，而不是变量某一个时刻的值
			fmt.Println(e) // 4 4 4 4
		})
	}
	for _, v := range f {
		v()
	}
}
