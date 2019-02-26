package main

import (
	"fmt"
	"math"
	   )

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func Sqrt(x float64) (float64, error) {
	return 0, nil
}

func (e ErrNegativeSqrt) Error() string {

}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
	
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-1))
	
	fmt.Sprint(e)
}
