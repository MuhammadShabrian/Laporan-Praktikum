package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	hasilBagi := 0
	for x >= y {
		x -= y
		hasilBagi++
	}
	fmt.Println(hasilBagi)
}
