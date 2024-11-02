package main

import (
	"fmt"
)

func main() {
	var fahrenheit float64

	// Meminta input dari pengguna
	fmt.Print("Masukkan suhu dalam Fahrenheit: ")
	fmt.Scanln(&fahrenheit)

	// Mengonversi Fahrenheit ke Celsius
	celsius := (fahrenheit - 32) * 5 / 9

	// Menampilkan hasil konversi
	fmt.Printf("Suhu dalam Celsius adalah: %.2f\n", celsius)
}
