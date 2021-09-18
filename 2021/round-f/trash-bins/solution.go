package main

import (
	"fmt"
)

func solve() {
	n := 0
	fmt.Scan(&n)

	t := make([]bool, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&t[i])
	}

	c := 0
	for i := 0; i < n; i++ {
		if !t[i] {
			r := 0
			for j := i + 1; j < n; j++ {
				if t[j] {
					r = j - i
					break
				}
			}

			l := 0
			for j := i - 1; j >= 0; j-- {
				if t[j] {
					l = i - j
					break
				}
			}

			if r == 0 {
				c += l
			} else if l == 0 {
				c += r
			} else if r > l {
				c += l
			} else {
				c += r
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
