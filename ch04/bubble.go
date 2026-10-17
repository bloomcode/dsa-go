// Package ch04: Bubble sort and quadratic code.
package ch04

// ==== EX 4.1 | Bubble sort | Core
// PROMPT: Write `BubbleSort(arr []int)` that sorts in place by repeatedly comparing neighbours and swapping them when out of order.
// EXAMPLE: `[4 2 7 1 3]` -> `[1 2 3 4 7]`
// BIGO: O(N^2) time, O(1) space. About N^2/2 comparisons plus up to N^2/2 swaps.
func BubbleSort(arr []int) {
	panic("todo")
}

// ==== END

// ==== EX 4.2 | Bubble sort with early exit | Practice
// PROMPT: Write `BubbleSortOptimized(arr []int) int` that sorts in place, shrinks the unsorted range each pass, stops as soon as a pass makes no swaps, and returns the number of passes.
// EXAMPLE: `[1 2 3 4]` -> `1` pass; `[4 3 2 1]` -> `3` passes
// BIGO: Worst case still O(N^2); best case (already sorted) O(N).
// NOTE: The largest unsorted value "bubbles" to its final spot each pass, which is why the range can shrink.
func BubbleSortOptimized(arr []int) int {
	panic("todo")
}

// ==== END

// ==== EX 4.3 | Duplicate check, quadratic | Core
// PROMPT: Write `HasDuplicateQuadratic(arr []int) bool` with two nested loops comparing every pair.
// EXAMPLE: `[1 5 3 9 1 4]` -> `true`; `[1 2 3]` -> `false`
// BIGO: O(N^2) time, O(1) space.
func HasDuplicateQuadratic(arr []int) bool {
	panic("todo")
}

// ==== END

// ==== EX 4.4 | Duplicate check, linear | Core
// PROMPT: Write `HasDuplicateLinear(arr []int) bool` using one loop and a `map[int]bool` of values already seen.
// EXAMPLE: Same results as 4.3
// BIGO: O(N) time, O(N) space. Trading memory for speed.
func HasDuplicateLinear(arr []int) bool {
	panic("todo")
}

// ==== END

// ==== EX 4.5 | Greatest number, two ways | Practice
// PROMPT: Write `GreatestQuadratic(arr []int) (int, bool)` (a value is greatest if no other value beats it, checked with nested loops) and `GreatestLinear(arr []int) (int, bool)` (one pass tracking the best so far). Return false for an empty slice.
// EXAMPLE: `[3 9 2 7]` -> `9, true`
// BIGO: O(N^2) versus O(N), for the same answer.
func GreatestQuadratic(arr []int) (int, bool) {
	panic("todo")
}

func GreatestLinear(arr []int) (int, bool) {
	panic("todo")
}

// ==== END

// ==== EX 4.6 | Is it sorted? | Core
// PROMPT: Write `IsSorted(arr []int) bool` that checks ascending order in one pass.
// EXAMPLE: `[1 2 2 5]` -> `true`; `[1 3 2]` -> `false`
// BIGO: O(N) worst case, can stop at the first out-of-order pair. The library version is `slices.IsSorted`.
func IsSorted(arr []int) bool {
	panic("todo")
}

// ==== END
