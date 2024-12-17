package main

import "fmt"

func main() {
	var username, password string
	var Perobaangagal int
	fmt.Scan(&username, &password)
	for username != "Admin" || password != "Admin" {
		Perobaangagal++
		fmt.Scan(&username, &password)
	}
	fmt.Println(Perobaangagal, "Percobaan gagal login")
}
