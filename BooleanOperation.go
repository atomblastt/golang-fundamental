package main

import "fmt"

func main() {
	pointA := 90
	pointB := 80

	var resultPointA bool = pointA > 60
	var resultPointb bool = pointB > 50

	var result bool = resultPointA && resultPointb

	fmt.Println("result =", result)
}
