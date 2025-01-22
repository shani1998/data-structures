package linkedlist

import (
	"fmt"
	"testing"
)

func constructList(Val []interface{}) *SinglyNode {
	list := SinglyLinkedList{}
	for _, num := range Val {
		list.InsertLast(NewSinglyNode(num))
	}
	return list.head
}

func compare(expectedNodes []interface{}, head *SinglyNode) (int, interface{}, interface{}, error) {
	curr, count := head, 0
	for curr != nil {
		if curr.Val != expectedNodes[count] {
			return count, curr.Val, expectedNodes[count], fmt.Errorf("")
		}
		curr = curr.Next
		count++
	}
	return 0, nil, nil, nil
}

func TestLinkedList_DeleteFirst(t *testing.T) {
	type fields struct {
		head *SinglyNode
		curr *SinglyNode
		len  int
	}
	tests := []struct {
		name          string
		fields        fields
		expectedNodes []interface{}
		expectedLen   int
	}{
		{
			"There is no element in the list",
			fields{
				head: nil,
				curr: nil,
				len:  0,
			},
			nil,
			0,
		},
		{
			"There is one element in the list",
			fields{
				head: constructList([]interface{}{"node-0"}),
				curr: nil,
				len:  1,
			},
			[]interface{}{},
			0,
		},
		{
			"There is multiple element in the list",
			fields{
				head: constructList([]interface{}{"node-0", "node-1", "node-2"}),
				curr: nil,
				len:  3,
			},
			[]interface{}{"node-1", "node-2"},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &SinglyLinkedList{
				head: tt.fields.head,
				len:  tt.fields.len,
			}
			l.DeleteFirst()
			if l.len != tt.expectedLen {
				t.Errorf("length mismatch after delete, expected:%d, got:%d", l.len, tt.expectedLen)
			}
			pos, currNode, expectedNode, err := compare(tt.expectedNodes, l.head)
			if err != nil {
				t.Fatalf("missing node at pos %v, expected:%v, got:%v", pos, currNode, expectedNode)
			}
		})
	}
}

func TestLinkedList_DeleteLast(t *testing.T) {
	type fields struct {
		head *SinglyNode
		curr *SinglyNode
		len  int
	}
	tests := []struct {
		name          string
		fields        fields
		expectedNodes []interface{}
		expectedLen   int
	}{
		{
			"There is no element in the list",
			fields{
				head: nil,
				curr: nil,
				len:  0,
			},
			nil,
			0,
		},
		{
			"There is one element in the list",
			fields{
				head: constructList([]interface{}{"node-0"}),
				curr: nil,
				len:  1,
			},
			[]interface{}{},
			0,
		},
		{
			"There is multiple element in the list",
			fields{
				head: constructList([]interface{}{"node-0", "node-1", "node-2"}),
				curr: nil,
				len:  3,
			},
			[]interface{}{"node-0", "node-1"},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &SinglyLinkedList{
				head: tt.fields.head,
				len:  tt.fields.len,
			}
			l.DeleteLast()
			if l.len != tt.expectedLen {
				t.Errorf("length mismatch after delete, expected:%d, got:%d", l.len, tt.expectedLen)
			}
			pos, currNode, expectedNode, err := compare(tt.expectedNodes, l.head)
			if err != nil {
				t.Fatalf("missing node at pos %v, expected:%v, got:%v", pos, currNode, expectedNode)
			}
		})
	}
}

