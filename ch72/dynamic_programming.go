// Package ch72: Dynamic programming.
package ch72


// ==== EX 72.1 | Fibonacci three ways | Core
// PROMPT: Write `FibNaive(n int) int`, `FibMemo(n int, memo map[int]int) int` and `FibBottomUp(n int) int`, with fib(0) = 0 and fib(1) = 1.
// EXAMPLE: All three return `55` for n = 10; only the last two finish quickly for n = 80
// BIGO: Naive O(2^N); memoised O(N) time and space; bottom-up O(N) time and O(1) space.
// NOTE: fib(92) is the largest that fits in an int64. Maps are reference types, so the callee's writes to `memo` are visible to the caller.
func FibNaive(n int) int {
	panic("todo")
}

func FibMemo(n int, memo map[int]int) int {
	panic("todo")
}

func FibBottomUp(n int) int {
	panic("todo")
}

// ==== END

// ==== EX 72.2 | One recursive call per level | Core
// PROMPT: A version that finds the rest's minimum and its maximum with two separate recursive calls doubles the work at every level. Write `MinMax(nums []int, i int) (int, int)` (call it with `i = 0`) returning the smallest and largest of `nums[i:]` for a non-empty slice, with a single recursive call per level.
// EXAMPLE: `[4 -2 9 0]` -> `-2, 9`
// BIGO: From O(2^N) to O(N), just by keeping the one recursive result and reading both parts from it.
func MinMax(nums []int, i int) (int, int) {
	panic("todo")
}

// ==== END

// ==== EX 72.3 | Tilings with memoization | Core
// PROMPT: Rewrite 71.15 as `TilingsMemo(n int, memo map[int]int) int`.
// EXAMPLE: `TilingsMemo(10, map[int]int{})` -> `28`; `TilingsMemo(60, ...)` -> `5582216355` at once
// BIGO: O(N): each length is solved once.
func TilingsMemo(n int, memo map[int]int) int {
	panic("todo")
}

// ==== END

// ==== EX 72.4 | Dice totals with memoization | Core
// PROMPT: Rewrite 71.17 as `DiceWaysMemo(dice, total int, memo map[[2]int]int) int`, keyed by `[2]int{dice, total}`.
// EXAMPLE: `DiceWaysMemo(2, 7, map[[2]int]int{})` -> `6`; `(20, 70)` -> `189456975899496`
// BIGO: O(D * T * 6) for D dice and total T: each (dice, total) pair is solved once, trying six faces.
func DiceWaysMemo(dice, total int, memo map[[2]int]int) int {
	panic("todo")
}

// ==== END

// ==== EX 72.5 | Count the calls | Practice
// PROMPT: Write `CountFibCalls(n int) (naive, memo int)`: how many times each Fibonacci version is invoked. Count with closures that increment captured counters.
// EXAMPLE: `CountFibCalls(10)` -> `177, 19`
// BIGO: Shows overlapping subproblems directly: naive calls grow about 1.6x per step of n, memo calls by 2.
func CountFibCalls(n int) (naive, memo int) {
	panic("todo")
}

// ==== END

// ==== EX 72.6 | Fewest coins, bottom-up | Stretch
// PROMPT: Write `MinCoins(coins []int, amount int) int` filling a table `best[0..amount]` from the bottom up. Return -1 if the amount cannot be made.
// EXAMPLE: `[1 5 10 25]`, 63 -> `6`; `[2]`, 3 -> `-1`
// BIGO: O(amount * number of coins) time, O(amount) space.
// NOTE: Greedy (largest coin first) fails for coins like `[1 3 4]` and amount 6: greedy gives 3 coins, the answer is 2. Use `math.MaxInt` as "impossible" only if you test for it before adding 1.
func MinCoins(coins []int, amount int) int {
	panic("todo")
}

// ==== END
