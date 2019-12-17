package main

import "fmt"

func main() {

	var input int
	fmt.Print("Input Nilai Bintang = ")
	fmt.Scan(&input)
	for i := 0; i < input; i++ {
		if i == 0 {
			for i := 0; i <= input; i++ {
				for j := i; j >= 1; j-- {
					fmt.Print("*")
				}
				fmt.Println("*")
			}

			for i := input; i >= 1; i-- {
				for j := i; j >= 1; j-- {
					fmt.Print("*")
				}
				fmt.Println()
			}
		}
		return
	}
}
