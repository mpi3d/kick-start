package main

import "fmt"

func solve() {
	var r, c int
	fmt.Scan(&r, &c)
	g, l := make([]int, r*c), 0
	for i := 0; i < r*c; i++ {
		fmt.Scan(&g[i])
	}
	for row := 0; row < r; row++ {
		for column := 0; column < c; column++ {
			if g[row*c+column] == 1 {
				up, right, down, left := 0, 0, 0, 0
				for rowUp := row; rowUp >= 0; rowUp-- {
					if g[rowUp*c+column] == 1 {
						up++
					} else {
						break
					}
				}
				for columnRight := column; columnRight < c; columnRight++ {
					if g[row*c+columnRight] == 1 {
						right++
					} else {
						break
					}
				}
				for rowDown := row; rowDown < r; rowDown++ {
					if g[rowDown*c+column] == 1 {
						down++
					} else {
						break
					}
				}
				for columnLeft := column; columnLeft >= 0; columnLeft-- {
					if g[row*c+columnLeft] == 1 {
						left++
					} else {
						break
					}
				}

				for rowUp := up; rowUp > 1; rowUp-- {
					for columnLeft := left; columnLeft > 1; columnLeft-- {
						if rowUp < columnLeft {
							if rowUp*2 == columnLeft {
								l++
							}
						} else {
							if columnLeft*2 == rowUp {
								l++
							}
						}
					}
					for columnRight := right; columnRight > 1; columnRight-- {
						if rowUp < columnRight {
							if rowUp*2 == columnRight {
								l++
							}
						} else {
							if columnRight*2 == rowUp {
								l++
							}
						}
					}
				}
				for rowDown := down; rowDown > 1; rowDown-- {
					for columnLeft := left; columnLeft > 1; columnLeft-- {
						if rowDown < columnLeft {
							if rowDown*2 == columnLeft {
								l++
							}
						} else {
							if columnLeft*2 == rowDown {
								l++
							}
						}
					}
					for columnRight := right; columnRight > 1; columnRight-- {
						if rowDown < columnRight {
							if rowDown*2 == columnRight {
								l++
							}
						} else {
							if columnRight*2 == rowDown {
								l++
							}
						}
					}
				}
			}
		}
	}
	fmt.Println(l)
}

func main() {
	var t int
	fmt.Scan(&t)
	for i := 1; i <= t; i++ {
		fmt.Print("Case #", i, ": ")
		solve()
	}
}
