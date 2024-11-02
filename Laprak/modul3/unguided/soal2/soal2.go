package main

import (
	"fmt"
	"math"
)

func main() {
	var r int

	fmt.Print("Masukkan jari-jari bola (bilangan bulat): ")
	_, err := fmt.Scanf("%d", &r)
	if err != nil || r <= 0 {
		fmt.Println("Error: Masukan tidak valid. Pastikan jari-jari adalah bilangan bulat positif.")
		return
	}

	volume := (4.0 / 3.0) * math.Pi * math.Pow(float64(r), 3)
	luas := 4 * math.Pi * math.Pow(float64(r), 2)

	fmt.Printf("Volume bola: %.4f\n", volume)
	fmt.Printf("Luas kulit bola: %.4f\n", luas)
}
