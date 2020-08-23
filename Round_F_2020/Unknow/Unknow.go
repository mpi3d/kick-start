package main

import "fmt"

func solve() {
	var r int
	fmt.Scan(&r)
	fmt.Print(r)
}

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		fmt.Printf("Case #%d: ", i)
		solve()
		fmt.Println()
	}
}
