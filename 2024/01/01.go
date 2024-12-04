package _1

import (
	"2024/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func Run(inputFile string) {
	// Step 1, read numbers into columns
	col1, col2 := readFile(inputFile)

	// Step 2, sort columns
	sort.Ints(col1)
	sort.Ints(col2)

	// Step 3, calculate distance
	dist := calcDistance(col1, col2)
	fmt.Printf("The distance is %d\n", dist)

	// Step 4, calculate similarity score
	simil := calcSimil(col1, col2)
	fmt.Printf("The similarity score is %d\n", simil)
}

func readFile(fileName string) (col1 []int, col2 []int) {
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
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		if len(col1) <= len(col2) {
			col1 = append(col1, num)
		} else {
			col2 = append(col2, num)
		}
	}

	return col1, col2
}

func calcDistance(col1 []int, col2 []int) int {
	if len(col1) != len(col2) {
		log.Fatal("Input is not valid; column lengths do not match")
	}

	if len(col1) == 0 || len(col2) == 0 {
		log.Fatal("Input is not valid; at least one pair of numbers is required")
	}

	dist := 0
	for i := 0; i < len(col1); i++ {
		dist += utils.AbsInt(col1[i] - col2[i])
	}

	return dist
}

func calcSimil(col1 []int, col2 []int) int {
	similScore := 0
	for i := 0; i < len(col1); i++ {
		similScore += col1[i] * utils.AppearsTimes(col1[i], col2)
	}

	return similScore
}
