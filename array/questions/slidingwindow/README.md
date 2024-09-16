# Sliding Window Problem Pattern

The sliding window technique is a powerful algorithmic approach used to solve problems involving subarrays or substrings within a given array or string. This pattern is particularly useful when the problem requires finding the maximum or minimum subarray/substring that satisfies a given condition.

## How to Identify Sliding Window Problems

Sliding window problems typically have the following characteristics:
- The problem involves an array or string.
- You need to find a subarray or substring that satisfies a particular condition (e.g., maximum sum, minimum length, contains specific characters, etc.).
- The subarray/substring's size or its contents can change dynamically as you iterate through the array/string.

Common problem types include:
- Finding the maximum or minimum sum of a subarray of a fixed size.
- Longest substring with at most `k` distinct characters.
- Smallest substring containing all characters of another string.

## General Steps for Solving Sliding Window Problems

1. **Define the Window**:
    - Identify what constitutes the "window" in your problem. This is the subarray or substring you are looking to optimize.

2. **Initialize Pointers**:
    - Use two pointers (usually called `left` and `right`) to represent the window's boundaries.

3. **Expand the Window**:
    - Move the `right` pointer to expand the window until the window contains the desired elements or satisfies the condition.

4. **Contract the Window**:
    - Once the window satisfies the condition, move the `left` pointer to minimize the window size while still satisfying the condition.

5. **Track the Optimal Solution**:
    - Keep track of the best (maximum/minimum) solution encountered during the window's expansion and contraction.

6. **Edge Cases**:
    - Handle edge cases, such as empty arrays or strings, and ensure your solution works for all possible inputs.
### References:
- https://medium.com/@timpark0807/leetcode-is-easy-sliding-window-c44c11cc33e1
- https://medium.com/leetcode-patterns/leetcode-pattern-2-sliding-windows-for-strings-e19af105316b