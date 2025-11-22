# Go DSA Interview Prep Toolkit

A curated collection of data structures, algorithms, and problem solutions implemented in Go to help you prepare confidently for coding interviews.

## Why This Repo

- Practical, bite-sized implementations with clear naming and idiomatic Go code.
- Topic-driven layout (arrays, graphs, dynamic programming, heaps, etc.) that mirrors common interview patterns.
- Reference tests and examples so you can validate your understanding as you go.
- Ready-to-run Go modules and a Makefile for streamlined practice sessions.

## Getting Started

### Requirements

- Go 1.20 or newer
- Make (optional, for convenience targets)

### Clone and Explore

```bash
git clone https://github.com/shani1998/data-structures.git
cd data-structures
```

### Run Everything

```bash
go test ./...
# or
make test
```

### Format the Code

```bash
make fmt
```

## Directory Overview

| Folder | What you will find |
| --- | --- |
| `array/` | Sliding window, backtracking, dynamic programming, and classic array questions. |
| `graph/` | Traversal templates (BFS/DFS), adjacency representations, and graph interview problems. |
| `hashing/` | Hash map patterns, frequency counting, and common interview puzzles. |
| `heap/` | Priority queue implementation and a max-heap backed by arrays. |
| `linkedlist/` | Singly and doubly linked list utilities, plus test-backed problems. |
| `queue/`, `stack/` | Core abstract data type implementations with unit tests. |
| `tree/`, `trie/` | Binary tree traversals, BST helpers, segment tree snippets, and trie utilities. |
| `examples/` | Small reference programs demonstrating how to wire the data structures together. |

Use the directory names as search anchors during practice sessions. Each subfolder mirrors a pattern you are likely to meet in interviews.

## Documentation Map

- `array/` - [Array & Sequence Patterns](array/README.md)
- `graph/` - [Graph Algorithms Playbook](graph/README.md)
- `hashing/` - [Hashing & Frequency Patterns](hashing/README.md)
- `heap/` - [Heaps & Priority Queues](heap/README.md)
- `linkedlist/` - [Linked List Systems](linkedlist/README.md)
- `queue/` - [Queue Patterns & Scheduling](queue/README.md)
- `stack/` - [Stack Techniques & Recursion Tools](stack/README.md)
- `tree/` - [Tree Structures & Traversals](tree/README.md)
- `trie/` - [Prefix Tree Overview](trie/readme.md)

## Suggested Study Path

1. Warm up with the foundational ADTs (`stack/`, `queue/`, `linkedlist/`).
2. Practice traversal templates in `tree/` and `graph/` - they unlock many medium problems.
3. Dive into `array/` and `dp/` folders to cover sliding window, prefix sums, and dynamic programming staples.
4. Challenge yourself with the `questions/` subfolders that mimic real interview prompts.
5. Revisit topics weekly; run the tests to confirm muscle memory.

## Real-World Applications

| Concept | Everyday example |
| --- | --- |
| Stack | Undo operations, browser history backtracking, recursion call stacks. |
| Queue | Background job scheduling, recently deleted photos, waitlists. |
| Linked List | Browser navigation, playlist sequencing, circular turn order in board games. |
| Tree | File systems, e-commerce category hierarchies, autocomplete tries. |
| Graph | Ride sharing route matching, dependency resolution, social network friend suggestions. |

## How to Contribute

- Fork the repository and create a feature branch.
- Add clear problem statements or comments when introducing new solutions.
- Include unit tests (`go test ./...`) when you add or refactor code.
- Run `make fmt` before opening a pull request to keep the style consistent.
- Share learnings by updating this README or adding notes in `resources.md` files where relevant.

## Additional Resources

- `graph/resources.md` - curated articles on graph theory fundamentals.
- `array/dp/README.md` and `array/slidingwindow/README.md` - topic-focused notes.
- `examples/README.md` - walkthroughs for combining multiple data structures.

If this toolkit helps you ace an interview, consider opening an issue with your story or adding the problem you solved. Happy studying!
