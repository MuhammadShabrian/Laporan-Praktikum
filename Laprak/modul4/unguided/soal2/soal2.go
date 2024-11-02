package main

import "fmt"

func main() {
	var bmi, tinggiBadan, beratBadan float64

	fmt.Print("Masukkan nilai BMI: ")
	fmt.Scan(&bmi)

	fmt.Print("Masukkan tinggi badan (dalam meter): ")
	fmt.Scan(&tinggiBadan)

	beratBadan = bmi * (tinggiBadan * tinggiBadan)

	fmt.Printf("Berat badan seseorang: %.0f kg\n", beratBadan)
}
