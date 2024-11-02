package main

import (
	"fmt"
	"math"
)

func distance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}
func main() {
	var ax, ay, bx, by, cx, cy float64
	fmt.Scan(&ax, &ay)
	fmt.Scan(&bx, &by)
	fmt.Scan(&cx, &cy)
	ab := distance(ax, ay, bx, by)
	bc := distance(bx, by, cx, cy)
	ca := distance(cx, cy, ax, ay)
	longest := math.Max(ab, math.Max(bc, ca))
	fmt.Printf("%.2f\n", longest)
}
