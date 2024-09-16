package recursion

// SumArray calculates the sum of all elements in the array.
func SumArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return nums[0] + SumArray(nums[1:])
}

// Recursions tree for the SumArray function:
/*
SumArray([]int{1, 2, 3, 4, 5})
│
├── 1 + SumArray([]int{2, 3, 4, 5})
│    ├── 2 + SumArray([]int{3, 4, 5})
│    │    ├── 3 + SumArray([]int{4, 5})
│    │    │    ├── 4 + SumArray([]int{5})
│    │    │    │    ├── 5 + SumArray([]int{})
│    │    │    │    │    └── return 0
│    │    │    │    └── return 5
│    │    │    └── return 9
│    │    └── return 12
│    └── return 14
└── return 15
*/
