// Package ch3_1: Simple sorts: bubble, selection and insertion.
package ch3_1


// ==== EX 3.1.1 | Bubble sort | Core
// PROMPT: Write `BubbleSort(arr []int)` that sorts in place by repeatedly comparing neighbours and swapping them when out of order.
// EXAMPLE: `[4 2 7 1 3]` -> `[1 2 3 4 7]`
// BIGO: O(N^2) time, O(1) space. About N^2/2 comparisons plus up to N^2/2 swaps.
func BubbleSort(arr []int) {
	panic("todo")
}

// ==== END

// ==== EX 3.1.2 | Bubble sort with early exit | Practice
// PROMPT: Write `BubbleSortOptimized(arr []int) int` that sorts in place, shrinks the unsorted range each pass, stops as soon as a pass makes no swaps, and returns the number of passes.
// EXAMPLE: `[1 2 3 4]` -> `1` pass; `[4 3 2 1]` -> `3` passes
// BIGO: Worst case still O(N^2); best case (already sorted) O(N).
// NOTE: The largest unsorted value "bubbles" to its final spot each pass, which is why the range can shrink.
func BubbleSortOptimized(arr []int) int {
	panic("todo")
}

// ==== END

// ==== EX 3.1.3 | Count equal pairs, quadratic | Core
// PROMPT: Write `EqualPairsQuadratic(nums []int) int` counting the pairs of positions `i < j` with `nums[i] == nums[j]`, using two nested loops.
// EXAMPLE: `[1 2 1 1 2]` -> `4`
// BIGO: O(N^2) time, O(1) space.
func EqualPairsQuadratic(nums []int) int {
	panic("todo")
}

// ==== END

// ==== EX 3.1.4 | Count equal pairs, linear | Core
// PROMPT: Write `EqualPairsLinear(nums []int) int` with one loop and a map counting how many times each value has appeared so far: each new value pairs with every earlier copy.
// EXAMPLE: Same results as 3.1.3: `[1 2 1 1 2]` -> `4`
// BIGO: O(N) time, O(N) space. Trading memory for speed.
func EqualPairsLinear(nums []int) int {
	panic("todo")
}

// ==== END

// ==== EX 3.1.5 | Largest pair sum, two ways | Practice
// PROMPT: For at least two numbers, write `LargestPairSumQuadratic(nums []int) int` (try every pair of positions) and `LargestPairSumLinear(nums []int) int` (one pass keeping the two largest values seen).
// EXAMPLE: `[3 9 1 7]` -> `16`
// BIGO: O(N^2) versus O(N), for the same answer.
func LargestPairSumQuadratic(nums []int) int {
	panic("todo")
}

func LargestPairSumLinear(nums []int) int {
	panic("todo")
}

// ==== END

// ==== EX 3.1.6 | Is it sorted? | Core
// PROMPT: Write `IsSorted(arr []int) bool` that checks ascending order in one pass.
// EXAMPLE: `[1 2 2 5]` -> `true`; `[1 3 2]` -> `false`
// BIGO: O(N) worst case, can stop at the first out-of-order pair. The library version is `slices.IsSorted`.
func IsSorted(arr []int) bool {
	panic("todo")
}

// ==== END

// ==== EX 3.1.7 | Selection sort | Core
// PROMPT: Write `SelectionSort(arr []int)`: on each pass find the index of the smallest unsorted value and swap it into place (one swap per pass at most).
// EXAMPLE: `[4 2 7 1 3]` -> `[1 2 3 4 7]`
// BIGO: O(N^2) time, O(1) space. About N^2/2 comparisons but at most N - 1 swaps.
func SelectionSort(arr []int) {
	panic("todo")
}

// ==== END

// ==== EX 3.1.8 | Bubble vs selection: count the work | Practice
// PROMPT: Write `BubbleCounts(arr []int) (comparisons, swaps int)` and `SelectionCounts(arr []int) (comparisons, swaps int)`, each sorting a copy. Compare them on a reversed slice of 10.
// EXAMPLE: reversed 10 elements -> bubble `45, 45`, selection `45, 5`
// BIGO: Both O(N^2). Selection does roughly half the total steps; Big O ignores that constant, which is why both share a category.
func BubbleCounts(arr []int) (comparisons, swaps int) {
	panic("todo")
}

func SelectionCounts(arr []int) (comparisons, swaps int) {
	panic("todo")
}

