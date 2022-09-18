package main

import (
	"fmt"
)

func solve() int {
	n := 0
	fmt.Scan(&n)

	c := 0
	m := -1
	p := -1
	for i, v := 0, 0; i < n; i++ {
		fmt.Scan(&v)
		if v > m {
			m = v
			p = v
		} else {
			if p > v {
				c++
			}
			p = -1
		}
	}
	if p > -1 {
		c++
	}

	return c
}

func main() {
	t := 0
	fmt.Scan(&t)

	for i := 1; i <= t; i++ {
		fmt.Printf("Case #%v: %v\n", i, solve())
	}
}
