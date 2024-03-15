package main

import "fmt"

func main() {
	type identityNumber string

	var noKtp identityNumber = "1234567890"

	fmt.Println(noKtp)
}
