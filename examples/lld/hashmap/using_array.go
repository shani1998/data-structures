package main

import (
	"fmt"
)

// FastSet implements an O(1) integer set for values in [0, N-1].
// It uses a dense-sparse array with timestamping for O(1) clear operations.
type FastSet struct {
	size       int   // number of elements currently in the set
	values     []int // dense array of current elements
	indices    []int // indices[x] = position of x in values[]
	timestamps []int // timestamps[x] = version when x was added
	timestamp  int   // current version number (increments on Clear)
	capacity   int   // maximum N
}

// NewFastSet initializes a new FastSet with capacity N.
func NewFastSet(N int) *FastSet {
	return &FastSet{
		values:     make([]int, N),
		indices:    make([]int, N),
		timestamps: make([]int, N),
		timestamp:  1,
		capacity:   N,
	}
}

// Insert adds x to the set if it’s not already present.
func (fs *FastSet) Insert(x int) {
	if x < 0 || x >= fs.capacity {
		return // out of bounds, ignore
	}
	if fs.timestamps[x] != fs.timestamp {
		fs.timestamps[x] = fs.timestamp
		fs.indices[x] = fs.size
		fs.values[fs.size] = x
		fs.size++
	}
}

// Remove deletes x from the set if it’s present.
func (fs *FastSet) Remove(x int) {
	if x < 0 || x >= fs.capacity {
		return
	}
	if fs.timestamps[x] != fs.timestamp {
		return // x not present
	}

	// Swap with last element in dense array
	idx := fs.indices[x]
	lastIdx := fs.size - 1
	lastVal := fs.values[lastIdx]

	fs.values[idx] = lastVal
	fs.indices[lastVal] = idx
	fs.size--

	fs.timestamps[x] = 0 // optional: mark stale
}

// Contains checks if x is in the set.
func (fs *FastSet) Contains(x int) bool {
	if x < 0 || x >= fs.capacity {
		return false
	}
	return fs.timestamps[x] == fs.timestamp
}

// Clear empties the set in O(1) time.
func (fs *FastSet) Clear() {
	fs.size = 0
	fs.timestamp++
	if fs.timestamp == 0 { // handle overflow wrap-around
		for i := range fs.timestamps {
			fs.timestamps[i] = 0
		}
		fs.timestamp = 1
	}
}

// Values returns a slice of all elements in the set (read-only view).
func (fs *FastSet) Values() []int {
	return fs.values[:fs.size]
}

// Size returns the number of elements in the set.
func (fs *FastSet) Size() int {
	return fs.size
}

// DebugPrint prints internal state for debugging (optional)
func (fs *FastSet) DebugPrint() {
	fmt.Printf("values: %v\n", fs.values[:fs.size])
	fmt.Printf("indices: %v\n", fs.indices[:fs.capacity])
	fmt.Printf("timestamps: %v (current=%d)\n", fs.timestamps, fs.timestamp)
}

func main() {
	N := 10
	set := NewFastSet(N)

	set.Insert(3)
	set.Insert(5)
	set.Insert(7)
	fmt.Println("After inserts:", set.Values()) // [3 5 7]

	fmt.Println("Contains 5?", set.Contains(5)) // true
	fmt.Println("Contains 2?", set.Contains(2)) // false

	set.Remove(5)
	fmt.Println("After removing 5:", set.Values()) // [3 7]

	set.Insert(2)
	fmt.Println("After inserting 2:", set.Values()) // [3 7 2]

	set.Clear()
	fmt.Println("After clearing:", set.Values()) // []

	set.Insert(1)
	set.Insert(9)
	fmt.Println("After inserting 1 and 9:", set.Values()) // [1 9]
	fmt.Println("Contains 1?", set.Contains(1))           // true
	fmt.Println("Contains 9?", set.Contains(9))           // true
}
