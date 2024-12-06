package search

import (
	"2024/types"
	"strconv"
)

const (
	Horizontal Direction = iota
	Vertical
	DiagonalFromLeft
	DiagonalFromRight
	All
)

// TODO: multiple directions at once
func Search2D[T types.BasicType](haystack [][]T, needle T, d Direction, includeReverse bool) int {
	needleStr := strconv.Itoa(int(needle))
	needleLen := len(needleStr)
	count := 0

	for i, haystackRow := range haystack {
		if len(haystackRow) < needleLen {
			continue
		}

		for j := range haystackRow {
			getLine(haystack, d, i, j, needleLen)
			// compare
		}
	}

	return count
}

func getLine[T types.BasicType](haystack [][]T, d Direction, row int, col int, lineLen int) []T {
	if row >= len(haystack) {
		return []T{}
	}

	switch d {
	case Horizontal:
		return getHorizontalLine(haystack[row], col, lineLen)
	//case Vertical:
	//	return getVerticalLine(haystack, row, col, lineLen)
	//case DiagonalFromLeft:
	//
	default:
		return []T{}
	}
}

func getHorizontalLine[T types.BasicType](haystack []T, start int, lineLen int) []T {
	end := lineLen + start
	if len(haystack)-start < lineLen {
		end = len(haystack)
	}

	return haystack[start:end]
}
