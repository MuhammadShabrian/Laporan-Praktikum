package main

import "fmt"

func main() {
	var n, hitung int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		if i%2 != 0 {
			hitung++
		}
	}
	fmt.Println("terdapat", hitung, "bilangan ganjil")
}
