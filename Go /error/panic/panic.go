package main

import "fmt"

// 捕获panic
func dopanic() {
	defer func() {
		err := recover()
		fmt.Println(err)
	}()
	panic("error!!")
}

// 捕获其他异常
func dootherpanic() {
	defer func() {
		err := recover()
		fmt.Println(err)
	}()
	a := 0
	b := 1 / a
	fmt.Println(b)
}

// 无异常处理
func donopanic() {
	defer func() {
		err := recover()
		if err, ok := err.(error); ok {
			fmt.Println(err)
		} else {
			fmt.Println("no error!")
		}
	}()
	a := 0
	fmt.Println(a)
}

func main() {
	dopanic()
	dootherpanic()
	donopanic()
}
