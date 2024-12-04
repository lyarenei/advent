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

// GetCol gets a nth col, in a 2D array.
func GetCol[T any](input [][]T, at int) []T {
	var col []T
	for _, row := range input {
		col = append(col, row[at])
	}

	return col
}

// GetDiagonalLTR gets a diagonal line in a 2D array.
// Anchor specifies start on nth col. If horizontal is set to false, it specifies row instead.
// It is not possible to start "inside", only full diagonals starting from 0 to len() of the array is supported.
func GetDiagonalLTR[T any](input [][]T, anchor int, horizontal bool) []T {
	var diag []T

	// Horizontal, anchor specifies starting col
	if horizontal {
		for x := 0; x < len(input); x++ {
			y := x + anchor
			if y >= len(input[x]) {
				break
			}

			diag = append(diag, input[x][y])
		}

		return diag
	}

	// Vertical, anchor specifies starting row
	y := 0
	for x := anchor; x < len(input); x++ {
		if y >= len(input[x]) {
			break
		}

		diag = append(diag, input[x][y])
		y++
	}

	return diag
}
