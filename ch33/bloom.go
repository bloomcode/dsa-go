// Package ch33: Bloom filters.
// Hashes come from SHA-256 so results are the same on every run and every machine.
package ch33

import (
	"crypto/sha256"
	"encoding/binary"
)

func twoHashes(item string) (uint64, uint64) {
	d := sha256.Sum256([]byte(item))
	return binary.BigEndian.Uint64(d[0:8]), binary.BigEndian.Uint64(d[8:16]) | 1
}

// ==== EX 33.1 | k hash functions from two | Core
// PROMPT: Write `BitPositions(item string, k, m int) []int` returning k bit positions in [0, m) using double hashing: `(h1 + i * h2) mod m` for i in 0..k-1, where h1 and h2 are the first two 8-byte words of one `sha256.Sum256` digest, read with `binary.BigEndian.Uint64`.
// EXAMPLE: `BitPositions("cat", 3, 32)` -> three positions in 0..31, the same on every run
// BIGO: O(k) after one hash computation. Kirsch and Mitzenmacher showed this behaves as well as k independent hashes.
// NOTE: With `uint64`, `h1 + i*h2` wraps around modulo 2^64 by definition and `%` is unsigned, so the result is never negative: no special remainder function is needed.
func BitPositions(item string, k, m int) []int {
	panic("todo")
}

// ==== END

// ==== EX 33.2 | A Bloom filter | Core
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

// ==== EX 33.3 | Predict the false-positive rate | Core
// PROMPT: Write `FalsePositiveRate(n, m, k int) float64` = `(1 - e^(-k n / m))^k`, the chance a never-added item looks present after n insertions.
// EXAMPLE: `FalsePositiveRate(1000, 10_000, 7)` -> about `0.0082` (under 1%)
// BIGO: O(1).
func FalsePositiveRate(n, m, k int) float64 {
	panic("todo")
}

// ==== END

// ==== EX 33.4 | Size a filter for a target error | Core
// PROMPT: Write `OptimalParams(n int, p float64) (m, k int)`: bits `m = ceil(-n ln p / (ln 2)^2)` and hash count `k = round((m / n) ln 2)`.
// EXAMPLE: `OptimalParams(1_000_000, 0.01)` -> `9585059, 7`: about 1.2 MB for a million items at 1% error
// BIGO: O(1). Each tenfold cut in p costs only about 4.8 more bits per item.
func OptimalParams(n int, p float64) (m, k int) {
	panic("todo")
}

// ==== END

// ==== EX 33.5 | Measure the real error | Core
// PROMPT: Write `MeasuredFPRate(n, m, k, probes int) float64`: add the strings "in-0" .. "in-{n-1}", then probe "out-0" .. "out-{probes-1}" (never added) and return the fraction reported present.
// EXAMPLE: `MeasuredFPRate(1000, 10_000, 7, 20_000)` lands close to the 13.3 prediction (within about 0.005)
// BIGO: O((n + probes) * k).
func MeasuredFPRate(n, m, k, probes int) float64 {
	panic("todo")
}

// ==== END

// ==== EX 33.6 | Duplicate detection in a stream | Core
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

// ==== EX 33.7 | Union of two filters | Practice
// PROMPT: Add `(b *BloomFilter) Union(o *BloomFilter) (*BloomFilter, error)` returning a new filter whose bits are the OR of both, for filters built with the same m and k (an error otherwise).
// EXAMPLE: filter A holds "x", filter B holds "y" -> the union might contain both
// BIGO: O(m / 64). Handy for merging filters built on separate machines.
func (b *BloomFilter) Union(o *BloomFilter) (*BloomFilter, error) {
	panic("todo")
}

// ==== END

// ==== EX 33.8 | A counting Bloom filter that can delete | Stretch
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
