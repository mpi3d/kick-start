package main

import (
	"fmt"
)

func solve() int {
	n := 0
	fmt.Scan(&n)

	return (n-1)/5 + 1
}

func main() {
	t := 0
	fmt.Scan(&t)

	for i := 1; i <= t; i++ {
		fmt.Printf("Case #%v: %v\n", i, solve())
	}
}
