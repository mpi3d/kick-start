package main

import (
	"fmt"
)

func solve() int {
	n, m := 0, 0
	fmt.Scan(&n, &m)

	c := 0
	for i, b := 0, 0; i < n; i++ {
		fmt.Scan(&b)
		c += b
	}

	return c % m
}

func main() {
	t := 0
	fmt.Scan(&t)

	for i := 1; i <= t; i++ {
		fmt.Printf("Case #%v: %v\n", i, solve())
	}
}
