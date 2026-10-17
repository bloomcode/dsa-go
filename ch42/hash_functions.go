// Package ch42: Hash functions.
package ch42

import (
	"math/rand/v2"
)

// ==== EX 42.1 | A naive string hash and its collisions | Core
// PROMPT: Write `SumHash(s string, m int) int` that adds the byte values and takes `% m`. Show that every anagram collides.
// EXAMPLE: `SumHash("stop", 101) == SumHash("pots", 101) && SumHash("pots", 101) == SumHash("tops", 101)` -> `true`
// BIGO: O(len(s)). Ignoring position throws away most of the information in the key.
func SumHash(s string, m int) int {
	panic("todo")
}

// ==== END

// ==== EX 42.2 | Polynomial string hashing | Core
// PROMPT: Write `PolyHash(s string, m int) int` with base 31, computing `(c0 * 31^(L-1) + c1 * 31^(L-2) + ...) % m` with Horner's rule, so character order matters.
// EXAMPLE: `PolyHash("stop", 1_000_000_007) != PolyHash("pots", 1_000_000_007)` -> `true`
// BIGO: O(len(s)).
// NOTE: Indexing a string `s[i]` gives bytes; `for _, r := range s` would give runes. For hashing, bytes are what you want.
func PolyHash(s string, m int) int {
	panic("todo")
}

// ==== END

// ==== EX 42.3 | The division method and a bad table size | Core
// PROMPT: Write `DivisionHash(k, m int) int` = `k mod m` and `BucketsUsed[K any](keys []K, m int, hash func(K, int) int) int`. Hash the keys 0, 8, 16, ..., 7992 into m = 64 and into m = 61 and count how many buckets each uses.
// EXAMPLE: m = 64 -> only `8` buckets used; m = 61 (a prime) -> all `61` used
// BIGO: O(1) per key. Keys that share a factor with m crowd into few buckets; a prime m not near a power of two avoids the most common patterns.
// NOTE: Go's `%` keeps the sign of the dividend, so `((k % m) + m) % m` keeps negative keys in range.
func DivisionHash(k, m int) int { panic("todo") }

func BucketsUsed[K any](keys []K, m int, hash func(K, int) int) int {
	panic("todo")
}

// ==== END

// ==== EX 42.4 | Measure a hash function's spread | Core
// PROMPT: Write `MaxBucket[K any](keys []K, m int, hash func(K, int) int) int` returning the size of the fullest bucket. A good function keeps it close to len(keys) / m.
// EXAMPLE: 10,000 words into 1,000 buckets with `PolyHash` -> a fullest bucket around 20 to 25; with `SumHash`, more than twice that, because letter sums crowd into a narrow band of values
// BIGO: O(N). The fullest bucket bounds the worst lookup time with chaining.
func MaxBucket[K any](keys []K, m int, hash func(K, int) int) int {
	panic("todo")
}

// ==== END

// ==== EX 42.5 | Universal hashing | Core
// PROMPT: Write `type UniversalHash struct` with `NewUniversalHash(m int, rng *rand.Rand)` that picks random a in [1, p - 1] and b in [0, p - 1] once, with p = 2^61 - 1, and `Hash(k int) int` = `((a * k + b) mod p) mod m`.
// EXAMPLE: two instances with different seeds send the same key to different buckets; any fixed pair of keys collides with probability about 1 / m
// BIGO: O(1) per key. Because the function is chosen at random after the keys are fixed, no input can be bad for most choices.
// NOTE: `a * k` needs up to 122 bits. `bits.Mul64` returns the product as high and low words, and because 2^61 is 1 mod p, the reduction is a few shifts and adds.
const P = 1<<61 - 1

func MulModP(a, b uint64) uint64 {
	panic("todo")
}

type UniversalHash struct {
	m    int
	a, b uint64
}

func NewUniversalHash(m int, rng *rand.Rand) *UniversalHash {
	panic("todo")
}

func (u *UniversalHash) Hash(k int) int {
	panic("todo")
}

// ==== END

// ==== EX 42.6 | Adversarial keys break a fixed function | Practice
// PROMPT: Write `AdversarialKeys(m, count int) []int` returning keys that all collide under `k % m` (multiples of m), and `Fullest(keys []int, fn func(int) int) int`. Compare the division method with `UniversalHash`.
// EXAMPLE: m = 1,000, 5,000 keys -> division method `5000` in one bucket; universal hash only about `10`
// BIGO: Fixed function: O(N) per lookup on these keys. Universal: O(1 + N/m) expected on any keys.
func AdversarialKeys(m, count int) []int {
	panic("todo")
}

func Fullest(keys []int, fn func(int) int) int {
	panic("todo")
}

// ==== END

// ==== EX 42.7 | A hash table that re-randomizes when it grows | Practice
// PROMPT: Write `type RandomizedHashTable struct` with chaining over `int` keys and `string` values: `Put`, `Get` (returning `value, ok`), `Len` and `LongestChain`. When the load factor passes 1, double the bucket count and draw a fresh `UniversalHash` before rehashing.
// EXAMPLE: insert 10,000 adversarial keys -> every value is found and the fullest bucket stays small
// BIGO: O(1) expected per operation, O(1) amortized for resizing.
type kv struct {
	key   int
	value string
}

type RandomizedHashTable struct {
	rng     *rand.Rand
	buckets [][]kv
	h       *UniversalHash
	count   int
}

func NewRandomizedHashTable(rng *rand.Rand) *RandomizedHashTable {
	panic("todo")
}

func (t *RandomizedHashTable) newTable(capacity int) {
	panic("todo")
}

func (t *RandomizedHashTable) Put(key int, value string) {
	panic("todo")
}

func (t *RandomizedHashTable) Get(key int) (string, bool) {
	panic("todo")
}

func (t *RandomizedHashTable) Len() int { panic("todo") }

func (t *RandomizedHashTable) LongestChain() int {
	panic("todo")
}

// ==== END

// ==== EX 42.8 | The multiplication method | Stretch
// PROMPT: Write `MultHash(k, m int) int` = `floor(m * frac(k * A))` with A = 0.6180339887498949 (Knuth's golden-ratio choice). Unlike the division method it works well even when m is a power of two.
// EXAMPLE: the keys 0, 8, ..., 7992 into m = 64 -> all `64` buckets used, where `k % 64` used only 8
// BIGO: O(1).
func MultHash(k, m int) int {
	panic("todo")
}

// ==== END
