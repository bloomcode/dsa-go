// Package ch83: Compact sets: bit vectors and Bloom filters.
package ch83

import (
	"crypto/sha256"
	"encoding/binary"
	"math/big"
)

// ==== EX 83.1 | A set as a boolean array | Core
// PROMPT: Write `type BoolArraySet struct` for integers 0..universe-1 backed by a `[]bool`, with `Add`, `Remove`, `Contains` and `Len`.
// EXAMPLE: universe 10: add 3, add 7, remove 3 -> `Contains(7)` is `true`, `Contains(3)` is `false`, `Len()` is `1`
// BIGO: O(1) per operation, but a whole byte per possible value: Go does not pack `[]bool` into bits.
type BoolArraySet struct {
	flags []bool
	count int
}

func NewBoolArraySet(universe int) *BoolArraySet {
	panic("todo")
}

func (s *BoolArraySet) Add(x int) {
	panic("todo")
}

func (s *BoolArraySet) Remove(x int) {
	panic("todo")
}

func (s *BoolArraySet) Contains(x int) bool { panic("todo") }

func (s *BoolArraySet) Len() int { panic("todo") }

// ==== END

// ==== EX 83.2 | A bit vector with masks | Core
// PROMPT: Write `type BitVector struct` on a `Words []uint64` of ceil(size / 64) words. Value x lives in word `x >> 6`, bit `x & 63`. Implement `Add` (OR with the mask), `Remove` (AND NOT with the mask, Go's `&^`) and `Contains` (AND, then test).
// EXAMPLE: `bv := NewBitVector(100); bv.Add(42)` -> `bv.Contains(42)` is `true`; `bv.Words[0]` is `4398046511104` (bit 42 of word 0)
// BIGO: O(1) per operation, and 1 bit per possible value: 8 times less than a `[]bool`.
// NOTE: Go's `&^` (AND NOT, "bit clear") does in one operator what other languages write as `& ~mask`.
type BitVector struct {
	Size  int
	Words []uint64
}

func NewBitVector(size int) *BitVector {
	panic("todo")
}

func (b *BitVector) Add(x int) { panic("todo") }

func (b *BitVector) Remove(x int) { panic("todo") }

func (b *BitVector) Contains(x int) bool { panic("todo") }

// ==== END

// ==== EX 83.3 | Count the members | Core
// PROMPT: Add `Count() int` that counts set bits across all words with `bits.OnesCount64`.
// EXAMPLE: add 1, 5, 99 -> `Count()` is `3`
// BIGO: O(size / 64), and `OnesCount64` compiles to a single POPCNT instruction where the CPU has one. Keep a running counter instead if you need O(1).
func (b *BitVector) Count() int {
	panic("todo")
}

// ==== END

// ==== EX 83.4 | Union, intersection and difference | Core
// PROMPT: Add `Union`, `Intersection` and `Difference`, each returning a new `*BitVector` of the same size by combining the words with `|`, `&` and `&^`.
// EXAMPLE: {1, 2, 3} and {2, 3, 4} -> union `{1, 2, 3, 4}`, intersection `{2, 3}`, difference `{1}`
// BIGO: O(size / 64): sixty-four set members combined per word operation.
func (b *BitVector) combine(o *BitVector, op func(x, y uint64) uint64) *BitVector {
	panic("todo")
}

func (b *BitVector) Union(o *BitVector) *BitVector {
	panic("todo")
}

func (b *BitVector) Intersection(o *BitVector) *BitVector {
	panic("todo")
}

func (b *BitVector) Difference(o *BitVector) *BitVector {
	panic("todo")
}

// ==== END

// ==== EX 83.5 | Iterate over members | Core
// PROMPT: Add `Members() []int` returning members in increasing order, skipping zero words entirely and peeling off the lowest set bit with `bits.TrailingZeros64` and `w & (w - 1)`.
// EXAMPLE: add 9, 2, 63 -> `Members()` is `[2 9 63]`
// BIGO: O(size / 64 + members).
func (b *BitVector) Members() []int {
	panic("todo")
}

// ==== END

