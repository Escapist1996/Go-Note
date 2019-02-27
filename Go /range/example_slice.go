package A

import "fmt"

func Swap(a, b *int) (int, int) {
	return *b, *a
}
func main() {
	fmt.Println("hello world")
	// a, b := 5, 10
	// Swap(&a, &b)
	// fmt.Println(a, b)
	// slice := []int{1, 2, 3, 4, 5}

	// slice1 := slice[:]
	// slice2 := slice[0:2]
	// slice3 := slice[:3]
	// // slice2[1] = 3
	// slice4 := append(slice2, 3)

	// fmt.Println(slice1)
	// fmt.Println(slice2)
	// fmt.Println(slice3)
	// fmt.Println(slice4)
}
