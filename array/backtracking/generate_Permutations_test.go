package backtracking

import (
	"reflect"
	"testing"
)

func TestGeneratePermutations(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Test case 1",
			args: args{nums: []int{1, 2, 3}},
			want: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 2, 1}, {3, 1, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GeneratePermutations(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GeneratePermutations() = %v, want %v", got, tt.want)
			}
		})
	}
}
