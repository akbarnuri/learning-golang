package main

import "fmt"

// func main() {
// 	var firstName, lastName string = "Akbar", "Nuc Ichfan"
// 	var Age int = 23

// 	fmt.Println(firstName, lastName, Age)
// }

// func main() {
// 	var name string
// 	name = "akbar"

// 	var age int = 23

// 	fmt.Printf("Halo, perkenalkan namaku %s dan aku berumur %d tahun\n", name, age)
// }

func main() {
	name, age, isSingle := "Akbar", 23, true
	fmt.Println(name, age, isSingle)
}
