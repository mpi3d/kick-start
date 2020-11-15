package main

import "fmt"

func solve() {
	var n, k, s int
	fmt.Scan(&n, &k, &s)
	if 2*(k-s)+n < n+k {
		fmt.Print(2*(k-s) + n)
	} else {
		fmt.Print(n + k)
	}

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
