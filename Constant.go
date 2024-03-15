package main

import "fmt"

func main() {
	const realName = "Thomas Barone"
	const realAge int8 = 28

	const (
		firstName = "Thomas"
		age       = 28
	)

	fmt.Println(realName)
	fmt.Println(realAge)

	fmt.Println("nama ", firstName, " age ", age)
}
