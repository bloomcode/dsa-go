// Package ch19: Counting memory: space complexity.
package ch19


// ==== EX 19.1 | Uppercase: new array vs in place | Core
// PROMPT: Write `UppercaseCopy(words []string) []string` returning a new slice and `UppercaseInPlace(words []string)` modifying the original. State the space complexity of each.
// EXAMPLE: `[amy bob]` -> `[AMY BOB]` either way
// BIGO: Both O(N) time. The copy is O(N) extra space; in place is O(1) extra beyond the new strings themselves.
func UppercaseCopy(words []string) []string {
	panic("todo")
}

func UppercaseInPlace(words []string) {
	panic("todo")
}

// ==== END

// ==== EX 19.2 | Duplicates: three time/space trade-offs | Core
// PROMPT: Implement three duplicate checks and fill in a time/space table: nested loops (`DupNested`), a map (`DupMap`), and sort-in-place (`DupSortInPlace`, which reorders the input).
// EXAMPLE: All three: `[3 1 4 1]` -> `true`
// BIGO: Nested: O(N^2) time, O(1) space. Map: O(N) time, O(N) space. Sort in place: O(N log N) time and O(log N) stack for `slices.Sort`. Choose by which resource is scarcer.
func DupNested(arr []int) bool {
	panic("todo")
}

func DupMap(arr []int) bool {
	panic("todo")
}

func DupSortInPlace(arr []int) bool {
	panic("todo")
}

// ==== END

// ==== EX 19.3 | The hidden cost of recursion | Core
// PROMPT: Write `SumToRecursive(n int) int` and `SumToIterative(n int) int`. Run both with n = 10,000,000 and compare memory: Go grows the goroutine stack instead of failing, so the recursive version succeeds but uses hundreds of MB of stack.
// EXAMPLE: `SumToIterative(50_000)` -> `1250025000`; `SumToRecursive(1_000_000)` -> `500000500000` without a crash
// BIGO: Both O(N) time. Recursive uses O(N) stack space; iterative uses O(1).
// NOTE: The default stack ceiling is 1 GB on 64-bit systems (`debug.SetMaxStack` changes it). Past it, the program dies with "goroutine stack exceeds limit", which no `recover` can catch.
func SumToRecursive(n int) int {
	panic("todo")
}

func SumToIterative(n int) int {
	panic("todo")
}

// ==== END

// ==== EX 19.4 | Reverse an array without extra space | Core
// PROMPT: Write `ReverseCopy(arr []int) []int` (new slice) and compare with your 1.16 `ReverseInPlace` and the library's `slices.Reverse`. Which would you use for a 1 GB slice?
// EXAMPLE: `[1 2 3]` -> `[3 2 1]`
// BIGO: Copy: O(N) time, O(N) space. In place: O(N) time, O(1) space.
func ReverseCopy(arr []int) []int {
	panic("todo")
}

// ==== END

// ==== EX 19.5 | Palindrome in O(1) space | Practice
// PROMPT: Building a reversed copy costs O(N) memory. Write `IsPalindromeClean(s string) bool` that ignores case and anything that is not a letter or digit, using two byte indices and no new string (ASCII input).
// EXAMPLE: `"A man, a plan, a canal: Panama"` -> `true`
// BIGO: O(N) time, O(1) extra space. For full Unicode, decode runes from both ends with `utf8.DecodeRuneInString` and `utf8.DecodeLastRuneInString`.
func IsPalindromeClean(s string) bool {
	panic("todo")
}

// ==== END
