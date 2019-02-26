package main

import (
	"fmt"
)

type Animal interface {
	Speak() string
}

type Dog struct {
	Age interface{}
}

func (d Dog) Speak() string {
	return "WangWang~"
}

type Cat struct {
}

func (c Cat) Speak() string {
	return "Miao~"
}

type Llama struct {
}

func (l Llama) Speak() string {
	return "???"
}

func main() {
	animals := []Animal{Dog{}, Cat{}, Llama{}}
	for _, animals := range animals {
		fmt.Println(animals.Speak())
	}
}
