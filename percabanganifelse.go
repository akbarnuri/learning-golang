package main

import (
	"fmt"
)

func main() {
	var num int

	// Meminta input dari pengguna
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&num)

	// Mengecek apakah angka genap atau ganjil
	if num%2 == 0 {
		fmt.Println(num, "=> genap")
	} else {
		fmt.Println(num, "=> ganjil")
	}
}
