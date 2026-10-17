// Package ch41: Hash tables.
package ch41


// ==== EX 41.1 | Build a hash table with chaining | Core
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

// ==== EX 41.2 | Build a word from letter tiles | Core
// PROMPT: Write `CanBuild(word, tiles string) bool`: each character of `tiles` can be used once. Count the tiles in a map, then spend them on the word.
// EXAMPLE: `"abba"`, `"aabbc"` -> `true`; `"abc"`, `"ab"` -> `false`
// BIGO: O(N + M): count the tiles once, then spend them. Nested loops would be O(N * M).
func CanBuild(word, tiles string) bool {
	panic("todo")
}

// ==== END

// ==== EX 41.3 | Values in exactly one list | Core
// PROMPT: Write `InExactlyOne(a, b []int) []int` returning the values that appear in one slice but not the other: first those from `a` in `a`'s order, then those from `b` in `b`'s order, each once.
// EXAMPLE: `[1 2 3]`, `[2 3 4]` -> `[1 4]`
// BIGO: O(N + M) with two sets, versus O(N * M) checking each value against the other list.
func InExactlyOne(a, b []int) []int {
	panic("todo")
}

// ==== END

// ==== EX 41.4 | Distinct visitors per day | Core
// PROMPT: Visit `i` was on day `days[i]` by user `users[i]`. Write `DistinctPerDay(days []int, users []string) map[int]int` returning each day's number of different users.
// EXAMPLE: `[1 1 1 2]`, `[ana bo ana ana]` -> `map[1:2 2:1]`
// BIGO: O(N): one set of visitors per day, one insert per visit.
func DistinctPerDay(days []int, users []string) map[int]int {
	panic("todo")
}

// ==== END

// ==== EX 41.5 | Same letter pattern | Practice
// PROMPT: Write `SamePattern(a, b string) bool`: `true` when the bytes of `a` can be renamed one-to-one to give `b` (equal letters stay equal, different letters stay different).
// EXAMPLE: `"egg"`, `"add"` -> `true`; `"foo"`, `"bar"` -> `false`
// BIGO: O(N) with two lookup tables, one for each direction.
func SamePattern(a, b string) bool {
	panic("todo")
}

// ==== END

// ==== EX 41.6 | Every position of each value | Core
// PROMPT: Write `Positions(items []string) map[string][]int` returning each value's indexes, in increasing order.
// EXAMPLE: `[a b a]` -> `map[a:[0 2] b:[1]]`
// BIGO: O(N): one append per item.
func Positions(items []string) map[string][]int {
	panic("todo")
}

// ==== END

// ==== EX 41.7 | Word frequency | Core
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

// ==== EX 41.8 | Group anagrams | Practice
// PROMPT: Write `GroupAnagrams(words []string) [][]string` returning groups of words that are anagrams of each other, groups in first-seen order. Use the sorted letters as the map key.
// EXAMPLE: `[eat tea tan ate nat bat]` -> `[[eat tea ate] [tan nat] [bat]]`
// BIGO: O(N * K log K) for N words of length K.
func GroupAnagrams(words []string) [][]string {
	panic("todo")
}

// ==== END

// ==== EX 41.9 | Two sum with a lookup table | Core
// PROMPT: Write `TwoSumIndices(nums []int, target int) (int, int, bool)` returning the indices of two numbers that add to target.
// EXAMPLE: `[2 7 11 15]`, target `9` -> `0, 1, true`
// BIGO: O(N) time, O(N) space; nested loops would be O(N^2).
func TwoSumIndices(nums []int, target int) (int, int, bool) {
	panic("todo")
}

// ==== END

// ==== EX 41.10 | Invert a lookup table | Practice
// PROMPT: Write `Invert(menu map[string]int) map[int][]string` turning item -> price into price -> sorted items, so you can look up everything at a price in one step.
// EXAMPLE: `map[burger:5 fries:2 soda:2]` -> `map[2:[fries soda] 5:[burger]]`
// BIGO: O(N log N) with the sort of each group; every later lookup by price is O(1).
// NOTE: A hash table is one-directional: fast from key to value, slow from value to key. Build a second table when you need both.
func Invert(menu map[string]int) map[int][]string {
	panic("todo")
}

// ==== END
