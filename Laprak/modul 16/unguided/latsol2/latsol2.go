package main

import "fmt"

func main() {
	var x string
	var n int
	fmt.Scan(&x)
	fmt.Scan(&n)
	strings := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&strings[i])
	}

	ditemukan := false
	for _, s := range strings {
		if s == x {
			ditemukan = true
			break
		}
	}
	if ditemukan {
		fmt.Println("string ditemukan.")
	} else {
		fmt.Println("string tidak ditemukan.")
	}

	posisi := -1
	for i, s := range strings {
		if s == x {
			posisi = i + 1
			break
		}
	}
	if posisi != -1 {
		fmt.Printf("string ditemukan pada posisi ke-%d.\n", posisi)
	} else {
		fmt.Println("string tidak ditemukan pada posisi manapun.")
	}

	count := 0
	for _, s := range strings {
		if s == x {
			count++
		}
	}
	fmt.Printf("jumlah string x: %d\n", count)
	if count >= 2 {
		fmt.Println("ada sedikitnya dua string x dalam kumpulan data.")
	} else {
		fmt.Println("tidak ada sedikitnya dua string x dalam kumpulan data.")
	}
}
