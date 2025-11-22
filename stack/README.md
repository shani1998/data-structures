# Stack Techniques & Recursion Tools

Stacks capture last-in-first-out behavior, powering expression parsing, backtracking, and recursion elimination.

## Interview Focus Areas

- Expression handling: parentheses validation, infix-to-postfix conversion, and calculator problems.
- Monotonic stacks: next greater element, largest rectangle in histogram, daily temperatures.
- Recursion to iteration: simulating call stacks for tree traversals or DFS.
- Stack + auxiliary structure combos: min/max stacks, two-stack queues, and undo-redo pipes.

## Real-World Applications

- Text editors: undo histories, brace matching, and syntax highlighting.
- Compilers: parsing expressions, managing function calls, and symbol tables.
- Browser navigation: back/forward stacks and tab session restoration.

## What's in This Folder

- `stack.go`: generic LIFO implementation using slices.
- `stack_test.go`: edge case coverage for empty pops, pushes, and peek semantics.
- `questions/`: applied problems such as balanced parentheses and other stack-based puzzles.

## Practice Game Plan

1. Implement a min-stack and add tests to maintain O(1) min retrieval.
2. Recreate tree traversals (`tree/dfs.go`) iteratively using the stack for practice.
3. Solve canonical problems like "largest rectangle in histogram" and compare with monotonic queue approaches.
4. Build a simple expression evaluator that supports +, -, *, /, and parentheses.

## Stretch Topics

- Investigate stack-overflow vs heap allocation behaviors in Go to understand recursion limits.
- Combine with `queue/` to implement a deque supporting amortized O(1) operations.
- Explore concurrency-safe stack patterns using channels and mutexes.
