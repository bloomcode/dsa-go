// Package ch32: Bit vectors: one bit per member.
package ch32

import (
	"math/big"
)

// ==== EX 32.1 | A set as a boolean array | Core
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

// ==== EX 32.2 | A bit vector with masks | Core
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

// ==== EX 32.3 | Count the members | Core
// PROMPT: Add `Count() int` that counts set bits across all words with `bits.OnesCount64`.
// EXAMPLE: add 1, 5, 99 -> `Count()` is `3`
// BIGO: O(size / 64), and `OnesCount64` compiles to a single POPCNT instruction where the CPU has one. Keep a running counter instead if you need O(1).
func (b *BitVector) Count() int {
	panic("todo")
}

// ==== END

// ==== EX 32.4 | Union, intersection and difference | Core
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

// ==== EX 32.5 | Iterate over members | Core
// PROMPT: Add `Members() []int` returning members in increasing order, skipping zero words entirely and peeling off the lowest set bit with `bits.TrailingZeros64` and `w & (w - 1)`.
// EXAMPLE: add 9, 2, 63 -> `Members()` is `[2 9 63]`
// BIGO: O(size / 64 + members).
func (b *BitVector) Members() []int {
	panic("todo")
}

// ==== END

// ==== EX 32.6 | Estimate the space | Core
// PROMPT: Write `SpaceReport(n int) map[string]int` estimating the bytes used by a `map[int]struct{}` of 0..n-1, a `[]bool` of length n, and a `BitVector` of size n. Model the map as Go 1.24's Swiss table: 9 bytes per slot (an 8-byte key plus 1 control byte), with the slot count the next power of two at or above n * 8/7. Each slice adds a 24-byte header.
// EXAMPLE: n = 1,000,000 -> map about 19 MB, `[]bool` `1000024` bytes, bit vector `125024` bytes
// BIGO: O(N) each, but the constants differ by a factor of about 150.
// NOTE: These are models, not measurements. `runtime.ReadMemStats` before and after building each structure gives the real figures.
func SpaceReport(n int) map[string]int {
	panic("todo")
}

// ==== END

// ==== EX 32.7 | Find duplicates in a bounded range | Core
// PROMPT: Write `FirstDuplicate(nums []int, universe int) (int, bool)` returning the first value seen twice, using a `BitVector` as the "seen" set, with ok false when there is none.
// EXAMPLE: `FirstDuplicate([]int{4, 9, 1, 9, 4}, 10)` -> `9, true`; no repeats -> `0, false`
// BIGO: O(N) time, universe / 8 bytes of space, no matter how many numbers there are.
func FirstDuplicate(nums []int, universe int) (int, bool) {
	panic("todo")
}

// ==== END

// ==== EX 32.8 | A built-in type as a growable bit set | Practice
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

// ==== EX 32.9 | Missing numbers from a range | Practice
// PROMPT: Write `Missing(nums []int, n int) []int` returning every value in 0..n-1 that does not appear in `nums`, marking a `BitVector` and then scanning it.
// EXAMPLE: `Missing([]int{0, 3, 1, 6}, 7)` -> `[2 4 5]`
// BIGO: O(N + n) time, n / 8 bytes of space.
func Missing(nums []int, n int) []int {
	panic("todo")
}

// ==== END
