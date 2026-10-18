// Package ch1_2: Memory as a cost.
package ch1_2


// ==== EX 1.2.1 | Scale scores: a copy or in place | Core
// PROMPT: Write `ScaledCopy(scores []int, factor int) []int` returning a new slice with every score multiplied by `factor`, and `ScaleInPlace(scores []int, factor int)` that changes the original slice. State the extra space each one needs.
// EXAMPLE: `[10 20], 3` -> `[30 60]` either way
// BIGO: Both O(N) time. The copy needs O(N) extra space; in place needs O(1).
func ScaledCopy(scores []int, factor int) []int {
	panic("todo")
}

func ScaleInPlace(scores []int, factor int) {
	panic("todo")
}

// ==== END

// ==== EX 1.2.2 | Most frequent value: three trade-offs | Core
// PROMPT: Find the most frequent value of a non-empty slice three ways, breaking ties toward the smaller value: `ModeNested(nums []int) int` (count each value with an inner loop), `ModeCounts(nums []int) int` (a map of counts) and `ModeSorted(nums []int) int` (sort the slice in place, then find the longest run).
// EXAMPLE: All three: `[4 1 4 2 1 4]` -> `4`
// BIGO: Nested loops: O(N^2) time, O(1) space. Counts table: O(N) time, O(N) space. Sort in place: O(N log N) time and no table of your own. Choose by which resource is scarcer.
func ModeNested(nums []int) int {
	panic("todo")
}

func ModeCounts(nums []int) int {
	panic("todo")
}

func ModeSorted(nums []int) int {
	panic("todo")
}

// ==== END

// ==== EX 1.2.3 | The call stack is memory too | Core
// PROMPT: Write `SumListRecursive(nums []int, i int) int` (the sum of `nums[i:]`, without reslicing) and `SumListIterative(nums []int) int`. Run both on a 1,000,000-item slice and compare memory: Go grows the goroutine stack instead of failing, so the recursive version succeeds but uses tens of MB of stack.
// EXAMPLE: `SumListIterative(0..49,999)` -> `1249975000`; `SumListRecursive` over 0..999,999 -> `499999500000` without a crash
// BIGO: Both O(N) time. The recursive version holds one stack frame per element, O(N) space; the loop uses O(1).
// NOTE: The default stack ceiling is 1 GB on 64-bit systems (`debug.SetMaxStack` changes it). Past it, the program dies with "goroutine stack exceeds limit", which no `recover` can catch.
func SumListRecursive(nums []int, i int) int {
	panic("todo")
}

func SumListIterative(nums []int) int {
	panic("todo")
}

// ==== END

// ==== EX 1.2.4 | Rotate an array without extra space | Core
// PROMPT: Rotate right by `k`: the last `k` items move to the front. Write `RotateCopy(arr []int, k int) []int` returning a new slice, and `RotateInPlace(arr []int, k int)` that reverses the whole slice, then the first `k` items, then the rest. Which would you use for a 1 GB slice?
// EXAMPLE: `[1 2 3 4 5], 2` -> `[4 5 1 2 3]`
// BIGO: Copy: O(N) time, O(N) space. Three reversals: O(N) time, O(1) space.
func RotateCopy(arr []int, k int) []int {
	panic("todo")
}

func RotateInPlace(arr []int, k int) {
	panic("todo")
}

// ==== END

// ==== EX 1.2.5 | Palindrome in O(1) space | Practice
// PROMPT: Building a reversed copy costs O(N) memory. Write `IsPalindromeClean(s string) bool` that ignores case and anything that is not a letter or digit, using two byte indices and no new string (ASCII input).
// EXAMPLE: `"A man, a plan, a canal: Panama"` -> `true`
// BIGO: O(N) time, O(1) extra space. For full Unicode, decode runes from both ends with `utf8.DecodeRuneInString` and `utf8.DecodeLastRuneInString`.
func IsPalindromeClean(s string) bool {
	panic("todo")
}

// ==== END
