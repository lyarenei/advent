package utils

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
