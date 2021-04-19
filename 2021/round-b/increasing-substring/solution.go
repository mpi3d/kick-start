package main

import "fmt"

func solve() {
	var n int
	var s string
	var r []int
	fmt.Scan(&n, &s)

	c := 0
	b := rune(0)
	for _, l := range s {
		if b < l {
			c++
		} else {
			c = 1
		}
		b = l
		r = append(r, c)
	}
	for _, value := range r {
		fmt.Printf(" %d", value)
	}
}

func main() {
	var t int
	fmt.Scan(&t)
	for i := 1; i <= t; i++ {
		fmt.Print("Case #", i, ": ")
		solve()
		fmt.Println()
	}
}
