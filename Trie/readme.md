A `Trie` or a prefix tree is a special data structure used to store strings that can be visualized like a graph. 
It consists of nodes and edges. Each node consists of at max n children and edges connect each parent
![alt text](https://assets.leetcode.com/users/images/370e0b41-022a-4d94-84b8-85292bc1eaef_1613351652.1921413.png)
Figure 1 is an example when a trie stores [ape,apple,are,art, ...]

- Each path from the root to leaves forms a word.
- Each node except for the root node contains a value.
- All the descendants of a node share a common prefix associated to that node. For example, are and art share ar as the prefix.

There are two operations provided by a trie: inserting a new string, and searching for a given string.

##### `Advantages`:
- The advantage of using a trie is that, regardless of the number of strings stored in it,
  the time complexity for both inserting and searching is always O(L), where L is the length of the input string.
- This is obviously faster than BST(which takes O(L log n) time, where n is the number of words).
  This is also faster than Hashing because there are no collisions of different keys in a trie.
- we can easily print all words in alphabetical order which is not easily possible with hashing.


##### `Applications`:
- Dictionary Suggestions OR Autocomplete dictionary, Search Engines.
- Searching Contact from Mobile Contact list OR Phone Directory.
- We can efficiently do prefix search (or auto-complete) with Trie.
- Ex: If we search for word "tiny", then it auto suggest words starting with same characters like "tine", "tin", "tinny" etc.

##### `Limitations`:
- The main disadvantage of tries is that they need a lot of memory for storing the strings. 
  For each node we have too many node pointers(equal to number of characters of the alphabet). 
  However, this can turn into an advantage when working with a set of strings where a lot of strings have common prefixes. 
  For example- house, housekey, housekeep, housekeeper, etc.