// ==== END

// ==== EX 3.1.9 | Differences between neighbours | Core
// PROMPT: Write `NeighbourGaps(nums []int) []int` returning `nums[i+1] - nums[i]` for each pair of neighbours, in order.
// EXAMPLE: `[3 7 4 10]` -> `[4 -3 6]`
// BIGO: N - 1 steps, which is still O(N): Big O drops constants and the minus one.
func NeighbourGaps(nums []int) []int {
	panic("todo")
}

// ==== END

// ==== EX 3.1.10 | Second largest in one pass | Practice
// PROMPT: Write `SecondLargest(arr []int) (int, bool)` returning the second-largest distinct value in a single pass and without sorting; false if there is none.
// EXAMPLE: `[5 1 9 7 9]` -> `7, true`; `[3 3]` -> `0, false`
// BIGO: O(N) time, O(1) space. Sorting first would be O(N log N).
func SecondLargest(arr []int) (int, bool) {
	panic("todo")
}

// ==== END

// ==== EX 3.1.11 | K smallest by partial selection sort | Stretch
// PROMPT: Write `KSmallest(arr []int, k int) []int` that runs only the first k passes of selection sort on a copy and returns the first k values.
// EXAMPLE: `[8 3 5 1 9]`, k = 2 -> `[1 3]`
// BIGO: O(kN). Better than a full O(N^2) sort when k is small.
func KSmallest(arr []int, k int) []int {
	panic("todo")
}

// ==== END

// ==== EX 3.1.12 | Insertion sort | Core
// PROMPT: Write `InsertionSort(arr []int)`: for each index from 1, lift the value out ("temp"), shift larger values on its left one slot right, and drop the value into the gap.
// EXAMPLE: `[4 2 7 1 3]` -> `[1 2 3 4 7]`
// BIGO: Worst O(N^2), average about N^2/2, best (already sorted) O(N). O(1) space.
func InsertionSort(arr []int) {
	panic("todo")
}

// ==== END

// ==== EX 3.1.13 | Best, average, worst case in numbers | Practice
// PROMPT: Write `InsertionSortSteps(arr []int) int` that sorts a copy and returns the count of comparisons plus shifts. Run it on a sorted, a random, and a reversed slice of 100.
// EXAMPLE: sorted 100 -> `99`; reversed 100 -> `9900`; random lands in between
// BIGO: Shows O(N) best case and O(N^2) worst case for the same algorithm.
// NOTE: When data is usually nearly sorted, insertion sort can beat selection sort, despite sharing O(N^2).
func InsertionSortSteps(arr []int) int {
	panic("todo")
}

// ==== END

// ==== EX 3.1.14 | Does the list hold every wanted value? | Core
// PROMPT: Write `ContainsAll(values, wanted []int) bool` with a linear search for each wanted value that `break`s on a match, returning `false` as soon as one is missing.
// EXAMPLE: `[4 8 15 16]`, `[8 16]` -> `true`; `[4 8]`, `[8 5]` -> `false`
// BIGO: O(N * M) worst case. Stopping early only helps the best and average case: a missing first value ends the work after one scan.
func ContainsAll(values, wanted []int) bool {
	panic("todo")
}

// ==== END

// ==== EX 3.1.15 | Insert one value into a sorted array | Core
// PROMPT: Write `InsertSorted(arr []int, value int) []int` that performs a single insertion-sort step: append, then shift left until the value is in place. Return the new slice.
// EXAMPLE: `[1 4 9]`, value `5` -> `[1 4 5 9]`
// BIGO: O(N) worst case, O(1) if the value belongs at the end.
// NOTE: `append` may return a slice backed by new memory, so the caller must use the return value: `s = InsertSorted(s, v)`.
func InsertSorted(arr []int, value int) []int {
	panic("todo")
}

// ==== END

// ==== EX 3.1.16 | Sort a nearly sorted list | Practice
// PROMPT: Every value in `nums` is at most `k` positions away from where it belongs. Write `SortNearlySorted(nums []int, k int)` that sorts in place with insertion sort, and explain why it beats O(N^2) here.
// EXAMPLE: `[2 1 3 5 4], 1` -> `[1 2 3 4 5]`
// BIGO: O(N * k): each value shifts at most k places, so for small k insertion sort is close to linear.
func SortNearlySorted(nums []int, k int) {
	panic("todo")
}

// ==== END
