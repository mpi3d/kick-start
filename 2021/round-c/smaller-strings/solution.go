package main

import (
	"fmt"
	"strings"
)

func solve() {
	var n, k int
	var s string
	var p []rune
	var j, l rune

	fmt.Scan(&n, &k)
	fmt.Scan(&s)

	p = []rune(strings.Repeat("a", n))
	l = rune(95 + k)
	c := 0

	fmt.Print()

	for {
		if string(p) >= s {
			break
		}
		c++
		for i := n / 2; i > -1; i++ {
			if p[i] > l {
				p[i], p[len(p)-i-1] = 97, 97
			} else {
				j = p[i] + 1
				p[i], p[len(p)-i-1] = j, j
				break
			}
		}
	}
	fmt.Print(c)
}

func main() {
	var t int
	fmt.Scan(&t)
	t++
	for i := 1; i < t; i++ {
		fmt.Print("Case #", i, ": ")
		solve()
		fmt.Println()
	}
}
