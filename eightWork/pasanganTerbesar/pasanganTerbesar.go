package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var arrHasil []int
var getNumb []string

func main() {
	var maxValue int
	fmt.Print("Input Number : ")
	getInput := isInput()
	_, errs := strconv.Atoi(getInput)
	if errs == nil {

		// fungsi delSpace di panggil untuk menghapus spasi
		delSpace := delSpace(getInput)
		getNumb = strings.Split(delSpace, "")

		// fungsi nya untuk mendapat kan array
		arrHasil = getArray(getNumb)

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

func delSpace(input string) string {
	delele := strings.Replace(input, " ", "", -1)
	return delele
}

func getArray(input []string) []int {
	for i := 1; i < len(input); i++ {
		join := input[i-1] + input[i]
		isNumber, err := strconv.Atoi(join)
		if err != nil {
			break
		} else {
			arrHasil = append(arrHasil, isNumber)

		}

	}
	return arrHasil
}
