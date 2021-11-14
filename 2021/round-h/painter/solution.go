package main

import (
	"fmt"
)

var t, n int
var p string

func color(l byte) [3]bool {
	if l == 'R' {
		return [3]bool{true, false, false}
	}
	if l == 'Y' {
		return [3]bool{false, true, false}
	}
	if l == 'B' {
		return [3]bool{false, false, true}
	}
	if l == 'O' {
		return [3]bool{true, true, false}
	}
	if l == 'P' {
		return [3]bool{true, false, true}
	}
	if l == 'G' {
		return [3]bool{false, true, true}
	}
	if l == 'A' {
		return [3]bool{true, true, true}
	}
	return [3]bool{false, false, false}
}

func solve() int {
	fmt.Scan(&n, &p)

	c := 0
	r, y, b := false, false, false
	for i := 0; i < n; i++ {
		a := color(p[i])

		if a[0] {
			if !r {
				c++
				r = true
			}
		} else {
			r = false
		}

		if a[1] {
			if !y {
				c++
				y = true
			}
		} else {
			y = false
		}

		if a[2] {
			if !b {
				c++
				b = true
			}
		} else {
			b = false
		}
	}

	return c
}

func main() {
	fmt.Scan(&t)

	for i := 1; i <= t; i++ {
		fmt.Printf("Case #%d: %d\n", i, solve())
	}
}
