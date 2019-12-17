package main

import "fmt"

func main() {
	var input int
	var sym string

	isGanjil := func(angka int) (valid bool) {
		if angka <= 0 {
			valid = false
		} else {
			for i := 0; i < angka*2; i++ {
				if i%2 == 1 {
					// sym = ""
					if i < angka*2 {
						sym = " "
					}
					fmt.Print(i, sym, " ")
				}
			}
			valid = true
		}
		return
	}

	fmt.Print("Input Nilai : ")
	fmt.Scanln(&input)

	getInput := isGanjil(input)
	isValid := getInput
	if !isValid {
		fmt.Println("Angka Tidak Valid")
	}

}
