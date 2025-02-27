package recursion

import (
	"reflect"
	"testing"
)

func Test_subsets(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "Test case 1",
			nums: []int{1, 2, 3},
			want: [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
		{
			name: "Test case 2",
			nums: []int{0},
			want: [][]int{{}, {0}},
		},
		{
			name: "Test case 3",
			nums: []int{},
			want: [][]int{{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := subsets(tt.nums)
			if len(got) != len(tt.want) {
				t.Errorf("subsets() length = %v, want length %v", len(got), len(tt.want))
				return
			}
			for _, expectedSubset := range tt.want {
				found := false
				for _, resultSubset := range got {
					if reflect.DeepEqual(resultSubset, expectedSubset) {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("subsets() missing expected subset %v", expectedSubset)
				}
			}
		})
	}
}
