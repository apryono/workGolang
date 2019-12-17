package main

import "fmt"

func main() {
	// arrAngka := [5][4][3]int{{{}}}
	var arrAngka [5][4][3]int

	for i := 0; i < len(arrAngka); i++ {
		for j := 0; j < len(arrAngka[i]); j++ {
			for k := 0; k < len(arrAngka[i][j]); k++ {
				arrAngka[i][j][k] = (i * 100) + (j * 10) + k
			}
		}
	}
	fmt.Println(arrAngka)
}
