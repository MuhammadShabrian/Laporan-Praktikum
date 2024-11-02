package main

import "fmt"

func main() {
	var input1 int
	var result int = 1

	fmt.Scan(&input1)

	for x := 1; x <= input1; x++ {
		result = result * x
	}
	fmt.Println(result)

}
