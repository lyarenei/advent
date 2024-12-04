package _2

import (
	"2024/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Run(inputFile string) {
	reports := readFile(inputFile)
	safeReports := calcSafeReports(reports, isReportSafe)
	fmt.Printf("The number of safe reports is: %d\n", safeReports)

	safeReportsWithCorrection := calcSafeReports(reports, isReportSafeWithTolerance)
	fmt.Printf("The number of safe reports with correction is: %d\n", safeReportsWithCorrection)
}

func readFile(fileName string) [][]int {
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

	var arr [][]int
	for scanner.Scan() {
		splitLine := strings.Split(scanner.Text(), " ")
		numArray := utils.StringToIntArray(splitLine)
		arr = append(arr, numArray)
	}

	return arr
}

func calcSafeReports(reports [][]int, cb func([]int) bool) int {
	safeReports := 0
	for _, report := range reports {
		if cb(report) {
			safeReports++
		}
	}

	return safeReports
}

func isReportSafe(reportLevels []int) bool {
	return isIncreasing(reportLevels) || isDecreasing(reportLevels)
}

func isReportSafeWithTolerance(reportLevels []int) bool {
	return isIncreasingWithTolerance(reportLevels) || isDecreasingWithTolerance(reportLevels)
}

func isIncreasingWithTolerance(reportLevels []int) bool {
	for i, j := 0, 1; i < len(reportLevels)-1; i, j = i+1, j+1 {
		if j > len(reportLevels)-1 {
			j = len(reportLevels) - 1
		}

		if !checkConditions(reportLevels[i], reportLevels[j], 1, 3) {
			reducedReports := utils.RemoveAt(reportLevels, i)
			if !isIncreasing(reducedReports) {
				reducedReports = utils.RemoveAt(reportLevels, j)
				return isIncreasing(reducedReports)
			} else {
				return true
			}
		}
	}

	return true
}

func isIncreasing(reportLevels []int) bool {
	for i, j := 0, 1; i < len(reportLevels)-1; i, j = i+1, j+1 {
		if j > len(reportLevels)-1 {
			j = len(reportLevels) - 1
		}

		if !checkConditions(reportLevels[i], reportLevels[j], 1, 3) {
			return false
		}
	}

	return true
}

func isDecreasingWithTolerance(reportLevels []int) bool {
	for i, j := 0, 1; i < len(reportLevels)-1; i, j = i+1, j+1 {
		if j > len(reportLevels)-1 {
			j = len(reportLevels) - 1
		}

		if !checkConditions(reportLevels[j], reportLevels[i], 1, 3) {
			reducedReports := utils.RemoveAt(reportLevels, i)
			if !isDecreasing(reducedReports) {
				reducedReports = utils.RemoveAt(reportLevels, j)
				return isDecreasing(reducedReports)
			} else {
				return true
			}
		}
	}

	return true
}

func isDecreasing(reportLevels []int) bool {
	for i, j := 0, 1; i < len(reportLevels)-1; i, j = i+1, j+1 {
		if j > len(reportLevels)-1 {
			j = len(reportLevels) - 1
		}

		if !checkConditions(reportLevels[j], reportLevels[i], 1, 3) {
			return false
		}
	}

	return true
}

// checkConditions checks if number A is lower than B.
// Additionally, it checks if the difference is greater than specified limits.
func checkConditions(a int, b int, minDiff int, maxDiff int) bool {
	diff := b - a
	if diff <= 0 || diff < minDiff || diff > maxDiff {
		return false
	}

	return true
}
