package main

import (
	"fmt"
)

func main() {
	var input1, input2 int
	var result int

	fmt.Scan(&input1, &input2)

	for x := 0; x < input2; x++ {
		if x == 0 {
			result = input1
		} else {
			result = result * input1
		}
	}
	fmt.Println(result)

}
