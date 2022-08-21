package main

import (
	"fmt"
)

func solve() []int {
	n := 0
	fmt.Scan(&n)

	r := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&r[i])
	}

	m := make([]int, n)
	for i, v := range r {
		n := -1
		x := 2 * v
		for j, w := range r {
			if i == j {
				continue
			}
			if n < w && w <= x {
				n = w
			}
		}
		m[i] = n
	}

	return m
}

func main() {
	t := 0
	fmt.Scan(&t)

	for i := 1; i <= t; i++ {
		s := fmt.Sprint(solve())
		fmt.Printf("Case #%v: %v\n", i, s[1:len(s)-1])
	}
}