func TestLinkedList_DeleteNode(t *testing.T) {
	type fields struct {
		head *SinglyNode
		curr *SinglyNode
		len  int
	}

	tests := []struct {
		name          string
		fields        fields
		key           interface{}
		expectedNodes []interface{}
		expectedLen   int
	}{
		{
			"There is no element in the list",
			fields{
				head: nil,
				curr: nil,
				len:  0,
			},
			nil,
			nil,
			0,
		},
		{
			"Delete a node from the list having len 1",
			fields{
				head: constructList([]interface{}{"node-0"}),
				curr: nil,
				len:  1,
			},
			"node-0",
			[]interface{}{},
			0,
		},
		{
			"Delete first node from the list",
			fields{
				head: constructList([]interface{}{"node-0", "node-1", "node-2"}),
				curr: nil,
				len:  3,
			},
			"node-0",
			[]interface{}{"node-1", "node-2"},
			2,
		},
		{
			"Delete intermediate node from the list",
			fields{
				head: constructList([]interface{}{"node-0", "node-1", "node-2", "node-3"}),
				curr: nil,
				len:  4,
			},
			"node-2",
			[]interface{}{"node-0", "node-1", "node-3"},
			3,
		},
		{
			"Delete last node from the list",
			fields{
				head: constructList([]interface{}{"node-0", "node-1", "node-2", "node-3"}),
				curr: nil,
				len:  4,
			},
			"node-3",
			[]interface{}{"node-0", "node-1", "node-2"},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &SinglyLinkedList{
				head: tt.fields.head,
				len:  tt.fields.len,
			}
			l.DeleteNode(tt.key)
			if l.len != tt.expectedLen {
				t.Errorf("length mismatch after delete, expected:%d, got:%d", l.len, tt.expectedLen)
			}
			pos, currNode, expectedNode, err := compare(tt.expectedNodes, l.head)
			if err != nil {
				t.Fatalf("missing node at pos %v, expected:%v, got:%v", pos, currNode, expectedNode)
			}
		})
	}
}

func TestLinkedList_InsertFirst(t *testing.T) {
	type fields struct {
		head *SinglyNode
		curr *SinglyNode
		len  int
	}
	tests := []struct {
		name          string
		fields        fields
		node          *SinglyNode
		expectedNodes []interface{}
		expectedLen   int
	}{
		{
			"insert in the empty list",
			fields{
				head: nil,
				curr: nil,
				len:  0,
			},
			constructList([]interface{}{"node-0"}),
			[]interface{}{"node-0"},
			1,
		},
		{
			"insert in the list having multiple elements",
			fields{
				head: constructList([]interface{}{"node-0", "node-1", "node-2"}),
				curr: nil,
				len:  3,
			},
			constructList([]interface{}{"node-4"}),
			[]interface{}{"node-4", "node-0", "node-1", "node-2"},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &SinglyLinkedList{
				head: tt.fields.head,
				len:  tt.fields.len,
			}
			l.InsertFirst(tt.node)
			if l.len != tt.expectedLen {
				t.Errorf("length mismatch after delete, expected:%d, got:%d", l.len, tt.expectedLen)
			}
			pos, currNode, expectedNode, err := compare(tt.expectedNodes, l.head)
			if err != nil {
				t.Fatalf("missing node at pos %v, expected:%v, got:%v", pos, currNode, expectedNode)
			}
		})
	}
}

func TestLinkedList_InsertLast(t *testing.T) {
	type fields struct {
		head *SinglyNode
		curr *SinglyNode
		len  int
	}
	tests := []struct {
		name          string
		fields        fields
		node          *SinglyNode
		expectedNodes []interface{}
		expectedLen   int
	}{
		{
			"insert in the empty list",
			fields{
				head: nil,
				curr: nil,
				len:  0,
			},
			constructList([]interface{}{"node-0"}),
			[]interface{}{"node-0"},
			1,
		},
		{
			"insert in the list having multiple elements",
			fields{
				head: constructList([]interface{}{"node-0", "node-1", "node-2"}),
				curr: nil,
				len:  3,
			},
			constructList([]interface{}{"node-4"}),
			[]interface{}{"node-0", "node-1", "node-2", "node-4"},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &SinglyLinkedList{
				head: tt.fields.head,
				len:  tt.fields.len,
			}
			l.InsertLast(tt.node)
			if l.len != tt.expectedLen {
				t.Errorf("length mismatch after delete, expected:%d, got:%d", l.len, tt.expectedLen)
			}
			pos, currNode, expectedNode, err := compare(tt.expectedNodes, l.head)
			if err != nil {
				t.Fatalf("missing node at pos %v, expected:%v, got:%v", pos, currNode, expectedNode)
			}
		})
	}
}
