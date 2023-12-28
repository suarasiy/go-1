package main

import "fmt"

func main() {
	var nilai32 int32 = 129
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai64)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	var name = "Sidiq"
	var s byte = name[0]
	var firstNameLetter string = string(s)

	fmt.Println(name)
	fmt.Println(firstNameLetter)
}
