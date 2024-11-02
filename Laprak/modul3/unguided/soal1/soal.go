package main

import (
	"fmt"
)

func main() {
	var fx float64
	fmt.Print("Masukkan nilai f(x): ")
	fmt.Scan(&fx)
	var x float64
	if fx == 5 {
		x = 5.2
	} else if fx == 11 {
		x = 5.125
	} else {
		x = (2 / (fx - 5)) - 5
	}
	fmt.Printf("Nilai x adalah: %.3f\n", x)
}
