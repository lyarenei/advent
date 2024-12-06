package search

import (
	"2024/types"
	"reflect"
	"testing"
)

func Test_getLine(t *testing.T) {
	type args[T types.BasicType] struct {
		haystack [][]T
		d        Direction
		row      int
		col      int
		len      int
	}

	type testCase[T types.BasicType] struct {
		name string
		args args[T]
		want []T
	}

	var testArray = [][]int{
		{0, 1, 2, 3},
		{4, 5, 6, 7},
		{8, 9, 10, 11},
		{12, 13, 14, 15},
	}

	tests := []testCase[int]{
		{
			name: "Horizontal (0,0)",
			args: args[int]{
				haystack: testArray,
				d:        Horizontal,
				row:      0,
				col:      0,
				len:      3,
			},
			want: []int{0, 1, 2},
		},
		{
			name: "Horizontal (0,1)",
			args: args[int]{
				haystack: testArray,
				d:        Horizontal,
				row:      0,
				col:      1,
				len:      3,
			},
			want: []int{1, 2, 3},
		},
		{
			name: "Horizontal (0,2)",
			args: args[int]{
				haystack: testArray,
				d:        Horizontal,
				row:      0,
				col:      2,
				len:      3,
			},
			want: []int{2, 3},
		},
		{
			name: "Horizontal (0,3)",
			args: args[int]{
				haystack: testArray,
				d:        Horizontal,
				row:      0,
				col:      3,
				len:      3,
			},
			want: []int{3},
		},
		{
			name: "Horizontal (0,4)",
			args: args[int]{
				haystack: testArray,
				d:        Horizontal,
				row:      0,
				col:      4,
				len:      3,
			},
			want: []int{},
		},
		{
			name: "Horizontal (3,1)",
			args: args[int]{
				haystack: testArray,
				d:        Horizontal,
				row:      3,
				col:      1,
				len:      3,
			},
			want: []int{13, 14, 15},
		},
		{
			name: "Horizontal (4,4)",
			args: args[int]{
				haystack: testArray,
				d:        Horizontal,
				row:      4,
				col:      4,
				len:      3,
			},
			want: []int{},
		},
		{
			name: "Vertical (0,0)",
			args: args[int]{
				haystack: testArray,
				d:        Vertical,
				row:      0,
				col:      0,
				len:      3,
			},
			want: []int{0, 4, 8, 12},
		},
		{
			name: "Vertical (0,1)",
			args: args[int]{
				haystack: testArray,
				d:        Vertical,
				row:      0,
				col:      1,
				len:      3,
			},
			want: []int{1, 5, 9, 13},
		},
		{
			name: "Vertical (0,2)",
			args: args[int]{
				haystack: testArray,
				d:        Vertical,
				row:      0,
				col:      2,
				len:      3,
			},
			want: []int{2, 6, 10, 14},
		},
		{
			name: "Vertical (0,3)",
			args: args[int]{
				haystack: testArray,
				d:        Vertical,
				row:      0,
				col:      3,
				len:      3,
			},
			want: []int{3, 7, 11, 15},
		},
		{
			name: "Vertical (0,4)",
			args: args[int]{
				haystack: testArray,
				d:        Vertical,
				row:      0,
				col:      4,
				len:      3,
			},
			want: []int{},
		},
		{
			name: "Vertical (3,1)",
			args: args[int]{
				haystack: testArray,
				d:        Vertical,
				row:      3,
				col:      1,
				len:      3,
			},
			want: []int{13},
		},
		{
			name: "Vertical (4,4)",
			args: args[int]{
				haystack: testArray,
				d:        Vertical,
				row:      4,
				col:      4,
				len:      3,
			},
			want: []int{},
		},
		{
			name: "Diagonal L (0,0)",
			args: args[int]{
				haystack: testArray,
				d:        DiagonalFromLeft,
				row:      0,
				col:      0,
				len:      3,
			},
			want: []int{0, 5, 10, 15},
		},
		{
			name: "Diagonal L (0,1)",
			args: args[int]{
				haystack: testArray,
				d:        DiagonalFromLeft,
				row:      0,
				col:      1,
				len:      3,
			},
			want: []int{1, 6, 11},
		},
		{
			name: "Diagonal L (0,2)",
			args: args[int]{
				haystack: testArray,
				d:        DiagonalFromLeft,
				row:      0,
				col:      2,
				len:      3,
			},
			want: []int{2, 7},
		},
		{
			name: "Diagonal L (0,3)",
			args: args[int]{
				haystack: testArray,
				d:        DiagonalFromLeft,
				row:      0,
				col:      3,
				len:      3,
			},
			want: []int{3},
		},
		{
			name: "Diagonal L (0,4)",
			args: args[int]{
				haystack: testArray,
				d:        DiagonalFromLeft,
				row:      0,
				col:      4,
				len:      3,
			},
			want: []int{},
		},
		{
			name: "Diagonal (3,1)",
			args: args[int]{
				haystack: testArray,
				d:        DiagonalFromLeft,
				row:      3,
				col:      1,
				len:      3,
			},
			want: []int{13},
		},
		{
			name: "Diagonal (4,4)",
			args: args[int]{
				haystack: testArray,
				d:        Vertical,
				row:      4,
				col:      4,
				len:      3,
			},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLine(tt.args.haystack, tt.args.d, tt.args.row, tt.args.col, tt.args.len); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
