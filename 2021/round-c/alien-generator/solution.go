package main

import (
	"fmt"
)

func solve() {
	var g, c, m, i, k int
	fmt.Scan(&g)
	c, k = 0, 0
	for {
		i, m = 1, 0
		for {
			m += k + i
			if m == g {
				c++
				break
			} else if m > g {
				break
			}
			i++
		}
		k++
		if k == g {
			break
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
