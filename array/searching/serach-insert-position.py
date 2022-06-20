'''Given a sorted array and a target value, return the index if the target is found.
If not, return the index where it would be if it were inserted in order.'''


def searchInsert(self, nums: List[int], target: int) -> int:
    def helper(nums, low, high):
        if low > high:
            return low
        mid = (low + high) // 2
        if nums[mid] == target:
            return mid
        if nums[mid] > target:
            return helper(nums, low, mid - 1)
        else:
            return helper(nums, mid + 1, high)

    return (helper(nums, 0, len(nums) - 1))
