// Package ch23: Randomness as a tool.
// Every function that needs randomness takes a *rand.Rand from math/rand/v2, so results are
// reproducible in tests: rand.New(rand.NewPCG(seed, 0)).
package ch23

import (
	"iter"
	"math/rand/v2"
)

// ==== EX 23.1 | Randomized quicksort | Core
// PROMPT: Write `RandomQuicksort(arr []int, rng *rand.Rand) []int` returning a new sorted slice, choosing each pivot uniformly at random with `rng.IntN` and splitting into less, equal and more.
// EXAMPLE: `RandomQuicksort(0..9 in order, rand.New(rand.NewPCG(1, 0)))` -> `[0 1 2 3 4 5 6 7 8 9]`
// BIGO: O(N log N) expected on every input, including sorted ones. The worst case still exists but needs extraordinarily bad luck.
func RandomQuicksort(arr []int, rng *rand.Rand) []int {
	panic("todo")
}

// ==== END

// ==== EX 23.2 | Reproducible randomness with seeds | Core
// PROMPT: Write `SampleWithSeed(seed uint64, population []int, k int) []int` that returns the same k-item sample every time for the same seed, using its own `rand.New(rand.NewPCG(seed, 0))` and a partial shuffle of a copy.
// EXAMPLE: `SampleWithSeed(42, 0..99, 3)` equals itself on every call
// BIGO: O(N) to copy, O(k) swaps.
// NOTE: The top-level functions of `math/rand/v2` (`rand.IntN`) are randomly seeded on every run and cannot be seeded at all, so tests need their own generator.
func SampleWithSeed(seed uint64, population []int, k int) []int {
	panic("todo")
}

// ==== END

// ==== EX 23.3 | A linear congruential generator | Core
// PROMPT: Write `type LCG struct` with `Next() uint64` applying `x = (a*x + c) % m` (defaults a = 1103515245, c = 12345, m = 2^31 via `NewLCG(seed)`). Then write `LCGPeriod(seed, a, c, m uint64) int` and find the period of the tiny generator a = 5, c = 3, m = 16.
// EXAMPLE: first three values of `LCG{X: 1, A: 5, C: 3, M: 16}` -> `8, 11, 10`; its period is `16`
// BIGO: O(1) per number. This is a PRNG: fully determined by the seed, which is why it is reproducible and why it is not secure.
type LCG struct{ X, A, C, M uint64 }

func NewLCG(seed uint64) *LCG { panic("todo") }

func (g *LCG) Next() uint64 {
	panic("todo")
}

func LCGPeriod(seed, a, c, m uint64) int {
	panic("todo")
}

// ==== END

// ==== EX 23.4 | Fisher-Yates shuffle | Core
// PROMPT: Write `FisherYates(arr []int, rng *rand.Rand)` that shuffles in place: for i from the last index down to 1, swap arr[i] with arr[j] for a random j in 0..i.
// EXAMPLE: `FisherYates([]int{1, 2, 3, 4, 5}, rand.New(rand.NewPCG(7, 0)))` leaves some ordering of 1..5, the same one every time for seed 7
// BIGO: O(N) time, O(1) extra space, and every permutation is equally likely. `rng.Shuffle` is the library version.
// NOTE: `rng.IntN(i + 1)` returns 0..i. Using `IntN(i)` (never swapping an item with itself) is a classic bug that rules out some orders.
func FisherYates(arr []int, rng *rand.Rand) {
	panic("todo")
}

// ==== END

// ==== EX 23.5 | Prove the naive shuffle is biased | Practice
// PROMPT: The naive shuffle swaps each position with any position 0..n-1. Write `NaiveDistribution(n int) map[string]int` and `FisherYatesDistribution(n int) map[string]int` that enumerate every possible sequence of random choices recursively and count the resulting permutations (keyed like "012").
// EXAMPLE: For n = 3: Fisher-Yates gives each of the 6 permutations exactly once; naive spreads 27 outcomes as `map[012:4 021:5 102:5 120:5 201:4 210:4]`
// BIGO: O(n^n) and O(n!) sequences: exact enumeration, no randomness needed.
// NOTE: 27 equally likely outcomes cannot split evenly over 6 permutations, so the naive shuffle must be biased.
func key(a []int) string {
	panic("todo")
}

func NaiveDistribution(n int) map[string]int {
	panic("todo")
}

func FisherYatesDistribution(n int) map[string]int {
	panic("todo")
}

// ==== END

// ==== EX 23.6 | Randomize before building a BST | Core
// PROMPT: Write `BSTHeight(keys []int) int` that inserts keys in order into a plain BST, stored as `left` and `right` index slices, and returns its height. Compare sorted keys with the same keys shuffled.
// EXAMPLE: `BSTHeight(0..199 in order)` -> `200`; the same keys shuffled with seed 3 -> `14`
// BIGO: Height O(N) for sorted input, O(log N) expected after shuffling.
func BSTHeight(keys []int) int {
	panic("todo")
}

// ==== END

// ==== EX 23.7 | Reservoir sampling | Practice
// PROMPT: Write `Reservoir(stream iter.Seq[int], k int, rng *rand.Rand) []int` that picks k items uniformly from a stream of unknown length in one pass: keep the first k, then replace a random slot with item i (0-based) with probability k / (i + 1).
// EXAMPLE: `Reservoir(seq of 0..999999, 3, rng)` -> 3 values, each equally likely to be any of the million
// BIGO: O(N) time, O(k) memory. `iter.Seq` (Go 1.23) lets the stream be a generator that never builds the whole list.
func Reservoir(stream iter.Seq[int], k int, rng *rand.Rand) []int {
	panic("todo")
}

// ==== END

// ==== EX 23.8 | Load balancing with random choices | Core
// PROMPT: Write `MaxLoadRandom(jobs, servers int, rng *rand.Rand) int` assigning each job to a random server, and `MaxLoadTwoChoices` picking two random servers and using the less loaded one. Return the busiest server's load.
// EXAMPLE: With 10,000 jobs on 1,000 servers, two choices gives a clearly lower maximum than one choice
// BIGO: O(J) for J jobs. The "power of two choices" drops the expected maximum excess from about log N / log log N to log log N.
func MaxLoadRandom(jobs, servers int, rng *rand.Rand) int {
	panic("todo")
}

func MaxLoadTwoChoices(jobs, servers int, rng *rand.Rand) int {
	panic("todo")
}

// ==== END

// ==== EX 23.9 | Weighted random choice | Practice
// PROMPT: Write `MakeWeightedPicker(weights []float64, rng *rand.Rand) func() int` that precomputes prefix sums and returns a function picking index i with probability weights[i] / total, using `sort.Search` for the first prefix sum greater than `rng.Float64() * total`.
// EXAMPLE: weights `[1 3 6]` -> over 10,000 picks index 2 comes up about 60% of the time
// BIGO: O(N) setup, O(log N) per pick.
func MakeWeightedPicker(weights []float64, rng *rand.Rand) func() int {
	panic("todo")
}

// ==== END
