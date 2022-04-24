package main

import (
	"fmt"
)

func nextPalindrome(n []uint8) []uint8 {
	l := len(n)
	for i := l / 2; i < l; i++ {
		r := l - 1 - i
		if n[i] < 9 {
			n[i], n[r] = n[i]+1, n[r]+1
			return n
		}
		n[i], n[r] = 0, 0
	}

	n = append(n, 1)
	n[0] = 1
	return n
}

func palindromeToInt(n []uint8) int {
	l, r := len(n), 0
	for i := 0; i < l; i++ {
		r = r*10 + int(n[i])
	}

	return r
}

func solve() int {
	a, c := 0, 0
	fmt.Scan(&a)

	n := []uint8{0}
	for {
		n = nextPalindrome(n)
		i := palindromeToInt(n)

		if i > a {
			break
		}

		if a%i == 0 {
			c++
		}
	}

	return c
}

func main() {
	t := 0
	fmt.Scan(&t)

	for i := 1; i <= t; i++ {
		fmt.Printf("Case #%d: %d\n", i, solve())
	}
}
