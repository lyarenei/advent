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

func TestGetDiagonalLTR(t *testing.T) {
	type args[T any] struct {
		input      [][]T
		anchor     int
		horizontal bool
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "Start at 0th col",
			args: args[int]{
				input:      testArray,
				anchor:     0,
				horizontal: true,
			},
			want: []int{0, 5, 10, 15},
		},
		{
			name: "Start at 0th row",
			args: args[int]{
				input:      testArray,
				anchor:     0,
				horizontal: false,
			},
			want: []int{0, 5, 10, 15},
		},
		{
			name: "Start at 1st col",
			args: args[int]{
				input:      testArray,
				anchor:     1,
				horizontal: true,
			},
			want: []int{1, 6, 11},
		},
		{
			name: "Start at 1st row",
			args: args[int]{
				input:      testArray,
				anchor:     1,
				horizontal: false,
			},
			want: []int{4, 9, 14},
		},
		{
			name: "Start at 2nd col",
			args: args[int]{
				input:      testArray,
				anchor:     2,
				horizontal: true,
			},
			want: []int{2, 7},
		},
		{
			name: "Start at 2nd row",
			args: args[int]{
				input:      testArray,
				anchor:     2,
				horizontal: false,
			},
			want: []int{8, 13},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDiagonalLTR(tt.args.input, tt.args.anchor, tt.args.horizontal); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDiagonalLTR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDiagonalRTL(t *testing.T) {
	type args[T any] struct {
		input      [][]T
		anchor     int
		horizontal bool
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "Start at last col",
			args: args[int]{
				input:      testArray,
				anchor:     3,
				horizontal: true,
			},
			want: []int{3, 6, 9, 12},
		},
		{
			name: "Start at 0th row",
			args: args[int]{
				input:      testArray,
				anchor:     0,
				horizontal: false,
			},
			want: []int{3, 6, 9, 12},
		},
		{
			name: "Start at last -1 col",
			args: args[int]{
				input:      testArray,
				anchor:     2,
				horizontal: true,
			},
			want: []int{2, 5, 8},
		},
		{
			name: "Start at 1st row",
			args: args[int]{
				input:      testArray,
				anchor:     1,
				horizontal: false,
			},
			want: []int{7, 10, 13},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDiagonalRTL(tt.args.input, tt.args.anchor, tt.args.horizontal); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDiagonalRTL() = %v, want %v", got, tt.want)
			}
		})
	}
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
