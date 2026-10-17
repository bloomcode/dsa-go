// Package ch22: Searching sorted data.
package ch22

// ==== EX 22.1 | Ordered array insert | Core
// PROMPT: Create `type OrderedArray struct{ Items []int }` that is always sorted, with `NewOrderedArray(values ...int)`. Write `Insert(value int)`: scan from the left to find the first element larger than value, grow by one, shift everything from there one slot right and write the value. Do not use `slices.Insert` or a sort.
// EXAMPLE: `[3 17 80 202]` -> `Insert(75)` -> `[3 17 75 80 202]`
// BIGO: O(N). Comparisons up to the spot plus shifts after it always total about N + 2 steps.
// NOTE: Insertion is slower than in a plain array (which appends in O(1)), the price paid for O(log N) search.
type OrderedArray struct{ Items []int }

func NewOrderedArray(values ...int) *OrderedArray {
	panic("todo")
}

func (o *OrderedArray) Insert(value int) {
	panic("todo")
}

// ==== END

// ==== EX 22.2 | Linear search with early exit | Core
// PROMPT: Write `LinearSearchOrdered(arr []int, target int) int` for a sorted slice. Return the index, or `-1` as soon as you pass a value larger than the target.
// EXAMPLE: `[3 17 75 80 202]`, target `22` -> `-1` after inspecting only 3 elements
// BIGO: O(N) worst case, but it often stops early.
func LinearSearchOrdered(arr []int, target int) int {
	panic("todo")
}

// ==== END

// ==== EX 22.3 | Binary search, iterative | Core
// PROMPT: Write `BinarySearch(arr []int, target int) int` on a sorted slice using `lower` and `upper` bounds and a midpoint. Return the index or `-1`.
// EXAMPLE: `[3 17 75 80 202]`, target `80` -> `3`
// BIGO: O(log N). Each step halves the range: 100 elements need at most 7 steps, 1,000,000 need 20.
// NOTE: Write the midpoint as `lower + (upper-lower)/2` so it cannot overflow, whatever the slice size.
func BinarySearch(arr []int, target int) int {
	panic("todo")
}

// ==== END

// ==== EX 22.4 | Binary search, recursive | Practice
// PROMPT: Write `BinarySearchRecursive(arr []int, target int) int` with the same contract, delegating to a helper that takes `lower` and `upper`.
// EXAMPLE: `[1 3 5 7 9]`, target `1` -> `0`
// BIGO: O(log N) time, O(log N) call-stack space.
func BinarySearchRecursive(arr []int, target int) int {
	panic("todo")
}

func search(arr []int, target, lower, upper int) int {
	panic("todo")
}

// ==== END

// ==== EX 22.5 | Insertion point (lower bound) | Core
// PROMPT: Write `LowerBound(arr []int, target int) int`: the first index whose value is `>= target` (or `len(arr)` if none). You will reuse it in 2.6. Compare with `sort.SearchInts`, which is exactly this.
// EXAMPLE: `[10 20 20 30]`, target `20` -> `1`; target `25` -> `3`; target `99` -> `4`
// BIGO: O(log N).
// NOTE: The range is half-open `[lo, hi)`, so `hi` starts at `len(arr)`, not `len(arr) - 1`.
func LowerBound(arr []int, target int) int {
	panic("todo")
}

// ==== END

// ==== EX 22.6 | Insert using binary search for the position | Practice
// PROMPT: Add `func (o *OrderedArray) InsertFast(value int)` that finds the position with your `LowerBound` from 22.5 instead of a linear scan, then shifts.
// EXAMPLE: Same result as `Insert`, with only about log2(N) comparisons
// BIGO: Still O(N) overall because the shifts dominate; the comparisons drop to O(log N).
func (o *OrderedArray) InsertFast(value int) {
	panic("todo")
}

// ==== END

// ==== EX 22.7 | First and last occurrence | Practice
// PROMPT: Write `FirstAndLast(arr []int, target int) (int, int)` returning the first and last index of target in a sorted slice with duplicates, or `-1, -1`. Use two binary searches.
// EXAMPLE: `[1 2 2 2 5]`, target `2` -> `1, 3`
// BIGO: O(log N). A linear scan outward from one match would degrade to O(N) when everything matches.
func FirstAndLast(arr []int, target int) (int, int) {
	panic("todo")
}

// ==== END

// ==== EX 22.8 | Delete from an ordered array | Core
// PROMPT: Add `func (o *OrderedArray) Delete(value int) bool`: find it with binary search, shift the rest left, shrink by one and return `true`; return `false` if absent.
// EXAMPLE: `[2 4 6 8]` -> `Delete(4)` -> `[2 6 8]`
// BIGO: O(log N) to find + O(N) to shift = O(N).
func (o *OrderedArray) Delete(value int) bool {
	panic("todo")
}

// ==== END

// ==== EX 22.9 | Count the steps: linear vs binary | Practice
// PROMPT: Write `SearchSteps(n int) (linear, binary int)` that searches the slice `0..n-1` for a missing value (`n`) both ways and counts the elements each one inspects.
// EXAMPLE: `SearchSteps(100)` -> `100, 7`; `SearchSteps(10_000)` -> `10000, 14`
// BIGO: Linear grows 1:1 with N; binary grows by one step each time N doubles.
func SearchSteps(n int) (linear, binary int) {
	panic("todo")
}

// ==== END

// ==== EX 22.10 | Integer square root by binary search | Stretch
// PROMPT: Write `IntSqrt(n int) int` returning the largest r with `r*r <= n`, without `math.Sqrt`. Binary-search the answer between 1 and n.
// EXAMPLE: `IntSqrt(27)` -> `5`; `IntSqrt(36)` -> `6`; `IntSqrt(math.MaxInt64)` -> `3037000499`
// BIGO: O(log N). Binary search works on any sorted range of candidate answers, not only on slices.
// NOTE: Test `mid <= n/mid` rather than `mid*mid <= n`: for large n the square overflows and turns negative.
func IntSqrt(n int) int {
	panic("todo")
}

// ==== END
