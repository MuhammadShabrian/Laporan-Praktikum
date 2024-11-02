package main

import (
	"fmt"
	"math"
)

func main() {
	var radius float64

	// Meminta input dari pengguna
	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scanln(&radius)

	// Menghitung luas lingkaran
	luas := math.Pi * radius * radius

	// Menampilkan hasil
	fmt.Printf("Luas lingkaran dengan jari-jari %.2f adalah %.1f\n", radius, luas)
}
