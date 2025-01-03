package questions

import (
	"reflect"
	"testing"

	"github.com/shani1998/data-structures/linkedlist"
)

// Helper function to create a singly linked list from a slice of integers
func createList(nums []int64) *linkedlist.SinglyNode {
	if len(nums) == 0 {
		return nil
	}

	head := &linkedlist.SinglyNode{Val: nums[0]}
	current := head
	for _, num := range nums[1:] {
		current.Next = &linkedlist.SinglyNode{Val: num}
		current = current.Next
	}
	return head
}

func TestMergeTwoLists(t *testing.T) {
	type args struct {
		list1 *linkedlist.SinglyNode
		list2 *linkedlist.SinglyNode
	}
	tests := []struct {
		name string
		args args
		want *linkedlist.SinglyNode
	}{
		{
			name: "Test with two empty lists",
			args: args{
				list1: nil,
				list2: nil,
			},
			want: nil,
		},
		{
			name: "Test with one empty list",
			args: args{
				list1: createList([]int64{1, 3, 5}),
				list2: nil,
			},
			want: createList([]int64{1, 3, 5}),
		},
		{
			name: "Test with one empty list reversed",
			args: args{
				list1: nil,
				list2: createList([]int64{2, 4, 6}),
			},
			want: createList([]int64{2, 4, 6}),
		},
		{
			name: "Test with two non-empty lists",
			args: args{
				list1: createList([]int64{1, 3, 5}),
				list2: createList([]int64{2, 4, 6}),
			},
			want: createList([]int64{1, 2, 3, 4, 5, 6}),
		},
		{
			name: "Test with lists having some overlapping values",
			args: args{
				list1: createList([]int64{1, 5, 9}),
				list2: createList([]int64{2, 5, 8}),
			},
			want: createList([]int64{1, 2, 5, 5, 8, 9}),
		},
		{
			name: "Test with lists having all smaller or larger elements",
			args: args{
				list1: createList([]int64{1, 2, 3}),
				list2: createList([]int64{4, 5, 6}),
			},
			want: createList([]int64{1, 2, 3, 4, 5, 6}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MergeTwoLists(tt.args.list1, tt.args.list2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
