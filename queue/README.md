# Queue Patterns & Scheduling

Queues model FIFO flows. Interviews use them for BFS traversals, rate limiting, and buffering problems.

## Interview Focus Areas

- Implementations: array-based circular buffers vs linked-list-backed queues.
- Variants: double-ended queues, monotonic queues, and priority queues for scheduler questions.
- BFS scaffolding: level-order traversals in trees/graphs and multi-source propagation.
- Concurrency considerations: producer-consumer coordination and blocking queues.

## Real-World Applications

- Task schedulers: print queues, job dispatchers, and message brokers.
- Rate limiting: API request throttling and sliding-window counters.
- UI interactions: undo/redo pipelines, command buffers, and event loops.

## What's in This Folder

- `queue.go`: idiomatic Go queue implementation with enqueue/dequeue helpers.
- `queue_test.go`: unit tests covering underflow, overflow, and mixed operation sequences.

## Practice Game Plan

1. Extend the queue to a double-ended version and mirror the tests.
2. Implement a bounded blocking queue using channels to simulate concurrency constraints.
3. Use the queue in `tree/` BFS functions to reinforce cross-folder practice.
4. Solve classic problems like "rotting oranges" or "minimum steps to reach end" using this queue.

## Stretch Topics

- Explore monotonic queues for sliding window maxima/minima problems.
- Build a rate limiter middleware that uses fixed-window or token-bucket logic.
- Combine with `heap/` to simulate priority scheduling in operating systems.
