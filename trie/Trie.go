package trie

type node struct {
	children map[rune]*node
	isEnd    bool
}

// Trie represents a prefix tree for runes.
type Trie struct {
	root *node
	size int
}

func newNode() *node {
	return &node{children: make(map[rune]*node)}
}

// NewTrie constructs an empty Trie.
func NewTrie() *Trie {
	return &Trie{root: newNode()}
}

// Insert adds the provided word to the trie.
func (t *Trie) Insert(word string) {
	if t == nil {
		return
	}

	current := t.root
	for _, ch := range word {
		next, exists := current.children[ch]
		if !exists {
			next = newNode()
			current.children[ch] = next
		}
		current = next
	}

	if !current.isEnd {
		current.isEnd = true
		t.size++
	}
}

// Contains reports whether the exact word exists in the trie.
func (t *Trie) Contains(word string) bool {
	if t == nil {
		return false
	}

	current := t.root
	for _, ch := range word {
		next, exists := current.children[ch]
		if !exists {
			return false
		}
		current = next
	}
	return current.isEnd
}

// HasPrefix reports whether any word in the trie starts with the prefix.
func (t *Trie) HasPrefix(prefix string) bool {
	if t == nil {
		return false
	}

	current := t.root
	for _, ch := range prefix {
		next, exists := current.children[ch]
		if !exists {
			return false
		}
		current = next
	}
	return true
}

// Size returns the number of distinct words stored in the trie.
func (t *Trie) Size() int {
	if t == nil {
		return 0
	}
	return t.size
}

// LongestCommonPrefix walks the trie while every node has a single child.
func (t *Trie) LongestCommonPrefix() string {
	if t == nil || t.size == 0 {
		return ""
	}

	current := t.root
	prefix := make([]rune, 0)

	for len(current.children) == 1 && !current.isEnd {
		for ch, next := range current.children {
			prefix = append(prefix, ch)
			current = next
			break
		}
	}

	return string(prefix)
}
