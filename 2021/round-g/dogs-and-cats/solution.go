package main

import (
	"fmt"
)

var t, n, d, c, m int
var s string

func solve() string {
	fmt.Scan(&n, &d, &c, &m, &s)

	for a := range s {
		if s[a] == 67 {
			if c == 0 {
				for a < len(s) {
					if s[a] == 68 {
						return "NO"
					}
					a++
				}
			}
			c--
		} else {
			if d == 0 {
				return "NO"
			}
			d--
			c += m
		}
	}

	return "YES"
}

func main() {
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		fmt.Printf("Case #%d: %s\n", i+1, solve())
	}
}
