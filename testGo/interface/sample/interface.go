package main

import "fmt"

type I interface {
	Set(int)
	Get() int
}

type S struct {
	age int
}

func dosomething(v interface{}) {
}
func (s S) Get() int {
	return s.age
}

func (s *S) Set(age int) {
	s.age = age
}

func f(i I) {
	i.Set(10)
	fmt.Println(i.Get())
}

func main() {
	s := &S{}
	var i I
	i = s
	i.Set(20)
	// 类型断言
	if t, ok := i.(*S); ok {
		fmt.Println("s implements I :", t)
	}

	fmt.Println(i.Get())
}
