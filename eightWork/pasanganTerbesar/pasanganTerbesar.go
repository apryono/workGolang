package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var arrHasil []int
	var maxValue int
	fmt.Print("Input Number : ")
	getInput := isInput()
	_, errs := strconv.Atoi(getInput)
	if errs == nil {
		delSpace := strings.Replace(getInput, " ", "", -1)
		getNumb := strings.Split(delSpace, "")

		for i := 1; i < len(getNumb); i++ {
			join := getNumb[i-1] + getNumb[i]
			isNumber, err := strconv.Atoi(join)
			if err != nil {
				break
			} else {
				arrHasil = append(arrHasil, isNumber)

			}

		}
		fmt.Println(arrHasil)
		maxValue = arrHasil[0]

		for _, value := range arrHasil {
			if value > maxValue {
				maxValue = value
			}
		}

		fmt.Println("The Biggest Value : ", maxValue)
	} else {
		fmt.Println("Harus Angka Ya :), Dan Tidak Pakai Space :(")
		main()
	}
}

func isInput() string {
	getData := bufio.NewScanner(os.Stdin)
	getData.Scan()
	data := getData.Text()
	return data
}
