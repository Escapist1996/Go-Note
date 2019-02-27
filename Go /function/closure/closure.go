package main

import "fmt"

// 闭包

func add() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

func main() {
	add := add()
	for i := 0; i < 10; i++ {
		fmt.Println(add(i))
	}
}
