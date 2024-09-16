package questions

import "testing"

func TestMinDiffHeightTowers(t *testing.T) {
	type args struct {
		arr []int
		n   int
		k   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case 1",
			args: args{
				arr: []int{1, 5, 8, 10},
				n:   4,
				k:   2,
			},
			want: 5,
		},
		{
			name: "Test case 2",
			args: args{
				arr: []int{3, 9, 12, 16, 20},
				n:   5,
				k:   3,
			},
			want: 11,
		},
		{
			name: "Test case 3",
			args: args{
				arr: []int{2, 6, 3, 4, 7, 2, 10, 3, 2, 1},
				n:   10,
				k:   5,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinDiffHeightTowers(tt.args.arr, tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("MinDiffHeightTowers() = %v, want %v", got, tt.want)
			}
		})
	}
}
