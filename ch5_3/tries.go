// Package ch5_3: Tries.
package ch5_3


// ==== EX 5.3.1 | Trie node and insertion | Core
// PROMPT: Create `type TrieNode struct{ Children map[byte]*TrieNode; End bool }` and `type Trie struct{ Root *TrieNode }` with `NewTrie()`. Write `Insert(word string)` that walks or creates one child per byte, then marks the last node as the end of a word.
// EXAMPLE: insert "pin", "pit" -> the root has one child `'p'`, whose `'i'` child has children `'n'` and `'t'`
// BIGO: O(K) for a K-character word, regardless of how many words are stored.
// NOTE: Map iteration order is random, so every method below that lists words visits children in sorted key order to give the same answer every run.
type TrieNode struct {
	Children   map[byte]*TrieNode
	End        bool
	Popularity int
}

type Trie struct{ Root *TrieNode }

func newNode() *TrieNode { panic("todo") }

func NewTrie() *Trie { panic("todo") }

func (t *Trie) Insert(word string) {
	panic("todo")
}

func sortedKeys(n *TrieNode) []byte {
	panic("todo")
}

// ==== END

// ==== EX 5.3.2 | Search a prefix | Core
// PROMPT: Add `Search(prefix string) *TrieNode` returning the node where the prefix ends, or nil if it is not in the trie.
// EXAMPLE: After "pin" -> `Search("pi")` is a node; `Search("po")` is `nil`
// BIGO: O(K).
func (t *Trie) Search(prefix string) *TrieNode {
	panic("todo")
}

// ==== END

// ==== EX 5.3.3 | Is it a whole word? | Core
// PROMPT: Add `ContainsWord(word string) bool`: the prefix must exist and its node must be marked as an end.
// EXAMPLE: After "pin" -> `ContainsWord("pin")` is `true`, `ContainsWord("pi")` is `false`
// BIGO: O(K).
func (t *Trie) ContainsWord(word string) bool {
	panic("todo")
}

// ==== END

// ==== EX 5.3.4 | Collect every word | Core
// PROMPT: Add `CollectAllWords(node *TrieNode) []string` that recursively gathers every word below node (suffixes relative to it), building each in one `[]byte` that you append to and trim back.
// EXAMPLE: After "pin", "pit", "sun" -> `CollectAllWords(t.Root)` is `[pin pit sun]`
// BIGO: O(M) nodes beneath the starting node (plus the length of the words copied out).
func (t *Trie) CollectAllWords(node *TrieNode) []string {
	panic("todo")
}

// ==== END

// ==== EX 5.3.5 | Autocomplete | Core
// PROMPT: Add `Autocomplete(prefix string) []string`: search to the prefix's node, collect the words below it and prepend the prefix.
// EXAMPLE: Words "pin", "pinch", "pit", "sun" -> `Autocomplete("pin")` -> `[pin pinch]`
// BIGO: O(K + M): reach the prefix, then visit only the matching subtree.
func (t *Trie) Autocomplete(prefix string) []string {
	panic("todo")
}

// ==== END

// ==== EX 5.3.6 | Print every key | Practice
// PROMPT: Add `AllKeys() string` returning every character stored in the trie, depth first, in sorted child order.
// EXAMPLE: After "pie", "pin" -> `"pien"`
// BIGO: O(M) nodes.
func (t *Trie) AllKeys() string {
	panic("todo")
}

// ==== END

// ==== EX 5.3.7 | Longest prefix every word shares | Practice
// PROMPT: Add `LongestCommonPrefix() string`: walk down from the root while the current node has exactly one child and no word ends there, collecting the bytes. Return `""` for an empty trie.
// EXAMPLE: Words "pine", "pint", "pinto" -> `"pin"`; words "pin", "pinch" -> `"pin"`
// BIGO: O(P) where P is the length of the shared prefix: only one path is ever walked.
func (t *Trie) LongestCommonPrefix() string {
	panic("todo")
}

