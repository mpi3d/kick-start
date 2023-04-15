package main

import (
	"fmt"
)

type digit struct {
	c bool
	n [10]*digit
}

func solve() string {
	d := [26]int{}
	for i := range d {
		fmt.Scan(&d[i])
	}
	n := 0
	fmt.Scan(&n)
	s := make([][]byte, n)
	for i := range s {
		fmt.Scan(&s[i])
	}

	p := new(digit)
	for _, s := range s {
		q := p
		for _, c := range s {
			i := d[c-'A']
			if q.n[i] == nil {
				q.n[i] = new(digit)
			}
			q = q.n[i]
		}
		if q.c {
			return "YES"
		}
		q.c = true
	}
	return "NO"
}

func main() {
	t := 0
	fmt.Scan(&t)

	for i := 1; i <= t; i++ {
		fmt.Printf("Case #%v: %v\n", i, solve())
	}
}
