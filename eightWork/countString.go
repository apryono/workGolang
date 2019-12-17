package main

import (
	"bufio"
	f "fmt"
	"os"
	s "strings"
)

func main() {
	sym := s.Repeat("=", 50)
	f.Print("Input Huruf x dan o : ")
	word := s.ToLower(getWord())

	checkXO := s.ContainsAny(word, "abcdefghijklmnpqrstuvwyz1234567890")

	if checkXO {
		f.Println("Sorry, Mohon hanya menginput x dan o")
	} else {

		countCX := s.Count(word, "x")
		countCO := s.Count(word, "o")

		if countCO == countCX {
			f.Println(sym)
			f.Println("TRUE, Karna Jumlah x dan o sama")
		} else {
			f.Println(sym)
			f.Println("FALSE, Karna Jumlah x dan o beda")
			f.Println("Ulang Input Lagi Deh ")
			f.Println(sym)
			main()
		}

		f.Println(sym)
		f.Println("Karena Jumlah : ")
		f.Println("X : ", countCX)
		f.Println("O : ", countCO)
		f.Println()
	}

	main()

}

func getWord() string {
	getData := bufio.NewScanner(os.Stdin)
	getData.Scan()
	data := getData.Text()
	return data
}
