package recursion

import "testing"

func TestSumArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Test case 1",
			nums: []int{1, 2, 3, 4, 5},
			want: 15,
		},
		{
			name: "Test case 2",
			nums: []int{1, 2, 3},
			want: 6,
		},
		{
			name: "Test case 3",
			nums: []int{1, 2},
			want: 3,
		},
		{
			name: "Test case 4",
			nums: []int{1},
			want: 1,
		},
		{
			name: "Test case 5",
			nums: []int{},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumArray(tt.nums); got != tt.want {
				t.Errorf("SumArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
