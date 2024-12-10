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
	compactFs(decodedFs)
	checksum := calculateChecksum(decodedFs)
	fmt.Printf("The filesystem checksum is %d\n", checksum)
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

func (b *block) IsFull() bool {
	return len(b.Values) == b.MaxLength
}

func (b *block) IsEmpty() bool {
	return len(b.Values) == 0
}

func (b *block) Append(ID int) {
	b.Values = append(b.Values, ID)
}

func (b *block) RemoveLast() int {
	lastID := b.Values[len(b.Values)-1]
	b.Values = b.Values[:len(b.Values)-1]
	return lastID
}

func decodeFs(encodedFs string) map[int]*block {
	blocks := make(map[int]*block)
	blockID := 0

	for i := 0; i < len(encodedFs); i += 2 {
		end := i + 2
		if end >= len(encodedFs) {
			end = len(encodedFs)
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

		blocks[blockID] = &block{
			ID:        blockID,
			Values:    blockValues,
			MaxLength: fileSize + freeSpace,
		}

		blockID++
	}

	return blocks
}

func compactFs(decodedFs map[int]*block) {
	left := 0
	right := len(decodedFs) - 1
	for left < right {
		dstBlock := decodedFs[left]
		srcBlock := decodedFs[right]

		if dstBlock.IsFull() {
			left++
			continue
		}

		if srcBlock.IsEmpty() {
			right--
			continue
		}

		fileBlock := srcBlock.RemoveLast()
		dstBlock.Append(fileBlock)
	}
}

func calculateChecksum(blocks map[int]*block) int {
	globalPos := 0
	checksum := 0

	keys := maps.Keys(blocks)
	slices.Sort(keys)
	for _, k := range keys {
		blk := blocks[k]
		if blk.IsEmpty() {
			break
		}

		for _, val := range blk.Values {
			checksum += globalPos * val
			globalPos++
		}
	}

	return checksum
}
