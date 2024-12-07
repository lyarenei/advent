package utils

import (
	"2024/types"
	"golang.org/x/exp/slices"
)

func Search2D[T types.BasicType](haystack [][]T, needle []T, direction Direction, includeReverse bool) int {
	needleLen := len(needle)
	count := 0

	for i, haystackRow := range haystack {
		if len(haystackRow) < needleLen {
			continue
		}

		for j := range haystackRow {
			line := GetLine(haystack, direction, i, j, needleLen)
			if slices.Equal(line, needle) {
				count++
			}

			revNeedle := slices.Clone(needle)
			slices.Reverse(revNeedle)
			if includeReverse && slices.Equal(line, revNeedle) {
				count++
			}
		}
	}

	return count
}

func Search2DSimple[T comparable](haystack [][]T, needle T) int {
	count := 0
	for _, haystackRow := range haystack {
		for _, item := range haystackRow {
			if item == needle {
				count++
			}
		}
	}

	return count
}

func FindIdx2D[T comparable](haystack [][]T, needle T) (int, int) {
	for i, haystackRow := range haystack {
		for j, item := range haystackRow {
			if item == needle {
				return i, j
			}
		}
	}

	return -1, -1
}
