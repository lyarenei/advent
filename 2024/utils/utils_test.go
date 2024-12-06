package utils

import (
	"reflect"
	"testing"
)

var testArray = [][]int{
	{0, 1, 2, 3},
	{4, 5, 6, 7},
	{8, 9, 10, 11},
	{12, 13, 14, 15},
}

func TestSubslice2D(t *testing.T) {
	type args[T any] struct {
		arr      [][]T
		startRow int
		startCol int
		edgeSize int
	}

	type testCase[T any] struct {
		name string
		args args[T]
		want [][]T
	}

	tests := []testCase[int]{
		{
			name: "Simple (0,0)",
			args: args[int]{
				arr:      testArray,
				startRow: 0,
				startCol: 0,
				edgeSize: 3,
			},
			want: [][]int{
				{0, 1, 2},
				{4, 5, 6},
				{8, 9, 10},
			},
		},
		{
			name: "Size 1 (0,0)",
			args: args[int]{
				arr:      testArray,
				startRow: 0,
				startCol: 0,
				edgeSize: 1,
			},
			want: [][]int{{0}},
		},
		{
			name: "Size 4 (0,0)",
			args: args[int]{
				arr:      testArray,
				startRow: 0,
				startCol: 0,
				edgeSize: 4,
			},
			want: testArray,
		},
		{
			name: "Size 2 (2,2)",
			args: args[int]{
				arr:      testArray,
				startRow: 2,
				startCol: 2,
				edgeSize: 2,
			},
			want: [][]int{
				{10, 11},
				{14, 15},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Subslice2D(tt.args.arr, tt.args.startRow, tt.args.startCol, tt.args.edgeSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Subslice2D() = %v, want %v", got, tt.want)
			}
		})
	}
}
