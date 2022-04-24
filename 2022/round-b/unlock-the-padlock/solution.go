package main

import (
	"fmt"
	"math"
)

func leftRight(n []int) []int {
	s, r, l := len(n), -1, -1
	for i := s - 1; i >= 0; i-- {
		if n[i] != 0 {
			r = i
			break
		}
	}

	if r == -1 {
		return nil
	}

	for i := 0; i < s; i++ {
		if n[i] != 0 {
			l = i
			break
		}
	}

	return n[l : r+1]
}

func rotation(n []int, d int) int {
	v, r, l := n[0], 0, 0
	if v < d-v {
		l = -v
	} else {
		l = d - v
	}

	v = n[len(n)-1]
	if v < d-v {
		r = -v
	} else {
		r = d - v
	}

	if math.Abs(float64(l)) < math.Abs(float64(r)) {
		return l
	}

	return r
}

func rotate(n []int, d int, r int) []int {
	for i, v := range n {
		n[i] = (v + r) % d
	}

	return n
}

func solve() int {
	n, d := 0, 0
	fmt.Scan(&n, &d)

	v := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&v[i])
	}

	c := 0
	for {
		v = leftRight(v)
		if v == nil {
			return c
		}

		r := rotation(v, d)
		c += int(math.Abs(float64(r)))
		v = rotate(v, d, r)
	}
}

func main() {
	t := 0
	fmt.Scan(&t)

	for i := 1; i <= t; i++ {
		fmt.Printf("Case #%d: %d\n", i, solve())
	}
}
