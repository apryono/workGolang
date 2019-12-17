package main

import "fmt"

func main() {

	var getNumber int
	fmt.Print("Masukkan Nilai Bintang = ")
	fmt.Scan(&getNumber)

	for i := 0; i <= getNumber; i++ {
		if i <= getNumber/2 {
			for j := i; j >= 1; j-- {
				fmt.Print("*")
			}
		} else if i > getNumber/2 {
			for k := getNumber; k >= i; k-- {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
}
