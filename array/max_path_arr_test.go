package main

import "testing"

func Test_maxSumPath(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
		m    int
		n    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case 1",
			args: args{
				arr1: []int{2, 3, 7, 10, 12},
				arr2: []int{1, 5, 7, 8},
				m:    5,
				n:    4,
			},
			want: 35,
		},
		{
			name: "Test case 2",
			args: args{
				arr1: []int{10, 12},
				arr2: []int{5, 7, 9},
				m:    2,
				n:    3,
			},
			want: 22,
		},
		{
			name: "Test case 3",
			args: args{
				arr1: []int{2, 3, 7, 10, 12, 15, 30, 34},
				arr2: []int{1, 5, 7, 8, 10, 15, 16, 19},
				m:    8,
				n:    8,
			},
			want: 122,
		},
		{
			name: "Test case 4",
			args: args{
				arr1: []int{1, 2, 3},
				arr2: []int{3, 4, 5},
				m:    3,
				n:    3,
			},
			want: 15,
		},
		{
			name: "Test case 5",
			args: args{
				arr1: []int{1, 2, 3},
				arr2: []int{4, 5, 6},
				m:    3,
				n:    3,
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSumPath(tt.args.arr1, tt.args.arr2, tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("maxSumPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
