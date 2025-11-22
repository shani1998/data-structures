# Linked List Systems

Linked lists train pointer discipline and are a gateway to more complex structures (LRU caches, skip lists, adjacency lists).

## Interview Focus Areas

- Pointer manipulation: reversing, swapping pairs, merging sorted lists, and detecting cycles.
- Variants: singly, doubly, and circular lists with their tradeoffs and use cases.
- Slow-fast pointer tricks: cycle detection, middle node discovery, palindrome checks.
- Memory considerations: dummy head strategies, in-place modifications, and tail handling.

## Real-World Applications

- Navigation history: browser back/forward stacks, music playlists, undo-redo systems.
- Cache design: implementing LRU/LFU caches that require O(1) insert/delete from arbitrary positions.
- Memory allocators: free lists for custom allocators and OS memory management.

## What's in This Folder

- `singly_linkedlist.go`: foundational node definitions and helpers.
- `singly_linkedlist_problems.go`: interview staples like reversal, middle element, and cycle detection.
- `singly_linkedlist_test.go`: tests covering edge cases and pointer correctness.
- `doubly_linked_list.go`: bidirectional list implementation for cache-style problems.
- `questions/`: targeted exercises including Floyd's cycle detection and list manipulations.
- `reverse_list.go`: focused implementation for iterative/recursive reversal patterns.

## Practice Game Plan

1. Re-create the core functions without looking, then diff against the existing code.
2. Implement copy-with-random-pointer and compare to typical interview expectations.
3. Integrate with `hashing/` to build an LRU cache; write benchmarks to validate O(1) operations.
4. Adapt solutions to handle gigantic lists (simulate streaming) to highlight memory patterns.

## Stretch Topics

- Explore skip lists and compare to balanced BST performance.
- Implement a lock-free linked list using Go's atomics for concurrency practice.
- Pair this knowledge with `tree/` BST rotations to understand structural transformations.
