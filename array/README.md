# Array & Sequence Patterns

Arrays are the backbone of most interview problems. Mastering sliding windows, prefix sums, and transformation tricks will help you solve everything from subarray sums to combinatorics.

## Interview Focus Areas

- Pattern matching: sliding window, two-pointer sweeps, prefix/suffix scans, and Kadane-style DP.
- Search tricks: binary search on answers, rotated arrays, and peak/valley identification.
- Backtracking on arrays: generate permutations/combinations with pruning and lexicographic order.
- Optimization themes: frequency maps, monotonic structures, and in-place transformations to tighten space.

## Real-World Applications

- Streaming analytics: rolling averages, anomaly detection, longest repeating subsequences.
- Recommendation engines: ranking top-k items, merging sorted feeds, deduplicating streams.
- Scheduling and planning: booking windows, demand forecasting, and resource allocation timelines.

## What's in This Folder

- `kadane_algorithm.go`: classic maximum subarray sum using Kadane's algorithm.
- `backtracking/`: parenthesis generation, permutations, and combinatorial search templates.
- `dp/`: climbing stairs, grid DP, and notes for adapting to path-counting variations.
- `questions/`: interview-style problems like best time to buy/sell stock, kth largest element, and merge intervals.
- `searching/`: binary search variations for sorted, rotated, and near-sorted arrays.
- `slidingwindow/`: longest substring, character replacement, subarray sum, and template README for sliding window theory.
- `sorting/`: comparison sort implementations (quick, merge, insertion) plus binary array sorting tricks.

## Practice Game Plan

1. Drill the core patterns: start with `slidingwindow/` and `searching/` to cement templates.
2. Move to `questions/` and time yourself to simulate interview pressure.
3. Tweak existing implementations (e.g., change constraints or add edge cases) to build intuition.
4. Finish by whiteboarding variants without the codebase open to ensure pattern recall.

## Stretch Topics

- Try ordered statistics (selection algorithms) and binary indexed trees for range queries.
- Combine array patterns with other structures: heap + array for streaming top-k, trie + array for autocomplete.
- Translate Go solutions into another language you use in interviews to double-check understanding.
