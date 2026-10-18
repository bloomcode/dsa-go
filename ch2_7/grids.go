// Package ch2_7: Grids and matrices.
package ch2_7


// ==== EX 2.7.1 | Transpose a grid | Core
// PROMPT: Write `Transpose(grid [][]int) [][]int` that returns a new grid whose row `c` is column `c` of `grid` (rectangular; an empty, non-nil slice for an empty grid).
// EXAMPLE: `[[1 2 3] [4 5 6]]` -> `[[1 4] [2 5] [3 6]]`
// BIGO: O(R*C) time and space: every cell is copied once to its mirrored position.
func Transpose(grid [][]int) [][]int {
	panic("todo")
}

// ==== END

// ==== EX 2.7.2 | Spiral order | Core
// PROMPT: Write `SpiralOrder(grid [][]int) []int` that returns the values of a rectangular grid clockwise from the top-left corner, spiralling inward, as a non-nil slice. Keep four bounds (top, bottom, left, right) and shrink them after each ring.
// EXAMPLE: `[[1 2 3] [4 5 6] [7 8 9]]` -> `[1 2 3 6 9 8 7 4 5]`
// BIGO: O(R*C) time, O(1) extra space beyond the output: four shrinking bounds visit each cell exactly once.
func SpiralOrder(grid [][]int) []int {
	panic("todo")
}

// ==== END

// ==== EX 2.7.3 | Staircase search in a sorted grid | Core
// PROMPT: Every row and every column of `grid` is sorted ascending. Write `FindInSortedGrid(grid [][]int, target int) [2]int` that returns `[row col]` of `target`, or `[-1 -1]` when it is absent. Start at the top-right corner: move left when the value is too big, down when it is too small.
// EXAMPLE: `[[1 4 7 11] [2 5 8 12] [3 6 9 16] [10 13 14 17]], 9` -> `[2 2]`
// BIGO: O(R + C) time, O(1) space: from the top-right corner every comparison rules out a whole row or a whole column.
func FindInSortedGrid(grid [][]int, target int) [2]int {
	panic("todo")
}

// ==== END

// ==== EX 2.7.4 | Diagonal zigzag | Practice
// PROMPT: Write `ZigzagOrder(grid [][]int) []int` that returns the values of a rectangular grid diagonal by diagonal (a non-nil slice), where diagonal `d` holds the cells with `row + col == d`. Even diagonals run from bottom-left to top-right, odd ones from top-right to bottom-left.
// EXAMPLE: `[[1 2 3] [4 5 6] [7 8 9]]` -> `[1 2 4 7 5 3 6 8 9]`
// BIGO: O(R*C) time, O(1) extra space beyond the output: diagonal `d` holds the cells with `row + col == d`, and its row range follows from the grid size.
func ZigzagOrder(grid [][]int) []int {
	panic("todo")
}

// ==== END

// ==== EX 2.7.5 | Zero out rows and columns | Practice
// PROMPT: Write `ZeroRowsCols(grid [][]int)` that, for every cell holding `0`, sets its whole row and column to `0`, in place with O(1) extra space. Mark rows and columns in the first column and first row, and remember separately whether those two need clearing.
// EXAMPLE: `[[1 2 3] [4 0 6] [7 8 9]]` -> `[[1 0 3] [0 0 0] [7 0 9]]`
// BIGO: O(R*C) time, O(1) extra space: the first row and column record which columns and rows to clear, and two flags remember whether they need clearing themselves.
func ZeroRowsCols(grid [][]int) {
	panic("todo")
}

// ==== END

// ==== EX 2.7.6 | Rotate a square grid in place | Practice
// PROMPT: Write `RotateClockwise(grid [][]int)` that turns a square grid a quarter turn clockwise in place: transpose it by swapping across the main diagonal, then reverse every row.
// EXAMPLE: `[[1 2 3] [4 5 6] [7 8 9]]` -> `[[7 4 1] [8 5 2] [9 6 3]]`
// BIGO: O(N^2) time, O(1) space: a clockwise turn is a transpose followed by reversing each row, and both are swaps.
func RotateClockwise(grid [][]int) {
	panic("todo")
}

// ==== END

// ==== EX 2.7.7 | Validate a sudoku board | Practice
// PROMPT: `board` is 9 rows of 9 digits, `0` for an empty cell. Write `ValidBoard(board [9][9]int) bool`: `true` when no digit 1-9 repeats in any row, column or 3 by 3 box (the board need not be solvable). Keep a bit mask for each row, column and box.
// EXAMPLE: an empty board -> `true`; `5` at `(0, 0)` and `(1, 1)` -> `false` (same box)
// BIGO: O(1) for a fixed 9 by 9 board (O(N^2) for N by N): one pass, with a bit mask per row, column and box.
func ValidBoard(board [9][9]int) bool {
	panic("todo")
}

// ==== END
