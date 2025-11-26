# DSA Interview Cheatsheet

Fast revision for the final week. For each problem theme below, rehearse every approach (brute-force to optimal) and revisit the linked implementation.

## Quick Commands

- `go test ./...` runs every unit test; use after practice refactors.
- `make test` mirrors the above; keep it muscle memory before mocks.
- `make fmt` keeps Go code clean before committing new variants.

## Arrays & Strings (`array/`)

| Theme | Approaches to Drill | Repo Anchor |
| --- | --- | --- |
| Longest substring without repeats | 1) O(n^2) brute window + set, 2) Sliding window with set, 3) Sliding window with map of last seen index | `slidingwindow/longest_substring.go` |
| Subarray sum equals target | 1) Prefix sums with nested loop, 2) Hash map storing prefix counts, 3) Sliding window for non-negative arrays | `slidingwindow/subarray_sum.go` |
| Max subarray / stock profit | 1) Enumerate all subarrays, 2) Dynamic running sum (Kadane), 3) Track min-prefix for profit | `kadane_algorithm.go`, `questions/best_time_to_buy_sell_stock.go` |
| Search in rotated array | 1) Linear scan, 2) Modified binary search, 3) Binary search via pivot detection | `searching/search_in_sorted_rotated_array.go` |
| Sorting classics | 1) Bubble/insertion for intuition, 2) Merge sort divide-and-conquer, 3) Quick sort partition variations | `sorting/` |

Supplement: practise two-pointer tricks on `questions/merge_intervals.go` and `questions/min_diff.go`.

## Backtracking & Recursion (`array/backtracking/`)

| Theme | Approach Ladder | Repo Anchor |
| --- | --- | --- |
| Parentheses generation | 1) Generate all strings + validate, 2) DFS with counts, 3) Catalan DP reasoning for counts | `generate_parentheses.go` |
| Permutations | 1) Use map of remaining elements, 2) Swap in-place with recursion, 3) Next permutation iterative variation | `generate_permutations.go` |
| Kth lexicographic string | 1) Generate all, 2) DFS with pruning using factorial counts, 3) Direct combinatorics math | `kth_lexico_string.go` |

## Dynamic Programming (`array/dp/`)

| Theme | Approach Ladder | Repo Anchor |
| --- | --- | --- |
| Climbing stairs / Fibonacci | 1) Naive recursion, 2) Memoization, 3) Tabulation, 4) O(1) space iteration | `1_climbing_stairs.go` |
| Grid paths / obstacle DP | 1) DFS with memo, 2) 2D table, 3) 1D rolling array, 4) Combinatorics shortcut (if applicable) | `grid/` |
| Kadane-style DP | 1) Naive enumeration, 2) Rolling max ending, 3) Track global + local best | `kadane_algorithm.go` |

Keep the `array/dp/README.md` open for transitions and optimization reminders.

## Graphs (`graph/`)

| Theme | Approaches | Repo Anchor |
| --- | --- | --- |
| Connected components / islands | 1) Recursive DFS, 2) Iterative DFS with stack, 3) BFS queue, 4) Disjoint set union | `questions/number_of_islands.go`, `bfs.go`, `dfs.go` |
| Reorder routes / topological ideas | 1) Traverse with parent tracking, 2) Build adjacency and count wrong directions, 3) Convert to topological ordering | `questions/city_edge_reordering.go` |
| Cycle detection | 1) DFS with visited stack, 2) BFS Kahn's algorithm, 3) Union-find detection | `questions/cycle_detection.go` |
| Shortest path in unweighted graph | 1) DFS for small graphs, 2) BFS for minimum edges, 3) Multi-source BFS variant | `bfs.go` |

Add weighted variants later with Dijkstra (pair with `heap/`).

## Hashing & Maps (`hashing/`)

| Theme | Approaches | Repo Anchor |
| --- | --- | --- |
| Longest consecutive sequence | 1) Sort and scan, 2) Hash set marking starts, 3) Union-find segments | `longest-consecutive-sequence.go` |
| Longest non-repeating string | 1) Sliding window with set, 2) Map storing last seen, 3) Fixed-size array for ASCII | `longest-non-repeating-string.go` |
| Top-k frequent elements | 1) Sort by frequency, 2) Min-heap of size k, 3) Bucket sort by count | `top-k-frequent-elements.go`, `heap/priority_queue.go` |

