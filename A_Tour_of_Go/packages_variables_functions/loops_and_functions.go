package main

import (
	"fmt"
	"math"
)

// eps ... an allowable error
const eps = 0.001

// Sqrt ... function to calculate sqrt
func Sqrt(x float64) float64 {
	z := 1.0
	for {
		tmp := z
		for i := 0; i < 10; i++ {
			z -= (z*z - x) / (2 * z)
		}
		if math.Abs(z-tmp) < eps {
			return z
		}
	}
}

func main() {
	fmt.Println(Sqrt(2))
}
