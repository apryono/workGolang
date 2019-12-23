package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type robot struct {
	x      int
	y      int
	xMax   int
	yMax   int
	energi int
}

var rbt = robot{
	x:      0,
	y:      0,
	xMax:   10,
	yMax:   10,
	energi: 100,
}

var getInput, input string

var result int

func main() {
	sym := strings.Repeat("=", 50)

	fmt.Println(sym)
	fmt.Printf("Robot Koordinat (%v,%v), Energi %v \n", rbt.x, rbt.y, rbt.energi)
	fmt.Println(sym)
	fmt.Println("Robot bisa diperintahkan untuk bergerak\nke arah Utara [ U ], Selatan [ S ], Timur [ T ], dan Barat [ B ] ")
	fmt.Println(sym)
	drawVector(rbt.x, rbt.y, rbt.xMax, rbt.yMax)

	for {
		fmt.Print("Masukkan Perintah >> ")
		input = getCommand()
		// getInput = strings.ToLower(input)

		resultInput := strings.ContainsAny(input, "acdefghijklmnopqrvwxyz0123456789")

		if resultInput {
			fmt.Println("Mohon Di input hanya karakter U,T,B, dan S")
		} else {
			data := strings.Split(input, "")

			for _, result := range data {
				if result == "u" && rbt.yMax != rbt.y {
					rbt.y++
					rbt.energi -= 5
				} else if result == "t" && rbt.xMax != rbt.x {
					rbt.x++
					rbt.energi -= 5
				} else if result == "b" && rbt.x != 0 {
					rbt.x--
					rbt.energi -= 5
				} else if result == "s" && rbt.y != 0 {
					rbt.y--
					rbt.energi -= 5
				} else {
					fmt.Println("Robot tidak boleh keluar dari arena")
				}
				if rbt.energi == 0 {
					fmt.Println("Maaf, Energi Habis")
					return
				}
			}
		}

	}

}

func drawVector(x int, y int, yMax int, xMax int) {
	for pY := 0; pY <= yMax; pY++ {
		if pY == (yMax - y) {
			for pX := 0; pX < xMax; pX++ {
				if pX == x {
					fmt.Print(" |0| ")
				} else {
					fmt.Print("  .  ")
				}
			}

		} else {
			fmt.Print(strings.Repeat("  .  ", 10))
		}
		fmt.Print("\n")
	}
}

func getCommand() string {
	getData := bufio.NewScanner(os.Stdin)
	getData.Scan()
	data := getData.Text()
	fixData := strings.ToLower(data)
	return fixData
}

// func showCoordinat(x int, y int, energi int) {
// 	fmt.Printf("\nRobot Koordinat (%v,%v), Energi %v %%\n", rbt.x, rbt.y, rbt.energi)
// }
