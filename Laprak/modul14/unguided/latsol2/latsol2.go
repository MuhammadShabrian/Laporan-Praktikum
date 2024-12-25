package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n <= 1 {
		fmt.Println("bukan prima")
	} else {
		isPrime := true
		for i := 2; i < n; i++ {
			if n%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Println("prima")
		} else {
			fmt.Println("bukan prima")
		}
	}
}
