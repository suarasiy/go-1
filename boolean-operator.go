package main

import "fmt"

func main() {
	var nilaiLulus = 80
	var absenseiLulus = 13
	var nilai = 75
	var absensi = 12

	if nilai >= nilaiLulus {
		fmt.Println("Nilai lulus")
	} else {
		fmt.Println("Tidak lulus")
	}
	fmt.Println("overal lulus ? ", nilai >= nilaiLulus && absensi >= absenseiLulus)
	fmt.Println("End program.")
}
