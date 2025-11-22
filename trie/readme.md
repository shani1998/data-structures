# Prefix Tree (Trie)

A trie (prefix tree) stores strings by sharing common prefixes across sibling branches. Each edge represents a character and every path from the root to a terminal node describes one word.

![Trie example](https://assets.leetcode.com/users/images/370e0b41-022a-4d94-84b8-85292bc1eaef_1613351652.1921413.png)

## Key Properties

- Each path from the root to a marked leaf/node forms a word.
- Each node (except the root) holds a character value and optional termination flag.
- Descendants of any node share the prefix associated with that node.

## Core Operations

- `Insert`: follow or create edges for each character and mark the terminal node.
- `Search`: walk the edges and confirm the terminal marker at the end of the word.
- `StartsWith`: like search but without requiring a terminal marker.

## Advantages

- Predictable O(L) insert/search where L is the word length, independent of total word count.
- Supports alphabetical iteration and prefix queries without extra sorting.
- Avoids hash collisions and offers deterministic traversal order.

## Applications

- Autocomplete engines and search suggestions.
- Contact search on mobile devices and directory lookups.
- Spell-checkers, predictive text, and DNA sequence prefix matching.
- IP routing tables (compressed tries like radix/patricia trees).

## Limitations

- Higher memory footprint compared to hash maps due to child pointers per alphabet symbol.
- Requires careful handling of sparse alphabets; consider compression or ternary search trees for efficiency.

## In This Repository

- `Trie.go`: core data structure implementation in Go.
- `longest_common_prefix.go`: leverages trie logic for prefix computations.
- `questions/`: interview-style problems to practice trie-based solutions.

