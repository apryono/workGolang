package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var isInt = regexp.MustCompile("^[0-9]+$")
var sym = strings.Repeat("=", 40)
var name, major string
var age int
var arrNames []string
var arrAges []int
var arrMajors []string
var dataMenu string

func main() {

	menus(dataMenu)
}

func menus(input string) int {
	// optNumb := inputNumber(input)
	for {
		fmt.Println(sym)
		fmt.Println("Main Menu")
		fmt.Println(sym)
		fmt.Println(`
	Silahkan Pilih Menu 
	1. Add Mahasiswa
	2. Delete Mahasiswa
	3. View Mahasiswa
	4. Exit
	`)
		fmt.Println(sym)
		fmt.Print("Masukkan Menu yang dipilih : ")

		optWord := inputWord()
		optNumb := inputNumber(optWord)

		switch optNumb {
		case -1:
			fmt.Print("")
		case 1:
			addMahasiswa()
			break
		case 2:
			fmt.Println("Pilihan Kedua")
		case 3:
			viewMahasiswa()
		case 4:
			os.Exit(0)
		default:
			fmt.Println("Input sesuai pilihan ")
		}

	}

}

func inputWord() string {
	getData := bufio.NewScanner(os.Stdin)
	getData.Scan()
	data := getData.Text()
	return data
}
func inputNumber(data string) int {
	if isInt.MatchString(data) {
		result, _ := strconv.Atoi(data)
		return result
	} else {
		fmt.Println("Masukkan Angka yang benar")
		return -1
	}
}

func addMahasiswa() {
	fmt.Println(sym)
	fmt.Println("Add Mahasiswa")
	fmt.Println(sym)

	addName(name)
	addAge(age)
	addMajor(major)

}

func addName(name string) string {
	for {
		fmt.Println("karakter Lebih dari 2 dan Kurang dari 21 ")
		fmt.Print("Masukkan Nama : ")
		name = inputWord()
		if len(name) > 2 && len(name) < 21 {
			arrNames = append(arrNames, name)
			break
		}
		fmt.Println("Error : Karakter harus lebih dari 2 dan kurang dari 21")
	}
	fmt.Println("Nama Anda: ", name)
	return name
}

func addAge(age int) int {
	for {
		fmt.Print("Umur Anda ( Umur > 17 ): ")
		optWord := inputWord()
		optNumb := inputNumber(optWord)

		if optNumb >= 17 {
			arrAges = append(arrAges, age)
			break
		}
		fmt.Println("Error : Umur harus lebih dari 17 tahun")
	}
	return age
}

func addMajor(major string) string {
	for {
		fmt.Print("Jurusan (Karakter <= 10): ")
		major = inputWord()

		if len(major) <= 10 {
			arrMajors = append(arrMajors, major)
			break
		}
		fmt.Println("ERROR :Karakter harus Di bawah atau sama dengan 10")
	}
	return major
}

func viewMahasiswa() {
	for {
		fmt.Println(sym)
		fmt.Println("View Mahasiswa")
		fmt.Println(sym)
		fmt.Println(`
		Pilih Menu View 
		1. View All
		2. View By Index
		3. Back to Menu
		`)
		fmt.Print("Masukkan Pilihan Anda : ")
		optWord := inputWord()
		optNumb := inputNumber(optWord)

		switch optNumb {
		case -1:
			fmt.Println("")
		case 1:
			fmt.Println("Data yang Tersedia : ")
			viewAllMahasiswa()
		case 2:
			viewMahasiswaByIndex()
		case 3:
			menus(dataMenu)
		default:
			fmt.Println("Error : Pilihan hanya sesuai menu tersedia")
		}

	}

}

func viewAllMahasiswa() {
	for i := 0; i < len(arrNames); i++ {
		fmt.Println(sym)
		fmt.Println(i+1, ".")
		fmt.Println(sym)
		fmt.Println("Nama : ", arrNames[i])
		fmt.Println("Umur : ", arrAges[i])
		fmt.Println("Jurusan : ", arrMajors[i])
	}
}

func viewMahasiswaByIndex() {
	var inputIndex int
	fmt.Println("Masukkan Nilai Index : ")
	fmt.Scanln(&inputIndex)
	fmt.Println("Hasil Data Sesuai index : ")
	fmt.Println("Nama : ", arrNames[inputIndex])
	fmt.Println("Umur : ", arrAges[inputIndex])
	fmt.Println("Jurusan : ", arrMajors[inputIndex])
}
