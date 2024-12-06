package _4

import (
	"2024/utils"
	search2 "2024/utils/search"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Run(inputFile string) {
	arr := readFile(inputFile)
	num := search("XMAS", arr)
	fmt.Printf("Number of XMAS occurences in crossword: %d\n", num)

	needle := []rune{'X', 'M', 'A', 'S'}
	newNum := search2.Search2D(arr, needle, search2.Horizontal, true)
	newNum += search2.Search2D(arr, needle, search2.Vertical, true)
	newNum += search2.Search2D(arr, needle, search2.DiagonalFromLeft, true)
	newNum += search2.Search2D(arr, needle, search2.DiagonalFromRight, true)
	fmt.Printf("Number of XMAS occurences in crossword (using new search): %d\n", newNum)
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

func search(s string, in [][]rune) int {
	// 1. - left to right
	// 2. - right to left
	totalCount := horizontalSearch(s, in)

	// 3. | top to down
	// 4. | bottom to top
	totalCount += verticalSearch(s, in)

	// 5. \ left top to right bottom
	// 6. \ right bottom to left top
	totalCount += diagonalSearchLTR(s, in)

	// 7. / right top to left bottom
	// 8. / left bottom to right top
	totalCount += diagonalSearchRTL(s, in)

	return totalCount
}

func horizontalSearch(s string, in [][]rune) int {
	count := 0
	for _, runeLine := range in {
		stringLine := string(runeLine)
		count += strings.Count(stringLine, s)
		count += strings.Count(stringLine, utils.ReverseString(s))
	}

	return count
}

func verticalSearch(s string, in [][]rune) int {
	count := 0
	width := len(in[0])
	for i := 0; i < width; i++ {
		col := utils.GetCol(in, i)
		stringColLine := string(col)
		count += strings.Count(stringColLine, s)
		count += strings.Count(stringColLine, utils.ReverseString(s))
	}

	return count
}

func diagonalSearchLTR(s string, in [][]rune) int {
	count := 0
	width := len(in[0])
	for i := 0; i < width; i++ {
		diag := utils.GetDiagonalLTR(in, i, true)
		stringDiagLine := string(diag)
		count += strings.Count(stringDiagLine, s)
		count += strings.Count(stringDiagLine, utils.ReverseString(s))
	}

	// Skip first col, because it's the same as in loop above (0,0)
	for i := 1; i < width; i++ {
		diag := utils.GetDiagonalLTR(in, i, false)
		stringDiagLine := string(diag)
		count += strings.Count(stringDiagLine, s)
		count += strings.Count(stringDiagLine, utils.ReverseString(s))
	}

	return count
}

func diagonalSearchRTL(s string, in [][]rune) int {
	count := 0
	width := len(in[0])
	for i := 0; i < width; i++ {
		diag := utils.GetDiagonalRTL(in, i, true)
		stringDiagLine := string(diag)
		count += strings.Count(stringDiagLine, s)
		count += strings.Count(stringDiagLine, utils.ReverseString(s))
	}

	// Skip first col, because it's the same as in loop above (0,0)
	for i := 1; i < width; i++ {
		diag := utils.GetDiagonalRTL(in, i, false)
		stringDiagLine := string(diag)
		count += strings.Count(stringDiagLine, s)
		count += strings.Count(stringDiagLine, utils.ReverseString(s))
	}

	return count
}
