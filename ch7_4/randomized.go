// Package ch7_4: Randomized algorithms.
package ch7_4

import (
	"iter"
	"math/rand/v2"
)

// ==== EX 7.4.1 | Randomized quicksort | Core
// PROMPT: Write `RandomQuicksort(arr []int, rng *rand.Rand) []int` returning a new sorted slice, choosing each pivot uniformly at random with `rng.IntN` and splitting into less, equal and more.
// EXAMPLE: `RandomQuicksort(0..9 in order, rand.New(rand.NewPCG(1, 0)))` -> `[0 1 2 3 4 5 6 7 8 9]`
// BIGO: O(N log N) expected on every input, including sorted ones. The worst case still exists but needs extraordinarily bad luck.
func RandomQuicksort(arr []int, rng *rand.Rand) []int {
	panic("todo")
}

// ==== END

// ==== EX 7.4.2 | Reproducible randomness with seeds | Core
// PROMPT: Write `SampleWithSeed(seed uint64, population []int, k int) []int` that returns the same k-item sample every time for the same seed, using its own `rand.New(rand.NewPCG(seed, 0))` and a partial shuffle of a copy.
// EXAMPLE: `SampleWithSeed(42, 0..99, 3)` equals itself on every call
// BIGO: O(N) to copy, O(k) swaps.
// NOTE: The top-level functions of `math/rand/v2` (`rand.IntN`) are randomly seeded on every run and cannot be seeded at all, so tests need their own generator.
func SampleWithSeed(seed uint64, population []int, k int) []int {
	panic("todo")
}

// ==== END

// ==== EX 7.4.3 | A linear congruential generator | Core
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

// ==== EX 7.4.4 | Fisher-Yates shuffle | Core
// PROMPT: Write `FisherYates(arr []int, rng *rand.Rand)` that shuffles in place: for i from the last index down to 1, swap arr[i] with arr[j] for a random j in 0..i.
// EXAMPLE: `FisherYates([]int{1, 2, 3, 4, 5}, rand.New(rand.NewPCG(7, 0)))` leaves some ordering of 1..5, the same one every time for seed 7
// BIGO: O(N) time, O(1) extra space, and every permutation is equally likely. `rng.Shuffle` is the library version.
// NOTE: `rng.IntN(i + 1)` returns 0..i. Using `IntN(i)` (never swapping an item with itself) is a classic bug that rules out some orders.
func FisherYates(arr []int, rng *rand.Rand) {
	panic("todo")
}

// ==== END

// ==== EX 7.4.5 | Prove the naive shuffle is biased | Practice
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

// ==== EX 7.4.6 | Randomize before building a BST | Core
// PROMPT: Write `BSTHeight(keys []int) int` that inserts keys in order into a plain BST, stored as `left` and `right` index slices, and returns its height. Compare sorted keys with the same keys shuffled.
// EXAMPLE: `BSTHeight(0..199 in order)` -> `200`; the same keys shuffled with seed 3 -> `14`
// BIGO: Height O(N) for sorted input, O(log N) expected after shuffling.
func BSTHeight(keys []int) int {
	panic("todo")
}

// ==== END

// ==== EX 7.4.7 | Reservoir sampling | Practice
// PROMPT: Write `Reservoir(stream iter.Seq[int], k int, rng *rand.Rand) []int` that picks k items uniformly from a stream of unknown length in one pass: keep the first k, then replace a random slot with item i (0-based) with probability k / (i + 1).
// EXAMPLE: `Reservoir(seq of 0..999999, 3, rng)` -> 3 values, each equally likely to be any of the million
// BIGO: O(N) time, O(k) memory. `iter.Seq` (Go 1.23) lets the stream be a generator that never builds the whole list.
func Reservoir(stream iter.Seq[int], k int, rng *rand.Rand) []int {
	panic("todo")
}

// ==== END

// ==== EX 7.4.8 | Load balancing with random choices | Core
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

// ==== EX 7.4.9 | Weighted random choice | Practice
// PROMPT: Write `MakeWeightedPicker(weights []float64, rng *rand.Rand) func() int` that precomputes prefix sums and returns a function picking index i with probability weights[i] / total, using `sort.Search` for the first prefix sum greater than `rng.Float64() * total`.
// EXAMPLE: weights `[1 3 6]` -> over 10,000 picks index 2 comes up about 60% of the time
// BIGO: O(N) setup, O(log N) per pick.
func MakeWeightedPicker(weights []float64, rng *rand.Rand) func() int {
	panic("todo")
}

// ==== END

// ==== EX 7.4.10 | Estimate pi by throwing darts | Core
// PROMPT: Write `EstimatePi(samples int, rng *rand.Rand) float64`: pick random points in the unit square and return 4 times the fraction that land inside the quarter circle.
// EXAMPLE: `EstimatePi(100_000, rand.New(rand.NewPCG(0, 0)))` -> about `3.14` (within 0.02)
// BIGO: O(S) samples; the error shrinks like 1 / sqrt(S), so 100 times more samples buy one more correct digit.
func EstimatePi(samples int, rng *rand.Rand) float64 {
	panic("todo")
}

// ==== END

