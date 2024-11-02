package main

import "fmt"

func main() {
	var totalBelanja, diskon int

	fmt.Print("Masukkan total belanja awal: ")
	fmt.Scan(&totalBelanja)

	fmt.Print("Masukkan besaran diskon (dalam persen): ")
	fmt.Scan(&diskon)

	totalSetelahDiskon := totalBelanja - (totalBelanja * diskon / 100)

	fmt.Printf("Total belanja setelah diskon: %d\n", totalSetelahDiskon)
}
