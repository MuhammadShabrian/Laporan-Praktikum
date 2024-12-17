package main

import "fmt"

func main() {
	var b int
	var jumlahFaktor int
	fmt.Scan(&b)
	if b <= 0 {
	} else {
		fmt.Printf("Bilangan: %d\n", b)
		fmt.Print("Faktor: ")
		for f := 1; f <= b; f++ {
			if b%f == 0 {
				fmt.Print(f, " ")
				jumlahFaktor++
			}
		}
		fmt.Println()
		if jumlahFaktor == 2 {
			fmt.Println("Prima: True")
		} else {
			fmt.Println("Prima: False")
		}
	}
}
