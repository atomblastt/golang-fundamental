package main

import "fmt"

func main() {
	a := 10
	b := 20

	tambah := a + b
	kurang := b - a
	kali := a * a
	bagi := b / a

	fmt.Println("tambah = ", tambah)
	fmt.Println("kurang = ", kurang)
	fmt.Println("kali = ", kali)
	fmt.Println("bagi = ", bagi)

	var i = 100
	i += b
	fmt.Println("argument singkat tambah = ", i)

	i++
	fmt.Println("unary operator tambah = ", i)
}
