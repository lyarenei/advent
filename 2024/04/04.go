package _4

import (
	"2024/utils"
	"bufio"
	"fmt"
	"os"
)

func Run(inputFile string) {
	arr := readFile(inputFile)

	needle := []rune{'X', 'M', 'A', 'S'}
	num := utils.Search2D(arr, needle, utils.Horizontal, true)
	num += utils.Search2D(arr, needle, utils.Vertical, true)
	num += utils.Search2D(arr, needle, utils.DiagonalFromLeft, true)
	num += utils.Search2D(arr, needle, utils.DiagonalFromRight, true)
	fmt.Printf("Number of XMAS occurences in crossword: %d\n", num)

	xNum := searchX(arr, []rune{'M', 'A', 'S'})
	fmt.Printf("Number of X-MAS occurences in crossword: %d\n", xNum)
}

func readFile(fileName string) [][]rune {
	file, err := os.Open(fileName)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var arr [][]rune
	for scanner.Scan() {
		arr = append(arr, []rune(scanner.Text()))
	}

	return arr
}

func searchX(haystack [][]rune, needle []rune) int {
	edgeSize := len(needle)
	count := 0
	for i := 0; i < len(haystack); i++ {
		for j := 0; j < len(haystack[i]); j++ {
			subSlice := utils.Subslice2D(haystack, i, j, edgeSize)
			if len(subSlice) < edgeSize {
				break
			}

			onDiagLeft := utils.Search2D(subSlice, needle, utils.DiagonalFromLeft, true)
			onDiagRight := utils.Search2D(subSlice, needle, utils.DiagonalFromRight, true)
			if onDiagLeft == 1 && onDiagRight == 1 {
				count++
			}
		}
	}

	return count
}
