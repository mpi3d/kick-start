package main

import (
	"fmt"
	"sort"
)

func sum(a []int) int {
	s := 0
	for _, v := range a {
		s += v
	}
	return s
}

func median(a []int) float64 {
	l := len(a)
	if l%2 == 0 {
		l /= 2
		return float64(a[l-1]+a[l]) / 2
	}
	return float64(a[l/2])
}

func solve() float64 {
	n, m := 0, 0
	fmt.Scan(&n, &m)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	sort.Ints(a)

	i := len(a) - m + 1
	return float64(sum(a[i:])) + median(a[:i])
}

func main() {
	t := 0
	fmt.Scan(&t)

	for i := 1; i <= t; i++ {
		fmt.Printf("Case #%d: %f\n", i, solve())
	}
}
