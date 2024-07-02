package questions

import "testing"

func Test_increasingTriplet(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test case 1, all are in ascending order",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
			},
			want: true,
		},
		{
			name: "Test case 2, all are in descending order",
			args: args{
				nums: []int{5, 4, 3, 2, 1},
			},
			want: false,
		},
		{
			name: "Test case 3, all are in random order",
			args: args{
				nums: []int{2, 1, 5, 0, 4, 6},
			},
			want: true,
		},
		{
			name: "Test case 4, all are in random order",
			args: args{
				nums: []int{1, 2, 1, 1, 3},
			},
			want: true,
		},
		{
			name: "Test case 5, all are in random order",
			args: args{
				nums: []int{1, 1, 3},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := increasingTriplet(tt.args.nums); got != tt.want {
				t.Errorf("increasingTriplet() = %v, want %v", got, tt.want)
			}
		})
	}
}
