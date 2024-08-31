package main

import "fmt"

func main() {
	const (
		firstName string = "Akbar"
		lastName         = "Ichfan"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

	//error
	// firstName = "wakwao"
	// lastName = "jkaa"
}
