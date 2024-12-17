package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)
	for number > 0 {
		digit := number % 10
		fmt.Println(digit)
		number = number / 10
	}
}
