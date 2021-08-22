package main

import (
	"fmt"
	"sort"
)

type array struct {
	char  rune
	count int
}

type byCount []array

func (a byCount) Len() int {
	return len(a)
}
func (a byCount) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a byCount) Less(i, j int) bool {
	return a[i].count > a[j].count
}

func solve() {
	var s string

	fmt.Scan(&s)

	if len(s) < 2 {
		fmt.Print("IMPOSSIBLE")
		return
	}

	r, t, u := []rune(s), []rune(s), []rune(s)
	c := make(map[rune]int)

	for _, v := range r {
		c[v]++
	}

	f := make([]array, 0, len(c))
	for k, v := range c {
		f = append(f, array{k, v})
	}

	sort.Sort(byCount(f))

	if f[0].count > len(s)/2 {
		fmt.Print("IMPOSSIBLE")
		return
	}

	for len(f) > 0 {
		if len(f) == 1 {
			i0, i1 := 0, 0
			for i0 = range r {
				if f[0].char == r[i0] {
					break
				}
			}
			for i1 = range t {
				if f[0].char != t[i1] && f[0].char != u[i1] {
					break
				}
			}

			t[i0], t[i1] = t[i1], t[i0]
			break
		}

		i0, i1 := 0, 0
		for i0 = range r {
			if f[0].char == r[i0] {
				break
			}
		}
		for i1 = range r {
			if f[1].char == r[i1] {
				break
			}
		}

		t[i0], t[i1] = t[i1], t[i0]
		r[i0], r[i1] = 0, 0

		f[0].count--
		f[1].count--

		if f[0].count == 0 {
			if f[1].count == 0 {
				f = f[2:]
			} else {
				f = f[1:]
			}
		} else if f[1].count == 0 {
			f[1] = f[0]
			f = f[1:]
		}

		sort.Sort(byCount(f))
	}

	fmt.Print(string(t))
}

func main() {
	var t int
	fmt.Scan(&t)
	t++
	for i := 1; i < t; i++ {
		fmt.Print("Case #", i, ": ")
		solve()
		fmt.Println()
	}
}
