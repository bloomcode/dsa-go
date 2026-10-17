// Package ch24: Caches and eviction policies.
package ch24

import (
	"container/list"
	"math/rand/v2"
)

// ==== EX 24.1 | LRU cache from the standard library | Core
// PROMPT: Write `type Cache[K comparable, V any] interface{ Get(K) (V, bool); Put(K, V) }`, then `type LRUCache[K, V]` that evicts the least recently used key, built from a `*list.List` (most recent at the front) and a `map[K]*list.Element`.
// EXAMPLE: capacity 2: put a, put b, get a, put c -> b is evicted; `Get("b")` is `0, false`
// BIGO: O(1) per operation.
type Cache[K comparable, V any] interface {
	Get(K) (V, bool)
	Put(K, V)
}

type lruEntry[K comparable, V any] struct {
	key   K
	value V
}

type LRUCache[K comparable, V any] struct {
	capacity int
	order    *list.List
	items    map[K]*list.Element
}

func NewLRU[K comparable, V any](capacity int) *LRUCache[K, V] {
	panic("todo")
}

func (c *LRUCache[K, V]) Get(k K) (V, bool) {
	panic("todo")
}

func (c *LRUCache[K, V]) Put(k K, v V) {
	panic("todo")
}

// ==== END

// ==== EX 24.2 | LRU from a hash table and a doubly linked list | Core
// PROMPT: Build `type LinkedLRU[K, V]` without `container/list`: a map from key to node, plus a doubly linked list with head and tail sentinel nodes. Most recent sits next to head; evict from next to tail.
// EXAMPLE: Same behaviour as 24.1
// BIGO: O(1) per Get and Put: the map finds the node, the links let it move without shifting.
// NOTE: Sentinel nodes remove every "is the list empty?" special case from insert and unlink. `container/list` uses one sentinel internally for the same reason.
type node[K comparable, V any] struct {
	key        K
	value      V
	prev, next *node[K, V]
}

type LinkedLRU[K comparable, V any] struct {
	capacity   int
	items      map[K]*node[K, V]
	head, tail *node[K, V]
}

func NewLinkedLRU[K comparable, V any](capacity int) *LinkedLRU[K, V] {
	panic("todo")
}

func (c *LinkedLRU[K, V]) unlink(n *node[K, V]) {
	panic("todo")
}

func (c *LinkedLRU[K, V]) pushFront(n *node[K, V]) {
	panic("todo")
}

func (c *LinkedLRU[K, V]) Get(k K) (V, bool) {
	panic("todo")
}

func (c *LinkedLRU[K, V]) Put(k K, v V) {
	panic("todo")
}

// ==== END

// ==== EX 24.3 | Measure a hit rate | Core
// PROMPT: Write `HitRate[K comparable](c Cache[K, bool], requests []K) float64` that, for each requested key, counts a hit when Get finds it and otherwise Puts it. Return hits / requests.
// EXAMPLE: `HitRate(NewLRU[string, bool](2), []string{"a", "b", "a", "a", "c", "b"})` -> `0.333...` (2 hits in 6)
// BIGO: O(R) for R requests.
func HitRate[K comparable](c Cache[K, bool], requests []K) float64 {
	panic("todo")
}

// ==== END

// ==== EX 24.4 | FIFO eviction for comparison | Core
// PROMPT: Write `type FIFOCache[K, V]` that evicts the oldest inserted key regardless of use, keeping insertion order in a slice queue. Compare its hit rate with LRU on a stream that reuses a few hot keys.
// EXAMPLE: hot keys reused constantly -> LRU's hit rate is higher than FIFO's
// BIGO: O(1) amortized per operation.
type FIFOCache[K comparable, V any] struct {
	capacity int
	data     map[K]V
	order    []K
}

func NewFIFO[K comparable, V any](capacity int) *FIFOCache[K, V] {
	panic("todo")
}

func (c *FIFOCache[K, V]) Get(k K) (V, bool) {
	panic("todo")
}

func (c *FIFOCache[K, V]) Put(k K, v V) {
	panic("todo")
}

// ==== END

// ==== EX 24.5 | Random eviction in O(1) | Core
// PROMPT: Write `type RandomCache[K, V]` that evicts a uniformly random key. Keep keys in a slice plus a map of key to index; remove by moving the last key into the victim's slot.
// EXAMPLE: capacity 3 never holds more than 3 keys; which key goes depends on the seed
// BIGO: O(1) per operation.
type RandomCache[K comparable, V any] struct {
	capacity int
	rng      *rand.Rand
	data     map[K]V
	keys     []K
	pos      map[K]int
}

func NewRandomCache[K comparable, V any](capacity int, rng *rand.Rand) *RandomCache[K, V] {
	panic("todo")
}

func (c *RandomCache[K, V]) Len() int { panic("todo") }

func (c *RandomCache[K, V]) Get(k K) (V, bool) {
	panic("todo")
}

func (c *RandomCache[K, V]) Put(k K, v V) {
	panic("todo")
}

// ==== END

// ==== EX 24.6 | LRU's worst case, fixed by randomness | Practice
// PROMPT: Write `CyclicRequests(capacity, rounds int) []int` that cycles through capacity + 1 keys. Show LRU's hit rate on it is 0, while `RandomCache` keeps a healthy share of hits.
// EXAMPLE: capacity 10, 100 rounds -> LRU `0`; random eviction well above `0.5`
// BIGO: O(R). Every request evicts the very key that will be needed next, the textbook LRU pathology.
func CyclicRequests(capacity, rounds int) []int {
	panic("todo")
}

// ==== END

// ==== EX 24.7 | A bounded memoizing wrapper | Practice
// PROMPT: Write `type Memo[K comparable, V any] struct` with `NewMemo(fn func(K) V, maxsize int)` and `Call(k K) V`, caching results in your `LRUCache` and counting `Hits` and `Misses`.
// EXAMPLE: wrap a slow square, call it with 1, 2, 1 -> Hits `1`, Misses `2`
// BIGO: O(1) bookkeeping per call; memory is capped at maxsize entries, unlike an unbounded map memo.
type Memo[K comparable, V any] struct {
	fn           func(K) V
	cache        *LRUCache[K, V]
	Hits, Misses int
}

func NewMemo[K comparable, V any](fn func(K) V, maxsize int) *Memo[K, V] {
	panic("todo")
}

func (m *Memo[K, V]) Call(k K) V {
	panic("todo")
}

// ==== END

// ==== EX 24.8 | Cache-friendly traversal order | Core
// PROMPT: Write `SumRowsFirst(grid [][]int) int`, `SumColumnsFirst(grid [][]int) int`, and `BlockedTranspose(m [][]int, block int) [][]int` that copies the matrix tile by tile. All must give correct results; time them on a 4000 by 4000 grid.
// EXAMPLE: both sums of `[[1 2] [3 4]]` -> `10`; `BlockedTranspose([[1 2 3] [4 5 6]], 2)` -> `[[1 4] [2 5] [3 6]]`
// BIGO: All O(R * C) steps. Row-first walks memory in order (spatial locality); column-first jumps to a different row on every step and is often several times slower.
func SumRowsFirst(grid [][]int) int {
	panic("todo")
}

func SumColumnsFirst(grid [][]int) int {
	panic("todo")
}

func BlockedTranspose(m [][]int, block int) [][]int {
	panic("todo")
}

// ==== END
