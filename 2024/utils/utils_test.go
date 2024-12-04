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
			name: "Start at 0th row",
			args: args[int]{
				input:      testArray,
				anchor:     0,
				horizontal: true,
			},
			want: []int{0, 5, 10, 15},
		},
		{
			name: "Start at 0th col",
			args: args[int]{
				input:      testArray,
				anchor:     0,
				horizontal: false,
			},
			want: []int{0, 5, 10, 15},
		},
		{
			name: "Start at 1st row",
			args: args[int]{
				input:      testArray,
				anchor:     1,
				horizontal: true,
			},
			want: []int{1, 6, 11},
		},
		{
			name: "Start at 1st col",
			args: args[int]{
				input:      testArray,
				anchor:     1,
				horizontal: false,
			},
			want: []int{4, 9, 14},
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
