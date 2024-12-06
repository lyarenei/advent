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
	if row >= len(haystack) || col >= len(haystack[row]) {
		return []T{}
	}

	switch d {
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
	if end >= len(haystack) {
		end = len(haystack) - 1
	}

	line := make([]T, 0, end)
	for i := start; i <= end; i++ {
		line = append(line, haystack[i][col])
	}

	return line
}

func getDiagonalLineFromLeft[T types.BasicType](haystack [][]T, startRow int, startCol int, lineLen int) []T {
	rowEnd := startRow + lineLen
	if rowEnd >= len(haystack[startRow]) {
		rowEnd = len(haystack) - 1
	}

	colEnd := startCol + lineLen
	if colEnd >= len(haystack) {
		colEnd = len(haystack) - 1
	}

	line := make([]T, 0)
	for x, y := startRow, startCol; x <= rowEnd && y <= colEnd; x, y = x+1, y+1 {
		line = append(line, haystack[x][y])
	}

	return line
}

func getDiagonalLineFromRight[T types.BasicType](haystack [][]T, startRow int, startCol int, lineLen int) []T {
	endRow := startRow + lineLen
	if endRow >= len(haystack[startRow]) {
		endRow = len(haystack) - 1
	}

	endCol := startCol - lineLen
	if endCol < 0 {
		endCol = 0
	}

	line := make([]T, 0)
	for row, col := startRow, startCol; row <= endRow && col >= endCol; row, col = row+1, col-1 {
		line = append(line, haystack[row][col])
	}

	return line
}
