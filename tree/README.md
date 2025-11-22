# Tree Structures & Traversals

Trees appear in every MAANG interview set: binary trees, BSTs, tries, segment trees, and more. This folder collects templates and problems to build reusable muscle memory.

## Interview Focus Areas

- Traversal fluency: preorder, inorder, postorder, level-order, and their iterative variants.
- Binary Search Trees (BST): insert/delete, validation, successor/predecessor logic, and balancing considerations.
- Segment trees and range queries: lazy propagation, point updates, and building from arrays.
- Tree DP and recursion: bottom-up vs top-down strategies for path sums, diameter, and deepest leaves.

## Real-World Applications

- Compilers and interpreters: abstract syntax trees and expression evaluation.
- File systems and configuration trees: hierarchy traversal, permission propagation.
- Gaming/AI: decision trees, minimax, and behavior trees for NPC logic.

## What's in This Folder

- `binary_tree.go` & `traversal.go`: reusable traversal helpers with recursive and iterative patterns.
- `bfs.go` & `dfs.go`: level-order vs depth-first templates that connect to graph practices.
- `bst/`: BST-specific helpers, search, insert, delete, and validation routines.
- `segmenttree/`: range query scaffolding for sum/min/max trackers.
- `questions/`: interview problems including deepest leaves sum, diameter, and balanced checks.

## Practice Game Plan

1. Implement tree traversals iteratively using stacks/queues to prove mastery of both approaches.
2. Solve `questions/deepest_leaves_sum.go` and extend it to track multiple aggregates (e.g., widths, counts).
3. Add deletion logic to the BST implementation and test edge cases (two-child removals, duplicates).
4. Build a segment tree for range minimum query and compare with sparse table approaches.

## Stretch Topics

- Dive into AVL and Red-Black tree rotations to understand self-balancing strategies.
- Explore binary lifting and Lowest Common Ancestor (LCA) techniques for ancestor queries.
- Combine tree knowledge with `trie/` to build prefix trees for autocomplete at scale.
