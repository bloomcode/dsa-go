// Package ch21: Mergesort and stable sorting.
package ch21


// ==== EX 21.1 | Merge two sorted arrays | Core
// PROMPT: Write `Merge(a, b []int) []int` returning one sorted slice from two sorted slices, using two indices that advance through them. No sorting functions.
// EXAMPLE: `Merge([]int{1, 4, 9}, []int{2, 3, 10, 11})` -> `[1 2 3 4 9 10 11]`
// BIGO: O(N + M) time, O(N + M) space.
// NOTE: Use `<=` when comparing so equal items from a come first. That one character is what makes mergesort stable (21.7).
func Merge(a, b []int) []int {
	panic("todo")
}

// ==== END

// ==== EX 21.2 | Count the comparisons in a merge | Core
// PROMPT: Write `MergeComparisons(a, b []int) int` returning how many comparisons a merge makes. Find inputs that need the fewest and the most comparisons.
// EXAMPLE: `MergeComparisons([1 2 3], [4 5 6])` -> `3`; `MergeComparisons([1 3 5], [2 4 6])` -> `5`
// BIGO: Between min(N, M) and N + M - 1 comparisons.
func MergeComparisons(a, b []int) int {
	panic("todo")
}

// ==== END

// ==== EX 21.3 | Mergesort, top-down | Core
// PROMPT: Write `Mergesort(arr []int) []int` returning a new sorted slice: split in half (re-slicing costs nothing), sort each half recursively, then `Merge` them. Base case: length 0 or 1.
// EXAMPLE: `Mergesort([]int{38, 27, 43, 3, 9, 82, 10})` -> `[3 9 10 27 38 43 82]`
// BIGO: O(N log N) in every case: log2 N levels of splitting, O(N) merging per level. O(N) extra space.
func Mergesort(arr []int) []int {
	panic("todo")
}

// ==== END

// ==== EX 21.4 | Mergesort with one shared buffer | Practice
// PROMPT: Write `MergesortInPlace(arr []int)` that sorts arr itself, recursing on index ranges and merging into one auxiliary slice allocated once, instead of allocating at every level.
// EXAMPLE: `a := []int{5, 2, 4, 6, 1, 3}; MergesortInPlace(a)` -> a is `[1 2 3 4 5 6]`
// BIGO: O(N log N) time, O(N) extra space in total.
func MergesortInPlace(arr []int) {
	panic("todo")
}

// ==== END

// ==== EX 21.5 | Mergesort, bottom-up | Core
// PROMPT: Write `MergesortBottomUp(arr []int) []int` with no recursion: merge runs of width 1 into width 2, then 4, 8, and so on, until one run covers the slice. Swap two buffers between passes.
// EXAMPLE: `MergesortBottomUp([]int{5, 1, 4, 2, 8, 0, 2})` -> `[0 1 2 2 4 5 8]`
// BIGO: O(N log N) time, O(N) space, and no call stack at all.
func MergesortBottomUp(arr []int) []int {
	panic("todo")
}

// ==== END

// ==== EX 21.6 | Count inversions while merging | Practice
// PROMPT: An inversion is a pair i < j with arr[i] > arr[j]. Write `CountInversions(arr []int) int` by adding, every time an item from the right half is taken first, the number of items still waiting in the left half.
// EXAMPLE: `CountInversions([]int{2, 4, 1, 3, 5})` -> `3`; a reversed slice of 5 -> `10`
// BIGO: O(N log N), versus O(N^2) for checking every pair.
func CountInversions(arr []int) int {
	panic("todo")
}

// ==== END

// ==== EX 21.7 | Stability, with records | Core
// PROMPT: Write `MergesortBy[T any](records []T, key func(T) int) []T` that sorts by key and keeps records with equal keys in their original order. Show that quicksort from Chapter 13 does not promise this.
// EXAMPLE: records `{ann 3} {bo 1} {cy 3} {di 1}` by the number -> `[{bo 1} {di 1} {ann 3} {cy 3}]`
// BIGO: O(N log N).
// NOTE: The standard library's `slices.SortStableFunc` gives the same guarantee; plain `slices.SortFunc` does not.
func MergesortBy[T any](records []T, key func(T) int) []T {
	panic("todo")
}

// ==== END

// ==== EX 21.8 | Hybrid: insertion sort for small pieces | Stretch
// PROMPT: Write `HybridSort(arr []int, cutoff int) []int`: mergesort that switches to insertion sort when a piece has cutoff items or fewer. Benchmark it against plain mergesort in Chapter 22.
// EXAMPLE: `HybridSort([]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, 4)` -> `[0 1 2 3 4 5 6 7 8 9]`
// BIGO: Still O(N log N); insertion sort wins on tiny inputs because its constant factor is small. Go's own `slices.Sort` does the same for short runs.
func HybridSort(arr []int, cutoff int) []int {
	panic("todo")
}

// ==== END
