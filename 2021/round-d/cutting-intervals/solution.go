package main

import (
	"fmt"
	"sort"
)

func solve() {
	var n, c int
	var l, u int
	fmt.Scan(&n, &c)

	d := make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Scan(&l, &u)
		for b := l + 1; b < u; b++ {
			d[b]++
		}
	}

	var b []int
	for _, v := range d {
		b = append(b, v)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(b)))

	if len(b) < c {
		c = len(b)
	}
	for i := 0; i < c; i++ {
		n += b[i]
	}

	fmt.Print(n)
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
