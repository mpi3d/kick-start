package main

import (
	"fmt"
	"math"
)

var t, n int
var s []bool

func l(i int) float64 {
	if i == 0 {
		return math.Inf(-1)
	}

	if s[i-1] {
		return float64(i - 1)
	}

	return l(i - 1)
}

func r(i int) float64 {
	if i == n-1 {
		return math.Inf(1)
	}

	if s[i+1] {
		return float64(i + 1)
	}

	return r(i + 1)
}

func f(i int) int {
	if s[i] {
		return 0
	}

	return int(math.Min(float64(i)-l(i), r(i)-float64(i)))
}

func solve() int {
	fmt.Scan(&n)

	s = make([]bool, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}

	a := 0
	for i := 0; i < n; i++ {
		a += f(i)
	}

	return a
}

func main() {
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		fmt.Printf("Case #%d: %d\n", i+1, solve())
	}
}
