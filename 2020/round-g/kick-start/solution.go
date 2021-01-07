package main

import (
	"fmt"
	"strings"
)

func solve() {
	var text string
	var i int
	kickIndexes, startIndexes := make([]int, 0), make([]int, 0)
	fmt.Scan(&text)
	subtext := text
	c := 0
	for {
		i = strings.Index(subtext, "KICK")
		if i == -1 {
			break
		}
		subtext = subtext[i+1:]
		kickIndexes = append(kickIndexes, i+c)
		c = c + i + 1
	}
	subtext = text
	c = 0
	for {
		i = strings.Index(subtext, "START")
		if i == -1 {
			break
		}
		subtext = subtext[i+1:]
		startIndexes = append(startIndexes, i+c)
		c = c + i + 1
	}
	g := 0
	for _, kickIndex := range kickIndexes {
		for _, startIndex := range startIndexes {
			if kickIndex < startIndex {
				g++
			}
		}
	}
	fmt.Print(g)
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
