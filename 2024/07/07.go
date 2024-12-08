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
		res := isSolvable(
			[]rune{'*', '+', '|'},
			calibration.Operands[0],
			calibration.Operands[1],
			calibration.Operands[2:],
			calibration.Result,
		)

		if len(res) > 0 {
			//fmt.Printf("Calibration for %d: %v\n", calibration.Result, string(res))
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

func isSolvable(operators []rune, leftOp, rightOp int, nextOps []int, wantResult int) []rune {
	if len(operators) == 0 {
		return []rune{}
	}

	newResult := -1
	for _, operator := range operators {
		switch operator {
		case '*':
			newResult = leftOp * rightOp
		case '+':
			newResult = leftOp + rightOp
		case '|':
			concat := fmt.Sprintf("%d%d", leftOp, rightOp)
			newResult = utils.StringToInt(concat)
		default:
			// noop
		}

		if newResult == wantResult && len(nextOps) == 0 {
			return []rune{operator}
		}

		if len(nextOps) > 0 {
			nextOperators := isSolvable(operators, newResult, nextOps[0], nextOps[1:], wantResult)
			if len(nextOperators) > 0 {
				return append([]rune{operator}, nextOperators...)
			}
		}
	}

	return []rune{}
}
