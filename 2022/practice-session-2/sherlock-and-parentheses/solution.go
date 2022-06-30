package main

import (
	"fmt"
	"math"
)

func solve() int {
	l, r := 0, 0
	fmt.Scan(&l, &r)

	b := int(math.Min(float64(l), float64(r)))
	c := 0
	for i := 1; i <= b; i++ {
		c += i
	}

	return c
}

func main() {
	t := 0
	fmt.Scan(&t)

	for i := 1; i <= t; i++ {
		fmt.Printf("Case #%d: %d\n", i, solve())
	}
}
