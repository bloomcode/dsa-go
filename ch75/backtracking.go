// Package ch75: Backtracking.
package ch75


// ==== EX 75.1 | Every ordering | Core
// PROMPT: Write `Orderings(items []int) [][]int` that returns every ordering of the distinct `items`. Build one path: at each step try every unused item in input order, recurse, then undo the choice.
// EXAMPLE: `[1 2 3]` -> `[[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]]`
// BIGO: O(N * N!) time: N! orderings of N items each, plus O(N) for the path and the used flags.
func Orderings(items []int) [][]int {
	panic("todo")
}

// ==== END

// ==== EX 75.2 | Letters for a keypad number | Core
// PROMPT: On a phone keypad `2` is `abc`, `3` `def`, `4` `ghi`, `5` `jkl`, `6` `mno`, `7` `pqrs`, `8` `tuv` and `9` `wxyz`. Write `KeypadLetters(digits string) []string` that returns every letter string the digits could type, in keypad order (an empty, non-nil slice for no digits). Choose a letter for each digit in turn, recursing on the rest.
// EXAMPLE: `"23"` -> `["ad" "ae" "af" "bd" "be" "bf" "cd" "ce" "cf"]`
// BIGO: O(4^N * N) for N digits: up to four letters per digit, and each finished string costs O(N) to build.
func KeypadLetters(digits string) []string {
	panic("todo")
}

// ==== END

// ==== EX 75.3 | Combinations that hit a target | Practice
// PROMPT: `candidates` are distinct positive integers, each usable any number of times. Write `CombinationsTo(candidates []int, target int) [][]int` (`target > 0`) that returns every combination summing to `target` (a non-nil slice), each in ascending order and the list in dictionary order. Sort, then recurse from the current candidate onward, stopping once a candidate exceeds what is left.
// EXAMPLE: `[3 4 5], 11` -> `[[3 3 5] [3 4 4]]`
// BIGO: Exponential in the worst case: sorting lets a branch stop as soon as the next candidate overshoots, and only taking candidates at or after the last one avoids repeats.
func CombinationsTo(candidates []int, target int) [][]int {
	panic("todo")
}

// ==== END

// ==== EX 75.4 | Queens that do not attack | Practice
// PROMPT: Write `QueenPlacements(n int) int` that counts the ways to put `n` queens on an `n` by `n` board so that no two share a row, column or diagonal. Place one queen per row; track the used columns, `row - col` diagonals and `row + col` diagonals in boolean slices.
// EXAMPLE: `4` -> `2`
// BIGO: O(N!) worst case, far less in practice: one queen per row, and sets of used columns and diagonals reject a square in O(1).
func QueenPlacements(n int) int {
	panic("todo")
}

// ==== END

// ==== EX 75.5 | Subsets when values repeat | Practice
// PROMPT: Write `UniqueSubsets(items []int) [][]int` that returns every distinct subset of `items`, which may repeat values. Sort first; record the current path at every call, and when choosing the next item at one depth skip a value equal to the one just tried. Each subset is ascending, listed in the order this walk finds them.
// EXAMPLE: `[2 1 2]` -> `[[] [1] [1 2] [1 2 2] [2] [2 2]]`
// BIGO: O(N * 2^N) worst case: sorting puts equal values side by side, and skipping a repeat at the same depth stops the same subset being built twice.
// NOTE: Compare 71.18, which assumes distinct items.
func UniqueSubsets(items []int) [][]int {
	panic("todo")
}

// ==== END

// ==== EX 75.6 | Split digits into an address | Practice
// PROMPT: Write `AddressSplits(digits string) []string` that returns every way to insert three dots into `digits` so that the four parts are each `0` to `255` with no leading zero (`"0"` is fine, `"01"` is not), joined with `.` (a non-nil slice). Recurse part by part, trying lengths 1, 2 and 3, and stop early when too many or too few digits are left.
// EXAMPLE: `"101023"` -> `["1.0.10.23" "1.0.102.3" "10.1.0.23" "10.10.2.3" "101.0.2.3"]`
// BIGO: O(1) for a fixed four parts of up to three digits: at most 3^4 splits are tried, each checked in constant time.
func AddressSplits(digits string) []string {
	panic("todo")
}

// ==== END

// ==== EX 75.7 | Sudoku solver | Stretch
// PROMPT: `board` is 9 rows of 9 digits, `0` for empty. Write `SolveSudoku(board *[9][9]int) bool` that fills the empty cells in place so every row, column and 3 by 3 box holds 1 to 9 once, returning `true`, or returns `false` (board unchanged) when no filling exists. Keep a bit mask of used digits per row, column and box; always fill the empty cell with the fewest candidates next, and undo on failure.
// EXAMPLE: a full grid with 45 cells blanked -> `true`, and every row, column and box holds 1 to 9
// BIGO: Exponential in the worst case, fast on real puzzles: bit masks per row, column and box make each candidate check O(1), and filling the most constrained cell first prunes hardest.
// NOTE: 27.7 checks a board; this one fills it.
func SolveSudoku(board *[9][9]int) bool {
	panic("todo")
}

// ==== END
