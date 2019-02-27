package main

import "fmt"

func main() {
	a := 1
	print(a)
	defer print(a)
	a = 2

	defer print(function(a))
}

func function(num int) int {
	return num
}
func print(num int) {
	fmt.Println(num)
}
