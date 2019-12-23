package main

import (
	"bufio"
	"errors"
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
func (h *heroes) attackHeroes(h2 *heroes) (hpValid bool, heroesValid bool, err error) {
	var hPHero int
	// hero tidak dapat menyerang diri sendiri

	if h == h2 {
		var err = errors.New("Can't Attack Your Self")
		return true, false, err
	}
	// jika hp hero = 0 tidak dapat menyerang
	if h.hpHero == 0 {
		var err = errors.New("Can't Attack With Zero HP")
		return false, true, err
	}
	// jika hp hero = tidak dapat menyerang
	if h2.hpHero == 0 {
		var err = errors.New("Can't Attack With Zero HP")
		return false, true, err
	}

	// hp hero dikurangi dengan damage hero yang attack
	hPHero = h2.hpHero - h.damageHero
	if hPHero > 0 {
		h2.hpHero = hPHero
	} else {
		h2.hpHero = 0
	}
	return true, true, err
}

// method untuk healing heroes
func (h *heroes) healHeroes(h2 *heroes) (hpValid bool, err error) {
	if h.hpHero == 0 {
		var err = errors.New("Can't Healing With Zero HP")
		return false, err
	}
	if h2.hpHero == 0 {
		var err = errors.New("Can't Healing With Zero HP")
		return false, err
	}

	h2.hpHero = h2.hpHero + (h.damageHero / 2)
	return true, err
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
	3. Healing
	4. Exit
	`)
		fmt.Println(sym2)
		fmt.Print("Masukkan Menu yang dipilih : ")
		getInput = inputWord()
		getNumb, _ = strconv.Atoi(getInput)

		switch getNumb {
		case 1:
			viewAllHeroes()
			// viewPlayer()
		case 2:
			err := attackPlayer()
			if err != nil {
				fmt.Println(err)
			}

		case 3:
			err := healthPlayer()
			if err != nil {
				fmt.Println(err)
			}
		case 4:
			os.Exit(0)
		default:
			fmt.Println("Input sesuai pilihan ")
		}
	}

}

func attackPlayer() error {
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
			fmt.Println("HP Heroes : ", listHero[attacker].hpHero)
			fmt.Println("Damage Heroes : ", listHero[attacker].damageHero)
			fmt.Println(sym2)
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
			fmt.Println("HP Heroes : ", listHero[defender].hpHero)
			fmt.Println("Damage Heroes : ", listHero[defender].damageHero)
			fmt.Println(sym2)
		} else {
			fmt.Println("Nothing! Player Heroes ")
			break
		}

		hp, player, err := listHero[attacker].attackHeroes(&listHero[defender])

		if err != nil {
			return fmt.Errorf("-+[+]+- Failed, Cause : %w -+[+]+-", err)
		} else if !hp {
			return fmt.Errorf("-+[+]+- Failed, Cause : %w -+[+]+-", hp) // akan di tampilkan setelah input kedua nya
		} else if !player {
			return fmt.Errorf("-+[+]+- Failed, Cause : %w -+[+]+-", player)
		}
		menus(dataMenu)

	}
	return nil
}

func healthPlayer() error {
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
			fmt.Println("HP Heroes : ", listHero[healerHeroes].hpHero)
			fmt.Println("Damage Heroes : ", listHero[healerHeroes].damageHero)
			fmt.Println(sym2)
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
			fmt.Println("HP HEroes : ", listHero[saveHero].hpHero)
			fmt.Println("Damage Heroes : ", listHero[saveHero].damageHero)
			fmt.Println(sym2)
		} else {
			fmt.Println("Nothing! Player Heroes ")
			break
		}

		hp, err := listHero[healerHeroes].healHeroes(&listHero[saveHero])
		if err != nil {
			return fmt.Errorf("-+[+]+- Failed, Cause : %w -+[+]+-", err)
		}
		if !hp {
			return fmt.Errorf("-+[+]+- Failed, Cause : %w -+[+]+-", err)
			// tidak dapat menyembuhkan jika memiliki hp = 0
		}
		menus(dataMenu)
	}
	return nil

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
			fmt.Println("HP Heroes :  ", listHero[i].hpHero)
			fmt.Println("Damage Heroes :  ", listHero[i].damageHero)
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
				fmt.Print("Nama Heroes =>> ")
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
