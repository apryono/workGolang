package main

import "fmt"

func main() {

	var input int
	fmt.Print("Input Nilai = ")
	fmt.Scan(&input)

	for i := input; i >= 1; i-- {
		for j := i; j >= 1; j-- {
			fmt.Print(j)
		}
		fmt.Println()
	}

}