// ==== END

// ==== EX 5.3.8 | Tries with values: ranked autocomplete | Practice
// PROMPT: Add `InsertRanked(word string, popularity int)` that stores the popularity on the end node, and `TopCompletions(prefix string, limit int) []string` returning the most popular completions first (ties alphabetical).
// EXAMPLE: (pinch, 9), (pin, 5), (pinball, 7) -> `TopCompletions("pin", 2)` -> `[pinch pinball]`
// BIGO: O(K + M + R log R) to collect and sort R matches.
func (t *Trie) InsertRanked(word string, popularity int) {
	panic("todo")
}

func (t *Trie) TopCompletions(prefix string, limit int) []string {
	panic("todo")
}

// ==== END

// ==== EX 5.3.9 | Delete a word | Stretch
// PROMPT: Add `Delete(word string) bool` that clears the end flag and prunes nodes that no longer lead to any word. Return true if the word was present.
// EXAMPLE: Words "pin", "pinch" -> delete "pinch" -> "pin" still found, the c-h nodes are gone
// BIGO: O(K).
func (t *Trie) Delete(word string) bool {
	panic("todo")
}

// ==== END

// WordNode is the trie node for 5.3.10-5.3.16: Next maps a letter to a child, Passes counts the words through it, Ends the words ending at it.
type WordNode struct {
	Next         map[byte]*WordNode
	Passes, Ends int
}

// WordTrie builds a WordNode trie holding words, repeats counted.
func WordTrie(words ...string) *WordNode {
	root := &WordNode{Next: map[byte]*WordNode{}}
	for _, word := range words {
		node := root
		node.Passes++
		for i := 0; i < len(word); i++ {
			child, ok := node.Next[word[i]]
			if !ok {
				child = &WordNode{Next: map[byte]*WordNode{}}
				node.Next[word[i]] = child
			}
			node = child
			node.Passes++
		}
		node.Ends++
	}
	return root
}

// ==== EX 5.3.10 | Count the words under each prefix | Core
// PROMPT: Write `PrefixCounts(words, prefixes []string) []int` that returns, for each prefix, how many of `words` (repeats included) start with it. Use the `WordTrie(words...)` helper: each `WordNode` has `Next` (letter -> child), `Passes` and `Ends`. Walk each prefix down from the root and read `Passes`.
// EXAMPLE: `["car" "card" "care" "dog"], ["car" "ca" "d" "x"]` -> `[3 3 1 0]`
// BIGO: O(total letters) to build, then O(K) per prefix of length K: the node at the end of the prefix already counts every word passing through it.
func PrefixCounts(words, prefixes []string) []int {
	panic("todo")
}

// ==== END

// ==== EX 5.3.11 | Sort words with a trie | Core
// PROMPT: Write `TrieSort(words []string) []string` that returns `words` in dictionary order, repeats kept, as a non-nil slice, by walking a trie: emit a node's word `Ends` times, then visit its children in letter order. Use the `WordTrie(words...)` helper: each `WordNode` has `Next` (letter -> child), `Passes` and `Ends`.
// EXAMPLE: `["pear" "apple" "pea" "app"]` -> `["app" "apple" "pea" "pear"]`
// BIGO: O(total letters) plus the sort of each node's children: a pre-order walk that takes children in letter order meets the words in dictionary order.
func TrieSort(words []string) []string {
	panic("todo")
}

// ==== END

// ==== EX 5.3.12 | Wildcard search | Practice
// PROMPT: Write `WildcardMatches(words []string, pattern string) []string` that returns, in dictionary order and without repeats, the words matching `pattern` (a non-nil slice), where `.` matches any one letter. Walk the trie: a letter follows one child, a dot tries every child in letter order. Use the `WordTrie(words...)` helper: each `WordNode` has `Next` (letter -> child), `Passes` and `Ends`.
// EXAMPLE: `["bad" "bed" "bid" "bud" "bead"], "b.d"` -> `["bad" "bed" "bid" "bud"]`
// BIGO: O(26^D * K) worst case for D dots in a pattern of length K, far less in practice: each dot branches only into children that exist.
func WildcardMatches(words []string, pattern string) []string {
	panic("todo")
}

