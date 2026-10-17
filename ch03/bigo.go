// Package ch03: Big O: how the step count grows.
package ch03

// ==== EX 3.1 | An O(1) algorithm | Core
// PROMPT: Write `IsLeapYear(year int) bool`. Explain why its step count does not depend on the input's size.
// EXAMPLE: `IsLeapYear(2024)` -> `true`; `1900` -> `false`; `2000` -> `true`
// BIGO: O(1). At most three checks, whatever the year.
func IsLeapYear(year int) bool {
	panic("todo")
}

// ==== END

// ==== EX 3.2 | An O(N) algorithm | Core
// PROMPT: Write `ArraySum(arr []int) int` with a loop, and say how many steps it takes for N elements.
// EXAMPLE: `[1 2 3 4]` -> `10`, 4 loop steps
// BIGO: O(N). One step per element.
func ArraySum(arr []int) int {
	panic("todo")
}

// ==== END

// ==== EX 3.3 | An O(log N) algorithm | Core
// PROMPT: Write `HalvingSteps(n int) int` returning how many times you can halve n (integer division) before it reaches 1.
// EXAMPLE: `HalvingSteps(8)` -> `3`; `HalvingSteps(1024)` -> `10`; `HalvingSteps(1)` -> `0`
// BIGO: O(log N). This is the definition of log2 N for powers of two; `bits.Len(uint(n)) - 1` computes it in one instruction.
func HalvingSteps(n int) int {
	panic("todo")
}

// ==== END

// ==== EX 3.4 | Doubling: grains on a chessboard | Practice
// PROMPT: One grain on square 1, two on square 2, four on square 3, and so on. Write `SquaresNeeded(grains int) int` returning which square first holds at least `grains`.
// EXAMPLE: `SquaresNeeded(1)` -> `1`; `SquaresNeeded(16)` -> `5`; `SquaresNeeded(17)` -> `6`
// BIGO: O(log N). The loop doubles a counter until it reaches N, the mirror image of halving.
// NOTE: Square 64 holds 2^63 grains, one more than `math.MaxInt64`, so an int counter overflows on the last square; use `uint64` there.
func SquaresNeeded(grains int) int {
	panic("todo")
}

// ==== END

// ==== EX 3.5 | Count nested-loop steps | Practice
// PROMPT: Write `AllPairsSteps(n int) int` (inner loop runs N times per outer step) and `UniquePairsSteps(n int) int` (inner loop starts at `i + 1`). Both return the number of inner-loop iterations. Classify both.
// EXAMPLE: `AllPairsSteps(4)` -> `16`; `UniquePairsSteps(4)` -> `6`
// BIGO: Both O(N^2). Unique pairs is N(N-1)/2, about half, and Big O ignores the 1/2.
func AllPairsSteps(n int) int {
	panic("todo")
}

func UniquePairsSteps(n int) int {
	panic("todo")
}

// ==== END
