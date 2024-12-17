package main

import "fmt"

func main() {
	var tipe string
	var jam, tarif int
	fmt.Scan(&tipe, &jam)

	switch {
	case tipe == "motor" && jam == 1 && jam < 2:
		tarif = 2000
	case tipe == "motor" && jam >= 2:
		tarif = jam * 2000
	case tipe == "mobil" && jam == 1 && jam < 2:
		tarif = 5000
	case tipe == "mobil" && jam >= 2:
		tarif = jam * 5000
	case tipe == "truk" && jam == 1 && jam <= 2:
		tarif = 8000
	case tipe == "truk" && jam >= 2:
		tarif = jam * 8000
	}
	fmt.Print("Rp ", tarif)
}
