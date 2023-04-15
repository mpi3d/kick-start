package main

import (
	"fmt"
)

func solve() string {
	n := 0
	fmt.Scan(&n)
	s := make([]int, n)
	for i := range s {
		fmt.Scan(&s[i])
	}

	r := []int{}
	q := 0
	for _, v := range s {
		if v == q {
			continue
		}
		q = v
		for _, w := range r {
			if w == v {
				return "IMPOSSIBLE"
			}
		}
		r = append(r, v)
	}

	t := fmt.Sprint(r)
	return t[1 : len(t)-1]
}

func main() {
	t := 0
	fmt.Scan(&t)

	for i := 1; i <= t; i++ {
		fmt.Printf("Case #%v: %v\n", i, solve())
	}
}
