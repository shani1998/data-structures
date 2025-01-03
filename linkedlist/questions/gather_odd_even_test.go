package questions

import (
	"reflect"
	"testing"

	"github.com/shani1998/data-structures/linkedlist"
)

func TestOddEvenList(t *testing.T) {
	type args struct {
		head *linkedlist.SinglyNode
	}
	tests := []struct {
		name string
		args args
		want *linkedlist.SinglyNode
	}{
		{
			name: "Test with empty list",
			args: args{
				head: nil,
			},
			want: nil,
		},
		{
			name: "Test with one node",
			args: args{
				head: &linkedlist.SinglyNode{Val: 1},
			},
			want: &linkedlist.SinglyNode{Val: 1},
		},
		{
			name: "Test with two nodes",
			args: args{
				head: &linkedlist.SinglyNode{Val: 1, Next: &linkedlist.SinglyNode{Val: 2}},
			},
			want: &linkedlist.SinglyNode{
				Val: 1,
				Next: &linkedlist.SinglyNode{
					Val: 2,
				},
			},
		},
		{
			name: "Test with five nodes",
			args: args{
				head: createList([]int64{1, 2, 3, 4, 5}),
			},
			want: createList([]int64{1, 3, 5, 2, 4}),
		},
		{
			name: "Test with six nodes",
			args: args{
				head: createList([]int64{1, 2, 3, 4, 5, 6}),
			},
			want: createList([]int64{1, 3, 5, 2, 4, 6}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := oddEvenList(tt.args.head)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OddEvenList() = %v, want %v", got, tt.want)
			}
		})
	}
}
