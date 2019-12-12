package main

import "fmt"

func main() {

	var jlh int
	fmt.Print("Masukkan Jumlah Nilai = ")
	fmt.Scan(&jlh)

	for i := jlh; i >= 1; i-- {
		for j := jlh; j >= 1; j-- {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
