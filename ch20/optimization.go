// Package ch20: Making code faster: lookups, patterns and greed.
package ch20


// ==== EX 20.1 | Two sum with a magical lookup | Core
// PROMPT: Write `HasTwoSum(nums []int, target int) bool` returning true if any two numbers add to target. Ask: "if I could look up any value in O(1), what would I look for?"
// EXAMPLE: `[2 0 4 1 7 9]`, 10 -> `true` (1 + 9)
// BIGO: O(N) time, O(N) space; the brute force is O(N^2).
func HasTwoSum(nums []int, target int) bool {
	panic("todo")
}

// ==== END

// ==== EX 20.2 | Sum swap | Practice
// PROMPT: Write `SumSwap(a, b []int) (i, j int, ok bool)` returning indices such that swapping a[i] and b[j] makes the two sums equal. Spot the pattern: the swap must shift each sum by half the difference.
// EXAMPLE: `[5 3 2 9 1]` (sum 20), `[1 12 5]` (sum 18) -> `2, 0, true`: swapping 2 and 1 gives 19 and 19
// BIGO: O(N + M) with a map of one slice's values.
func SumSwap(a, b []int) (int, int, bool) {
	panic("todo")
}

// ==== END

// ==== EX 20.3 | Greatest subarray sum (greedy) | Core
// PROMPT: Write `MaxSubarraySum(nums []int) int` in one pass: keep a running sum, reset it to 0 when it goes negative, and track the best seen.
// EXAMPLE: `[3 -4 4 -3 5 -9]` -> `6` (4 - 3 + 5)
// BIGO: O(N) time, O(1) space (Kadane's algorithm), versus O(N^2) for every subarray.
// NOTE: Initialise `best` with the first value, not 0, so an all-negative slice returns its largest value.
func MaxSubarraySum(nums []int) int {
	panic("todo")
}

// ==== END

// ==== EX 20.4 | Increasing triplet (greedy) | Practice
// PROMPT: Write `HasIncreasingTriplet(nums []float64) bool`: are there i < j < k with nums[i] < nums[j] < nums[k]? Track the lowest value and a "middle" value in one pass, starting both at `math.Inf(1)`.
// EXAMPLE: `[22 25 21 18 19.6 17 16 20.5]` -> `true` (18, 19.6, 20.5); `[5 4 3]` -> `false`
// BIGO: O(N) time, O(1) space.
func HasIncreasingTriplet(nums []float64) bool {
	panic("todo")
}

// ==== END

// ==== EX 20.5 | Best stock trade | Core
// PROMPT: Write `MaxProfit(prices []int) int`: the most money from buying once and selling later (0 if prices only fall).
// EXAMPLE: `[10 7 5 8 11 2 6]` -> `6` (buy 5, sell 11)
// BIGO: O(N) greedy, one pass tracking the lowest price so far.
func MaxProfit(prices []int) int {
	panic("todo")
}

// ==== END

// ==== EX 20.6 | Anagram check by changing the data structure | Core
// PROMPT: Write `AreAnagrams(a, b string) bool` by counting runes in a `map[rune]int`, instead of generating permutations (11.9) or sorting.
// EXAMPLE: `"listen"`, `"silent"` -> `true`
// BIGO: O(N + M). Sorting is O(N log N); generating anagrams is O(N!).
func AreAnagrams(a, b string) bool {
	panic("todo")
}

// ==== END

// ==== EX 20.7 | Sort values from a small known range | Practice
// PROMPT: Body temperatures are recorded to one decimal place between 97.0 and 99.0. Write `SortTemperatures(temps []float64) []float64` in O(N) by counting readings in a `[21]int`, then emitting them in order.
// EXAMPLE: `[98.6 98.0 97.1 99.0 98.9 97.8 98.5 98.2 98.0 97.1]` -> `[97.1 97.1 97.8 98 98 98.2 98.5 98.6 98.9 99]`
// BIGO: O(N + R) where R = 21 possible readings, so O(N). A known small range beats O(N log N) sorting.
// NOTE: Convert with `math.Round(t * 10)`, not `int(t * 10)`: 98.6 is not exact in binary, and truncation can give 985.
func SortTemperatures(temps []float64) []float64 {
	panic("todo")
}

// ==== END

// ==== EX 20.8 | Missing number in O(N) | Core
// PROMPT: Revisit 13.6. Write `MissingNumber(arr []int) int` using the formula 0 + 1 + ... + n = n(n + 1) / 2.
// EXAMPLE: `[5 2 4 1 0]` -> `3`
// BIGO: O(N) time, O(1) space, versus O(N log N) by sorting.
func MissingNumber(arr []int) int {
	panic("todo")
}

// ==== END

// ==== EX 20.9 | Longest consecutive run | Stretch
// PROMPT: Write `LongestConsecutive(nums []int) int` returning the length of the longest run of consecutive integers in an unsorted slice, in O(N). Only start counting from a number whose predecessor is absent.
// EXAMPLE: `[10 5 12 3 55 30 4 11 2]` -> `4` (2, 3, 4, 5)
// BIGO: O(N): each number is visited at most twice across all runs.
func LongestConsecutive(nums []int) int {
	panic("todo")
}

// ==== END
