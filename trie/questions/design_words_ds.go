package questions

/*
Design a data structure that supports adding new words and finding if a string matches any previously added string.

Implement the WordDictionary class:
WordDictionary() Initializes the object.
void addWord(word) Adds word to the data structure, it can be matched later.
bool search(word) Returns true if there is any string in the data structure that matches word or false otherwise. word may contain dots '.' where dots can be matched with any letter.


Example:
Input
["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
Output
[null,null,null,null,false,true,true,true]

LC: https://leetcode.com/problems/design-add-and-search-words-data-structure/description/
*/

type WordDictionary struct {
	children    [26]*WordDictionary // Children nodes for 'a' to 'z'
	isEndOfWord bool                // Indicates if this node represents the end of a word
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func NewTrieNode() *WordDictionary {
	return &WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
	current := this
	for _, ch := range word {
		index := ch - 'a' // Calculate index for 'a' to 'z'
		if current.children[index] == nil {
			current.children[index] = NewTrieNode()
		}
		current = current.children[index]
	}
	current.isEndOfWord = true
}

func (this *WordDictionary) Search(word string) bool {
	return searchHelper(word, 0, this)
}

func searchHelper(word string, index int, node *WordDictionary) bool {
	if node == nil {
		return false
	}
	if index == len(word) {
		return node.isEndOfWord
	}
	ch := rune(word[index])
	if ch == '.' {
		for _, child := range node.children {
			if searchHelper(word, index+1, child) {
				return true
			}
		}
		return false
	}
	return searchHelper(word, index+1, node.children[ch-'a'])
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
