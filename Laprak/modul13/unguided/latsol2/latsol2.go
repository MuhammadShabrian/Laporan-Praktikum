package main

import "fmt"

func main() {
	var number, jumlah float64
	fmt.Scan(&number)
	jumlah = number + 0.1
	for jumlah <= float64(int(number)+1) {
		fmt.Printf("%.1f\n", jumlah)
		jumlah += 0.1
	}
}
