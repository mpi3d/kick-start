package main

import (
	"fmt"
	"math"
)

var t, k int
var xMin, yMin, xMax, yMax, bottleXMin, bottleYMin, stepMin, step float64
var object [][4]float64

func solve() string {
	fmt.Scan(&k)

	object = make([][4]float64, k)
	xMin, yMin = math.Inf(1), math.Inf(1)
	xMax, yMax = math.Inf(-1), math.Inf(-1)
	for i := 0; i < k; i++ {
		fmt.Scan(&object[i][0], &object[i][1], &object[i][2], &object[i][3])

		xMin, yMin = math.Min(xMin, object[i][0]), math.Min(yMin, object[i][1])
		xMax, yMax = math.Max(xMax, object[i][2]), math.Max(yMax, object[i][3])
	}

	stepMin = math.Inf(1)
	for x := xMin; x <= xMax; x++ {
		for y := yMin; y <= yMax; y++ {
			step = 0
			for i := 0; i < k; i++ {
				if x < object[i][0] {
					step += object[i][0] - x
				} else if x > object[i][2] {
					step += x - object[i][2]
				}

				if y < object[i][1] {
					step += object[i][1] - y
				} else if y > object[i][3] {
					step += y - object[i][3]
				}
			}

			if step < stepMin {
				stepMin, bottleXMin, bottleYMin = step, x, y
			}
		}
	}

	return fmt.Sprintf("%.0f %.0f", bottleXMin, bottleYMin)
}

func main() {
	fmt.Scan(&t)
	for i := 1; i <= t; i++ {
		fmt.Printf("Case #%d: %s\n", i, solve())
	}
}
