package utils

import (
	"2024/types"
)

const (
	Horizontal Direction = iota
	Vertical
	DiagonalFromLeft
	DiagonalFromRight
)

func GetLine[T types.BasicType](haystack [][]T, direction Direction, row int, col int, lineLen int) []T {
	if row >= len(haystack) || col >= len(haystack[row]) {
		return []T{}
	}

	switch direction {
	case Horizontal:
		return getHorizontalLine(haystack[row], col, lineLen)
	case Vertical:
		return getVerticalLine(haystack, col, row, lineLen)
	case DiagonalFromLeft:
		return getDiagonalLineFromLeft(haystack, row, col, lineLen)
	case DiagonalFromRight:
		return getDiagonalLineFromRight(haystack, row, col, lineLen)
	default:
		return []T{}
	}
}

func getHorizontalLine[T types.BasicType](haystack []T, start int, lineLen int) []T {
	end := start + lineLen
	if end >= len(haystack) {
		end = len(haystack)
	}

	return haystack[start:end]
}

func getVerticalLine[T types.BasicType](haystack [][]T, col int, start int, lineLen int) []T {
	end := start + lineLen
	if end > len(haystack) {
		end = len(haystack)
	}

	line := make([]T, 0, end)
	for i := start; i < end; i++ {
		line = append(line, haystack[i][col])
	}

	return line
}

func getDiagonalLineFromLeft[T types.BasicType](haystack [][]T, startRow int, startCol int, lineLen int) []T {
	rowEnd := startRow + lineLen
	if rowEnd > len(haystack[startRow]) {
		rowEnd = len(haystack)
	}

	colEnd := startCol + lineLen
	if colEnd > len(haystack) {
		colEnd = len(haystack)
	}

	line := make([]T, 0)
	for x, y := startRow, startCol; x < rowEnd && y < colEnd; x, y = x+1, y+1 {
		line = append(line, haystack[x][y])
	}

	return line
}

func getDiagonalLineFromRight[T types.BasicType](haystack [][]T, startRow int, startCol int, lineLen int) []T {
	endRow := startRow + lineLen
	if endRow > len(haystack[startRow]) {
		endRow = len(haystack)
	}

	endCol := startCol - lineLen
	if endCol < 0 {
		endCol = 0
	}

	line := make([]T, 0)
	for row, col := startRow, startCol; row < endRow && col >= endCol; row, col = row+1, col-1 {
		line = append(line, haystack[row][col])
	}

	return line
}

// GetCol gets a nth col, in a 2D array.
func GetCol[T any](input [][]T, at int) []T {
	var col []T
	for _, row := range input {
		col = append(col, row[at])
	}

	return col
}

func MapSlice[I any, O any](slice []I, function func(I) O) []O {
	newSlice := make([]O, 0, len(slice))
	for _, item := range slice {
		newSlice = append(newSlice, function(item))
	}

	return newSlice
}

func Swap[T any](s []T, i int, j int) {
	s[i], s[j] = s[j], s[i]
}
