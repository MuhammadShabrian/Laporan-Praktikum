package main

import "fmt"

func main() {
	var beratParsel int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&beratParsel)

	kg := beratParsel / 1000
	gramSisa := beratParsel % 1000

	biayaKg := kg * 10000

	var biayaSisa int
	if kg > 10 {
		biayaSisa = 0
	} else {
		if gramSisa >= 500 {
			biayaSisa = gramSisa * 5
		} else {
			biayaSisa = gramSisa * 15
		}

		totalBiaya := biayaKg + biayaSisa
		fmt.Printf("Detail berat: %d kg + %d gr\n", kg, gramSisa)
		fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKg, biayaSisa)
		fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
	}
}
