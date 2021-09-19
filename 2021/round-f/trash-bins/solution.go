package main

import (
	"fmt"
	"math"
)

func l(n int, i int, s []bool) float64 {
	if i == 0 {
		return math.Inf(-1)
	}

	if s[i-1] {
		return float64(i - 1)
	}

	return l(n, i-1, s)
}

func r(n int, i int, s []bool) float64 {
	if i == n-1 {
		return math.Inf(1)
	}

	if s[i+1] {
		return float64(i + 1)
	}

	return r(n, i+1, s)
}

func f(n int, i int, s []bool) int {
	if s[i] {
		return 0
	}

	return int(math.Min(float64(i)-l(n, i, s), r(n, i, s)-float64(i)))
}

func solve() int {
	n := 0
	fmt.Scan(&n)

	s := make([]bool, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}

	a := 0
	for i := 0; i < n; i++ {
		a += f(n, i, s)
	}

	return a
}

func main() {
	var t int

	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		fmt.Printf("Case #%d: %d\n", i+1, solve())
	}
}
