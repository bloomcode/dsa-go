// Package ch7_1: Recursion.
package ch7_1

// ==== EX 7.1.1 | Count the digits | Core
// PROMPT: Write `DigitCount(n int) int` for `n >= 0` recursively: a number below 10 has one digit; otherwise it has one more digit than `n / 10`.
// EXAMPLE: `DigitCount(12345)` -> `5`; `DigitCount(0)` -> `1`
// BIGO: O(log N) calls: one per digit.
func DigitCount(n int) int {
	panic("todo")
}

// ==== END

// ==== EX 7.1.2 | Factorial | Core
// PROMPT: Write `Factorial(n int) int` recursively.
// EXAMPLE: `Factorial(5)` -> `120`; `Factorial(0)` -> `1`; `Factorial(20)` -> `2432902008176640000`
// BIGO: O(N).
// NOTE: 20! is the largest factorial that fits in a 64-bit int; use `math/big` beyond that.
func Factorial(n int) int {
	panic("todo")
}

// ==== END

// ==== EX 7.1.3 | Sum an array recursively | Core
// PROMPT: Write `SumRecursive(arr []int) int` that adds `arr[0]` to the sum of `arr[1:]`. Re-slicing is free in Go: it makes a new header, not a copy.
// EXAMPLE: `[1 2 3 4]` -> `10`
// BIGO: O(N) time and stack depth. Unlike Python's slicing, `arr[1:]` costs O(1).
func SumRecursive(arr []int) int {
	panic("todo")
}

// ==== END

// ==== EX 7.1.4 | Depth of nesting | Core
// PROMPT: Write `MaxDepth(nested []any) int` for a slice whose items are numbers or further `[]any` slices: a slice with no slices inside has depth 1, and each level of nesting adds one.
// EXAMPLE: `[1 [2 [3]] 4]` -> `3`; `[]` -> `1`
// BIGO: O(N) in the total number of items and lists. Recursion fits because the depth is unknown in advance.
func MaxDepth(nested []any) int {
	panic("todo")
}

// ==== END

// ==== EX 7.1.5 | Fast power | Practice
// PROMPT: Write `Power(x, n int) int` for n >= 0 by squaring: x^n = (x^(n/2))^2, times x if n is odd.
// EXAMPLE: `Power(2, 10)` -> `1024`; `Power(3, 0)` -> `1`
// BIGO: O(log N): each call halves n.
func Power(x, n int) int {
	panic("todo")
}

// ==== END

// ==== EX 7.1.6 | Headcount of an org chart | Core
// PROMPT: Using `type Person struct{ Name string; Reports []Person }`, write `Headcount(p Person) int` counting that person and everyone below them.
// EXAMPLE: ana -> [bo, cy -> [di]] -> `4`
// BIGO: O(N) in the number of people: each person is visited once.
type Person struct {
	Name    string
	Reports []Person
}

func Headcount(p Person) int {
	panic("todo")
}

// ==== END

// ==== EX 7.1.7 | Binary digits, recursively | Core
// PROMPT: Write `ToBinary(n int) string` for `n >= 0` returning its base-2 digits: below 2 it is the digit itself; otherwise it is `ToBinary(n / 2)` followed by the last bit.
// EXAMPLE: `ToBinary(5)` -> `"101"`; `ToBinary(0)` -> `"0"`
// BIGO: O(log N) calls: each one halves n.
func ToBinary(n int) string {
	panic("todo")
}

// ==== END

// ==== EX 7.1.8 | Running totals in place | Core
// PROMPT: Write `RunningTotals(arr []int, i int)` (call it with `i = 1`) that replaces each element with the sum of itself and everything before it, recursively and in place.
// EXAMPLE: `[1 2 3 4]` -> `[1 3 6 10]`
// BIGO: O(N) time and call-stack depth.
func RunningTotals(arr []int, i int) {
	panic("todo")
}

// ==== END

// ==== EX 7.1.9 | Palindrome check by indexes | Core
// PROMPT: Write `IsPalindromeRec(s string, lo, hi int) bool` (call it with `0` and `len(s) - 1`): compare the two ends, then recurse on the part between them by moving the indexes inward.
// EXAMPLE: `"level"` -> `true`; `"levels"` -> `false`
// BIGO: O(N): N/2 calls, and no copies because the indexes move instead of the string being sliced.
func IsPalindromeRec(s string, lo, hi int) bool {
	panic("todo")
}

// ==== END

// ==== EX 7.1.10 | Sum of the digits | Core
// PROMPT: Write `DigitSum(n int) int` for `n >= 0` recursively: the last digit plus the digit sum of `n / 10`.
// EXAMPLE: `DigitSum(9045)` -> `18`; `DigitSum(0)` -> `0`
// BIGO: O(log N) calls: one per digit.
func DigitSum(n int) int {
	panic("todo")
}

// ==== END

