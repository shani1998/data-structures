package Trie

import (
	"fmt"
	"strings"
)

//Node represent each character
type Node struct {
	//this is a single char stored for example letter a,b,c,d,etc
	char rune
	//store all children  of a node that is from a-z
	children [26]*Node
}

// Trie  is our actual tree that will hold all of our nodes
// the Root node will be nil
type Trie struct {
	root *Node
}

// NewTrieNode this will be used to initialize a new node with 26 children
//each child should first be initialized to nil
func NewTrieNode(ch rune) *Node {
	node := &Node{char: ch}
	for i := 0; i < 26; i++ {
		node.children[i] = nil
	}
	return node
}

// Insert inserts a word to the trie
func (t *Trie) Insert(str string) {
	//this will keep track of our current node
	//when transversing our tree  it should always
	// start at the top of our tree, i.e our root
	current := t.root

	///remove all spaces from the word
	///and convert it to lowercase
	strippedWord := strings.ToLower(strings.ReplaceAll(str, " ", ""))

	for _, ch := range strippedWord {
		// check if current already has a node created at our current node
		// if not create the node
		if current.children[ch-97] == nil {
			current.children[ch-97] = NewTrieNode(ch)
		}
		current = current.children[ch-97]
	}
}

func main() {
	trie := &Trie{root: NewTrieNode(0)}
	trie.Insert("app")
	trie.Insert("apple")
	fmt.Printf("%+v", trie.root.children)
	for v := range trie.root.children {
		fmt.Printf("%+v", v)
	}
}
