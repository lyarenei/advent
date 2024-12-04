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
