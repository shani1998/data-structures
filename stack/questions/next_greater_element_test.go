package questions

import (
	"reflect"
	"testing"
)

func TestNextGreaterElement1(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test NextGreaterElement | Testcase 1",
			args: args{
				nums: []int{4, 1, 2, 1, 3},
			},
			want: []int{0, 2, 3, 3, 0},
		},
		{
			name: "Test NextGreaterElement | Testcase 2",
			args: args{
				nums: []int{2, 1, 3, 4, 2},
			},
			want: []int{3, 3, 4, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NextGreaterElement1(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NextGreaterElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNextGreaterElement2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test NextGreaterElement | Testcase 1",
			args: args{
				nums: []int{4, 1, 2, 1, 3},
			},
			want: []int{0, 2, 3, 3, 0},
		},
		{
			name: "Test NextGreaterElement | Testcase 2",
			args: args{
				nums: []int{2, 1, 3, 4, 2},
			},
			want: []int{3, 3, 4, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NextGreaterElement2(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NextGreaterElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
