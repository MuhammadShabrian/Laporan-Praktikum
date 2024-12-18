package main

import (
	"fmt"
)

func main() {
	var target, donasi, totalDonasi, jumlahDonatur int
	fmt.Scan(&target)
	for totalDonasi < target {
		fmt.Scan(&donasi)
		jumlahDonatur++
		totalDonasi += donasi
		fmt.Printf("Donatur %d: Menyumbang %d. Total terkumpul: %d\n", jumlahDonatur, donasi, totalDonasi)
	}
	fmt.Printf("Target tercapai! Total donasi: %d dari %d donatur.\n", totalDonasi, jumlahDonatur)
}
