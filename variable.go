package main

import "fmt"

func main() {
	var name string

	name = "Sidiq Indrajati Yusuf"
	fmt.Println(name)

	name = "Constance Suara"
	fmt.Println(name)

	var age int8 = 23
	fmt.Println(age)

	country := "NYC"
	fmt.Println(country)

	country = "Japan"
	fmt.Println(country)

	var (
		lastName  = "Yusuf"
		firstName = "Sidiq"
	)
	fmt.Println(lastName, ',', firstName)
}