## Heaps & Priority Queues (`heap/`)

| Theme | Approach Options | Repo Anchor |
| --- | --- | --- |
| Kth largest / smallest | 1) Sort entire array, 2) Heap of size k, 3) Quickselect | `max_heap_using_array.go`, `priority_queue.go`, `array/questions/kth_largest_element.go` |
| Merge k sorted lists | 1) Sequential merge, 2) Divide and conquer, 3) Min-heap of heads | `priority_queue.go`, `linkedlist/` |
| Streaming median | 1) Re-sort slice per insert, 2) Two-heaps (max + min), 3) Order statistics tree (theory) | implement via `priority_queue.go` |

## Linked Lists (`linkedlist/`)

| Theme | Approach Ladder | Repo Anchor |
| --- | --- | --- |
| Reverse list | 1) Stack based, 2) Iterative pointer flip, 3) Recursive | `reverse_list.go` |
| Detect cycle | 1) Hash set of visited nodes, 2) Floyd slow-fast pointers, 3) Modify pointers (interview caveat) | `singly_linkedlist_problems.go`, `questions/cycle_detection_test.go` |
| Merge two sorted lists | 1) Create new list, 2) In-place pointer weaving, 3) Divide and conquer for k lists | `singly_linkedlist_problems.go` |

## Stacks & Queues (`stack/`, `queue/`)

| Theme | Approaches | Repo Anchor |
| --- | --- | --- |
| Valid parentheses | 1) Replace pairs iteratively, 2) Stack push/pop, 3) Counter for single type, 4) Two-stack for wildcard cases | `stack/questions/` |
| Min stack | 1) Stack of pairs (value, min), 2) Two stacks (values + mins), 3) Encode min via difference trick | `stack.go` |
| Queue via stacks | 1) Two-stack approach, 2) Recursion within stack | `queue/` + exercises |
| Sliding window maximum | 1) Heap per window, 2) Monotonic deque, 3) Block precomputation (sparse table) | `array/slidingwindow/` |

## Trees (`tree/`)

| Theme | Approaches | Repo Anchor |
| --- | --- | --- |
| Traversals | 1) Recursive, 2) Iterative with stack, 3) Morris traversal (inorder) | `binary_tree.go`, `dfs.go`, `bfs.go` |
| Lowest common ancestor | 1) Path tracing, 2) Recursive post-order, 3) Binary lifting (advanced) | `questions/` (add if missing) |
| Diameter / longest path | 1) BFS from every node (inefficient), 2) DFS returning height & diameter, 3) Double BFS (trees) | `questions/max_path_arr.go`, `tree/questions/` |
| Segment tree range queries | 1) Prefix sums for static, 2) Segment tree with lazy propagation, 3) Fenwick tree (add) | `segmenttree/` |

## Tries & Strings (`trie/`)

| Theme | Approaches | Repo Anchor |
| --- | --- | --- |
| Word search / autocomplete | 1) Brute force check all words, 2) Trie with DFS, 3) Compressed trie / DAWG (theory) | `Trie.go`, `questions/` |
| Longest common prefix | 1) Sort and compare ends, 2) Vertical scanning, 3) Trie depth traversal | `longest_common_prefix.go` |

## Mixed Patterns / Bonus

- Merge intervals (`array/questions/merge_intervals.go`): sort + merge, sweep line, difference array.
- Increasing triplet subsequence (`array/questions/increasing_triplet_subsequence.go`): O(n^3) brute, patience sorting style O(n log n), two-pass O(n).
- Graph + heap: shortest path variations (extend with Dijkstra using `heap/priority_queue.go`).

## Practice Loop

1. Select a theme, implement the brute approach from memory, then refactor to optimal.
2. Timebox: 10 mins for brute outline, 20 mins for optimal, 5 mins for tests (`go test ./...`).
3. Write down pitfalls encountered and revisit 48 hours later for spaced repetition.

## Last-Minute Reminder

- For each archetype above, keep at least two independent solutions ready (e.g., DP vs greedy) so you can pivot during interviews.
- Recite time/space for each approach; note when trade-offs flip (e.g., bucket sort beats heap).
- Maintain calm: narrate problem restatement, outline plan, code incrementally, run through sample inputs.
