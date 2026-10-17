// Package ch12: Dynamic programming: memoize or build up.
package ch12


// ==== EX 12.1 | Fibonacci three ways | Core
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

// ==== EX 12.2 | Fix the double recursive call | Core
// PROMPT: This version calls itself twice per level: `if arr[0] > maxBad(arr[1:]) { return arr[0] }; return maxBad(arr[1:])`. Write `MaxRecursive(arr []int) int` that calls itself once per level.
// EXAMPLE: `[1 2 3 4]` -> `4`
// BIGO: From O(2^N) to O(N), just by storing the recursive result in a variable.
func MaxRecursive(arr []int) int {
	panic("todo")
}

// ==== END

// ==== EX 12.3 | Staircase with memoization | Core
// PROMPT: Rewrite 11.8 as `StaircaseMemo(n int) int` with a `[]int` memo filled with -1 and a recursive closure.
// EXAMPLE: `StaircaseMemo(5)` -> `13`; `StaircaseMemo(50)` -> `10562230626642` at once
// BIGO: O(N).
func StaircaseMemo(n int) int {
	panic("todo")
}

// ==== END

// ==== EX 12.4 | Unique paths with memoization | Core
// PROMPT: Rewrite 11.10 as `UniquePathsMemo(rows, cols int, memo map[[2]int]int) int`, keyed by the array `[2]int{rows, cols}`.
// EXAMPLE: `UniquePathsMemo(3, 7, ...)` -> `28`; `(18, 18)` -> `2333606220`
// BIGO: O(rows * cols): each sub-grid is solved once.
func UniquePathsMemo(rows, cols int, memo map[[2]int]int) int {
	panic("todo")
}

// ==== END

// ==== EX 12.5 | Count the calls | Practice
// PROMPT: Write `CountFibCalls(n int) (naive, memo int)`: how many times each Fibonacci version is invoked. Count with closures that increment captured counters.
// EXAMPLE: `CountFibCalls(10)` -> `177, 19`
// BIGO: Shows overlapping subproblems directly: naive calls grow about 1.6x per step of n, memo calls by 2.
func CountFibCalls(n int) (naive, memo int) {
	panic("todo")
}

// ==== END

// ==== EX 12.6 | Fewest coins, bottom-up | Stretch
// PROMPT: Write `MinCoins(coins []int, amount int) int` filling a table `best[0..amount]` from the bottom up. Return -1 if the amount cannot be made.
// EXAMPLE: `[1 5 10 25]`, 63 -> `6`; `[2]`, 3 -> `-1`
// BIGO: O(amount * number of coins) time, O(amount) space.
// NOTE: Greedy (largest coin first) fails for coins like `[1 3 4]` and amount 6: greedy gives 3 coins, the answer is 2. Use `math.MaxInt` as "impossible" only if you test for it before adding 1.
func MinCoins(coins []int, amount int) int {
	panic("todo")
}

// ==== END
