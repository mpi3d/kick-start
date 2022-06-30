package main

import (
	"fmt"
)

func answer(s string) bool {
	m := make(map[rune]int)
	for _, v := range s {
		m[v]++
	}

	b := false
	for _, v := range m {
		if v%2 != 0 {
			if b {
				return false
			} else {
				b = true
			}
		}
	}

	return true
}

func solve() int {
	n, q := 0, 0
	fmt.Scan(&n, &q)

	s := ""
	fmt.Scan(&s)

	c := 0
	l, r := 0, 0
	for i := 0; i < q; i++ {
		fmt.Scan(&l, &r)
		if answer(s[l-1 : r]) {
			c++
		}
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
