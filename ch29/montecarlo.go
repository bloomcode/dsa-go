// Package ch29: Monte Carlo and Las Vegas algorithms.
package ch29

import (
	"math/rand/v2"
)

// ==== EX 29.1 | Estimate pi by throwing darts | Core
// PROMPT: Write `EstimatePi(samples int, rng *rand.Rand) float64`: pick random points in the unit square and return 4 times the fraction that land inside the quarter circle.
// EXAMPLE: `EstimatePi(100_000, rand.New(rand.NewPCG(0, 0)))` -> about `3.14` (within 0.02)
// BIGO: O(S) samples; the error shrinks like 1 / sqrt(S), so 100 times more samples buy one more correct digit.
func EstimatePi(samples int, rng *rand.Rand) float64 {
	panic("todo")
}

// ==== END

// ==== EX 29.2 | Averages from a random sample | Core
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

// ==== EX 29.3 | Monte Carlo versus Las Vegas | Core
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

// ==== EX 29.4 | The trial-division baseline | Core
// PROMPT: Write `IsPrimeTrial(n uint64) bool` checking odd divisors up to sqrt(n). Time it on a 12-digit prime and note why it is too slow for 100-digit numbers.
// EXAMPLE: `IsPrimeTrial(999_999_000_001)` -> `true` after about 500,000 divisions
// BIGO: O(sqrt N): exponential in the number of digits.
func IsPrimeTrial(n uint64) bool {
	panic("todo")
}

// ==== END

// ==== EX 29.5 | Fermat's little theorem, one base | Core
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

// ==== EX 29.6 | Fermat's primality test | Core
// PROMPT: Write `FermatTest(n uint64, k int, rng *rand.Rand) bool` trying k random bases in [2, n - 2] and answering "probably prime" (`true`) if none is a witness. Then show 561, a Carmichael number, fools every base coprime to it.
// EXAMPLE: `FermatTest(1_000_000_007, 20, rng)` -> `true`; `FermatTest(1_000_000_007 * 998_244_353, 20, rng)` -> `false`; `FermatWitness(2, 561)` -> `false` although 561 = 3 x 11 x 17
// BIGO: O(k log n). A Monte Carlo algorithm: "composite" is always right, "prime" can be wrong.
func FermatTest(n uint64, k int, rng *rand.Rand) bool {
	panic("todo")
}

// ==== END

// ==== EX 29.7 | Miller-Rabin, which Carmichael numbers cannot fool | Stretch
// PROMPT: Write `MillerRabin(n uint64) bool` with the bases 2, 3, 5, ..., 37 (the first 12 primes): write n - 1 = d x 2^s, and for each base check whether a^d is 1 or reaches n - 1 by repeated squaring. With these bases it is exact for every 64-bit n.
// EXAMPLE: `MillerRabin(561)` -> `false`; `MillerRabin(1<<61 - 1)` -> `true`
// BIGO: O(k log^3 n) bit operations; effectively instant for 64-bit numbers.
// NOTE: `bits.TrailingZeros64(n - 1)` gives s in one instruction.
var bases = []uint64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37}

func MillerRabin(n uint64) bool {
	panic("todo")
}

// ==== END

// ==== EX 29.8 | Simulate, then check against the formula | Practice
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
