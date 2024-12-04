package _3

import (
	"2024/utils"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strings"
)

func Run(inputFile string) {
	data := readFile(inputFile)
	instr := getInstructions(data)
	result := calculate(instr)
	fmt.Printf("The result is: %d\n", result)
}

func readFile(fileName string) string {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

func getInstructions(s string) []instruction {
	re := regexp.MustCompile(`(?m)(mul\(([1-9][0-9]{0,3}),([1-9][0-9]{0,3})\))|(do\(\))|(don't\(\))`)
	matches := re.FindAllStringSubmatch(s, -1)
	muls := make([]instruction, 0, len(matches))
	for i := range matches {
		match := matches[i]
		instrStr := match[0]
		var instr instruction
		if strings.HasPrefix(instrStr, "mul") {
			instr = mul{
				a: utils.StringToInt(match[2]),
				b: utils.StringToInt(match[3]),
			}
		} else if strings.HasPrefix(instrStr, "don't") {
			instr = dont{}
		} else if strings.HasPrefix(instrStr, "do") {
			instr = do{}
		}

		muls = append(muls, instr)
	}

	return muls
}

func calculate(instructions []instruction) uint64 {
	var result uint64
	enabled := true
	for _, instr := range instructions {
		if reflect.TypeOf(instr).Name() == "do" {
			enabled = true
		} else if reflect.TypeOf(instr).Name() == "dont" {
			enabled = false
		}

		if !enabled {
			continue
		}

		if reflect.TypeOf(instr).Name() == "mul" {
			result += instr.Execute().(uint64)
		}
	}

	return result
}
