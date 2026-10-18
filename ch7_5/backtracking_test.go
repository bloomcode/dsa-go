package ch7_5

import (
	"testing"

	"dsa/internal/check"
)

func TestEx7_5_1(t *testing.T) {
	check.Eq(t, Orderings([]int{1, 2, 3}), [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}})
	check.Eq(t, Orderings(nil), [][]int{{}})
	check.Eq(t, Orderings([]int{7}), [][]int{{7}})
	check.Eq(t, len(Orderings([]int{1, 2, 3, 4, 5})), 120)
}

func TestEx7_5_2(t *testing.T) {
	check.Eq(t, KeypadLetters("23"), []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"})
	check.Eq(t, KeypadLetters(""), []string{})
	check.Eq(t, KeypadLetters("7"), []string{"p", "q", "r", "s"})
	check.Eq(t, len(KeypadLetters("79")), 16)
	check.Eq(t, KeypadLetters("99")[15], "zz")
}

func TestEx7_5_3(t *testing.T) {
	check.Eq(t, CombinationsTo([]int{3, 4, 5}, 11), [][]int{{3, 3, 5}, {3, 4, 4}})
	check.Eq(t, CombinationsTo([]int{2}, 1), [][]int{})
	check.Eq(t, CombinationsTo([]int{1}, 3), [][]int{{1, 1, 1}})
	check.Eq(t, CombinationsTo([]int{5, 2}, 7), [][]int{{2, 5}})
	check.Eq(t, len(CombinationsTo([]int{1, 2}, 6)), 4)
}

func TestEx7_5_4(t *testing.T) {
	got := []int{}
	for _, n := range []int{4, 1, 2, 3, 5, 6, 8} {
		got = append(got, QueenPlacements(n))
	}
	check.Eq(t, got, []int{2, 1, 0, 0, 10, 4, 92})
}

func TestEx7_5_5(t *testing.T) {
	check.Eq(t, UniqueSubsets([]int{2, 1, 2}), [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}})
	check.Eq(t, UniqueSubsets(nil), [][]int{{}})
	check.Eq(t, UniqueSubsets([]int{3, 3}), [][]int{{}, {3}, {3, 3}})
	check.Eq(t, []int{len(UniqueSubsets([]int{1, 2, 3})), len(UniqueSubsets([]int{1, 1, 2, 2}))}, []int{8, 9})
}

func TestEx7_5_6(t *testing.T) {
	check.Eq(t, AddressSplits("101023"), []string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"})
	check.Eq(t, AddressSplits("0000"), []string{"0.0.0.0"})
	check.Eq(t, AddressSplits("1111"), []string{"1.1.1.1"})
	check.Eq(t, AddressSplits("123"), []string{})
	check.Eq(t, AddressSplits("256256256256"), []string{})
	check.Eq(t, AddressSplits("255255255255"), []string{"255.255.255.255"})
}

func TestEx7_5_7(t *testing.T) {
	solved := func(b [9][9]int) bool {
		for i := range 9 {
			row, col, box := 0, 0, 0
			for j := range 9 {
				row |= 1 << b[i][j]
				col |= 1 << b[j][i]
				box |= 1 << b[i/3*3+j/3][i%3*3+j%3]
			}
			if row != 0x3FE || col != 0x3FE || box != 0x3FE {
				return false
			}
		}
		return true
	}
	var puzzle [9][9]int
	blanks := 0
	for r := range 9 {
		for c := range 9 {
			if (r*7+c*3)%9 < 5 {
				blanks++
			} else {
				puzzle[r][c] = (r*3+r/3+c)%9 + 1
			}
		}
	}
	given := puzzle
	check.Eq(t, blanks, 45)
	check.True(t, SolveSudoku(&puzzle) && solved(puzzle), "puzzle solved")
	for r := range 9 {
		for c := range 9 {
			check.True(t, given[r][c] == 0 || given[r][c] == puzzle[r][c], "givens kept")
		}
	}
	var stuck [9][9]int
	for c := range 8 {
		stuck[0][c] = c + 1
	}
	stuck[1][8] = 9
	before := stuck
	check.True(t, !SolveSudoku(&stuck) && stuck == before, "unsolvable board left unchanged")
	var empty [9][9]int
	check.True(t, SolveSudoku(&empty) && solved(empty), "empty board filled")
}
