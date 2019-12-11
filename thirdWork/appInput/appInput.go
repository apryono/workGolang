package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(`
	-------------------------------
	    Aplikasi Validasi Form 
	-------------------------------
	`)

	var nama string
	check := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan Nama : ")
	check.Scan()
	nama = check.Text()
	// fmt.Scan(&nama)

	if len(nama) >= 5 && len(nama) <= 20 {
		var year int
		fmt.Print("Masukkan Tahun Lahir : ")
		fmt.Scan(&year)
		if year > 1900 {
			tahun := 2019 - year
			if tahun > 0 {
				if tahun > 17 {
					var gender string
					fmt.Print("Masukkan Jenis Kelamin : ")
					fmt.Scan(&gender)
					gens := strings.ToUpper(gender)
					if gens == "PRIA" || gens == "WANITA" {
						fmt.Println("Thank For Your Informations")
					} else {
						fmt.Println("Kelamin Salah")
					}
				} else {
					fmt.Println("Umur Anda Tidak Mencukupi")
				}
			} else {
				fmt.Println("Tahun tidak boleh minus")
			}
		} else {
			fmt.Println("Failed, Tahun tidak boleh di bawah 1900")
		}

	} else {
		fmt.Println("Failed, Karakter Harus Lebih dari 5 dan Batasnya 20")
	}
}
