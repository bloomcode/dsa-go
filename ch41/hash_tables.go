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

// ==== EX 41.11 | Longest consecutive run | Practice
// PROMPT: Write `LongestConsecutive(nums []int) int` that returns the length of the longest run of consecutive integers found anywhere in `nums` (any order, repeats allowed), in O(N): put the values in a set (`map[int]bool`) and count upward only from values whose predecessor is missing.
// EXAMPLE: `[10 4 20 1 3 2 11]` -> `4` (1, 2, 3, 4)
// BIGO: O(N) expected time and space: only a value with no predecessor starts a count, so every value is stepped over at most twice.
func LongestConsecutive(nums []int) int {
	panic("todo")
}

// ==== END

// ==== EX 41.12 | Does a stretch sum to zero? | Core
// PROMPT: Write `HasZeroSumRun(nums []int) bool`: `true` when some non-empty run of consecutive values sums to `0`. Keep a set of the prefix sums seen so far, starting with `0`; a repeated prefix sum means the values between the two sum to zero.
// EXAMPLE: `[4 2 -3 1 6]` -> `true` (`2, -3, 1`); `[4 2 1]` -> `false`
// BIGO: O(N) expected time and space: a stretch sums to zero exactly when two prefix sums are equal, and a set spots the repeat.
func HasZeroSumRun(nums []int) bool {
	panic("todo")
}

// ==== END

// ==== EX 41.13 | Count stretches summing to k | Practice
// PROMPT: Write `CountSumRuns(nums []int, k int) int` that counts the non-empty runs of consecutive values summing to `k` (values may be negative), in O(N). Keep a map from each prefix sum to how often it has occurred, starting with `{0: 1}`; at each step add the count of `total - k`.
// EXAMPLE: `[1 2 3 -2 2], 3` -> `4` (`1, 2`; `3`; `2, 3, -2`; `3, -2, 2`)
// BIGO: O(N) expected time and space: each prefix sum `p` pairs with every earlier prefix sum `p - k`, and a count table answers that in O(1).
func CountSumRuns(nums []int, k int) int {
	panic("todo")
}

// ==== END

// ==== EX 41.14 | Symmetric pairs | Core
// PROMPT: `pairs` holds distinct pairs `[a, b]` with `a != b`. Write `SymmetricPairs(pairs [][2]int) [][2]int` that returns every pair whose reverse `[b, a]` also appears, once per couple, as a non-nil slice: report the pair that came first, in the order the second one arrives.
// EXAMPLE: `[[1 2] [3 4] [5 6] [4 3] [2 1]]` -> `[[3 4] [1 2]]`
// BIGO: O(N) expected time and space: one pass, checking each pair's reverse against the pairs already seen.
func SymmetricPairs(pairs [][2]int) [][2]int {
	panic("todo")
}

// ==== END

// ==== EX 41.15 | Trace a journey from its tickets | Practice
// PROMPT: `tickets` are `[origin, destination]` legs of one journey that visits no place twice, in shuffled order. Write `TraceJourney(tickets [][2]string) []string` that returns the places in travel order (an empty, non-nil slice for no tickets). Map each origin to its destination; the start is the one origin that is nobody's destination.
// EXAMPLE: `[["MAA" "BOM"] ["DEL" "MAA"] ["BOM" "GOI"]]` -> `["DEL" "MAA" "BOM" "GOI"]`
// BIGO: O(N) expected time and space: one map from each origin to its destination, one set of destinations to find the start, then one walk.
func TraceJourney(tickets [][2]string) []string {
	panic("todo")
}

// ==== END

// ==== EX 41.16 | Do two arrays share a value? | Core
// PROMPT: Write `Disjoint(a, b []int) bool`: `true` when no value appears in both `a` and `b`. Build a set from `a` and check each value of `b` against it.
// EXAMPLE: `[1 5 9], [2 6]` -> `true`; `[1 5], [5]` -> `false`
// BIGO: O(N + M) expected time, O(N) space: build a set from one array and probe it with the other.
func Disjoint(a, b []int) bool {
	panic("todo")
}

// ==== END

// ==== EX 41.17 | Most points on one line | Stretch
// PROMPT: `points` are `[x, y]` integer pairs, possibly repeated. Write `MostOnALine(points [][2]int) int` that returns the largest number of them on one straight line. For each anchor point, count the other points by direction `(dx, dy)` reduced by their gcd with a fixed sign; copies of the anchor count on every line through it.
// EXAMPLE: `[[0 0] [1 1] [2 2] [3 1] [0 3]]` -> `3`
// BIGO: O(N^2) expected time, O(N) space: for each anchor point, every other point's direction, reduced by the gcd, is a hash key for the line through both.
// NOTE: Reducing `(dx, dy)` by their gcd and fixing the sign gives an exact key; a floating-point slope would make nearly equal slopes collide or differ.
func MostOnALine(points [][2]int) int {
	panic("todo")
}

// ==== END
