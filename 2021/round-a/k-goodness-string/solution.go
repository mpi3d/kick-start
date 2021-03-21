package main

import "fmt"

func solve() {
	var n, k, c int
	var s string
	fmt.Scan(&n, &k, &s)
	c = 0
	for i := 1; i <= n/2; i++ {
		if s[i-1] != s[n-i] {
			c++
		}
	}
	if c > k {
		fmt.Println(c - k)
	} else {
		fmt.Println(k - c)
	}
}

func main() {
	var t int
	fmt.Scan(&t)
	for i := 1; i <= t; i++ {
		fmt.Print("Case #", i, ": ")
		solve()
	}
}
