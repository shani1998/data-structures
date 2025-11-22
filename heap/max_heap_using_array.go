package heap

import "fmt"

/*
---------------------------------------------------------
                Max Heap (Array Representation)
---------------------------------------------------------

A Max-Heap is a complete binary tree where:
    - Every parent node is >= its children
    - The largest value is always at the root (index 0)

We store the heap in an array using index formulas:

    parent(i) = (i - 1) / 2
    left(i)   = 2*i + 1
    right(i)  = 2*i + 2

Example array heap:
    [15, 10, 8, 7, 2]

Represents the tree:
                  15
            10          8
         7     2

---------------------------------------------------------
*/

type MaxHeap struct {
	data []int // underlying array storing heap elements
}

/* Index helper functions */
func parent(i int) int { return (i - 1) / 2 }
func left(i int) int   { return 2*i + 1 }
func right(i int) int  { return 2*i + 2 }

/*
---------------------------------------------------------
 Insert(x)
---------------------------------------------------------

Adds a new number to the heap.

Steps:
    1. Append the new value at the end of the array.
    2. "Bubble up" → while child > parent, swap.

Result:
    Heap property restored.
*/

func (h *MaxHeap) Insert(value int) {
	h.data = append(h.data, value)
	h.bubbleUp(len(h.data) - 1)
}

/*
---------------------------------------------------------
 ExtractMax()
---------------------------------------------------------

Removes and returns the root element (largest value).

Steps:
    1. Save the root (max value)
    2. Replace root with last element in array
    3. Shrink array by removing last element
    4. Bubble down from root to restore heap structure

Time: O(log n)
*/

func (h *MaxHeap) ExtractMax() (int, bool) {
	if len(h.data) == 0 {
		return 0, false
	}

	maxVal := h.data[0]

	last := h.data[len(h.data)-1]
	h.data[0] = last
	h.data = h.data[:len(h.data)-1]

	if len(h.data) > 0 {
		h.bubbleDown(0)
	}

	return maxVal, true
}

/*
---------------------------------------------------------
 Peek()
---------------------------------------------------------

Returns the largest element without removing it.

Time: O(1)
*/

func (h *MaxHeap) Peek() (int, bool) {
	if len(h.data) == 0 {
		return 0, false
	}
	return h.data[0], true
}

/*
---------------------------------------------------------

	bubbleUp(index)

---------------------------------------------------------

Fixes heap by comparing a child with its parent and
swapping upward until heap property is restored.

Condition:

	If child > parent → swap.

Used after Insert().
*/
func (h *MaxHeap) bubbleUp(i int) {
	for i > 0 {
		p := parent(i)
		if h.data[i] <= h.data[p] {
			break
		}
		h.data[i], h.data[p] = h.data[p], h.data[i]
		i = p
	}
}

/*
---------------------------------------------------------

	bubbleDown(index)

---------------------------------------------------------

Fixes heap by comparing a parent with its children and
swapping downward with the larger child until heap
property is restored.

Used after ExtractMax().
*/
func (h *MaxHeap) bubbleDown(i int) {
	n := len(h.data)

	for {
		l := left(i)
		r := right(i)
		largest := i

		if l < n && h.data[l] > h.data[largest] {
			largest = l
		}
		if r < n && h.data[r] > h.data[largest] {
			largest = r
		}

		if largest == i {
			break
		}

		h.data[i], h.data[largest] = h.data[largest], h.data[i]
		i = largest
	}
}

/*
---------------------------------------------------------
 NewMaxHeap(arr)
---------------------------------------------------------

Builds a Max-Heap from an existing array in O(n).

Approach:
    - Start from last non-leaf node
    - Call bubbleDown() on each node

This efficiently heapifies the array bottom-up.
*/

func NewMaxHeap(arr []int) *MaxHeap {
	h := &MaxHeap{data: arr}

	for i := parent(len(arr) - 1); i >= 0; i-- {
		h.bubbleDown(i)
	}

	return h
}

/*
---------------------------------------------------------

	MAIN: Example Usage

---------------------------------------------------------
*/
func main() {

	heap := NewMaxHeap([]int{3, 10, 5, 6, 2, 8, 1})
	fmt.Println("Initial Max Heap:", heap.data)

	heap.Insert(15)
	fmt.Println("After inserting 15:", heap.data)

	maxVal, _ := heap.ExtractMax()
	fmt.Println("Extracted Max:", maxVal)
	fmt.Println("Heap after extract:", heap.data)

	peek, _ := heap.Peek()
	fmt.Println("Current Max:", peek)
}
