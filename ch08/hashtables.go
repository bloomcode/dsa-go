// Package ch08: Hash tables and constant-time lookup.
package ch08


// ==== EX 8.1 | Build a hash table with chaining | Core
// PROMPT: Build `type HashTable struct` for string keys and int values on a slice of buckets, each a slice of key-value entries, hashing keys with `hash/fnv`. Implement `Put`, `Get(key) (int, bool)`, `Delete(key) bool` and `Len()`. Double the bucket count when the load factor passes 0.75.
// EXAMPLE: `t.Put("bad", 7); t.Get("bad")` -> `7, true`; putting the same key again overwrites it
// BIGO: O(1) average for Put, Get and Delete; O(N) worst case if every key collides into one bucket. A resize is O(N), but rare, so O(1) amortized.
// NOTE: Go exposes no hash function for arbitrary keys, which is why this table is written for strings; the built-in map is the general-purpose version.
type entry struct {
	key   string
	value int
}

type HashTable struct {
	buckets [][]entry
	count   int
}

func NewHashTable() *HashTable { panic("todo") }

func (t *HashTable) index(key string) int {
	panic("todo")
}

func (t *HashTable) Put(key string, value int) {
	panic("todo")
}

func (t *HashTable) Get(key string) (int, bool) {
	panic("todo")
}

func (t *HashTable) Delete(key string) bool {
	panic("todo")
}

func (t *HashTable) Len() int { panic("todo") }

func (t *HashTable) Buckets() int { panic("todo") }

// ==== END

// ==== EX 8.2 | Is one array a subset of another? | Core
// PROMPT: Write `IsSubset(small, big []string) bool` returning true if every value in small appears in big.
// EXAMPLE: `[b d]`, `[a b c d]` -> `true`
// BIGO: O(N + M) with a hash table, versus O(N * M) with nested loops.
func IsSubset(small, big []string) bool {
	panic("todo")
}

// ==== END

// ==== EX 8.3 | Intersection in linear time | Core
// PROMPT: Write `Intersection(a, b []int) []int` returning values present in both, in the order they appear in a, each once.
// EXAMPLE: `[1 2 3 4 5]`, `[0 2 4 6 8]` -> `[2 4]`
// BIGO: O(N + M). Compare with 6.3's O(N * M).
func Intersection(a, b []int) []int {
	panic("todo")
}

// ==== END

// ==== EX 8.4 | First duplicate string | Core
// PROMPT: Write `FirstDuplicate(words []string) (string, bool)` returning the first word seen a second time.
// EXAMPLE: `[a b c d c e f]` -> `"c", true`
// BIGO: O(N).
func FirstDuplicate(words []string) (string, bool) {
	panic("todo")
}

// ==== END

// ==== EX 8.5 | The missing letter | Practice
// PROMPT: A string contains every letter a-z except one. Write `MissingLetter(s string) (byte, bool)` to find it.
// EXAMPLE: `"the quick brown box jumps over a lazy dog"` -> `'f', true`
// BIGO: O(N): one pass to record letters, then a fixed 26 checks.
// NOTE: For lowercase letters a `[26]bool` indexed by `c - 'a'` is a perfect hash table: no hashing at all.
func MissingLetter(s string) (byte, bool) {
	panic("todo")
}

// ==== END

// ==== EX 8.6 | First non-repeated character | Core
// PROMPT: Write `FirstUniqueChar(s string) (rune, bool)` returning the first character that appears exactly once.
// EXAMPLE: `"minimum"` -> `'n', true`
// BIGO: O(N): two passes, counting then checking.
func FirstUniqueChar(s string) (rune, bool) {
	panic("todo")
}

// ==== END

// ==== EX 8.7 | Word frequency | Core
// PROMPT: Write `WordCounts(text string) map[string]int` mapping each lowercase word to its count, and `MostCommonWord(text string) string` (ties go to the word that appears first).
// EXAMPLE: `"the cat and the hat"` -> `map[and:1 cat:1 hat:1 the:2]`, most common `"the"`
// BIGO: O(N) in the number of words. Map iteration order is random, so break ties by walking the words in order, not the map.
func WordCounts(text string) map[string]int {
	panic("todo")
}

func MostCommonWord(text string) string {
	panic("todo")
}

// ==== END

// ==== EX 8.8 | Group anagrams | Practice
// PROMPT: Write `GroupAnagrams(words []string) [][]string` returning groups of words that are anagrams of each other, groups in first-seen order. Use the sorted letters as the map key.
// EXAMPLE: `[eat tea tan ate nat bat]` -> `[[eat tea ate] [tan nat] [bat]]`
// BIGO: O(N * K log K) for N words of length K.
func GroupAnagrams(words []string) [][]string {
	panic("todo")
}

// ==== END

// ==== EX 8.9 | Two sum with a lookup table | Core
// PROMPT: Write `TwoSumIndices(nums []int, target int) (int, int, bool)` returning the indices of two numbers that add to target.
// EXAMPLE: `[2 7 11 15]`, target `9` -> `0, 1, true`
// BIGO: O(N) time, O(N) space; nested loops would be O(N^2).
func TwoSumIndices(nums []int, target int) (int, int, bool) {
	panic("todo")
}

// ==== END

// ==== EX 8.10 | Invert a lookup table | Practice
// PROMPT: Write `Invert(menu map[string]int) map[int][]string` turning item -> price into price -> sorted items, so you can look up everything at a price in one step.
// EXAMPLE: `map[burger:5 fries:2 soda:2]` -> `map[2:[fries soda] 5:[burger]]`
// BIGO: O(N log N) with the sort of each group; every later lookup by price is O(1).
// NOTE: A hash table is one-directional: fast from key to value, slow from value to key. Build a second table when you need both.
func Invert(menu map[string]int) map[int][]string {
	panic("todo")
}

// ==== END
