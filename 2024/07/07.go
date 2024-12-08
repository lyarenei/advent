package _7

import (
	"2024/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Calibration struct {
	Result   int
	Operands []int
}

func Run(inputFile string) {
	calibrations := readFile(inputFile)
	sum := 0
	for _, calibration := range calibrations {
		res := isSolvable([]rune{'*', '+'}, calibration.Operands, calibration.Result, 0)
		if len(res) > 0 {
			fmt.Printf("Calibration for %d: %v\n", calibration.Result, string(res))
			sum += calibration.Result
		}
	}

	fmt.Printf("Sum of all solvable calibrations: %d\n", sum)
}

func readFile(fileName string) []Calibration {
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

	calibrations := make([]Calibration, 0)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		operandsStr := strings.Trim(parts[1], " ")
		operands := strings.Split(operandsStr, " ")
		calibrations = append(calibrations, Calibration{
			Result:   utils.StringToInt(parts[0]),
			Operands: utils.StringToIntArray(operands),
		})
	}

	return calibrations
}

func isSolvable(operators []rune, operands []int, wantResult, interResult int) []rune {
	if len(operators) == 0 || len(operands) == 0 {
		return []rune{}
	}

	if len(operands) == 1 {
		newInterResult := -1
		for _, operator := range operators {
			switch operator {
			case '*':
				newInterResult = interResult * operands[0]
			case '+':
				newInterResult = interResult + operands[0]
			default:
				// noop
			}

			if newInterResult == wantResult {
				return []rune{operator}
			}
		}

		return []rune{}
	}

	newInterResult := -1
	for _, operator := range operators {
		switch operator {
		case '*':
			if interResult == 0 {
				newInterResult = operands[0] * operands[1]
			} else {
				newInterResult = interResult * operands[0]
			}
		case '+':
			if interResult == 0 {
				newInterResult = operands[0] + operands[1]
			} else {
				newInterResult = interResult + operands[0]
			}
		default:
			// noop
		}

		var nextOperators []rune
		if interResult == 0 {
			nextOperators = isSolvable(operators, operands[2:], wantResult, newInterResult)
		} else {
			nextOperators = isSolvable(operators, operands[1:], wantResult, newInterResult)
		}

		if len(nextOperators) == 0 && newInterResult == wantResult {
			return []rune{operator}
		}

		if len(nextOperators) > 0 {
			correctOps := []rune{operator}
			return append(correctOps, nextOperators...)
		}
	}

	return []rune{}
}
