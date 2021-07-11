package main

import (
	"fmt"
)

func solve() {
	var g [9]int

	fmt.Scan(&g[0], &g[1], &g[2])
	fmt.Scan(&g[3], &g[5])
	fmt.Scan(&g[6], &g[7], &g[8])

	d := make(map[int]int)
	if g[0]%2 == g[8]%2 {
		d[g[0]+(g[8]-g[0])/2]++
	}
	if g[1]%2 == g[7]%2 {
		d[g[1]+(g[7]-g[1])/2]++
	}
	if g[2]%2 == g[6]%2 {
		d[g[2]+(g[6]-g[2])/2]++
	}
	if g[3]%2 == g[5]%2 {
		d[g[3]+(g[5]-g[3])/2]++
	}

	m := 0
	for n, c := range d {
		if c > m {
			g[4] = n
			m = c
		}
	}

	c := 0
	if 2*g[1]-g[0] == g[2] {
		c++
	}
	if 2*g[4]-g[3] == g[5] {
		c++
	}
	if 2*g[7]-g[6] == g[8] {
		c++
	}
	if 2*g[3]-g[0] == g[6] {
		c++
	}
	if 2*g[4]-g[1] == g[7] {
		c++
	}
	if 2*g[5]-g[2] == g[8] {
		c++
	}
	if 2*g[4]-g[0] == g[8] {
		c++
	}
	if 2*g[4]-g[2] == g[6] {
		c++
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
