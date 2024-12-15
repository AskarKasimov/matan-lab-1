package main

import (
	"fmt"
	"math"
)

func bisectionMethod(f func(float64) float64, a, b, tol float64) (float64, float64, int) {
	if f(a) == 0 {
		return a, 0, 0
	}
	if f(b) == 0 {
		return b, 0, 0
	}
	iters := 0
	var x_i float64 = 0
	for {
		if (b - a) <= tol {
			break
		}
		if iters > int(math.Log2((b-a)/tol)) {
			break
		}
		iters += 1
		dx := (b - a) / 2
		x_i = a + dx
		if math.Signbit(f(a)) != math.Signbit(f(x_i)) {
			b = x_i
		} else {
			a = x_i
		}
	}
	return x_i, tol, iters
}

func f(x float64) float64 {
	return math.Pow(2, math.Pow(x, 2)) - 1/x
}

func main() {
	fmt.Println(bisectionMethod(f, 0.001, 1.0, 0.000000000000000005))
}
