package main

import "fmt"

func main() {
	var value32 int32 = 129
	var value64 int64 = int64(value32)
	var value8 int8 = int8(value32)

	name := "Thomas"
	var t byte = name[0]
	var tString = string(t)

	fmt.Println(value32)
	fmt.Println(value64)
	fmt.Println(value8)
	fmt.Println(tString)
}
