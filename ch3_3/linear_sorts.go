// Package ch3_3: Linear-time sorts.
package ch3_3


// ==== EX 3.3.1 | Counting sort | Core
// PROMPT: Every value in `nums` is an integer from `0` to `maxValue`. Write `CountingSort(nums []int, maxValue int) []int` that returns a new sorted slice by counting how often each value occurs, with no comparisons between values.
// EXAMPLE: `[3 0 2 3 1], 3` -> `[0 1 2 3 3]`
// BIGO: O(N + K) time and O(K) space for keys 0..K: one pass to count, one pass over the counts to write the output.
func CountingSort(nums []int, maxValue int) []int {
	panic("todo")
}

// ==== END

// ==== EX 3.3.2 | Sort to a given key order | Core
// PROMPT: `order` lists a few distinct keys, and every value in `values` is one of them. Write `SortByOrder(values, order []int) []int` that returns the values grouped in the order of `order`, using one counting pass and a map of counts.
// EXAMPLE: `[2 0 1 2 0], [2 0 1]` -> `[2 2 0 0 1]`
// BIGO: O(N + K) time, O(K) space for K keys: count each key once, then write the keys out in the order asked for.
func SortByOrder(values, order []int) []int {
	panic("todo")
}

// ==== END

// ==== EX 3.3.3 | Radix sort | Practice
// PROMPT: Write `RadixSort(nums []int) []int` that sorts non-negative integers and returns a new slice, digit by digit from the lowest: each pass distributes the values into ten buckets by the current digit, keeping their order, and reads the buckets back in order.
// EXAMPLE: `[329 57 640 4 1000 72]` -> `[4 57 72 329 640 1000]`
// BIGO: O(D * (N + 10)) time for D-digit keys, O(N) space: each digit pass is a stable distribution, so earlier digits stay in order within equal later ones.
func RadixSort(nums []int) []int {
	panic("todo")
}

// ==== END

// ==== EX 3.3.4 | Bucket sort | Practice
// PROMPT: Every value in `values` is a float in `[0, 1)`. Write `BucketSort(values []float64) []float64` that returns a new sorted slice: put each value `v` into bucket `int(v * n)` of `n = len(values)` buckets, sort each (small) bucket, and join them.
// EXAMPLE: `[0.42 0.07 0.93 0.5 0.31]` -> `[0.07 0.31 0.42 0.5 0.93]`
// BIGO: O(N) expected time for evenly spread values, O(N) space: each of the N buckets gets about one value, so sorting the buckets is cheap; the worst case is one full bucket.
func BucketSort(values []float64) []float64 {
	panic("todo")
}

// ==== END
