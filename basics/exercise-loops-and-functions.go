package main

import (
	"fmt"
)

func Abs(x float64) float64 {
	if x < 0 {
		return -x	
	}
	
	return x
}

func Sqrt(x float64) float64 {
	
	//Precision for square root approximation
	DELTA := 0.00001
	
	//initial z value
	z := x/2
	
	for {
		x_prime := z * z
		delta := Abs(x_prime - x)
		if delta < DELTA {
			return z	
		}
		z -= (z*z - x) / (2*z)
	}
}

func main() {
	fmt.Println(Sqrt(2))
}
