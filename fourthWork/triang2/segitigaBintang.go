package main

import "fmt"

func main() {

	var input int
	fmt.Print("Masukkan Jumlah = ")
	fmt.Scan(&input)

	for i := 0; i <= input; i++ {
		for j := i; j >= 1; j-- {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
