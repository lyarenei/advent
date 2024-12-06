package utils

import "strconv"

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func AppearsTimes(num int, col []int) int {
	times := 0
	for _, i := range col {
		if col[i] == num {
			times++
		}
	}

	return times
}

func StringToIntArray(strArr []string) []int {
	var numArr []int
	for _, s := range strArr {
		numArr = append(numArr, StringToInt(s))
	}

	return numArr
}

func StringToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return num
}

func RemoveAt(arr []int, index int) []int {
	// P E R F O R M A N C E
	ret := make([]int, 0, len(arr)-1)
	ret = append(ret, arr[:index]...)
	return append(ret, arr[index+1:]...)
}

func ReverseString(s string) string {
	// Trigger warning: not a performant solution (immutable strings)
	result := ""
	for _, v := range s {
		result = string(v) + result
	}

	return result
}

func Subslice2D[T any](arr [][]T, startRow int, startCol int, edgeSize int) [][]T {
	endRow := startRow + edgeSize
	if endRow >= len(arr[startRow]) {
		endRow = len(arr[startRow])
	}

	endCol := startCol + edgeSize
	if endCol >= len(arr) {
		endCol = len(arr)
	}

	subSlice := make([][]T, 0, edgeSize)
	for i := startRow; i < endRow; i++ {
		selectedRow := arr[i][startCol:endCol]
		subSlice = append(subSlice, selectedRow)
	}

	return subSlice
}
