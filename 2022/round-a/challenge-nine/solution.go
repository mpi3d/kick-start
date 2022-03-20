package main

import (
	"fmt"
	"math/big"
)

var (
	One  = big.NewInt(1)
	Nine = big.NewInt(9)
	Ten  = big.NewInt(10)
)

func solve() *big.Int {
	n := new(big.Int)
	fmt.Scan(n)

	c := new(big.Int)
	for i := new(big.Int).Set(n); i.Sign() > 0; i.Div(i, Ten) {
		c.Add(c, One)
	}

	s, t := new(big.Int).Mod(n, Nine), new(big.Int)
	s.Sub(Nine, s)
	c.Sub(c, One)
	if s.Cmp(Nine) == 0 {
		i := new(big.Int).Exp(Ten, c, nil)
		d := new(big.Int).Div(n, i)
		return t.Add(t.Mul(t.Mul(d, Nine), i), n)
	}

	for i := new(big.Int).Exp(Ten, c, nil); i.Sign() > 0; i.Div(i, Ten) {
		t.Mod(t.Div(n, i), Ten)
		if t.Cmp(s) > 0 {
			i.Mul(i, Ten)
			d := new(big.Int).Div(n, i)
			return t.Add(t.Mul(t.Add(t.Mul(d, Nine), s), i), n)
		}
	}

	return t.Add(t.Mul(n, Ten), s)
}

func main() {
	t := 0
	fmt.Scan(&t)

	for i := 1; i <= t; i++ {
		fmt.Printf("Case #%d: %s\n", i, solve())
	}
}
