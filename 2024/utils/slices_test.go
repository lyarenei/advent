package utils

import (
	"reflect"
	"testing"
)

func TestDeepClone(t *testing.T) {
	type args[T any] struct {
		s [][]T
	}

	type testCase[T any] struct {
		name string
		args args[T]
		want [][]T
	}

	tests := []testCase[int]{
		{
			name: "TC",
			args: args[int]{
				s: [][]int{
					{0, 1},
					{2, 3},
				},
			},
			want: [][]int{
				{0, 1},
				{2, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DeepClone(tt.args.s)
			got[1][1] = 9
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeepClone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setHorizontalLine(t *testing.T) {
	type args[T any] struct {
		row int
		col int
		val []T
	}

	type testCase[T any] struct {
		name string
		args args[T]
		want [][]T
	}

	tests := []testCase[int]{
		{
			name: "At (0,0), exact size",
			args: args[int]{
				row: 0,
				col: 0,
				val: []int{8, 8, 8, 8},
			},
			want: [][]int{
				{8, 8, 8, 8},
				{4, 5, 6, 7},
				{8, 9, 10, 11},
				{12, 13, 14, 15},
			},
		},
		{
			name: "At (0,0), len 2",
			args: args[int]{
				row: 0,
				col: 0,
				val: []int{8, 8},
			},
			want: [][]int{
				{8, 8, 2, 3},
				{4, 5, 6, 7},
				{8, 9, 10, 11},
				{12, 13, 14, 15},
			},
		},
		{
			name: "At (0,0), len 6",
			args: args[int]{
				row: 0,
				col: 0,
				val: []int{8, 8, 8, 8, 8, 8},
			},
			want: [][]int{
				{8, 8, 8, 8},
				{4, 5, 6, 7},
				{8, 9, 10, 11},
				{12, 13, 14, 15},
			},
		},
		{
			name: "At (1,1), len 4",
			args: args[int]{
				row: 1,
				col: 1,
				val: []int{8, 8, 8, 8},
			},
			want: [][]int{
				{0, 1, 2, 3},
				{4, 8, 8, 8},
				{8, 9, 10, 11},
				{12, 13, 14, 15},
			},
		},
	}

	for _, tt := range tests {
		testSlice := DeepClone(testArray)
		t.Run(tt.name, func(t *testing.T) {
			setHorizontalLine(testSlice, tt.args.row, tt.args.col, tt.args.val)
			if !reflect.DeepEqual(testSlice, tt.want) {
				t.Errorf("setHorizontalLine() = %v, want %v", testSlice, tt.want)
			}
		})
	}
}

func Test_setVerticalLine(t *testing.T) {
	type args[T any] struct {
		row int
		col int
		val []T
	}

	type testCase[T any] struct {
		name string
		args args[T]
		want [][]T
	}

	tests := []testCase[int]{
		{
			name: "At (0,0), exact size",
			args: args[int]{
				row: 0,
				col: 0,
				val: []int{8, 8, 8, 8},
			},
			want: [][]int{
				{8, 1, 2, 3},
				{8, 5, 6, 7},
				{8, 9, 10, 11},
				{8, 13, 14, 15},
			},
		},
		{
			name: "At (0,0), len 2",
			args: args[int]{
				row: 0,
				col: 0,
				val: []int{6, 6},
			},
			want: [][]int{
				{6, 1, 2, 3},
				{6, 5, 6, 7},
				{8, 9, 10, 11},
				{12, 13, 14, 15},
			},
		},
		{
			name: "At (0,0), len 6",
			args: args[int]{
				row: 0,
				col: 0,
				val: []int{8, 8, 8, 8, 8, 8},
			},
			want: [][]int{
				{8, 1, 2, 3},
				{8, 5, 6, 7},
				{8, 9, 10, 11},
				{8, 13, 14, 15},
			},
		},
		{
			name: "At (1,1), len 4",
			args: args[int]{
				row: 1,
				col: 1,
				val: []int{8, 8, 8, 8},
			},
			want: [][]int{
				{0, 1, 2, 3},
				{4, 8, 6, 7},
				{8, 8, 10, 11},
				{12, 8, 14, 15},
			},
		},
	}

	for _, tt := range tests {
		testSlice := DeepClone(testArray)
		t.Run(tt.name, func(t *testing.T) {
			setVerticalLine(testSlice, tt.args.row, tt.args.col, tt.args.val)
			if !reflect.DeepEqual(testSlice, tt.want) {
				t.Errorf("setVerticalLine() = %v, want %v", testSlice, tt.want)
			}
		})
	}
}
