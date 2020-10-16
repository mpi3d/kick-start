package main

import "fmt"

func solve() {
	var n int
	fmt.Scan(&n)
	fmt.Print(n)
}

func main() {
	var t int
	fmt.Scan(&t)
	t++
	for i := 1; i < t; i++ {
		fmt.Print("Case #", i, ":")
		solve()
		fmt.Println()
	}
}
