package main

import (
	"fmt"
	"strings"
)

func main() {
	var email string
	fmt.Println("Aplikasi Validasi Email")
	fmt.Print("Masukkan Email Anda : ")
	fmt.Scan(&email)
	lenght := len(email)
	// karakter harus lebih dari 5
	if lenght > 5 {
		// di dalam karakter harus ada "@"
		if strings.ContainsAny("@", email) {
			isiEmail := strings.Split(email, "@")
			// karakter sebelum @ harus lebih atau sama dengan 2
			if len(isiEmail[0]) >= 2 {
				if strings.ContainsAny("@", isiEmail[0]) {
					fmt.Println("Format Salah")
				} else {
					// isi email harus berisi karakter
					if isiEmail[1] != "" {
						// setelah "@" harus ada karakter "."
						if strings.ContainsAny(".", isiEmail[1]) {
							afterDot := strings.Split(isiEmail[1], ".")
							//karakter setelah . harus lebih atau sama dengan 2
							if len(afterDot[1]) >= 2 {
								fmt.Println("Success")
							} else {
								fmt.Println("Sorry Failed")
							}
						} else {
							fmt.Println("format email anda salah ex : example@gmail.com")
						}
					} else {
						fmt.Println("Failed")
					}
				}
			} else {
				fmt.Println("karakter setelah @ harus lebih atau sama dengan 2")
			}
		} else {
			fmt.Println("Tidak ada '@' ex : example@gmail.com")
		}
	} else {
		fmt.Println("Karakter Harus Lebih Dari 5")
	}
}