// ==== EX 7.1.11 | Longest word, recursively | Core
// PROMPT: Write `LongestWord(words []string, i int) string` (call it with `i = 0`) returning the longest word from position `i` on (the earliest one on a tie, `""` for none): compare `words[i]` with the longest word of the rest.
// EXAMPLE: `[go java py rust]` -> `"java"`
// BIGO: O(N) calls, one per word.
func LongestWord(words []string, i int) string {
	panic("todo")
}

// ==== END

// ==== EX 7.1.12 | Collapse repeated neighbours | Core
// PROMPT: Write `CollapseRuns(nums []int, i int) []int` (call it with `i = 0`) returning `nums` with each run of equal neighbours reduced to one value, recursively: keep `nums[i]` unless the next value equals it.
// EXAMPLE: `[1 1 2 2 2 3 1]` -> `[1 2 3 1]`
// BIGO: O(N) calls. Building each result by joining one value onto a new list copies, so O(N^2) values are copied in total.
func CollapseRuns(nums []int, i int) []int {
	panic("todo")
}

// ==== END

// ==== EX 7.1.13 | Collatz steps | Core
// PROMPT: Write `CollatzSteps(n int) int` for `n >= 1`: 0 steps from 1; otherwise one step to `n / 2` (n even) or `3n + 1` (n odd), plus the steps from there.
// EXAMPLE: `CollatzSteps(6)` -> `8`; `CollatzSteps(1)` -> `0`
// BIGO: No known bound in terms of N: each call does O(1) work, so the cost is the number of steps itself.
func CollatzSteps(n int) int {
	panic("todo")
}

// ==== END

// ==== EX 7.1.14 | Greatest common divisor | Core
// PROMPT: Write `GCD(a, b int) int` for `a, b >= 0` (not both 0) recursively: `GCD(a, 0)` is `a`; otherwise it equals `GCD(b, a % b)`.
// EXAMPLE: `GCD(48, 18)` -> `6`; `GCD(7, 0)` -> `7`
// BIGO: O(log(min(a, b))) calls: the remainders shrink at least as fast as the Fibonacci numbers grow.
func GCD(a, b int) int {
	panic("todo")
}

// ==== END

// ==== EX 7.1.15 | Tiling a path with short and long tiles | Practice
// PROMPT: A path of length `n` is covered with tiles of length 1 and 3, in order. Write `Tilings(n int) int` counting the ways: the first tile is either short (leaving `n - 1`) or long (leaving `n - 3`).
// EXAMPLE: `Tilings(4)` -> `3`; `Tilings(10)` -> `28`
// BIGO: About O(1.47^N) as written: the two calls overlap. Chapter 7.2 (7.2.3) brings it to O(N).
func Tilings(n int) int {
	panic("todo")
}

// ==== END

// ==== EX 7.1.16 | Every balanced bracket string | Practice
// PROMPT: Write `ParenStrings(n int) []string` returning every string of `n` `'('` and `n` `')'` that is balanced. Build them one character at a time: add `'('` while fewer than `n` are open, and `')'` while it closes something.
// EXAMPLE: `ParenStrings(2)` -> `[(()) ()()]`
// BIGO: Exponential: the number of results follows the Catalan numbers (1, 2, 5, 14, 42, ...), roughly 4^N.
func ParenStrings(n int) []string {
	panic("todo")
}

// ==== END

// ==== EX 7.1.17 | Ways to roll a total with dice | Practice
// PROMPT: Write `DiceWays(dice, total int) int` counting the ordered ways `dice` six-sided dice can add up to `total`: zero dice make 0 one way; otherwise try each face 1 to 6 for the first die and count the ways for the rest.
// EXAMPLE: `DiceWays(2, 7)` -> `6`; `DiceWays(1, 7)` -> `0`
// BIGO: O(6^D) as written for D dice. Memoized in 72.4.
func DiceWays(dice, total int) int {
	panic("todo")
}

// ==== END

// ==== EX 7.1.18 | All subsets | Stretch
// PROMPT: Write `Subsets(items []int) [][]int` returning every subset (the power set). Each subset of the rest appears once without and once with the first item.
// EXAMPLE: `[1 2]` -> `[[] [2] [1] [1 2]]` (any order)
// BIGO: O(2^N) subsets.
// NOTE: Build each "with" subset in a fresh slice. `append(s, x)` on a shared subset could write into the same backing array twice.
func Subsets(items []int) [][]int {
	panic("todo")
}

// ==== END

// ==== EX 7.1.19 | Last one standing in a circle | Practice
// PROMPT: `n` people sit in a circle in seats `1` to `n`. Counting from seat 1, every `k`-th person still seated leaves, and counting restarts from the next one. Write `LastStanding(n, k int) int` that returns the seat of the last person left. Recurse on `n - 1`: after one person leaves, the survivor's position in the smaller circle shifts by `k`.
// EXAMPLE: `5, 2` -> `3` (seats 2, 4, 1, 5 leave in that order)
// BIGO: O(N) time and O(N) stack: each call removes one person and renumbers the circle so the next count starts at 0.
func LastStanding(n, k int) int {
	panic("todo")
}

// ==== END
