package main

import (
	"fmt"
	"math"
)

var t int
var s, f string

func difference(a byte, b byte) float64 {
	min := math.Min(float64(a), float64(b))
	max := math.Max(float64(a), float64(b))
	return math.Min(max-min, 'z'-max+min-'a'+1)
}

func solve() int {
	fmt.Scan(&s, &f)

	c := 0.0
	for l := range s {
		m := math.Inf(1)
		for n := range f {
			d := difference(s[l], f[n])
			if d < m {
				m = d
			}
		}
		c += m
	}

	return int(c)
}

func main() {
	fmt.Scan(&t)

	for i := 1; i <= t; i++ {
		fmt.Printf("Case #%d: %d\n", i, solve())
	}
}
