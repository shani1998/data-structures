# Hashing & Frequency Patterns

Hash tables turn complex constraints into O(1) lookups. Interviews lean on them for deduplication, counting, and grouping problems.

## Interview Focus Areas

- Constant-time lookups: set membership, two-sum families, and amortized performance caveats.
- Frequency analysis: histograms, sliding windows with hash maps, and anagram grouping.
- Order preservation: leveraging ordered maps or auxiliary slices when iteration order matters.
- Collision handling: Go map internals, load factor intuition, and when to reach for custom hashing.

## Real-World Applications

- Caching layers: memoization, LRU/LFU caches, and API throttling counters.
- Text analytics: detecting longest unique substrings, keyword counts, and log aggregation.
- Distributed systems: consistent hashing for sharding, leader election bookkeeping.

## What's in This Folder

- `basics_map.go`: primer on Go map usage and idiomatic patterns.
- `longest-consecutive-sequence.go`: hash set technique for O(n) sequence detection.
- `longest-non-repeating-string.go`: sliding window with character counts.
- `top-k-frequent-elements.go`: frequency buckets and heap combo for top-k retrieval.

## Practice Game Plan

1. Rewrite classic two-sum, four-sum, and anagram checks using the patterns here.
2. Extend `top-k-frequent-elements.go` to support streaming input with a min-heap.
3. Inject intentional collisions (custom structs as keys) to understand equality semantics.
4. Combine hashing with rolling techniques (e.g., Rabin-Karp) for substring search.

## Stretch Topics

- Implement an LRU cache in Go using maps plus doubly linked lists from `linkedlist/`.
- Explore probabilistic structures like Bloom filters and Count-Min Sketch.
- Simulate distributed hash tables to understand partitioning and fault tolerance.
