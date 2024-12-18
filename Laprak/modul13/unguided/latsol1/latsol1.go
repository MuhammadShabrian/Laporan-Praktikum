package main

import "fmt"

func main() {
	var number, digitCount int
	fmt.Scan(&number)
	for number > 0 {
		digitCount++
		number = number / 10
	}
	fmt.Print(digitCount)
}
