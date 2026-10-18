package ch26

import (
	"testing"

	"dsa/internal/check"
)

func TestEx26_1(t *testing.T) {
	nums := []int{1, 3, 4, 6, 9}
	check.Eq(t, PairSumSorted(nums, 10), []int{0, 4})
	check.Eq(t, PairSumSorted(nums, 13), []int{2, 4})
	check.True(t, PairSumSorted(nums, 2) == nil, "no pair sums to 2")
	check.True(t, PairSumSorted(nil, 0) == nil, "empty input")
}

func TestEx26_2(t *testing.T) {
	check.Eq(t, SortedSquares([]int{-4, -1, 0, 3}), []int{0, 1, 9, 16})
	check.Eq(t, SortedSquares([]int{-3, -2}), []int{4, 9})
	check.Eq(t, SortedSquares([]int{2, 5}), []int{4, 25})
	check.Eq(t, SortedSquares([]int{}), []int{})
}

func TestEx26_3(t *testing.T) {
	check.Eq(t, SmallestGap([]int{5, 19, 31}, []int{12, 29, 1}), 2)
	check.Eq(t, SmallestGap([]int{4}, []int{4}), 0)
	check.Eq(t, SmallestGap([]int{-10, 50}, []int{7}), 17)
}

func TestEx26_4(t *testing.T) {
	for _, c := range []struct{ in, want []int }{
		{[]int{2, 1, 2, 3, 2, 4}, []int{1, 3, 4, 2, 2, 2}},
		{[]int{7, 7}, []int{7, 7}},
		{[]int{1, 2}, []int{1, 2}},
	} {
		value := 2
		if c.in[0] == 7 {
			value = 7
		} else if c.in[0] == 1 {
			value = 9
		}
		MoveToEnd(c.in, value)
		check.Eq(t, c.in, c.want)
	}
}

func TestEx26_5(t *testing.T) {
	got := []bool{IsMonotonic([]int{1, 2, 2, 5}), IsMonotonic([]int{3, 1, 2}), IsMonotonic(nil), IsMonotonic([]int{5, 5, 4, 0}), IsMonotonic([]int{1, 2, 1})}
	check.Eq(t, got, []bool{true, false, true, true, false})
}

func TestEx26_6(t *testing.T) {
	given := [][2]int{{6, 8}, {1, 3}, {2, 4}, {8, 9}}
	check.Eq(t, MergeIntervals(given), [][2]int{{1, 4}, {6, 9}})
	check.Eq(t, given, [][2]int{{6, 8}, {1, 3}, {2, 4}, {8, 9}})
	check.Eq(t, MergeIntervals(nil), [][2]int{})
	check.Eq(t, MergeIntervals([][2]int{{1, 10}, {2, 3}}), [][2]int{{1, 10}})
	check.Eq(t, MergeIntervals([][2]int{{4, 5}, {1, 2}}), [][2]int{{1, 2}, {4, 5}})
}

func TestEx26_7(t *testing.T) {
	check.Eq(t, TriplesToTarget([]int{4, -1, 2, 0, -3, 1}, 1), [][]int{{-3, 0, 4}, {-1, 0, 2}})
	check.Eq(t, TriplesToTarget([]int{0, 0, 0, 0}, 0), [][]int{{0, 0, 0}})
	check.Eq(t, TriplesToTarget([]int{1, 2}, 3), [][]int{})
	check.Eq(t, TriplesToTarget([]int{2, 2, 1, 1, 0}, 4), [][]int{{0, 2, 2}, {1, 1, 2}})
}

func TestEx26_8(t *testing.T) {
	check.Eq(t, []int{MostWater([]int{2, 7, 3, 6, 1, 5}), MostWater([]int{1, 1}), MostWater([]int{4}), MostWater([]int{3, 1, 1, 3})}, []int{20, 1, 0, 9})
}

func TestEx26_9(t *testing.T) {
	check.Eq(t, IntervalOverlaps([][2]int{{1, 4}, {6, 9}}, [][2]int{{3, 7}, {9, 12}}), [][2]int{{3, 4}, {6, 7}, {9, 9}})
	check.Eq(t, IntervalOverlaps([][2]int{{1, 2}}, [][2]int{{3, 4}}), [][2]int{})
	check.Eq(t, IntervalOverlaps(nil, [][2]int{{1, 5}}), [][2]int{})
	check.Eq(t, IntervalOverlaps([][2]int{{0, 10}}, [][2]int{{1, 2}, {4, 5}}), [][2]int{{1, 2}, {4, 5}})
}

func TestEx26_10(t *testing.T) {
	check.Eq(t, UnsortedStretch([]int{1, 3, 7, 5, 6, 9}), [2]int{2, 4})
	check.Eq(t, UnsortedStretch([]int{1, 2, 3}), [2]int{-1, -1})
	check.Eq(t, UnsortedStretch([]int{3, 2, 1}), [2]int{0, 2})
	check.Eq(t, UnsortedStretch([]int{1, 2, 4, 3, 3, 5}), [2]int{2, 4})
	check.Eq(t, UnsortedStretch([]int{2, 2}), [2]int{-1, -1})
}

func TestEx26_11(t *testing.T) {
	for _, c := range [][2][]int{
		{{3, -1, 4, -5, 9, -2}, {-1, -5, -2, 3, 4, 9}},
		{{}, {}},
		{{1, 2}, {1, 2}},
		{{0, -3, 0, -4, 7}, {-3, -4, 0, 0, 7}},
	} {
		NegativesFirst(c[0])
		check.Eq(t, c[0], c[1])
	}
}

func TestEx26_12(t *testing.T) {
	for _, c := range [][2][]int{
		{{1, 3, 2}, {2, 1, 3}},
		{{3, 2, 1}, {1, 2, 3}},
		{{1, 1, 5}, {1, 5, 1}},
		{{2, 4, 3, 1}, {3, 1, 2, 4}},
		{{}, {}},
	} {
		NextArrangement(c[0])
		check.Eq(t, c[0], c[1])
	}
}

func TestEx26_13(t *testing.T) {
	check.Eq(t, QuadsToTarget([]int{2, -1, 0, 3, -2, 1}, 2), [][]int{{-2, -1, 2, 3}, {-2, 0, 1, 3}, {-1, 0, 1, 2}})
	check.Eq(t, QuadsToTarget([]int{1, 1, 1, 1, 1}, 4), [][]int{{1, 1, 1, 1}})
	check.Eq(t, QuadsToTarget([]int{1, 2, 3}, 6), [][]int{})
}

func TestEx26_14(t *testing.T) {
	got := []int{TrappedWater([]int{3, 0, 2, 0, 4}), TrappedWater(nil), TrappedWater([]int{1, 2, 3}), TrappedWater([]int{5, 1, 1, 5}), TrappedWater([]int{4, 1, 3, 1, 2})}
	check.Eq(t, got, []int{7, 0, 0, 8, 3})
}
