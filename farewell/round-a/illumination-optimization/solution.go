package main

import (
	"fmt"
)

func solve() string {
	m, r, n := 0, 0, 0
	fmt.Scan(&m, &r, &n)
	x := make([]int, n)
	for i := range x {
		fmt.Scan(&x[i])
	}

	c := 0
	i, p, t := 0, -1, 0
	for t < m {
		c++
		t += r
		for x[i] <= t {
			i++
			if i == n {
				break
			}
		}
		i--
		if i == p {
			return "IMPOSSIBLE"
		}
		p = i
		t = x[i] + r
	}
	return fmt.Sprint(c)
}

func main() {
	t := 0
	fmt.Scan(&t)

	for i := 1; i <= t; i++ {
		fmt.Printf("Case #%v: %v\n", i, solve())
	}
}
