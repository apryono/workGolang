package main

func main() {

}

func combination(slice []int) [][]int {
	result := [][]int{[]int{}}
	for _, i := range slice {
		var mixCombination [][]int
		for _, j := range slice {
			mixCombination := append(mixCombination, append(i, j))
			
		}
	}
}
