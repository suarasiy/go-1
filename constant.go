package main

import "fmt"

func main() {
	const id int32 = 1
	const (
		firstName string = "Sidiq Indrajati"
		lastName         = "Yusuf"
	)
	var age = 23

	fmt.Println(id)
	fmt.Println(firstName, lastName)
	fmt.Println(age)
}
