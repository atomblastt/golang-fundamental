package main

import "fmt"

func main() {
	var names [2]string
	names[0] = "Thomas"
	names[1] = "Barone"
	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])

	var numbers = [4]int{
		10,
		20,
		30,
	}

	fmt.Println(numbers)
	fmt.Println(numbers[0])
	fmt.Println(numbers[1])
	fmt.Println(numbers[2])
	fmt.Println(len(numbers))

}
