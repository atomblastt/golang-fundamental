package main

import "fmt"

func main() {
	var firstName string
	firstName = "Thomas"
	fmt.Println("isi variable firstName = ", firstName)

	var angka int8
	angka = 13
	fmt.Println("isi variable angka = ", angka)

	var lastName = "Barone"
	fmt.Println("isi variable lastName = ", lastName)

	umur := 24
	fmt.Println("isi variable umur = ", umur)

	var (
		firstVar = "First"
		lastVar  = "Last"
	)
	fmt.Println(firstVar)
	fmt.Println(lastVar)

}
