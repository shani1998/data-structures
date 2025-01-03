package questions

import (
	"testing"

	"github.com/shani1998/data-structures/linkedlist"
)

func TestDetectCycle(t *testing.T) {
	type args struct {
		head *linkedlist.SinglyNode
	}
	tests := []struct {
		name string
		args args
		want *linkedlist.SinglyNode
	}{
		{
			name: "Empty list",
			args: args{
				head: nil,
			},
			want: nil,
		},
		{
			name: "Single node with no cycle",
			args: args{
				head: &linkedlist.SinglyNode{Val: 1},
			},
			want: nil,
		},
		{
			name: "Single node with a cycle",
			args: args{
				head: func() *linkedlist.SinglyNode {
					n1 := &linkedlist.SinglyNode{Val: 1}
					n1.Next = n1 // Cycle pointing to itself
					return n1
				}(),
			},
			want: func() *linkedlist.SinglyNode {
				n1 := &linkedlist.SinglyNode{Val: 1}
				return n1
			}(),
		},
		{
			name: "No cycle in the list",
			args: args{
				head: func() *linkedlist.SinglyNode {
					n1 := &linkedlist.SinglyNode{Val: 1}
					n2 := &linkedlist.SinglyNode{Val: 2}
					n3 := &linkedlist.SinglyNode{Val: 3}
					n1.Next = n2
					n2.Next = n3
					return n1
				}(),
			},
			want: nil,
		},
		{
			name: "Cycle exists in the list",
			args: args{
				head: func() *linkedlist.SinglyNode {
					n1 := &linkedlist.SinglyNode{Val: 1}
					n2 := &linkedlist.SinglyNode{Val: 2}
					n3 := &linkedlist.SinglyNode{Val: 3}
					n4 := &linkedlist.SinglyNode{Val: 4}
					n1.Next = n2
					n2.Next = n3
					n3.Next = n4
					n4.Next = n2 // Cycle starts at n2
					return n1
				}(),
			},
			want: func() *linkedlist.SinglyNode {
				n2 := &linkedlist.SinglyNode{Val: 2}
				return n2
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DetectCycle(tt.args.head)
			// We use pointer comparison here because `reflect.DeepEqual` is not suitable for cycles.
			if tt.want != nil && got.Val != tt.want.Val {
				t.Errorf("DetectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
