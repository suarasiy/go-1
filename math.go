package main

import "fmt"

func main() {
	var a int = 10
	var b int = 10
	var c int = a + b

	fmt.Println("hasil a + b : ", c)

	c -= b + 2
	fmt.Println("c -= b + 2 :", c)

	a++
	fmt.Println("a will be 11. output :", a)

	var positiveNumber int = +50
	var negativeNumber int = -50

	fmt.Println(positiveNumber)
	fmt.Println(negativeNumber)
}
