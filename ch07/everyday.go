// Package ch07: Big O in day-to-day code.
package ch07


// ==== EX 7.1 | Mean of the even numbers | Core
// PROMPT: Write `MeanOfEvens(numbers []int) (float64, bool)`, returning false if there are no even numbers.
// EXAMPLE: `[1 2 3 4 6]` -> `4, true`
// BIGO: O(N). The loop runs once per number, whatever happens inside.
func MeanOfEvens(numbers []int) (float64, bool) {
	panic("todo")
}

// ==== END

// ==== EX 7.2 | Word builder, two characters | Core
// PROMPT: Write `TwoLetterWords(chars string) []string` returning every string made of two different positions in `chars`.
// EXAMPLE: `"abc"` -> `[ab ac ba bc ca cb]`
// BIGO: O(N^2).
func TwoLetterWords(chars string) []string {
	panic("todo")
}

// ==== END

// ==== EX 7.3 | Word builder, three characters | Practice
// PROMPT: Extend to `ThreeLetterWords(chars string) []string` using three distinct positions. How does the Big O change?
// EXAMPLE: `"abc"` -> 6 words: `abc acb bac bca cab cba`
// BIGO: O(N^3). Each extra nested loop over the same input adds one to the exponent.
func ThreeLetterWords(chars string) []string {
	panic("todo")
}

// ==== END

// ==== EX 7.4 | Array sample | Core
// PROMPT: Write `Sample(arr []int) [3]int` returning the first, middle and last of a non-empty slice.
// EXAMPLE: `[1 2 3 4 5]` -> `[1 3 5]`
// BIGO: O(1). Three reads by index, no matter how long the slice is.
func Sample(arr []int) [3]int {
	panic("todo")
}

// ==== END

// ==== EX 7.5 | Average Celsius reading | Core
// PROMPT: Write `AverageCelsius(fahrenheit []float64) float64` that converts each reading with (F - 32) / 1.8 into a new slice, then averages it in a second loop.
// EXAMPLE: `[32 212]` -> `50`
// BIGO: O(N). Two loops one after another are 2N steps, still O(N); only nesting multiplies.
func AverageCelsius(fahrenheit []float64) float64 {
	panic("todo")
}

// ==== END

// ==== EX 7.6 | Clothing labels | Core
// PROMPT: Write `ClothingLabels(items []string) []string` producing a label for every item in sizes 1 to 5.
// EXAMPLE: `[Shirt]` -> `[Shirt Size: 1 ... Shirt Size: 5]`
// BIGO: O(N). The inner loop is a fixed 5, so it is 5N steps, not N^2.
func ClothingLabels(items []string) []string {
	panic("todo")
}

// ==== END

// ==== EX 7.7 | Count the ones | Core
// PROMPT: Write `CountOnes(rows [][]int) int` for a jagged slice of 0s and 1s.
// EXAMPLE: `[[0 1 1] [1 0] [1 1 1 0]]` -> `6`
// BIGO: O(N) where N is the total number of values, even though the loops are nested.
func CountOnes(rows [][]int) int {
	panic("todo")
}

// ==== END

// ==== EX 7.8 | Palindrome checker | Core
// PROMPT: Write `IsPalindrome(s string) bool` comparing bytes from both ends inward.
// EXAMPLE: `"racecar"` -> `true`; `"kayaks"` -> `false`
// BIGO: O(N) (N/2 comparisons).
func IsPalindrome(s string) bool {
	panic("todo")
}

// ==== END

// ==== EX 7.9 | Products of all pairs | Core
// PROMPT: Write `PairProducts(nums []int) []int` returning the product of every unique pair (each pair once).
// EXAMPLE: `[1 2 3 4]` -> `[2 3 4 6 8 12]`
// BIGO: O(N^2): N(N-1)/2 pairs.
func PairProducts(nums []int) []int {
	panic("todo")
}

// ==== END

// ==== EX 7.10 | Products across two arrays | Practice
// PROMPT: Write `CrossProducts(a, b []int) []int` returning `x * y` for every x in a and y in b.
// EXAMPLE: `[1 2]`, `[10 100 1000]` -> `[10 100 1000 20 200 2000]`
// BIGO: O(N * M). With two different inputs, name both sizes instead of writing O(N^2).
func CrossProducts(a, b []int) []int {
	panic("todo")
}

// ==== END

// ==== EX 7.11 | Password cracker | Practice
// PROMPT: Write `AllPasswords(length int, alphabet string) []string` that returns every string of that length, in order, by extending each shorter password with every letter.
// EXAMPLE: `AllPasswords(2, "ab")` -> `[aa ab ba bb]`
// BIGO: O(26^N) for the full alphabet. Exponential: each extra character multiplies the work by the alphabet size.
func AllPasswords(length int, alphabet string) []string {
	panic("todo")
}

// ==== END