// ==== EX 83.6 | Estimate the space | Core
// PROMPT: Write `SpaceReport(n int) map[string]int` estimating the bytes used by a `map[int]struct{}` of 0..n-1, a `[]bool` of length n, and a `BitVector` of size n. Model the map as Go 1.24's Swiss table: 9 bytes per slot (an 8-byte key plus 1 control byte), with the slot count the next power of two at or above n * 8/7. Each slice adds a 24-byte header.
// EXAMPLE: n = 1,000,000 -> map about 19 MB, `[]bool` `1000024` bytes, bit vector `125024` bytes
// BIGO: O(N) each, but the constants differ by a factor of about 150.
// NOTE: These are models, not measurements. `runtime.ReadMemStats` before and after building each structure gives the real figures.
func SpaceReport(n int) map[string]int {
	panic("todo")
}

// ==== END

// ==== EX 83.7 | Find duplicates in a bounded range | Core
// PROMPT: Write `FirstDuplicate(nums []int, universe int) (int, bool)` returning the first value seen twice, using a `BitVector` as the "seen" set, with ok false when there is none.
// EXAMPLE: `FirstDuplicate([]int{4, 9, 1, 9, 4}, 10)` -> `9, true`; no repeats -> `0, false`
// BIGO: O(N) time, universe / 8 bytes of space, no matter how many numbers there are.
func FirstDuplicate(nums []int, universe int) (int, bool) {
	panic("todo")
}

// ==== END

// ==== EX 83.8 | A built-in type as a growable bit set | Practice
// PROMPT: Go's standard library has no bit-set type, but a `big.Int` is an arbitrarily long run of bits. Write `UnionList(a, b []int) []int` and `CommonCount(a, b []int) int` with `big.Int`: `SetBit`, `Or`, `And`, `BitLen` and `Bit` to list members, and `bits.OnesCount` over `Bits()` to count them.
// EXAMPLE: `a = [1 5]`, `b = [5 7]` -> `UnionList` is `[1 5 7]`; `CommonCount` is `1`
// BIGO: Word-at-a-time operations, O(size / 64), and the set grows as needed.
// NOTE: `z.Or(x, y)` stores the result in z and returns it, so a fresh `new(big.Int)` as the receiver leaves both inputs unchanged.
func toBits(items []int) *big.Int {
	panic("todo")
}

func UnionList(a, b []int) []int {
	panic("todo")
}

func CommonCount(a, b []int) int {
	panic("todo")
}

// ==== END

// ==== EX 83.9 | Missing numbers from a range | Practice
// PROMPT: Write `Missing(nums []int, n int) []int` returning every value in 0..n-1 that does not appear in `nums`, marking a `BitVector` and then scanning it.
// EXAMPLE: `Missing([]int{0, 3, 1, 6}, 7)` -> `[2 4 5]`
// BIGO: O(N + n) time, n / 8 bytes of space.
func Missing(nums []int, n int) []int {
	panic("todo")
}

// ==== END

func twoHashes(item string) (uint64, uint64) {
	d := sha256.Sum256([]byte(item))
	return binary.BigEndian.Uint64(d[0:8]), binary.BigEndian.Uint64(d[8:16]) | 1
}

// ==== EX 83.10 | k hash functions from two | Core
// PROMPT: Write `BitPositions(item string, k, m int) []int` returning k bit positions in [0, m) using double hashing: `(h1 + i * h2) mod m` for i in 0..k-1, where h1 and h2 are the first two 8-byte words of one `sha256.Sum256` digest, read with `binary.BigEndian.Uint64`.
// EXAMPLE: `BitPositions("cat", 3, 32)` -> three positions in 0..31, the same on every run
// BIGO: O(k) after one hash computation. Kirsch and Mitzenmacher showed this behaves as well as k independent hashes.
// NOTE: With `uint64`, `h1 + i*h2` wraps around modulo 2^64 by definition and `%` is unsigned, so the result is never negative: no special remainder function is needed.
func BitPositions(item string, k, m int) []int {
	panic("todo")
}

// ==== END

// ==== EX 83.11 | A Bloom filter | Core
// PROMPT: Write `type BloomFilter struct` with an m-bit `[]uint64`, `Add(item string)` setting the item's k bits, and `MightContain(item string) bool` returning `true` only if all k bits are set.
// EXAMPLE: add "apple" and "pear" -> `MightContain("apple")` is `true`; `MightContain("grape")` is almost certainly `false`
// BIGO: O(k) per operation; m bits of memory regardless of how large the items are.
// NOTE: "No" is always right. "Yes" means "probably": other items may have set those same bits.
type BloomFilter struct {
	m, k  int
	bits  []uint64
	Count int
}

