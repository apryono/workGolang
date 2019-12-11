package main

import "fmt"

const phi = 3.14

func main() {
	var inputPilih int
ulang:
	fmt.Println("--- WELCOME TO SIMPLE APPLICATION ---")
	fmt.Println("Pilihlah sesuai keinginan anda :")
	fmt.Println("1. Hasil Hitungan Luas Dan Keliling Lingkaran")
	fmt.Println("2. Hasil Hitungan Luas Dan Keliling Persegi")
	fmt.Println("3. Hasil Hitungan Luas Permukaan dan Volume Sebuah Balok")
	fmt.Println("4. Any Number For Exit")
	fmt.Print("Tentukan Pilihan Anda : ")
	fmt.Scan(&inputPilih)

	switch inputPilih {
	//kurung kurawal pada case bisa di gunakan bisa tidak
	case 1:
		{
			// fmt.Println("Success 1")
			var jari, luas, keliling float32
			fmt.Print("Masukkan Nilai Jari-Jari nya : ")
			fmt.Scan(&jari)

			luas = jari * jari * phi
			keliling = 2 * phi * jari
			fmt.Println("Hasil Luas Lingkaran Adalah ", luas)
			fmt.Println("Hasil Keliling Lingkaran adalah ", keliling)
			fmt.Println(" ")
			fmt.Println("---- TERIMAKASIH SUDAH MEMILIH ----")
			goto ulang
		}
	case 2:
		{
			var sisi, luasPersegi, kelilingPersegi float32
			fmt.Print("Masukkan Nilai Sisi nya : ")
			fmt.Scan(&sisi)

			luasPersegi = sisi * sisi
			kelilingPersegi = 4 * sisi
			fmt.Println("Hasil Luas Persegi adalah : ", luasPersegi)
			fmt.Println("Hasil Keliling Persegi adalah : ", kelilingPersegi)
			fmt.Println(" ")
			fmt.Println("---- TERIMAKASIH SUDAH MEMILIH ----")
			goto ulang
		}
	case 3:
		{
			var lenght, large, high, luasPermukaan, volume float32
			fmt.Print("Masukkan Nilai Panjangnya : ")
			fmt.Scan(&lenght)
			fmt.Print("Masukkan Nilai Lebar : ")
			fmt.Scan(&large)
			fmt.Print("Masukkan Nilai Tinggi : ")
			fmt.Scan(&high)

			luasPermukaan = 2 * (lenght*large + lenght*high + large*high)
			volume = lenght * large * high
			fmt.Println("Hasil Luas Permukaan Balok adalah : ", luasPermukaan)
			fmt.Println("Hasil Volume Balok adalah : ", volume)
			fmt.Println(" ")
			fmt.Println("---- TERIMAKASIH SUDAH MEMILIH ----")
			var pilih string
			fmt.Println("Apakah ingin dilanjut ? Ketik yes Untuk Lanjut, no Untuk Done")
			fmt.Println("Jawaban : ")
			fmt.Scan(&pilih)
			if pilih == "yes" {
				goto ulang
			} else {
				inputPilih = 7
			}
		}
	default:
		fmt.Println("Done")
	}
}
