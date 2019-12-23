package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var name, getNames, getInput string
var hp, damage, getHp, getDamages, getNumb int

type heroes struct {
	nameHero   string
	hpHero     int
	damageHero int
}

var listHero []heroes
var sym = strings.Repeat("=", 10)
var sym2 = strings.Repeat("-", 68)
var dataMenu int

// method untuk attack heroes
func (h *heroes) attackHeroes(h2 *heroes) (hpValid bool, heroesValid bool) {
	var hPHero int
	// hero tidak dapat menyerang diri sendiri
	if h == h2 {
		return true, false
	}
	// jika hp hero = 0 tidak dapat menyerang
	if h.hpHero == 0 {
		return false, true
	}
	// jika hp hero = tidak dapat menyerang
	if h2.hpHero == 0 {
		return false, true
	}

	// hp hero dikurangi dengan damage hero yang attack
	hPHero = h2.hpHero - h.damageHero
	if hPHero > 0 {
		h2.hpHero = hPHero
	} else {
		h2.hpHero = 0
	}
	return true, true
}

// method untuk healing heroes
func (h *heroes) healHeroes(h2 *heroes) (hpValid bool) {
	if h.hpHero == 0 {
		return false
	}
	if h2.hpHero == 0 {
		return false
	}

	h2.hpHero = h2.hpHero + (h.damageHero / 2)
	return true
}

func main() {
	fmt.Println(sym2)
	fmt.Println(sym, "+ + [x]===[ WELCOME TO GAMES ARENA ]===[x] + +", sym)
	fmt.Println(sym2)
	addHero()       // untuk memanggil fungsi menambah hero
	menus(dataMenu) // untuk memanggil fungsi menu
	// fmt.Println(listHero)
}

func menus(input int) int {
	for {
		fmt.Println(sym2)
		fmt.Println(sym, "[+] + Main Menu Games + [+]", sym)
		fmt.Println(sym2)
		fmt.Println(`
	Silahkan Pilih Menu 
	1. View Player
	2. Attack
	3. Health
	4. Exit
	`)
		fmt.Println(sym2)
		fmt.Print("Masukkan Menu yang dipilih : ")
		getInput = inputWord()
		getNumb, _ = strconv.Atoi(getInput)

		switch getNumb {
		case 1:
			viewPlayer()
		case 2:
			attackPlayer()
		case 3:
			healthPlayer()
		case 4:
			os.Exit(0)
		default:
			fmt.Println("Input sesuai pilihan ")
		}
	}

}

func attackPlayer() {
	fmt.Println(sym2)
	fmt.Println(sym, "[+] + Choose Heroes Attacker + [+]", sym)
	fmt.Println(sym2)
	for {
		fmt.Print("Pilih Attacker : ")
		getInput = inputWord()
		attacker, _ := strconv.Atoi(getInput)

		if attacker >= 0 && attacker <= len(listHero) {
			fmt.Println(attacker, ".")
			fmt.Println("Name Heroes : ", listHero[attacker].nameHero)
			fmt.Println("Name HP : ", listHero[attacker].hpHero)
			fmt.Println("Name Damage : ", listHero[attacker].damageHero)
			fmt.Println(sym)
		} else {
			fmt.Println("Nothing! Player Heroes ")
			break
		}

		fmt.Print("Pilih Defender : ")
		getInput = inputWord()
		defender, _ := strconv.Atoi(getInput)

		if defender >= 0 && defender <= len(listHero) {
			fmt.Println(defender, ".")
			fmt.Println("Name Heroes : ", listHero[defender].nameHero)
			fmt.Println("Name HP : ", listHero[defender].hpHero)
			fmt.Println("Name Damage : ", listHero[defender].damageHero)
			fmt.Println(sym)
		} else {
			fmt.Println("Nothing! Player Heroes ")
			break
		}

		hp, player := listHero[attacker].attackHeroes(&listHero[defender])

		if !hp {
			fmt.Println("Can't Attack,  You Have Died") // akan di tampilkan setelah input kedua nya
		}
		if !player {
			fmt.Println("Can't Attack Your Self !")
		}
		menus(dataMenu)

	}

}

