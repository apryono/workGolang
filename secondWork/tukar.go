package main

import "fmt"

func main() {
	// ini untuk memanggil fungsi proses nya
	tukarPosisi()
}

func tukarPosisi() {
	var a, b int

	fmt.Printf("Masukkan Nilai a = ")
	fmt.Scan(&a)
	fmt.Printf("Masukkan Nilai b = ")
	fmt.Scan(&b)
	fmt.Println("Hasilnya menjadi sebagai berikut :")
	a = a - b
	b = a + b
	a = b - a
	fmt.Println("Hasil a = ", a)
	fmt.Println("Hasil b = ", b)
	//ini untuk memanggil fungsi mengulang proses nya atau tidak
	again()
}

func again() {
	var ingin string
	fmt.Println("Apakah ingin diulang? Jika Lanjut Ketik ya, no untuk berhenti")
	fmt.Scan(&ingin)
	if ingin == "ya" {
		tukarPosisi()
	} else if ingin == "no" {
		fmt.Println("Thanks")
	} else {
		fmt.Println("Silahkan Pilih ya atau no")
		again()
	}
}
