// Package ch2_6: Two pointers and intervals.
package ch2_6


// ==== EX 2.6.1 | Pair sum in a sorted array | Core
// PROMPT: `nums` is sorted ascending. Write `PairSumSorted(nums []int, target int) []int` that walks a pointer in from each end and returns the indices `[i j]`, `i < j`, of the first pair it finds summing to `target`, or `nil` when there is none.
// EXAMPLE: `[1 3 4 6 9], 10` -> `[0 4]`; `[1 3 4 6 9], 2` -> `nil`
// BIGO: O(N) time, O(1) space: a sum that is too small can only grow by moving the low pointer, one that is too big only shrinks by moving the high pointer.
func PairSumSorted(nums []int, target int) []int {
	panic("todo")
}

// ==== END

// ==== EX 2.6.2 | Squares of a sorted array | Core
// PROMPT: `nums` is sorted ascending and may hold negatives. Write `SortedSquares(nums []int) []int` that returns the squares in ascending order without sorting: compare the two ends and fill a new slice from the back.
// EXAMPLE: `[-4 -1 0 3]` -> `[0 1 9 16]`
// BIGO: O(N) time and space: the largest square sits at one of the two ends, so filling the output from the back needs no sort.
func SortedSquares(nums []int) []int {
	panic("todo")
}

// ==== END

// ==== EX 2.6.3 | Smallest difference across two arrays | Core
// PROMPT: Write `SmallestGap(a, b []int) int` that returns the smallest `|x - y|` with `x` from `a` and `y` from `b` (both non-empty). Sort copies of both and advance whichever index holds the smaller value.
// EXAMPLE: `[5 19 31], [12 29 1]` -> `2`
// BIGO: O(N log N + M log M) time for the sorts, then O(N + M) for the walk: moving the smaller value forward is the only move that can shrink the gap.
func SmallestGap(a, b []int) int {
	panic("todo")
}

// ==== END

// ==== EX 2.6.4 | Move a value to the end | Core
// PROMPT: Write `MoveToEnd(nums []int, value int)` that moves every copy of `value` to the end of `nums` in place, keeping the other values in their order. Use a read index and a write index.
// EXAMPLE: `[2 1 2 3 2 4], 2` -> `[1 3 4 2 2 2]`
// BIGO: O(N) time, O(1) space: the write pointer never passes the read pointer, so nothing unread is overwritten.
func MoveToEnd(nums []int, value int) {
	panic("todo")
}

// ==== END

// ==== EX 2.6.5 | Monotonic check | Core
// PROMPT: Write `IsMonotonic(nums []int) bool`: `true` when `nums` never goes down or never goes up (equal neighbours are fine; empty and single-value slices count).
// EXAMPLE: `[1 2 2 5]` -> `true`; `[3 1 2]` -> `false`
// BIGO: O(N) time, O(1) space: one pass that rules out each direction as soon as a step goes the other way.
func IsMonotonic(nums []int) bool {
	panic("todo")
}

// ==== END

// ==== EX 2.6.6 | Merge overlapping intervals | Core
// PROMPT: Each interval is `[start, end]`, in any order. Write `MergeIntervals(intervals [][2]int) [][2]int` that returns the merged intervals sorted by start (an empty, non-nil slice for no input); intervals that touch (`end == start`) merge too. Do not change the input.
// EXAMPLE: `[[6 8] [1 3] [2 4] [8 9]]` -> `[[1 4] [6 9]]`
// BIGO: O(N log N) time for the sort, O(N) for the output: once sorted by start, an interval can only overlap the last merged one.
func MergeIntervals(intervals [][2]int) [][2]int {
	panic("todo")
}

// ==== END

// ==== EX 2.6.7 | Triples to a target | Practice
// PROMPT: Write `TriplesToTarget(nums []int, target int) [][]int` that returns every distinct triple of values from `nums` (three different positions) summing to `target`, as a non-nil slice. Each triple is ascending and the slice is in ascending order. Sort a copy, fix the first value, and two-pointer the rest, skipping repeated values.
// EXAMPLE: `[4 -1 2 0 -3 1], 1` -> `[[-3 0 4] [-1 0 2]]`
// BIGO: O(N^2) time, O(1) extra space beyond the output: fix the smallest value, then a pair sum in the sorted rest is one two-pointer walk.
func TriplesToTarget(nums []int, target int) [][]int {
	panic("todo")
}

// ==== END