func healthPlayer() {
	fmt.Println(sym2)
	fmt.Println(sym, "[+] + Healer Heroes Player + [+]", sym)
	fmt.Println(sym2)
	fmt.Println("Choose 0, 1, 2 ")
	fmt.Println(sym)

	for {
		fmt.Print("Masukkan Healer : ")
		healing := inputWord()
		healerHeroes, _ := strconv.Atoi(healing)

		if healerHeroes >= 0 && healerHeroes <= len(listHero) {
			fmt.Println(healerHeroes, ".")
			fmt.Println("Name Heroes : ", listHero[healerHeroes].nameHero)
			fmt.Println("Name HP : ", listHero[healerHeroes].hpHero)
			fmt.Println("Name Damage : ", listHero[healerHeroes].damageHero)
			fmt.Println(sym)
		} else {
			fmt.Println("Nothing! Player Heroes ")
			break
		}

		fmt.Print("Masukkan Receiver : ")
		healing = inputWord()
		saveHero, _ := strconv.Atoi(healing)

		if saveHero >= 0 && saveHero <= len(listHero) {
			fmt.Println(saveHero, ".")
			fmt.Println("Name Heroes :  ", listHero[saveHero].nameHero)
			fmt.Println("Name HP : ", listHero[saveHero].hpHero)
			fmt.Println("Name Damage : ", listHero[saveHero].damageHero)
			fmt.Println(sym)
		} else {
			fmt.Println("Nothing! Player Heroes ")
			break
		}

		hp := listHero[healerHeroes].healHeroes(&listHero[saveHero])
		if !hp {
			fmt.Println("Can't Healing, You Have Died") // tidak dapat menyembuhkan jika memiliki hp = 0
		}
		menus(dataMenu)
	}

}

func viewPlayer() {
	fmt.Println(sym)
	fmt.Println("View Heroes Player")
	fmt.Println(sym)
	fmt.Println(`
	Pilih Menu View 
	1. View All Heroes
	2. Back to Menu
	`)
	fmt.Print("Masukkan Pilihan Anda : ")
	var input int
	fmt.Scan(&input)
	switch input {
	case 1:
		fmt.Println("Data yang Tersedia : ")
		viewAllHeroes()
	case 2:
		menus(dataMenu)
	default:
		fmt.Println("Error : Pilihan hanya sesuai menu tersedia")
	}

}

func viewAllHeroes() {
	fmt.Println(sym2)
	fmt.Println(sym, "[+] + View All Player + [+]", sym)
	fmt.Println(sym2)
	if len(listHero) != 0 {
		for i := 0; i < len(listHero); i++ {
			fmt.Println(sym)
			fmt.Println(i+1, ".")
			fmt.Println(sym)
			fmt.Println("Nama Heroes :  ", listHero[i].nameHero)
			fmt.Println("HP :  ", listHero[i].hpHero)
			fmt.Println("Damage :  ", listHero[i].damageHero)
		}
	} else {
		fmt.Println("Maaf, Data Tidak Tersedia !")
	}
}

func addHero() {
	fmt.Println(sym, "[+] + Add Heroes Player + [+]", sym)
	fmt.Println(sym2)
	for {
		if len(listHero) < 3 {
			for {
				fmt.Println("Karakter Lebih dari 2 dan Kurang dari 21 ")
				fmt.Print("Masukkan Nama =>> ")
				name = inputWord()
				if len(name) > 2 && len(name) < 21 {
					for {
						fmt.Println("HP Awal Harus Diatas 100")
						fmt.Print("Input HP =>> ")
						hp, _ = strconv.Atoi(inputWord())

						if hp >= 100 {
							for {
								fmt.Println("Damage minimal lebih dari 10")
								fmt.Print("Input Damage =>> ")
								damage, _ = strconv.Atoi(inputWord())

								if damage >= 10 {
									listHero = append(listHero, heroes{nameHero: name, hpHero: hp, damageHero: damage})
									fmt.Println(sym2)
									fmt.Println(sym, "[+] + Data Ditambahkan + [+]", sym)
									fmt.Println(sym2)
									break
								} else {
									fmt.Println("Input Angka Yang Benar ")
								}
							}
							break
						} else {
							fmt.Println("Input Angka Yang Benar ")
						}

					}
				}
				break
			}

		} else {
			fmt.Println("Maaf, Batas Input Data Maksimal 3 ")
			break
		}

	}

}

func inputWord() string {
	getData := bufio.NewScanner(os.Stdin)
	getData.Scan()
	data := getData.Text()
	return data
}
