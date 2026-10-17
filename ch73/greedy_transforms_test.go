package ch73

import (
	"testing"

	"dsa/internal/check"
)

func TestEx73_1(t *testing.T) {
	check.Eq(t, []bool{HasPairWithDifference([]int{1, 8, 4, 11}, 3), HasPairWithDifference([]int{1, 5}, 3)}, []bool{true, false})
	check.Eq(t, []bool{HasPairWithDifference([]int{3, 3}, 0), HasPairWithDifference([]int{3}, 0), HasPairWithDifference([]int{10, -2}, 12)}, []bool{true, false, true})
}

func TestEx73_2(t *testing.T) {
	i, j, ok := WindowWithSum([]int{2, 4, 3, 5, 1}, 8)
	check.Eq(t, []any{i, j, ok}, []any{2, 3, true})
	_, _, ok = WindowWithSum([]int{1, 2}, 7)
	check.Eq(t, ok, false)
	i, j, _ = WindowWithSum([]int{1, 2, 3}, 6)
	check.Eq(t, []int{i, j}, []int{0, 2})
}

func TestEx73_3(t *testing.T) {
	check.Eq(t, []bool{CanReachEnd([]int{2, 3, 1, 1, 4}), CanReachEnd([]int{3, 2, 1, 0, 4})}, []bool{true, false})
	check.Eq(t, []bool{CanReachEnd([]int{0}), CanReachEnd([]int{0, 1}), CanReachEnd([]int{1, 1, 1})}, []bool{true, false, true})
}

func TestEx73_4(t *testing.T) {
	check.Eq(t, RoomsNeeded([][2]int{{9, 11}, {10, 12}, {11, 13}, {9, 10}}), 2)
	check.Eq(t, []int{RoomsNeeded(nil), RoomsNeeded([][2]int{{1, 5}, {2, 6}, {3, 7}}), RoomsNeeded([][2]int{{1, 2}, {2, 3}})}, []int{0, 3, 1})
}

func TestEx73_5(t *testing.T) {
	check.Eq(t, MaxMeetings([][2]int{{1, 4}, {3, 5}, {0, 6}, {5, 7}, {8, 9}, {5, 9}}), 3)
	check.Eq(t, []int{MaxMeetings(nil), MaxMeetings([][2]int{{1, 10}, {2, 3}, {3, 4}, {4, 5}}), MaxMeetings([][2]int{{1, 2}, {1, 2}})}, []int{0, 3, 1})
}

func TestEx73_6(t *testing.T) {
	r := NewRangeSum([]int{3, 1, 4, 1, 5})
	check.Eq(t, []int{r.Total(1, 3), r.Total(0, 4), r.Total(2, 2), NewRangeSum([]int{-2, 2}).Total(0, 1)}, []int{6, 14, 4, 0})
}

func TestEx73_7(t *testing.T) {
	check.Eq(t, TopKFrequent([]string{"b", "a", "b", "c", "a", "b"}, 2), []string{"b", "a"})
	check.Eq(t, TopKFrequent([]string{"x", "y", "z"}, 2), []string{"x", "y"})
	check.Eq(t, TopKFrequent([]string{"q"}, 5), []string{"q"})
	check.Eq(t, TopKFrequent([]string{"m", "n", "n", "m", "o"}, 3), []string{"m", "n", "o"})
}

func TestEx73_8(t *testing.T) {
	check.Eq(t, []int{LoneValue([]int{4, 1, 2, 1, 2}), LoneValue([]int{7}), LoneValue([]int{0, 9, 9}), LoneValue([]int{-3, 5, 5})}, []int{4, 7, 0, -3})
}

func TestEx73_9(t *testing.T) {
	check.Eq(t, []int{LongestUniqueRun("abcadbe"), LongestUniqueRun("aaaa"), LongestUniqueRun(""), LongestUniqueRun("abba"), LongestUniqueRun("tmmzuxt")}, []int{5, 1, 0, 2, 5})
}