// ==== EX 2.6.8 | Most water between two walls | Practice
// PROMPT: `heights[i]` is a wall at position `i`. Write `MostWater(heights []int) int` that returns the largest `(j - i) * min(heights[i], heights[j])` over pairs `i < j` (`0` with fewer than two walls). Start with the widest pair and move the shorter wall inward.
// EXAMPLE: `[2 7 3 6 1 5]` -> `20`
// BIGO: O(N) time, O(1) space: the shorter wall limits every pair it is part of with a closer partner, so it can be dropped.
func MostWater(heights []int) int {
	panic("todo")
}

// ==== END

// ==== EX 2.6.9 | Overlaps between two interval lists | Practice
// PROMPT: `a` and `b` hold `[start, end]` intervals, each sorted and non-overlapping within itself. Write `IntervalOverlaps(a, b [][2]int) [][2]int` that returns every overlap `[max start, min end]` (a single shared point counts), in order, as a non-nil slice.
// EXAMPLE: `[[1 4] [6 9]], [[3 7] [9 12]]` -> `[[3 4] [6 7] [9 9]]`
// BIGO: O(N + M) time: the interval that ends first cannot overlap anything later in the other list, so its pointer moves on.
func IntervalOverlaps(a, b [][2]int) [][2]int {
	panic("todo")
}

// ==== END

// ==== EX 2.6.10 | Shortest stretch to sort | Practice
// PROMPT: Write `UnsortedStretch(nums []int) [2]int` that returns `[start end]`, the shortest index range which, once sorted, leaves all of `nums` sorted, or `[-1 -1]` when it already is. Sweep left to right tracking the maximum, then right to left tracking the minimum.
// EXAMPLE: `[1 3 7 5 6 9]` -> `[2 4]`; `[1 2 3]` -> `[-1 -1]`
// BIGO: O(N) time, O(1) space: a value below the running maximum must move, so the last such index is the right edge; the mirror pass finds the left edge.
func UnsortedStretch(nums []int) [2]int {
	panic("todo")
}

// ==== END

// ==== EX 2.6.11 | Negatives first, order kept | Practice
// PROMPT: Write `NegativesFirst(nums []int)` that moves the negative values before the others (zero counts as non-negative) in place, keeping the order within each group. Split the range in half, solve each half, then swap the middle blocks (non-negatives of the left, negatives of the right) with three reversals.
// EXAMPLE: `[3 -1 4 -5 9 -2]` -> `[-1 -5 -2 3 4 9]`
// BIGO: O(N log N) time, O(log N) stack: each level of the recursion moves every value at most a constant number of times with three reversals.
// NOTE: The one-pass swap partition is O(N) but not stable; keeping the order in place costs the log factor.
func NegativesFirst(nums []int) {
	panic("todo")
}

// ==== END

// ==== EX 2.6.12 | Next arrangement in order | Stretch
// PROMPT: Write `NextArrangement(nums []int)` that rearranges `nums` in place into the next larger ordering in dictionary order, or into ascending order when it is already the largest. Find the last position whose value is below its right neighbour, swap it with the smallest larger value to its right, then reverse the tail.
// EXAMPLE: `[1 3 2]` -> `[2 1 3]`; `[3 2 1]` -> `[1 2 3]`
// BIGO: O(N) time, O(1) space: the suffix after the pivot is descending, so one swap and one reversal give the next arrangement.
func NextArrangement(nums []int) {
	panic("todo")
}

// ==== END

// ==== EX 2.6.13 | Four values to a target | Stretch
// PROMPT: Write `QuadsToTarget(nums []int, target int) [][]int` that returns every distinct group of four values from `nums` (four different positions) summing to `target`, as a non-nil slice. Each group is ascending and the slice is in ascending order. Sort a copy, fix two values, and two-pointer the rest, skipping repeats at every level.
// EXAMPLE: `[2 -1 0 3 -2 1], 2` -> `[[-2 -1 2 3] [-2 0 1 3] [-1 0 1 2]]`
// BIGO: O(N^3) time, O(1) extra space beyond the output: two fixed values, then one two-pointer walk for the last pair.
func QuadsToTarget(nums []int, target int) [][]int {
	panic("todo")
}

// ==== END

// ==== EX 2.6.14 | Water trapped between bars | Stretch
// PROMPT: `heights[i]` is the height of a bar of width 1. Write `TrappedWater(heights []int) int` that returns how many units of water stay between the bars after rain. Walk two indices inward, always moving the lower side and adding its running maximum minus its height.
// EXAMPLE: `[3 0 2 0 4]` -> `7`
// BIGO: O(N) time, O(1) space: the side with the lower bar is bounded by its own running maximum, because the other side is known to be at least as tall.
func TrappedWater(heights []int) int {
	panic("todo")
}

// ==== END
