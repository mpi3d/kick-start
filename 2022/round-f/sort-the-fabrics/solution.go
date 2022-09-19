package main

import (
	"fmt"
	"sort"
	"strings"
)

type Color struct {
	color string
	id    int
}

type Colors []Color

func (cs Colors) Len() int {
	return len(cs)
}

func (cs Colors) Less(i, j int) bool {
	ci, cj := cs[i], cs[j]
	switch strings.Compare(ci.color, cj.color) {
	case -1:
		return true
	case 1:
		return false
	default:
		return ci.id < cj.id
	}
}

func (cs Colors) Swap(i, j int) {
	cs[i], cs[j] = cs[j], cs[i]
}

type Durability struct {
	durability int
	id         int
}

type Durabilities []Durability

func (ds Durabilities) Len() int {
	return len(ds)
}

func (ds Durabilities) Less(i, j int) bool {
	di, dj := ds[i], ds[j]
	switch i, j := di.durability, dj.durability; {
	case i < j:
		return true
	case i > j:
		return false
	default:
		return di.id < dj.id
	}
}

func (ds Durabilities) Swap(i, j int) {
	ds[i], ds[j] = ds[j], ds[i]
}

func solve() int {
	n := 0
	fmt.Scan(&n)

	cs := make(Colors, n)
	ds := make(Durabilities, n)
	for i := 0; i < n; i++ {
		c, d, u := "", 0, 0
		fmt.Scan(&c, &d, &u)
		cs[i] = Color{c, u}
		ds[i] = Durability{d, u}
	}

	sort.Sort(cs)
	sort.Sort(ds)

	c := 0
	for i := 0; i < n; i++ {
		if cs[i].id == ds[i].id {
			c++
		}
	}

	return c
}

func main() {
	t := 0
	fmt.Scan(&t)

	for i := 1; i <= t; i++ {
		fmt.Printf("Case #%v: %v\n", i, solve())
	}
}
