package queue

import (
	"reflect"
	"testing"
)

func TestQueue_Length(t *testing.T) {
	type fields struct {
		capacity int
		items    chan any
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queue{
				capacity: tt.fields.capacity,
				items:    tt.fields.items,
			}
			if got := q.Length(); got != tt.want {
				t.Errorf("Length() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Pop(t *testing.T) {
	type fields struct {
		capacity int
		items    chan any
	}
	tests := []struct {
		name    string
		fields  fields
		want    any
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queue{
				capacity: tt.fields.capacity,
				items:    tt.fields.items,
			}
			got, err := q.Pop()
			if (err != nil) != tt.wantErr {
				t.Errorf("Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pop() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Push(t *testing.T) {
	type fields struct {
		capacity int
		items    chan any
	}
	type args struct {
		val any
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queue{
				capacity: tt.fields.capacity,
				items:    tt.fields.items,
			}
			if err := q.Push(tt.args.val); (err != nil) != tt.wantErr {
				t.Errorf("Push() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
