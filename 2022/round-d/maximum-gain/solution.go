package main

import (
	"fmt"
)

func array(l int) []int {
	s := make([]int, l)
	for i := 0; i < l; i++ {
		fmt.Scan(&s[i])
	}
	return s
}

func reverse(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func min(a []int) int {
	m := a[0]
	for _, v := range a[1:] {
		if v < m {
			m = v
		}
	}
	return m
}

func sum(a []int) int {
	s := 0
	for _, v := range a {
		s += v
	}
	return s
}

func solve() int {
	n := 0
	fmt.Scan(&n)
	a := array(n)

	m := 0
	fmt.Scan(&m)
	b := array(m)

	am := 0
	ali := n / 2
	ari := ali
	if n%2 == 1 {
		am = a[ali]
		ari++
	}

	bm := 0
	bli := m / 2
	bri := bli
	if m%2 == 1 {
		bm = b[bli]
		bri++
	}

	al, ar, bl, br := a[:ali], a[ari:], b[:bli], b[bri:]
	ali, ari, bli, bri = len(al)-1, len(ar)-1, len(bl)-1, len(br)-1

	reverse(ar)
	reverse(br)

	k := 0
	fmt.Scan(&k)

	l := n + m
	for l > k {
		c := []int{}
		if am > 0 {
			c = append(c, am)
		}
		if bm > 0 {
			c = append(c, bm)
		}
		if ali >= 0 {
			c = append(c, al[ali])
		}
		if ari >= 0 {
			c = append(c, ar[ari])
		}
		if bli >= 0 {
			c = append(c, bl[bli])
		}
		if bri >= 0 {
			c = append(c, br[bri])
		}

		switch m := min(c); {
		case am == m:
			am = 0
		case bm == m:
			bm = 0
		case ali >= 0 && al[ali] == m:
			ali--
		case ari >= 0 && ar[ari] == m:
			ari--
		case bli >= 0 && bl[bli] == m:
			bli--
		case bri >= 0 && br[bri] == m:
			bri--
		}

		l--
	}

	return sum(al[:ali+1]) + sum(ar[:ari+1]) +
		sum(bl[:bli+1]) + sum(br[:bri+1]) + am + bm
}

func main() {
	t := 0
	fmt.Scan(&t)

	for i := 1; i <= t; i++ {
		fmt.Printf("Case #%d: %d\n", i, solve())
	}
}
