package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt int

func (e ErrNegativeSqrt) Error() string {
	s := fmt.Sprintf("%d", int(e))
	return "cannot Sqrt negative number: " + s
}

func Sqrt(x int) (float64, error) {

	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	return math.Sqrt(float64(x)), nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
