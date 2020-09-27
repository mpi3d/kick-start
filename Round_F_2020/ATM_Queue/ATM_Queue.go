package main

import "fmt"

func solve() {
	var n, x int
	fmt.Scan(&n, &x)
	a, b := make([]int, n), make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
		b[i] = i + 1
	}
	for len(a) > 0 {
		if d := a[0] - x; d <= 0 {
			fmt.Print(" ", b[0])
			a, b = a[1:], b[1:]
		} else {
			a, b = append(a[1:], d), append(b[1:], b[0])
		}

	}
}

func main() {
	var t int
	fmt.Scan(&t)
	t++
	for i := 1; i < t; i++ {
		fmt.Print("Case #", i, ":")
		solve()
		fmt.Println()
	}
}
