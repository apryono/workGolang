package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var email string
	fmt.Println("Aplikasi Validasi Email")
	fmt.Print("Masukkan Email Anda : ")
	fmt.Scan(&email)
	// karakter harus lebih dari 5

	// di dalam karakter harus ada "@"
	if valChar(email) {
		if strings.ContainsAny("@", email) {
			isiEmail := strings.Split(email, "@")
			// karakter sebelum @ harus lebih atau sama dengan 2
			if len(isiEmail[0]) >= 3 {
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
								// akhiran setelah . domain nya harus sesuai com, co, co.id, org
								var dom = [4]string{"com", "co", "co.id", "org"}
								for i := 0; i < len(dom); i++ {
									if afterDot[1] == dom[i] {
										fmt.Println("Success")
										os.Exit(0)
									}
								}
								fmt.Println("Domain hanya bisa berakhiran com, co, co.id, org ")
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
		fmt.Println("Tidak dapat menggunakan karakter /,!,$,#,&,$,%,*,(,)")
	}
}

func valChar(check string) bool {
	if len(check) > 5 {
		char := [...]string{"/", "#", "!", "&", "%", "$", "*", "(", ")", "+"}
		for j := 0; j < len(char); j++ {
			if strings.ContainsAny(check, char[j]) {
				return false
			}
		}
		return true
	} else {
		return false
	}
}
