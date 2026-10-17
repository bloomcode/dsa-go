// Package ch05: Selection sort and constant factors.
package ch05


// ==== EX 5.1 | Selection sort | Core
// PROMPT: Write `SelectionSort(arr []int)`: on each pass find the index of the smallest unsorted value and swap it into place (one swap per pass at most).
// EXAMPLE: `[4 2 7 1 3]` -> `[1 2 3 4 7]`
// BIGO: O(N^2) time, O(1) space. About N^2/2 comparisons but at most N - 1 swaps.
func SelectionSort(arr []int) {
	panic("todo")
}

// ==== END

// ==== EX 5.2 | Bubble vs selection: count the work | Practice
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

// ==== EX 5.3 | Every other element | Core
// PROMPT: Write `EveryOther(arr []int) []int` returning elements at indices 0, 2, 4 and so on, using a loop that steps by 2.
// EXAMPLE: `[10 11 12 13 14]` -> `[10 12 14]`
// BIGO: N/2 steps, which is still O(N): constants are dropped.
func EveryOther(arr []int) []int {
	panic("todo")
}

// ==== END

// ==== EX 5.4 | Second largest in one pass | Practice
// PROMPT: Write `SecondLargest(arr []int) (int, bool)` returning the second-largest distinct value in a single pass and without sorting; false if there is none.
// EXAMPLE: `[5 1 9 7 9]` -> `7, true`; `[3 3]` -> `0, false`
// BIGO: O(N) time, O(1) space. Sorting first would be O(N log N).
func SecondLargest(arr []int) (int, bool) {
	panic("todo")
}

// ==== END

// ==== EX 5.5 | K smallest by partial selection sort | Stretch
// PROMPT: Write `KSmallest(arr []int, k int) []int` that runs only the first k passes of selection sort on a copy and returns the first k values.
// EXAMPLE: `[8 3 5 1 9]`, k = 2 -> `[1 3]`
// BIGO: O(kN). Better than a full O(N^2) sort when k is small.
func KSmallest(arr []int, k int) []int {
	panic("todo")
}

// ==== END
