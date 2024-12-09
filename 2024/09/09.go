package _9

import (
	"2024/utils"
	"bufio"
	"fmt"
	"golang.org/x/exp/maps"
	"os"
	"slices"
)

func Run(inputFile string) {
	encodedFs := readFile(inputFile)
	decodedFs := decodeFs(encodedFs)

	keys := maps.Keys(decodedFs)
	slices.Sort(keys)
	for _, k := range keys {
		fmt.Printf("%d: %v\n", k, decodedFs[k])
	}
}

func readFile(fileName string) string {
	file, err := os.Open(fileName)
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	return scanner.Text()
}

type block struct {
	ID        int
	MaxLength int
	Values    []int
}

func decodeFs(encodedFs string) map[int]block {
	blocks := make(map[int]block)
	blockID := 0

	for i := 0; i < len(encodedFs); i += 2 {
		end := i + 2
		if end >= len(encodedFs) {
			end = len(encodedFs) - 1
		}

		encBlock := encodedFs[i:end]
		if len(encBlock) == 0 {
			break
		}

		fileSize := utils.StringToInt(string(encBlock[0]))
		freeSpace := 0
		if len(encBlock) == 2 {
			freeSpace = utils.StringToInt(string(encBlock[1]))
		}

		blockValues := make([]int, fileSize, fileSize+freeSpace)
		for j := 0; j < fileSize; j++ {
			blockValues[j] = blockID
		}

		blocks[blockID] = block{
			ID:        blockID,
			Values:    blockValues,
			MaxLength: fileSize + freeSpace,
		}

		blockID++
	}

	return blocks
}
