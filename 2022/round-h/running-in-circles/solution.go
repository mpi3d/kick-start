package main

import (
	"fmt"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func solve() int {
	l, n := 0, 0
	fmt.Scan(&l, &n)
	j := 0
	p, m := 0, ""
	for i := 0; i < n; i++ {
		d, c := 0, ""
		fmt.Scan(&d, &c)
		if m == "" {
			m = c
			j += d / l
			p = d % l
			continue
		}
		if m == c {
			p += d
			j += p / l
			p %= l
			continue
		}
		p = l - p
		p += d
		a := p/l - 1
		if a < 0 {
			p = l - p
			continue
		}
		j += a
		m = c
		p %= l
	}
	return j
}

func main() {
	t := 0
	fmt.Scan(&t)

	for i := 1; i <= t; i++ {
		fmt.Printf("Case #%v: %v\n", i, solve())
	}
}
