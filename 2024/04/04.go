package _4

import (
	"bufio"
	"fmt"
	"os"
)

func Run(inputFile string) {
	arr := readFile(inputFile)
	num := search("XMAS", arr)
	fmt.Printf("Number of XMAS occurences in crossword: %d\n", num)
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
	// 3. | top to down
	// 4. | bottom to top
	// 5. / left bottom to right top
	// 5. / right top to left bottom
	// 6. \ left top to right bottom
	// 7. \ right bottom to left top

	return 0
}
