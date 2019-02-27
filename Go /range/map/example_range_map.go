package main

import "fmt"

func main() {
	ages := map[string]int{"张": 11, "李": 22, "吴": 33}

	for n, age := range ages {
		fmt.Println(n, age)
	}
}
