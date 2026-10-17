// Package ch11: Thinking recursively.
package ch11

// ==== EX 11.1 | Double every element in place | Core
// PROMPT: Write `DoubleInPlace(arr []int)` recursively, modifying the slice rather than returning a new one ("repeatedly execute" category).
// EXAMPLE: `[1 2 3]` -> `[2 4 6]`
// BIGO: O(N). Because `arr[1:]` shares the backing array, writes through it change the caller's slice.
func DoubleInPlace(arr []int) {
	panic("todo")
}

// ==== END

// ==== EX 11.2 | Reverse a string | Core
// PROMPT: Write `Reverse(s string) string` top-down: assume `Reverse` already works on everything after the first byte.
// EXAMPLE: `"abcde"` -> `"edcba"`
// BIGO: O(N) calls, but each `+` builds a new string, so O(N^2) bytes are copied in total.
func Reverse(s string) string {
	panic("todo")
}

// ==== END

// ==== EX 11.3 | Count a character | Core
// PROMPT: Write `CountX(s string) int` returning how many times 'x' appears. Substrings of a Go string share its bytes, so `s[1:]` is O(1).
// EXAMPLE: `"axbxcxd"` -> `3`
// BIGO: O(N).
func CountX(s string) int {
	panic("todo")
}

// ==== END

// ==== EX 11.4 | Total characters across strings | Core
// PROMPT: Write `TotalChars(strs []string) int` returning the sum of all string lengths.
// EXAMPLE: `[ab c def ghij]` -> `10`
// BIGO: O(N) in the number of strings.
func TotalChars(strs []string) int {
	panic("todo")
}

// ==== END

// ==== EX 11.5 | Select the even numbers | Core
// PROMPT: Write `SelectEvens(nums []int) []int` returning a new slice of only the even numbers, recursively, building the result from the back with `append`.
// EXAMPLE: `[1 2 3 4 5 6]` -> `[2 4 6]`
// BIGO: O(N).
func SelectEvens(nums []int) []int {
	panic("todo")
}

// ==== END

// ==== EX 11.6 | Triangular numbers | Core
// PROMPT: Write `Triangular(n int) int` returning 1 + 2 + ... + n recursively.
// EXAMPLE: `Triangular(7)` -> `28`
// BIGO: O(N).
func Triangular(n int) int {
	panic("todo")
}

// ==== END

// ==== EX 11.7 | Index of the first x | Core
// PROMPT: Write `IndexOfX(s string, i int) int` returning the first index of 'x' at or after i, or -1.
// EXAMPLE: `IndexOfX("abcdefghijklmnopqrstuvwxyz", 0)` -> `23`
// BIGO: O(N).
func IndexOfX(s string, i int) int {
	panic("todo")
}

// ==== END

// ==== EX 11.8 | The staircase problem | Practice
// PROMPT: Write `Staircase(n int) int`: how many ways to climb n steps taking 1, 2 or 3 at a time. Think top-down: the last move was 1, 2 or 3 steps.
// EXAMPLE: `Staircase(3)` -> `4`; `Staircase(5)` -> `13`
// BIGO: About O(3^N) as written. Chapter 12 (12.3) brings it to O(N).
func Staircase(n int) int {
	panic("todo")
}

// ==== END

// ==== EX 11.9 | Anagram generation | Practice
// PROMPT: Write `Anagrams(s string) []string` returning every arrangement of the characters. For each anagram of `s[1:]`, insert `s[0]` at every position.
// EXAMPLE: `"abc"` -> `[abc bac bca acb cab cba]` (any order)
// BIGO: O(N!). Six letters already give 720 results.
func Anagrams(s string) []string {
	panic("todo")
}

// ==== END

// ==== EX 11.10 | Unique paths through a grid | Practice
// PROMPT: Write `UniquePaths(rows, cols int) int`: moving only right or down, how many paths lead from the top-left to the bottom-right cell?
// EXAMPLE: `UniquePaths(3, 7)` -> `28`
// BIGO: Exponential as written. Memoised in 12.4.
func UniquePaths(rows, cols int) int {
	panic("todo")
}

// ==== END

// ==== EX 11.11 | All subsets | Stretch
// PROMPT: Write `Subsets(items []int) [][]int` returning every subset (the power set). Each subset of the rest appears once without and once with the first item.
// EXAMPLE: `[1 2]` -> `[[] [2] [1] [1 2]]` (any order)
// BIGO: O(2^N) subsets.
// NOTE: Build each "with" subset in a fresh slice. `append(s, x)` on a shared subset could write into the same backing array twice.
func Subsets(items []int) [][]int {
	panic("todo")
}

// ==== END
