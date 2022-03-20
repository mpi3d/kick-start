package main

import (
	"fmt"
)

func solve() string {
	i, p := "", ""
	fmt.Scan(&i, &p)

	iLen := len(i)
	pLen := len(p)
	if pLen < iLen {
		return "IMPOSSIBLE"
	}

	count, pInc := 0, 0
	for iInc := 0; iInc < iLen; iInc++ {
		letter := i[iInc]
		for p[pInc] != letter {
			pLen--
			if pLen < iLen {
				return "IMPOSSIBLE"
			}
			pInc++
			count++
		}
		pInc++
	}

	count += pLen - iLen

	return fmt.Sprintf("%d", count)
}

func main() {
	t := 0
	fmt.Scan(&t)

	for i := 1; i <= t; i++ {
		fmt.Printf("Case #%d: %s\n", i, solve())
	}
}
