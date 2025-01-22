package backtracking

import (
	"reflect"
	"testing"
)

func TestGenerateParentheses(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test case 1",
			args: args{n: 1},
			want: []string{"()"},
		},
		{
			name: "Test case 2",
			args: args{n: 2},
			want: []string{"(())", "()()"},
		},
		{
			name: "Test case 3",
			args: args{n: 3},
			want: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateParentheses(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
