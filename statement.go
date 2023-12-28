package main

import "fmt"

func main() {
	var colorBanana string = "Yellow"
	var colorPaprika string = "Green"
	var yellow string = "Yellow"

	var bananaIsYellow bool = colorBanana == yellow
	var paprikaIsyellow bool = colorPaprika == yellow

	fmt.Println("Banana is yellow? ", bananaIsYellow)
	fmt.Println("Paprika is yellow? ", paprikaIsyellow)

	const (
		height int = 100
		width  int = 50
	)

	fmt.Println(height < width)
	fmt.Println(height > width)
	fmt.Println(height == width)
	fmt.Println(height != width)
}
