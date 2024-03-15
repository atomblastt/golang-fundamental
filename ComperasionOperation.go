package main

import "fmt"

func main() {
	var1 := "Thomas"
	var2 := "Thomas"

	var3 := 100
	var4 := 200

	var result bool = var1 == var2 // true
	var result2 bool = var3 > var4 // false
	fmt.Println(result)
	fmt.Println(result2)
}
