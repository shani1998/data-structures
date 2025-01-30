package slidingwindow

import "testing"

func Test_longestSubArraySumEqualTok(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test with simple case",
			args: args{
				arr: []int{1, -1, 5, -2, 3},
				k:   3,
			},
			want: 4, // [1, -1, 5, -2] has sum 3
		},
		{
			name: "Test with array having no valid subarray",
			args: args{
				arr: []int{1, 2, 3},
				k:   7,
			},
			want: 0, // No subarray with sum 7
		},
		{
			name: "Test with array having sum 0 as subarray",
			args: args{
				arr: []int{1, -1},
				k:   0,
			},
			want: 2, // The entire array has sum 0
		},
		{
			name: "Test with negative numbers",
			args: args{
				arr: []int{-1, 1, -1, 1},
				k:   0,
			},
			want: 4, // [-1, 1, -1, 1] has sum 0
		},
		{
			name: "Test with all zeros",
			args: args{
				arr: []int{0, 0, 0, 0, 0},
				k:   0,
			},
			want: 5, // The whole array has sum 0
		},
		{
			name: "Test with single element (not enough for subarray)",
			args: args{
				arr: []int{1},
				k:   1,
			},
			want: 0, // No subarray of length >= 2
		},
		{
			name: "Test with larger array and sum",
			args: args{
				arr: []int{5, 8, -3, -6, 7, 4, 1},
				k:   11,
			},
			want: 5, // [5, 8, -3, -6, 7] has sum 11
		},
		{
			name: "Test with zero sum in the middle",
			args: args{
				arr: []int{3, 2, -1, -1, 3},
				k:   4,
			},
			want: 5, // The whole array has sum 4
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubArraySumEqualTok(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("longestSubArraySumEqualTok() = %v, want %v", got, tt.want)
			}
		})
	}
}
