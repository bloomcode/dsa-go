// Package ch13: Partitioning: quicksort and quickselect.
package ch13


// ==== EX 13.1 | Partition | Core
// PROMPT: Write `Partition(arr []int, left, right int) int` using the rightmost element as pivot: a left pointer moves right past values smaller than the pivot, a right pointer moves left past values larger, swap, repeat; finally swap the pivot into the left pointer's spot. Return the pivot's final index.
// EXAMPLE: `[0 5 2 1 6 3]` -> pivot 3 ends at index 3, values left of it are all < 3, right of it all > 3
// BIGO: O(N) for one partition.
func Partition(arr []int, left, right int) int {
	panic("todo")
}

// ==== END

// ==== EX 13.2 | Quicksort | Core
// PROMPT: Write `Quicksort(arr []int)` that sorts in place: partition, then recurse on the parts left and right of the pivot. Base case: fewer than two elements. Recursing on sub-slices like `arr[:p]` and `arr[p+1:]` needs no index bounds.
// EXAMPLE: `[0 5 2 1 6 3]` -> `[0 1 2 3 5 6]`
// BIGO: O(N log N) average, O(N^2) worst case (pivot is always the smallest or largest, e.g. sorted input with a rightmost pivot).
func Quicksort(arr []int) {
	panic("todo")
}

// ==== END

// ==== EX 13.3 | Quickselect | Core
// PROMPT: Write `Quickselect(arr []int, k int) int` returning the k-th smallest value (k = 0 is the minimum) without fully sorting: partition a copy, then continue into only the side that holds index k.
// EXAMPLE: `[0 50 20 10 60 30]`, k = 1 -> `10`
// BIGO: O(N) average (N + N/2 + N/4 ... = 2N), O(N^2) worst case.
func Quickselect(arr []int, k int) int {
	panic("todo")
}

// ==== END

// ==== EX 13.4 | Duplicates via sorting | Core
// PROMPT: Write `HasDuplicatesSorted(arr []int) bool`: sort a copy with your quicksort, then compare neighbours.
// EXAMPLE: `[5 9 3 2 4 5 6]` -> `true`
// BIGO: O(N log N) time. It sits between the O(N^2) nested loops of 4.3 and the O(N)-space map of 4.4.
func HasDuplicatesSorted(arr []int) bool {
	panic("todo")
}

// ==== END

// ==== EX 13.5 | Greatest product of three | Practice
// PROMPT: Write `MaxProductOfThree(nums []int) int`. Sort a copy first; remember two large negatives can make a large positive.
// EXAMPLE: `[1 10 2 6 5 3]` -> `300`; `[-10 -10 1 3 2]` -> `300`
// BIGO: O(N log N).
func MaxProductOfThree(nums []int) int {
	panic("todo")
}

// ==== END

// ==== EX 13.6 | Missing number via sorting | Core
// PROMPT: A slice holds every integer 0..n except one. Write `MissingNumberSorted(arr []int) int` by sorting a copy and finding where index and value disagree.
// EXAMPLE: `[5 2 4 1 0]` -> `3`; `[0 1 2]` -> `3`
// BIGO: O(N log N). Chapter 20 (20.8) gets it to O(N).
func MissingNumberSorted(arr []int) int {
	panic("todo")
}

// ==== END

// ==== EX 13.7 | Median with quickselect | Practice
// PROMPT: Write `Median(nums []int) float64` using `Quickselect`. For an even count, average the two middle values.
// EXAMPLE: `[7 1 3]` -> `3`; `[4 1 3 2]` -> `2.5`
// BIGO: O(N) average.
func Median(nums []int) float64 {
	panic("todo")
}

// ==== END

// ==== EX 13.8 | Merge sort for comparison | Stretch
// PROMPT: Write `MergeSort(arr []int) []int` returning a new sorted slice: split in half, sort each half recursively, merge (reuse 6.5's idea).
// EXAMPLE: `[38 27 43 3 9 82 10]` -> `[3 9 10 27 38 43 82]`
// BIGO: O(N log N) in every case, but O(N) extra space, unlike quicksort's in-place partitioning.
func MergeSort(arr []int) []int {
	panic("todo")
}

// ==== END