// ==== END

// ==== EX 5.3.13 | Substring check with a suffix trie | Practice
// PROMPT: Write `SubstringChecks(text string, queries []string) []bool` that returns, for each query, whether it occurs in `text` (the empty string always does). Put every suffix of `text` in a trie, then walk each query from the root. Use the `WordTrie(words...)` helper: each `WordNode` has `Next` (letter -> child), `Passes` and `Ends`.
// EXAMPLE: `"banana", ["nan" "ana" "nab" ""]` -> `[true true false true]`
// BIGO: O(N^2) time and space to build for a text of length N, then O(M) per query of length M: every substring is a prefix of some suffix.
// NOTE: Suffix trees and suffix arrays get the build down to O(N) or O(N log N); the query idea is the same.
func SubstringChecks(text string, queries []string) []bool {
	panic("todo")
}

// ==== END

// ==== EX 5.3.14 | Shortest unique prefixes | Practice
// PROMPT: `words` are distinct and none is a prefix of another. Write `ShortestUniquePrefixes(words []string) []string` that returns, for each word in order, its shortest non-empty prefix that no other word starts with. Walk each word down until a node is passed by one word only. Use the `WordTrie(words...)` helper: each `WordNode` has `Next` (letter -> child), `Passes` and `Ends`.
// EXAMPLE: `["zebra" "dog" "duck" "dove"]` -> `["z" "dog" "du" "dov"]`
// BIGO: O(total letters) time and space: the `passes` count on each node says at once whether any other word shares the prefix so far.
func ShortestUniquePrefixes(words []string) []string {
	panic("todo")
}

// ==== END

// ==== EX 5.3.15 | Keypad words | Practice
// PROMPT: On a phone keypad, `2` is `abc`, `3` `def`, `4` `ghi`, `5` `jkl`, `6` `mno`, `7` `pqrs`, `8` `tuv` and `9` `wxyz`. Write `KeypadWords(digits string, words []string) []string` that returns, in dictionary order, the words from `words` whose letters type exactly `digits` (a non-nil slice). Walk the trie, trying each letter of the current digit. Use the `WordTrie(words...)` helper: each `WordNode` has `Next` (letter -> child), `Passes` and `Ends`.
// EXAMPLE: `"4663", ["good" "home" "gone" "hood" "hoof" "in"]` -> `["gone" "good" "home" "hood" "hoof"]`
// BIGO: O(4^K) worst case for K digits, usually far less: the trie cuts off every letter sequence that starts no word.
func KeypadWords(digits string, words []string) []string {
	panic("todo")
}

// ==== END

// ==== EX 5.3.16 | Words on a letter board | Stretch
// PROMPT: `board` holds equal-length strings. A word can be traced by starting on any cell and stepping to any of the 8 neighbouring cells, using each cell at most once. Write `BoardWords(board, words []string) []string` that returns, sorted and without repeats, the words from `words` that can be traced (a non-nil slice). Start a depth-first walk from every cell, following the trie in step with the path. Use the `WordTrie(words...)` helper: each `WordNode` has `Next` (letter -> child), `Passes` and `Ends`.
// EXAMPLE: `["cat" "ore" "bin"]`, `["car" "cat" "core" "tea" "rat" "bin" "rob" "cab" "tin" "ore" "rare"]` -> `["bin" "car" "cat" "core" "ore" "rat" "rob" "tea"]`
// BIGO: O(R * C * 8^L) worst case for words up to length L, in practice far less: the walk stops as soon as the path so far starts no word.
func BoardWords(board, words []string) []string {
	panic("todo")
}

// ==== END
