package main

import (
	"fmt"
)

func solve() {
	var n int
	fmt.Scan(&n)
	l := n * n
	coins := make([]int, l)
	for i := 0; i < l; i++ {
		fmt.Scan(&coins[i])
	}
	l = n*2 - 1
	sums := make([]int, l)
	c := 1
	for i := 0; i < l; i++ {
		sums[i] = 0
		for d := 0; d < c; d++ {
			if i < n {
				sums[i] += coins[(n-i-1)*n+d*n+d]
			} else {
				sums[i] += coins[(i-n+1)+d*n+d]
			}
		}
		if i < n-1 {
			c++
		} else {
			c--
		}
	}
	m := 0
	for _, v := range sums {
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
