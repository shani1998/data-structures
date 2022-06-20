package stack

import (
	"reflect"
	"testing"
)

func TestStack_Push(t *testing.T) {
	var testStack Stack
	tests := []struct {
		name           string
		item           any
		desireStackLen int
	}{
		{
			name:           "Insert integer into stack",
			item:           23,
			desireStackLen: 1,
		},
		{
			name:           "Insert string into stack",
			item:           "i am testing",
			desireStackLen: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testStack.Push(tt.item)
			if testStack.Length() != tt.desireStackLen {
				t.Errorf("expected stack of length %d, got length %d\n", tt.desireStackLen, testStack.Length())
			}
		})
	}
}

func TestStack_Pop(t *testing.T) {
	tests := []struct {
		name           string
		stack          Stack
		want           any
		wantErr        bool
		desireStackLen int
	}{
		{
			name:           "Failure:empty stack should return error",
			stack:          nil,
			want:           nil,
			wantErr:        true,
			desireStackLen: 0,
		},
		{
			name:           "Success: Pop from stack having data and test length",
			stack:          []any{"Hi", 123},
			want:           123,
			wantErr:        false,
			desireStackLen: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.stack.Pop()
			if (err != nil) != tt.wantErr {
				t.Errorf("Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pop() got = %v, want %v", got, tt.want)
			}
			if tt.stack.Length() != tt.desireStackLen {
				t.Errorf("expected stack of length %d, got length %d\n", tt.desireStackLen, tt.stack.Length())
			}
		})
	}
}

func TestStack_Length(t *testing.T) {
	tests := []struct {
		name  string
		stack Stack
		want  int
	}{
		{
			name:  "empty stack",
			stack: nil,
			want:  0,
		},
		{
			name:  "non empty data stack",
			stack: []any{"Hi", "How are you ?"},
			want:  2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.stack.Length(); got != tt.want {
				t.Errorf("Length() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_IsEmpty(t *testing.T) {
	tests := []struct {
		name  string
		stack Stack
		want  bool
	}{
		{
			name:  "empty stack",
			stack: nil,
			want:  true,
		},
		{
			name:  "non empty data stack",
			stack: []any{"Hi", "How are you ?"},
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.stack.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
