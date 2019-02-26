package main

import (
	"fmt"

	"../../interface"
)

func get(f file.File1) string {
	res := f.Read()
	return res
}

func main() {
	f := file.File1{}
	f.Write("aaaa")
	fmt.Println("get:", get(f))
}
