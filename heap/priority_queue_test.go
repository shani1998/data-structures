package heap

import (
	"container/heap"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	tests := []struct {
		name       string
		operations []struct {
			value    string
			priority int
		}
		expectedOutput []*Item
	}{
		{
			name: "Retrieve in Order of Priority",
			operations: []struct {
				value    string
				priority int
			}{
				{"operation-name-2", 2},
				{"operation-name-4", 2},
				{"operation-name-3", 4},
				{"operation-name-1", 3},
			},
			expectedOutput: []*Item{
				{value: "operation-name-2", priority: 2},
				{value: "operation-name-4", priority: 2},
				{value: "operation-name-1", priority: 3},
				{value: "operation-name-3", priority: 4},
			},
		},
		{
			name: "Mixed Priorities",
			operations: []struct {
				value    string
				priority int
			}{
				{"operation-name-5", 5},
				{"operation-name-6", 1},
			},
			expectedOutput: []*Item{
				{value: "operation-name-6", priority: 1}, // Lower priority first
				{value: "operation-name-5", priority: 5},
			},
		},
		{
			name: "Test Empty Queue Pop",
			operations: []struct {
				value    string
				priority int
			}{},
			expectedOutput: nil, // Expecting nil for an empty queue
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testPQ := NewPriorityQueue()

			for _, operation := range tt.operations {
				heap.Push(testPQ, &Item{
					value:    operation.value,
					priority: operation.priority,
				})
			}

			var actualOutput []*Item
			for testPQ.Len() > 0 {
				actualOutput = append(actualOutput, heap.Pop(testPQ).(*Item))
			}

			// Check if the output matches expected
			if len(actualOutput) != len(tt.expectedOutput) {
				t.Errorf("Expected %d items, got %d", len(tt.expectedOutput), len(actualOutput))
			}

			for i := range actualOutput {
				if actualOutput[i].value != tt.expectedOutput[i].value || actualOutput[i].priority != tt.expectedOutput[i].priority {
					t.Errorf("Expected %+v, got %+v", tt.expectedOutput[i], actualOutput[i])
				}
			}
		})
	}
}
