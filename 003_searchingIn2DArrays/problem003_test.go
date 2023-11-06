package searchingIn2DArrays

import (
	"testing"
)

func TestFind(t *testing.T) {
	type args struct {
		numbers [][]int
		target  int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty matrix",
			args: args{numbers: [][]int{{}}, target: 1},
			want: false},
		{
			name: "matrix with a single element equal to target",
			args: args{numbers: [][]int{{1}}, target: 1},
			want: true,
		},
		{
			name: "matrix with a single element not equal to target",
			args: args{numbers: [][]int{{2}}, target: 1},
			want: false,
		},
		{
			name: "matrix with multiple elements, target in the middle",
			args: args{numbers: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, target: 5},
			want: true,
		},
		{
			name: "matrix with multiple elements, target not in the matrix",
			args: args{numbers: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, target: 10},
			want: false,
		},
		{
			name: "empty matrix with target 0",
			args: args{numbers: [][]int{}, target: 0},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Find(tt.args.numbers, tt.args.target); got != tt.want {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}
