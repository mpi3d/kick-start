package main

import (
	"fmt"
	"strings"
)

var t, n int
var s string

func solve() string {
	fmt.Scan(&n, &s)

	c := s + " "
	for c != s {
		c = s
		s = strings.Replace(s, "01", "2", -1)
		s = strings.Replace(s, "12", "3", -1)
		s = strings.Replace(s, "23", "4", -1)
		s = strings.Replace(s, "34", "5", -1)
		s = strings.Replace(s, "45", "6", -1)
		s = strings.Replace(s, "56", "7", -1)
		s = strings.Replace(s, "67", "8", -1)
		s = strings.Replace(s, "78", "9", -1)
		s = strings.Replace(s, "89", "0", -1)
		s = strings.Replace(s, "90", "1", -1)
	}

	return s
}

func main() {
	fmt.Scan(&t)

	for i := 1; i <= t; i++ {
		fmt.Printf("Case #%d: %s\n", i, solve())
	}
}
