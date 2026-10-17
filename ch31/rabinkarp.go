// Package ch31: Substring search with rolling hashes.
// Text is lowercase ASCII; 'a'..'z' map to 0..25 so base 26 is exact.
package ch31

import "math/rand/v2"

const (
	Base  = 26
	Prime = 1_000_000_007
)

func code(ch byte) int { return int(ch - 'a') }

func powMod(b, e, m int) int {
	r := 1 % m
	b %= m
	for e > 0 {
		if e&1 == 1 {
			r = r * b % m
		}
		b = b * b % m
		e >>= 1
	}
	return r
}

// ==== EX 31.1 | Brute-force substring search | Core
// PROMPT: Write `BruteForceSearch(text, needle string) []int` returning every index where `needle` starts, comparing byte by byte at each position.
// EXAMPLE: `BruteForceSearch("abracadabra", "abra")` -> `[0 7]`
// BIGO: O(N * M) worst case, e.g. text `"aaaa...a"` and needle `"aa...ab"`.
func BruteForceSearch(text, needle string) []int {
	panic("todo")
}

// ==== END

// ==== EX 31.2 | Sliding a window in O(1) per step | Core
// PROMPT: Write `WindowSums(arr []int, k int) []int` returning the sum of every window of k consecutive values, adding the entering value and subtracting the leaving one.
// EXAMPLE: `WindowSums([]int{1, 2, 3, 4, 5}, 3)` -> `[6 9 12]`
// BIGO: O(N) instead of O(N * k). Rabin-Karp applies exactly this idea to a hash.
func WindowSums(arr []int, k int) []int {
	panic("todo")
}

// ==== END

// ==== EX 31.3 | Rolling base-26 hashes | Core
// PROMPT: Write `RollingHashes(text string, m int) []int` returning the hash of every length-m window, treating each window as a base-26 number mod Prime. Roll with: remove the leading digit times 26^(m-1), multiply by 26, add the new digit.
// EXAMPLE: `RollingHashes("abcd", 2)` -> `[1 28 55]` (ab = 0*26 + 1, bc = 1*26 + 2, cd = 2*26 + 3)
// BIGO: O(N) for all windows.
// NOTE: After subtracting the leading digit the value can go negative, and Go's `%` keeps the sign. Add Prime before reducing.
func RollingHashes(text string, m int) []int {
	panic("todo")
}

func rolling(text string, m, prime int) []int {
	panic("todo")
}

// ==== END

// ==== EX 31.4 | Rabin-Karp, verified (Las Vegas) | Core
// PROMPT: Write `RabinKarp(text, needle string, prime int) []int` that compares rolling hashes and confirms each hash match with a real comparison (`text[i:i+m] == needle`), so the answer is always correct.
// EXAMPLE: `RabinKarp("abracadabra", "abra", Prime)` -> `[0 7]`
// BIGO: O(N + M) expected; O(N * M) only if hash matches are spuriously frequent.
// NOTE: Slicing a string never copies, so `text[i:i+m] == needle` is an allocation-free comparison.
func RabinKarp(text, needle string, prime int) []int {
	panic("todo")
}

func search(text, needle string, prime int, verify bool) []int {
	panic("todo")
}

// ==== END

// ==== EX 31.5 | Spurious hits with a tiny modulus | Core
// PROMPT: Write `SpuriousHits(text, needle string, prime int) int` counting windows whose hash equals the needle's but whose text does not. Compare prime = 7 with prime = 10^9 + 7 on a long text.
// EXAMPLE: 10,000 random letters, needle "abcde" -> prime 7 gives over a thousand spurious hits; 10^9 + 7 gives `0`
// BIGO: O(N). Expected spurious hits are about N / prime, which is why the modulus is large.
func SpuriousHits(text, needle string, prime int) int {
	panic("todo")
}

// ==== END

// ==== EX 31.6 | Rabin-Karp without verification (Monte Carlo) | Practice
// PROMPT: Write `RabinKarpMonteCarlo(text, needle string, rng *rand.Rand) []int` that picks a random prime modulus between 10^9 and 2 x 10^9 for this run and reports every hash match without checking the text.
// EXAMPLE: on ordinary text it returns the same indices as 31.4, in strictly O(N + M) time
// BIGO: O(N + M) worst case, with a small chance of a false match. 31.4 trades that chance for a verification step.
func isPrime(n int) bool {
	panic("todo")
}

func RabinKarpMonteCarlo(text, needle string, rng *rand.Rand) []int {
	panic("todo")
}

// ==== END

// ==== EX 31.7 | Many needles of one length | Practice
// PROMPT: Write `type Hit struct{ Index int; Needle string }` and `FindAny(text string, needles []string) []Hit` for needles that all share one length: put their hashes in a map and slide one rolling hash over the text.
// EXAMPLE: `FindAny("thecatsatonthemat", []string{"cat", "mat", "dog"})` -> `[{3 cat} {14 mat}]`
// BIGO: O(N + K * M) for K needles, instead of K separate searches.
type Hit struct {
	Index  int
	Needle string
}

func FindAny(text string, needles []string) []Hit {
	panic("todo")
}

// ==== END

// ==== EX 31.8 | Longest repeated substring | Stretch
// PROMPT: Write `LongestRepeat(text string) string` returning the longest substring that occurs at least twice. Binary-search the length L; for each L, check for a repeated rolling hash (verifying matches).
// EXAMPLE: `LongestRepeat("banana")` -> `"ana"`; `LongestRepeat("abcd")` -> `""`
// BIGO: O(N log N) expected.
func LongestRepeat(text string) string {
	panic("todo")
}

func repeatOfLength(text string, n int) (string, bool) {
	panic("todo")
}

// ==== END
