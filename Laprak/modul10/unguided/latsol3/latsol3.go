package main

import "fmt"

func main() {
	var b int
	fmt.Scan(&b)
	if b <= 1 {
	} else {
		fmt.Printf("Bilangan: %d\n", b)
		fmt.Print("Faktor: ")
		for f := 1; f <= b; f++ {
			if b%f == 0 {
				fmt.Print(f, " ")
			}
		}
		fmt.Println()
	}
}
