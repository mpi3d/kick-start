package main

import (
	"fmt"
	"math/big"
)

func solve() {
	var z, pp, np, p, r int64
	fmt.Scan(&z)

	pp, np = 2, 2
	p, r = 0, 0
	for {
		pp = np
		for {
			np++
			if big.NewInt(np).ProbablyPrime(0) {
				break
			}
		}
		r = pp * np
		if r > z {
			fmt.Print(p)
			break
		}
		p = r
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
