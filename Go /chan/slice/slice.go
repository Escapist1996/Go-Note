package main

import "fmt"

func main() {
	array := [6]int{10, 20, 30, 40}
	slice := array[0:2]
	// newarray := append(slice, 50)
	newarray := append(append(append(slice, 50), 60), 70)
	newarray[1] += 1
	fmt.Println(newarray)
	fmt.Println(slice)
}

//切片的扩容，就是当切片添加元素时，切片容量不够了，就会扩容，扩容的大小遵循下面的原则：
// （如果切片的容量小于1024个元素，那么扩容的时候slice的cap就翻番，乘以2；
// 一旦元素个数超过1024个元素，增长因子就变成1.25，即每次增加原来容量的四分之一。）
// 如果扩容之后，还没有触及原数组的容量，那么，切片中的指针指向的位置，就还是原数组（这就是产生bug的原因）；
// 如果扩容之后，超过了原数组的容量，那么，Go就会开辟一块新的内存，把原来的值拷贝过来，这种情况丝毫不会影响到原数组。
// 建议尽量避免bug的产生。
