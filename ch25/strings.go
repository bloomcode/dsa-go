// Package ch25: Strings.
package ch25


// ==== EX 25.1 | Reverse the word order in place | Core
// PROMPT: `chars` holds words separated by single spaces. Write `ReverseWords(chars []byte)` that reverses the order of the words in place: reverse the whole slice, then reverse each word.
// EXAMPLE: `[]byte("sun is up")` -> `"up is sun"`
// BIGO: O(N) time, O(1) extra space: reversing everything puts the words in the right order but backwards; reversing each word fixes the letters.
func ReverseWords(chars []byte) {
	panic("todo")
}

// ==== END

// ==== EX 25.2 | One edit apart | Core
// PROMPT: Write `OneEditApart(a, b string) bool`: `true` when exactly one insertion, deletion or replacement of a byte turns `a` into `b` (equal strings are zero edits apart).
// EXAMPLE: `"cart", "card"` -> `true`; `"cart", "art"` -> `true`; `"cart", "cards"` -> `false`
// BIGO: O(N): one walk with two pointers; the first mismatch decides which pointer skips.
func OneEditApart(a, b string) bool {
	panic("todo")
}

// ==== END

// ==== EX 25.3 | First unique character | Core
// PROMPT: Write `FirstUniqueIndex(s string) int` returning the index of the first byte that appears exactly once in `s`, or -1.
// EXAMPLE: `"abacabad"` -> `3`; `"aabb"` -> `-1`
// BIGO: O(N): one pass counts, a second pass finds the first character with a count of one.
func FirstUniqueIndex(s string) int {
	panic("todo")
}

// ==== END

// ==== EX 25.4 | Letters needed to write any word | Core
// PROMPT: Write `LettersNeeded(words []string) string` returning, as one sorted string, the smallest bag of letters from which any single word can be spelled (each letter as many times as the word that needs it most).
// EXAMPLE: `[tool loot lot]` -> `"loot"`
// BIGO: O(C + A log A) for C characters in all and A distinct letters: one count table per word, keeping the largest count of each letter.
func LettersNeeded(words []string) string {
	panic("todo")
}

// ==== END

// ==== EX 25.5 | Reversed-word pairs | Core
// PROMPT: Write `ReversedPairs(words []string) [][2]string` for distinct words: whenever a word's reverse appeared earlier (and the word is not a palindrome), add `{earlierWord, word}`, in the order the second words appear.
// EXAMPLE: `[stop tops pots spot level]` -> `[[stop pots] [tops spot]]`
// BIGO: O(N * L) for N words of length L: one reversal and one set lookup per word.
func ReversedPairs(words []string) [][2]string {
	panic("todo")
}

// ==== END

// ==== EX 25.6 | Longest palindromic substring | Practice
// PROMPT: Write `LongestPalindrome(s string) string` returning the longest substring that reads the same both ways (the earliest one on a tie, `""` for an empty string). Expand outwards from every centre: each byte, and each gap between two bytes.
// EXAMPLE: `"noonabcbad"` -> `"abcba"`
// BIGO: O(N^2) time, O(1) space: 2N - 1 centres, each expanded at most N/2 steps.
func LongestPalindrome(s string) string {
	panic("todo")
}

// ==== END

// ==== EX 25.7 | Anagram windows | Practice
// PROMPT: Write `AnagramStarts(text, word string) []int` returning every index where `text` has a window of `len(word)` bytes that is an anagram of `word`. Slide the window one byte at a time, updating counts instead of recounting.
// EXAMPLE: `"tapatpta", "pat"` -> `[0 2 3 5]`
// BIGO: O(N * A) for a text of length N and A distinct letters: slide a fixed-size window, adding one character and removing one per step, then compare the two count tables.
func AnagramStarts(text, word string) []int {
	panic("todo")
}

// ==== END

// ==== EX 25.8 | Longest run after k changes | Practice
// PROMPT: Write `LongestAfterChanges(s string, k int) int`: the length of the longest substring that can be made all one byte by changing at most `k` of its bytes. Slide a window; it is valid while its length minus the count of its most common byte is at most `k`.
// EXAMPLE: `"aabcabb", 2` -> `5`
// BIGO: O(N): the window only grows or slides; the most common letter's count never needs to shrink for the answer to stay right.
func LongestAfterChanges(s string, k int) int {
	panic("todo")
}

// ==== END

// ==== EX 25.9 | Underscore every match | Practice
// PROMPT: Write `UnderscoreMatches(text, word string) string` that wraps every occurrence of `word` in `text` with `_` on each side, merging occurrences that overlap or touch into one wrapped stretch.
// EXAMPLE: `"lalala la", "la"` -> `"_lalala_ _la_"`
// BIGO: O(N * M) to find the matches of a word of length M, then O(N) to merge the spans and write the result.
func UnderscoreMatches(text, word string) string {
	panic("todo")
}

// ==== END

// ==== EX 25.10 | Smallest window holding every wanted letter | Stretch
// PROMPT: Write `SmallestWindow(text, wanted string) string` returning the shortest substring of `text` that contains every byte of `wanted`, counting repeats (the earliest on a tie, `""` if there is none). Grow the window until it covers `wanted`, then shrink it from the left while it still does.
// EXAMPLE: `"bxaxcaxb", "abc"` -> `"caxb"`
// BIGO: O(N + M): each index of the text enters and leaves the window at most once, and a counter of missing letters makes the validity check O(1).
func SmallestWindow(text, wanted string) string {
	panic("todo")
}

// ==== END