func NewBloomFilter(m, k int) *BloomFilter {
	panic("todo")
}

func (b *BloomFilter) Add(item string) {
	panic("todo")
}

func (b *BloomFilter) MightContain(item string) bool {
	panic("todo")
}

// ==== END

// ==== EX 83.12 | Predict the false-positive rate | Core
// PROMPT: Write `FalsePositiveRate(n, m, k int) float64` = `(1 - e^(-k n / m))^k`, the chance a never-added item looks present after n insertions.
// EXAMPLE: `FalsePositiveRate(1000, 10_000, 7)` -> about `0.0082` (under 1%)
// BIGO: O(1).
func FalsePositiveRate(n, m, k int) float64 {
	panic("todo")
}

// ==== END

// ==== EX 83.13 | Size a filter for a target error | Core
// PROMPT: Write `OptimalParams(n int, p float64) (m, k int)`: bits `m = ceil(-n ln p / (ln 2)^2)` and hash count `k = round((m / n) ln 2)`.
// EXAMPLE: `OptimalParams(1_000_000, 0.01)` -> `9585059, 7`: about 1.2 MB for a million items at 1% error
// BIGO: O(1). Each tenfold cut in p costs only about 4.8 more bits per item.
func OptimalParams(n int, p float64) (m, k int) {
	panic("todo")
}

// ==== END

// ==== EX 83.14 | Measure the real error | Core
// PROMPT: Write `MeasuredFPRate(n, m, k, probes int) float64`: add the strings "in-0" .. "in-{n-1}", then probe "out-0" .. "out-{probes-1}" (never added) and return the fraction reported present.
// EXAMPLE: `MeasuredFPRate(1000, 10_000, 7, 20_000)` lands close to the 13.3 prediction (within about 0.005)
// BIGO: O((n + probes) * k).
func MeasuredFPRate(n, m, k, probes int) float64 {
	panic("todo")
}

// ==== END

// ==== EX 83.15 | Duplicate detection in a stream | Core
// PROMPT: Write `StreamDuplicates(items []string, m, k int) []string` that reports items the filter says it has seen before. Then write `FalseAlarms(items []string, m, k int) int` that double-checks each report against an exact `map[string]bool` and counts how many reports were false.
// EXAMPLE: a stream of 5,000 unique items plus 100 repeats with m = 64,000, k = 6 -> all 100 repeats reported, and only a handful of false alarms
// BIGO: O(N * k) time, m bits of space. Real systems (web crawlers, databases) use the filter to skip most exact lookups.
func StreamDuplicates(items []string, m, k int) []string {
	panic("todo")
}

func FalseAlarms(items []string, m, k int) int {
	panic("todo")
}

// ==== END

// ==== EX 83.16 | Union of two filters | Practice
// PROMPT: Add `(b *BloomFilter) Union(o *BloomFilter) (*BloomFilter, error)` returning a new filter whose bits are the OR of both, for filters built with the same m and k (an error otherwise).
// EXAMPLE: filter A holds "x", filter B holds "y" -> the union might contain both
// BIGO: O(m / 64). Handy for merging filters built on separate machines.
func (b *BloomFilter) Union(o *BloomFilter) (*BloomFilter, error) {
	panic("todo")
}

// ==== END

// ==== EX 83.17 | A counting Bloom filter that can delete | Stretch
// PROMPT: Write `type CountingBloomFilter struct` with a slice of small counters instead of bits: `Add` increments, `Remove` decrements (only for items you added) and reports whether it did, and `MightContain` checks all counters are above zero.
// EXAMPLE: add "a", add "b", remove "a" -> `MightContain("a")` is `false` (barring collisions), `MightContain("b")` is `true`
// BIGO: O(k) per operation, at the cost of 4 to 8 bits per counter instead of 1 (here a `uint8` per counter).
type CountingBloomFilter struct {
	m, k   int
	counts []uint8
}

func NewCountingBloomFilter(m, k int) *CountingBloomFilter {
	panic("todo")
}

func (c *CountingBloomFilter) Add(item string) {
	panic("todo")
}

func (c *CountingBloomFilter) Remove(item string) bool {
	panic("todo")
}

func (c *CountingBloomFilter) MightContain(item string) bool {
	panic("todo")
}

// ==== END
