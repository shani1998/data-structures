package questions

import (
	"reflect"
	"testing"
)

func TestCountGreaterElementRight(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test case 1",
			args: args{
				nums: []int{10, 12, 8, 17, 3, 24, 19},
			},
			want: []int{4, 3, 3, 2, 2, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountGreaterElementRight(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountGreaterElementRight() = %v, want %v", got, tt.want)
			}
		})

	}
}

func TestCountGreaterRight(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test case 1",
			args: args{
				nums: []int{3, 24, 19},
			},
			want: []int{2, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countGreaterRight(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountGreaterElementRight() = %v, want %v", got, tt.want)
			}
		})

	}
}
