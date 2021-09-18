package main

import (
	"fmt"
)

func solve() {
	var d, n, k, h, s, e int

	fmt.Scan(&d, &n, &k)

	a := make([][3]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&h, &s, &e)
		a[i] = [3]int{h, s, e}
	}

	m := 0
	d++
	for i := 1; i < d; i++ {
		t := make([]int, k)
		for _, x := range a {
			if x[1] <= i && x[2] >= i {
				o := 0
				for l, y := range t {
					if y < t[o] {
						o = l
					}
				}
				if t[o] < x[0] {
					t[o] = x[0]
				}
			}
		}
		v := 0
		for _, x := range t {
			v += x
		}
		if v > m {
			m = v
		}
	}

	fmt.Print(m)
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
