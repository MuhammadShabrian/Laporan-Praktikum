package main

import "fmt"

func main() {
	var beratbadan, tinggibadan, bmi float64
	fmt.Print("Masukkan Berat Badan (kg): ")
	fmt.Scan(&beratbadan)
	fmt.Print("Masukkan Tinggi Badan (m): ")
	fmt.Scan(&tinggibadan)
	bmi = beratbadan / (tinggibadan * tinggibadan)
	fmt.Printf("BMI anda: %.2f", bmi)
}
