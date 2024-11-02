package main

import "fmt"

func main() {
	var tahun int
	var status bool
	fmt.Print("masukan tahun: ")
	fmt.Scan(&tahun)

	status = (tahun%4 == 0 && tahun%100 != 0) || tahun%400 == 0
	fmt.Printf("%t", status)
}
