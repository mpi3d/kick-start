package main

import (
	"fmt"
	"math"
)

func solve() float64 {
	r, a, b, c := 0, 0, 0, 0
	fmt.Scan(&r, &a, &b)

	for r > 0 {
		c += r * r
		r *= a
		c += r * r
		r /= b
	}

	return float64(c) * math.Pi
}

func main() {
	t := 0
	fmt.Scan(&t)

	for i := 1; i <= t; i++ {
		fmt.Printf("Case #%d: %f\n", i, solve())
	}
}
