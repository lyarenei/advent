package _3

import (
	"2024/utils"
	"fmt"
	"os"
	"regexp"
)

func Run(inputFile string) {
	data := readFile(inputFile)
	instr := getInstructions(data)
	result := calculate(instr)
	fmt.Printf("The result is: %d\n", result)
}

type mul struct {
	a int
	b int
}

func readFile(fileName string) string {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

func getInstructions(s string) []mul {
	re := regexp.MustCompile(`(?m)mul\(([1-9][0-9]{0,3}),([1-9][0-9]{0,3})\)`)
	matches := re.FindAllStringSubmatch(s, -1)
	muls := make([]mul, 0, len(matches))
	for i := range matches {
		match := matches[i]
		newMul := mul{
			a: utils.StringToInt(match[1]),
			b: utils.StringToInt(match[2]),
		}
		muls = append(muls, newMul)
	}

	return muls
}

func calculate(muls []mul) uint64 {
	var result uint64
	for _, m := range muls {
		result += uint64(m.a) * uint64(m.b)
	}

	return result
}
