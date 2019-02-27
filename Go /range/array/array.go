package main

import (
	"fmt"
	"reflect"
)

func main() {
	arrayA := [...]int{1, 2, 3, 4}
	arrayB := [...]int{1, 2, 3}

	fmt.Println(reflect.TypeOf(arrayA)) // Type : [4]int
	fmt.Println(reflect.TypeOf(arrayB)) // Type : [3]int

	// 数组可以指定索引和对应值的方式进行初始化
	arrayC := [4]int{0: 1, 2: 3, 3: 4}
	fmt.Println(arrayC) // 1 0 3 4
}
