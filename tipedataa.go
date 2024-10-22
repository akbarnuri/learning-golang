package main

import "fmt"

// func main() {
// 	var teks string = "ini adalah sebuah teks"
// 	fmt.Println(teks)
// }

// func main() {
// 	var teks string = `
// 	Ini adalah "teks".
// 	Dan ini dari 'baris baru'.
// 	Okeee...
// 	`
// 	fmt.Println(teks)
// }

func main() {
	var contohDecimal float32 = 1.66
	fmt.Printf("%f\n", contohDecimal)
	fmt.Printf("%.1f\n", contohDecimal)
	fmt.Printf("%.3f\n", contohDecimal)
}
