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
