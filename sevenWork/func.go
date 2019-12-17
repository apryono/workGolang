package main

import "fmt"

func main() {
	fmt.Println(isOdd(11))
}

func isOdd(angka int) (ganjil bool) {
	if angka <= 0 {
		ganjil = false
	} else {
		if angka%2 == 1 {
			ganjil = true
		} else {
			ganjil = false
		}
	}
	return
}