// ==== EX 7.4.11 | Averages from a random sample | Core
// PROMPT: Write `SampledMean(data []int, k int, rng *rand.Rand) float64` estimating the mean of a huge slice from k random positions (with replacement), and `ExactMean(data []int) float64` for comparison.
// EXAMPLE: one million values, k = 2,000 -> the estimate lands within about 2% of the true mean
// BIGO: O(k) instead of O(N); accuracy depends on k and the data's spread, not on N.
func SampledMean(data []int, k int, rng *rand.Rand) float64 {
	panic("todo")
}

func ExactMean(data []int) float64 {
	panic("todo")
}

// ==== END

// ==== EX 7.4.12 | Monte Carlo versus Las Vegas | Core
// PROMPT: In a slice where half the entries are 1, write `LasVegasFindOne(arr []int, rng *rand.Rand) int` (probe random positions until a 1 appears; always right, time varies) and `MonteCarloFindOne(arr []int, rng *rand.Rand, tries int) int` (give up after `tries` probes and return -1; time fixed, may fail).
// EXAMPLE: both return an index holding 1; Monte Carlo with 20 tries fails with probability (1/2)^20, about one in a million
// BIGO: Las Vegas: 2 probes expected. Monte Carlo: exactly `tries` probes at most.
func LasVegasFindOne(arr []int, rng *rand.Rand) int {
	panic("todo")
}

func MonteCarloFindOne(arr []int, rng *rand.Rand, tries int) int {
	panic("todo")
}

// ==== END

// ==== EX 7.4.13 | The trial-division baseline | Core
// PROMPT: Write `IsPrimeTrial(n uint64) bool` checking odd divisors up to sqrt(n). Time it on a 12-digit prime and note why it is too slow for 100-digit numbers.
// EXAMPLE: `IsPrimeTrial(999_999_000_001)` -> `true` after about 500,000 divisions
// BIGO: O(sqrt N): exponential in the number of digits.
func IsPrimeTrial(n uint64) bool {
	panic("todo")
}

// ==== END

// ==== EX 7.4.14 | Fermat's little theorem, one base | Core
// PROMPT: If n is prime, a^(n-1) mod n is 1 for every a not divisible by n. Write `PowMod(a, e, n uint64) uint64` by repeated squaring and `FermatWitness(a, n uint64) bool` returning `true` when a proves n composite.
// EXAMPLE: `FermatWitness(2, 15)` -> `true` (15 is composite); `FermatWitness(2, 13)` -> `false`
// BIGO: O(log n) multiplications.
// NOTE: For n near 10^18, `a * b % n` overflows 64 bits. `bits.Mul64` gives the full 128-bit product as two words and `bits.Rem64` reduces it, with no big integers needed.
func MulMod(a, b, n uint64) uint64 {
	panic("todo")
}

func PowMod(a, e, n uint64) uint64 {
	panic("todo")
}

func FermatWitness(a, n uint64) bool { panic("todo") }

// ==== END

// ==== EX 7.4.15 | Fermat's primality test | Core
// PROMPT: Write `FermatTest(n uint64, k int, rng *rand.Rand) bool` trying k random bases in [2, n - 2] and answering "probably prime" (`true`) if none is a witness. Then show 561, a Carmichael number, fools every base coprime to it.
// EXAMPLE: `FermatTest(1_000_000_007, 20, rng)` -> `true`; `FermatTest(1_000_000_007 * 998_244_353, 20, rng)` -> `false`; `FermatWitness(2, 561)` -> `false` although 561 = 3 x 11 x 17
// BIGO: O(k log n). A Monte Carlo algorithm: "composite" is always right, "prime" can be wrong.
func FermatTest(n uint64, k int, rng *rand.Rand) bool {
	panic("todo")
}

// ==== END

// ==== EX 7.4.16 | Miller-Rabin, which Carmichael numbers cannot fool | Stretch
// PROMPT: Write `MillerRabin(n uint64) bool` with the bases 2, 3, 5, ..., 37 (the first 12 primes): write n - 1 = d x 2^s, and for each base check whether a^d is 1 or reaches n - 1 by repeated squaring. With these bases it is exact for every 64-bit n.
// EXAMPLE: `MillerRabin(561)` -> `false`; `MillerRabin(1<<61 - 1)` -> `true`
// BIGO: O(k log^3 n) bit operations; effectively instant for 64-bit numbers.
// NOTE: `bits.TrailingZeros64(n - 1)` gives s in one instruction.
var bases = []uint64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37}

func MillerRabin(n uint64) bool {
	panic("todo")
}

// ==== END

// ==== EX 7.4.17 | Simulate, then check against the formula | Practice
// PROMPT: Write `BirthdaySimulation(people, trials int, rng *rand.Rand) float64` estimating the chance that two of `people` share a birthday (365 days), and `BirthdayExact(people int) float64` from the product formula.
// EXAMPLE: 23 people -> exact `0.507...`; 20,000 simulated trials land within 0.02 of it
// BIGO: O(trials x people). When no formula exists, simulation is how you get the number.
func BirthdaySimulation(people, trials int, rng *rand.Rand) float64 {
	panic("todo")
}

func BirthdayExact(people int) float64 {
	panic("todo")
}

// ==== END
