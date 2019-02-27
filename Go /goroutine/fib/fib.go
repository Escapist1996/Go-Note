package main

// func main() {
// 	var x chan int
// 	go func() {
// 		x <- 1
// 	}()
// 	<-x
// }

// func main() {
// 	go spinner(100 * time.Millisecond)
// 	const n = 45
// 	fibN := fib(n) // slow
// 	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
// }
// func spinner(delay time.Duration) {
// 	for {
// 		for _, r := range `-\|/` {
// 			fmt.Printf("\r%c", r)
// 			time.Sleep(delay)
// 		}
// 	}
// }
// func fib(x int) int {
// 	if x < 2 {
// 		return x
// 	}
// 	return fib(x-1) + fib(x-2)
// }
import "fmt"

type people interface {
	speak()
}

type student struct {
	name string
	age  int
}

func (stu student) speak() {
	fmt.Println("I am a student, I am ", stu.age)
}

func main() {
	var p people
	p = student{name: "RyuGou", age: 12} //应该改为 p = &student{name:"RyuGou", age:12}
	p.speak()
}
