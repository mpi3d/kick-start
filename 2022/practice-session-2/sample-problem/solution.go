package main

import (
	"fmt"
)

func solve() int {
	var n, m, b int
	fmt.Scan(&n, &m)

	c := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&b)
		c += b
	}

	return c % m
}

func main() {
	var t int
	fmt.Scan(&t)

	for i := 1; i <= t; i++ {
		fmt.Printf("Case #%d: %d\n", i, solve())
	}
}
