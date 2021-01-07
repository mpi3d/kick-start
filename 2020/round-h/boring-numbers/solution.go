package main

import (
	"fmt"
	"strconv"
	"strings"
)

func solve() {
	var c, l, r int
	c = 0
	fmt.Scan(&l, &r)
	for n := l; n <= r; n++ {
		boring := true
		for i, v := range strings.Split(strconv.Itoa(n), "") {
			d, _ := strconv.Atoi(v)
			if (i%2 != 0) != (d%2 == 0) {
				boring = false
				break
			}
		}
		if boring {
			c++
		}
	}
	fmt.Print(c)
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
