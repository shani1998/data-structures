package main

import "fmt"

// DenseSparseSet implements a set for integers in range [0, N-1]
type DenseSparseSet struct {
	N      int
	size   int
	dense  []int // keep elements in insertion order
	sparse []int // maps element to its index in dense
	inSet  []bool
}

// NewDenseSparseSet creates a new empty set of capacity N
func NewDenseSparseSet(N int) *DenseSparseSet {
	return &DenseSparseSet{
		N:      N,
		dense:  make([]int, N),
		sparse: make([]int, N),
		inSet:  make([]bool, N),
	}
}

// Insert adds x to the set if not already present
func (s *DenseSparseSet) Insert(x int) {
	if x < 0 || x >= s.N {
		fmt.Printf("Value %d out of range\n", x)
		return
	}
	if s.inSet[x] {
		return // already in set
	}
	s.dense[s.size] = x
	s.sparse[x] = s.size
	s.inSet[x] = true
	s.size++
}

// Remove deletes x from the set if it exists
func (s *DenseSparseSet) Remove(x int) {
	if x < 0 || x >= s.N {
		fmt.Printf("Value %d out of range\n", x)
		return
	}
	if !s.inSet[x] {
		return // not in set
	}

	// Swap with last element to keep dense packed
	last := s.dense[s.size-1]
	idx := s.sparse[x]

	s.dense[idx] = last
	s.sparse[last] = idx

	s.size--
	s.inSet[x] = false
}

// Contains checks if x is in the set
func (s *DenseSparseSet) Contains(x int) bool {
	if x < 0 || x >= s.N {
		return false
	}
	return s.inSet[x]
}

// Clear removes all elements (O(N))
func (s *DenseSparseSet) Clear() {
	for i := 0; i < s.size; i++ {
		s.inSet[s.dense[i]] = false
	}
	s.size = 0
}

// Iterate returns all elements currently in the set
func (s *DenseSparseSet) Iterate() []int {
	return s.dense[:s.size]
}

func main() {
	set := NewDenseSparseSet(10)

	set.Insert(3)
	set.Insert(5)
	set.Insert(7)
	fmt.Println("After inserts:", set.Iterate()) // [3 5 7]

	set.Insert(5)
	fmt.Println("After inserting 5 again:", set.Iterate()) // [3 5 7]

	fmt.Println("Contains 5?", set.Contains(5)) // true
	fmt.Println("Contains 9?", set.Contains(9)) // false

	set.Remove(5)
	fmt.Println("After removing 5:", set.Iterate()) // [3 7]

	set.Insert(2)
	fmt.Println("After inserting 2:", set.Iterate()) // [3 7 2]

	set.Clear()
	fmt.Println("After clearing:", set.Iterate()) // []

	set.Insert(9)
	set.Insert(1)
	fmt.Println("After inserting 9,1:", set.Iterate()) // [9 1]
}

// Implment all set functionality using dense-sparse technique in O(1) time complexity for insert, remove, and lookup, clear operations.
// where x is in the range [0, N)
type Set struct {
	dense  []int
	sparse []int
	size   int
}

func NewSet(N int) *Set {
	return &Set{
		dense:  make([]int, N),
		sparse: make([]int, N),
	}
}

func (s *Set) Insert(x int) {
	if s.lookup(x) {
		return
	}
	s.dense[s.size] = x
	s.sparse[x] = s.size
	s.size++
}

func (s *Set) Remove(x int) {
	if !s.lookup(x) {
		return
	}
	last := s.dense[s.size-1]
	idx := s.sparse[x]

	s.dense[idx] = last
	s.sparse[last] = idx

	s.size--
}

func (s *Set) lookup(x int) bool {
	idx := s.sparse[x]
	if idx >= s.size {
		return false
	}
	return s.dense[idx] == x
}

func (s *Set) Clear() {
	s.size = 0
}

func (s *Set) Iterate() []int {
	return s.dense[:s.size]
}