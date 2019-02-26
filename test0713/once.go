package main

import (
	"sync"
	)



var a string
var once sync.Once

func setup() {
	a = "Hello world"
}

func doprint() {
	once.setup()
	print(a)
}

func main() {
	go doprint()
	go doprint()
}
