package main

import "fmt"

func main() {
	ages := []string{"11", "12", "15", "16", "18", "20"}

	for i, age := range ages {
		fmt.Println(i, age)
	}
}
