package sorting

import (
	"reflect"
	"testing"
)

func Test_binSort(t *testing.T) {
	type args struct {
		arr []int
		n   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test case 1",
			args: args{
				arr: []int{1, 0, 1, 1, 0},
				n:   5,
			},
			want: []int{0, 0, 1, 1, 1},
		},
		{
			name: "Test case 2",
			args: args{
				arr: []int{1, 1, 1, 1, 1},
				n:   5,
			},
			want: []int{1, 1, 1, 1, 1},
		},
		{
			name: "Test case 3",
			args: args{
				arr: []int{0, 0, 0, 0, 0},
				n:   5,
			},
			want: []int{0, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binSort(tt.args.arr, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("binSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortColors(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test case 1",
			args: args{
				nums: []int{2, 0, 2, 1, 1, 0},
			},
		},
		{
			name: "Test case 2",
			args: args{
				nums: []int{2, 1, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.args.nums)
		})
	}
}
