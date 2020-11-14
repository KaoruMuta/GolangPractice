package main

import (
	"fmt"
	"math"
)

const eps = 0.001

type ErrNegativeSqrt float64

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	z := 1.0
	for {
		tmp := z
		for i := 0; i < 10; i++ {
			z -= (z*z - x) / (2 * z)
		}
		if math.Abs(z-tmp) < eps {
			return z, nil
		}
	}
}

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g", e)
}

func main() {
	if answer, err := Sqrt(2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(answer)
	}

	if answer, err := Sqrt(-2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(answer)
	}
}
