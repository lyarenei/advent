package _7

import (
	"2024/utils"
	"bufio"
	"os"
	"strings"
)

type Calibration struct {
	Result   int
	Operands []int
}

func Run(inputFile string) {
	readFile(inputFile)
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
