package main

import "fmt"

func main() {
	name := "Akbar Nur Ichfan"
	fmt.Println(name)

	name = "Akbar Ichfan"
	fmt.Println(name)

	var (
		firstName  = "Akbar"
		middleName = "Nur"
		lastName   = "Ichfan"
	)

	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}
