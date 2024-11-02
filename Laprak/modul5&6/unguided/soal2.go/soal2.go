package main

import (
	"fmt"
	"math"
)

func main() {
	var input1, input2, input3 int
	var result float64

	fmt.Scan(&input1)

	for x := 0; x < input1; x++ {
		fmt.Scan(&input2, &input3)
		result = (1.0 / 3.0) * math.Pi * math.Pow(float64(input2), 2) * float64(input3)
		fmt.Println(result)
	}

}
