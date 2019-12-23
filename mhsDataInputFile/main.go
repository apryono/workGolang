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
var name, major, getNames, getMajors string
var age, getAges int

var path = "dataMahasiswa.txt"

type student struct {
	arrNames  string
	arrAges   int
	arrMajors string
}

var mhs = []student{}

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

		optNumb, _ := strconv.Atoi(inputWord())

		switch optNumb {
		case 1:
			addMahasiswa()
			break
		case 2:
			deleteMahasiswa()
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

func addMahasiswa() {
	if len(mhs) > 4 {
		fmt.Println("Maaf, Batas Input Data Maksimal 5 ")
	} else {
		fmt.Println(sym)
		fmt.Println("Add Mahasiswa")
		fmt.Println(sym)

		addName()
		addAge()
		addMajor()
		mhs = append(mhs, student{arrNames: getNames, arrAges: getAges, arrMajors: getMajors})
		fmt.Println("Data Ditambahkan")
		writeToFile(mhs)
	}
}

func addName() string {
	for {
		fmt.Println("karakter Lebih dari 2 dan Kurang dari 21 ")
		fmt.Print("Masukkan Nama : ")
		name = inputWord()
		if len(name) > 2 && len(name) < 21 {
			getNames = name
			break
		}
		fmt.Println("Error : Karakter harus lebih dari 2 dan kurang dari 21")
	}
	fmt.Println("Nama Anda: ", name)
	return name
}

func addAge() int {
	for {
		fmt.Print("Umur Anda ( Umur > 17 ): ")
		optNumb, _ := strconv.Atoi(inputWord())

		if optNumb >= 17 {
			getAges = optNumb
			break
		}
		fmt.Println("Error : Umur harus lebih dari 17 tahun")
	}
	return age
}

func addMajor() string {
	for {
		fmt.Print("Jurusan (Karakter <= 10): ")
		major = inputWord()

		if len(major) <= 10 {
			getMajors = major
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
		optNumb, _ := strconv.Atoi(inputWord())

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
	fmt.Println(sym)
	fmt.Println("View All")
	fmt.Println(sym)
	if len(mhs) != 0 {
		for i := 0; i < len(mhs); i++ {
			fmt.Println(sym)
			fmt.Println(i+1, ".")
			fmt.Println(sym)
			fmt.Println("Nama : ", mhs[i].arrNames)
			fmt.Println("Umur : ", mhs[i].arrAges)
			fmt.Println("Jurusan : ", mhs[i].arrMajors)
		}
	} else {
		fmt.Println("Maaf, Data Tidak Tersedia !")
	}
}

func viewMahasiswaByIndex() {
	fmt.Println(sym)
	fmt.Println("View By Index")
	fmt.Println(sym)

	var inputIndex int
	fmt.Print("Masukkan Nilai Index : ")
	fmt.Scanln(&inputIndex)
	if inputIndex >= 0 && inputIndex < len(mhs)-1 {
		fmt.Println("Hasil Data Sesuai index : ")
		fmt.Println("Nama : ", mhs[inputIndex].arrNames)
		fmt.Println("Umur : ", mhs[inputIndex].arrAges)
		fmt.Println("Jurusan : ", mhs[inputIndex].arrMajors)
	} else {
		fmt.Println("Maaf, Data Tidak Tersedia !")
	}
}

func deleteMahasiswa() {

	if (len(mhs) - 1) < 0 {
		fmt.Println("Maaf, Data tidak ada untuk di delete")
	} else {
		fmt.Println(sym)
		fmt.Println("Delete Mahasiswa")
		fmt.Println(sym)

		mhs = mhs[:len(mhs)-1]

		fmt.Println("Data Telah Di Hapus")
		deleteFile()
	}

}

func writeToFile(mhs []student) {
	var file, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0664)
	defer file.Close()
	if err == nil {
		for _, mahasiswa := range mhs {
			strMhs := fmt.Sprintf("%-10v,%3v, %3v\n", mahasiswa.arrNames, mahasiswa.arrAges, mahasiswa.arrMajors)
			file.WriteString(strMhs)
		}
		file.Sync()
	}
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return true
}

func deleteFile() {
	var err = os.Remove(path)
	if isError(err) {
		writeToFile(mhs)
		return
	}
	fmt.Println("File berhasil di delete")
}
