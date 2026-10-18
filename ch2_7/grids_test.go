package ch2_7

import (
	"testing"

	"dsa/internal/check"
)

func TestEx2_7_1(t *testing.T) {
	check.Eq(t, Transpose([][]int{{1, 2, 3}, {4, 5, 6}}), [][]int{{1, 4}, {2, 5}, {3, 6}})
	check.Eq(t, Transpose([][]int{{7}}), [][]int{{7}})
	check.Eq(t, Transpose(nil), [][]int{})
	check.Eq(t, Transpose([][]int{{1}, {2}}), [][]int{{1, 2}})
}

func TestEx2_7_2(t *testing.T) {
	check.Eq(t, SpiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}), []int{1, 2, 3, 6, 9, 8, 7, 4, 5})
	check.Eq(t, SpiralOrder([][]int{{1, 2, 3}, {4, 5, 6}}), []int{1, 2, 3, 6, 5, 4})
	check.Eq(t, SpiralOrder([][]int{{1, 2}, {3, 4}, {5, 6}}), []int{1, 2, 4, 6, 5, 3})
	check.Eq(t, SpiralOrder([][]int{{1, 2, 3, 4}}), []int{1, 2, 3, 4})
	check.Eq(t, SpiralOrder([][]int{{1}, {2}, {3}}), []int{1, 2, 3})
	check.Eq(t, SpiralOrder(nil), []int{})
}

func TestEx2_7_3(t *testing.T) {
	grid := [][]int{{1, 4, 7, 11}, {2, 5, 8, 12}, {3, 6, 9, 16}, {10, 13, 14, 17}}
	check.Eq(t, FindInSortedGrid(grid, 9), [2]int{2, 2})
	check.Eq(t, FindInSortedGrid(grid, 15), [2]int{-1, -1})
	check.Eq(t, FindInSortedGrid(grid, 1), [2]int{0, 0})
	check.Eq(t, FindInSortedGrid(grid, 17), [2]int{3, 3})
	check.Eq(t, FindInSortedGrid(grid, 10), [2]int{3, 0})
	check.Eq(t, FindInSortedGrid(nil, 3), [2]int{-1, -1})
}

func TestEx2_7_4(t *testing.T) {
	check.Eq(t, ZigzagOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}), []int{1, 2, 4, 7, 5, 3, 6, 8, 9})
	check.Eq(t, ZigzagOrder([][]int{{1, 2, 3}, {4, 5, 6}}), []int{1, 2, 4, 5, 3, 6})
	check.Eq(t, ZigzagOrder([][]int{{1}, {2}, {3}}), []int{1, 2, 3})
	check.Eq(t, ZigzagOrder([][]int{{1, 2, 3}}), []int{1, 2, 3})
	check.Eq(t, ZigzagOrder(nil), []int{})
}

func TestEx2_7_5(t *testing.T) {
	for _, c := range []struct{ in, want [][]int }{
		{[][]int{{1, 2, 3}, {4, 0, 6}, {7, 8, 9}}, [][]int{{1, 0, 3}, {0, 0, 0}, {7, 0, 9}}},
		{[][]int{{0, 1}, {1, 1}}, [][]int{{0, 0}, {0, 1}}},
		{[][]int{{1, 1}, {1, 1}}, [][]int{{1, 1}, {1, 1}}},
		{[][]int{{5, 0, 7}}, [][]int{{0, 0, 0}}},
		{[][]int{{1, 2}, {3, 4}, {0, 6}}, [][]int{{0, 2}, {0, 4}, {0, 0}}},
	} {
		ZeroRowsCols(c.in)
		check.Eq(t, c.in, c.want)
	}
}

func TestEx2_7_6(t *testing.T) {
	for _, c := range []struct{ in, want [][]int }{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}},
		{[][]int{{1, 2}, {3, 4}}, [][]int{{3, 1}, {4, 2}}},
		{[][]int{{5}}, [][]int{{5}}},
		{[][]int{}, [][]int{}},
	} {
		RotateClockwise(c.in)
		check.Eq(t, c.in, c.want)
	}
}

func TestEx2_7_7(t *testing.T) {
	board := func(cells ...int) (b [9][9]int) {
		for i := 0; i < len(cells); i += 3 {
			b[cells[i]][cells[i+1]] = cells[i+2]
		}
		return b
	}
	var full [9][9]int
	for r := range 9 {
		for c := range 9 {
			full[r][c] = (r*3+r/3+c)%9 + 1
		}
	}
	check.True(t, ValidBoard(board()), "empty board")
	check.True(t, ValidBoard(full), "full valid board")
	check.True(t, ValidBoard(board(0, 0, 5, 4, 4, 5)), "same digit, different row, column and box")
	check.True(t, !ValidBoard(board(0, 0, 5, 1, 1, 5)), "repeat in a box")
	check.True(t, !ValidBoard(board(0, 4, 3, 8, 4, 3)), "repeat in a column")
	check.True(t, !ValidBoard(board(4, 0, 7, 4, 8, 7)), "repeat in a row")
	full[8][8] = full[8][7]
	check.True(t, !ValidBoard(full), "broken full board")
}
