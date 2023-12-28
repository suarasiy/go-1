package main

import "fmt"

func main() {
	type NoKTP string
	type test bool

	var noKtp NoKTP = "1"
	var something test = true

	fmt.Println(noKtp)
	fmt.Println(something)
}
