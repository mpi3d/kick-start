package main

import (
	"fmt"
	"unicode"
)

var v = []rune{'a', 'e', 'i', 'o', 'u'}

func solve() (string, string) {
	var k string
	fmt.Scan(&k)

	l := unicode.ToLower(rune(k[len(k)-1]))
	if l == 'y' {
		return k, "nobody"
	}

	for _, w := range v {
		if l == w {
			return k, "Alice"
		}
	}

	return k, "Bob"
}

func main() {
	var t int
	fmt.Scan(&t)

	for i := 1; i <= t; i++ {
		k, o := solve()
		fmt.Printf("Case #%d: %s is ruled by %s.\n", i, k, o)
	}
}
