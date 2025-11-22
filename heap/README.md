# Heaps & Priority Queues

Heaps power scheduling, streaming analytics, and greedy optimization. Interviewers love them for k-selection and merging tasks.

## Interview Focus Areas

- Heap mechanics: binary heap properties, push/pop operations, and zero-based indexing math.
- Priority queue tricks: merging k-sorted arrays, streaming medians, and task scheduling.
- Adapters: using Go's `container/heap` interface vs building bespoke array-backed heaps.
- Complexity awareness: when heapifying beats incremental insertion, and how to tune space usage.

## Real-World Applications

- Job schedulers: CPU task ordering, throttled background work, and rate limiting.
- Recommendation feeds: merging sorted ranking lists or selecting top-k items in a stream.
- Graph algorithms: Dijkstra, Prim, and A* rely on priority queues for frontier management.

## What's in This Folder

- `max_heap_using_array.go`: manual array-backed max heap with explanatory operations.
- `priority_queue.go`: Go `container/heap` wrapper tailored for interview-style usage.
- `priority_queue_test.go`: unit tests showing push, pop, update, and custom comparator patterns.

## Practice Game Plan

1. Implement a min-heap variant and mirror the tests to ensure confidence both ways.
2. Extend the priority queue with decrease-key logic for Dijkstra or Prim implementations.
3. Combine heaps with hashing from `hashing/` to implement an LRU or LFU cache.
4. Practice k-way merge and running median problems to reinforce core patterns.

## Stretch Topics

- Explore indexed priority queues and Fibonacci heaps for theoretical depth.
- Benchmark heap operations vs sorting for various data sizes to internalize trade-offs.
- Implement a scheduler simulator: push tasks with deadlines and prioritize based on slack.
