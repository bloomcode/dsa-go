package ch7_3

import (
	"testing"

	"dsa/internal/check"
)

func TestEx7_3_1(t *testing.T) {
	check.Eq(t, []bool{HasPairWithDifference([]int{1, 8, 4, 11}, 3), HasPairWithDifference([]int{1, 5}, 3)}, []bool{true, false})
	check.Eq(t, []bool{HasPairWithDifference([]int{3, 3}, 0), HasPairWithDifference([]int{3}, 0), HasPairWithDifference([]int{10, -2}, 12)}, []bool{true, false, true})
}

func TestEx7_3_2(t *testing.T) {
	i, j, ok := WindowWithSum([]int{2, 4, 3, 5, 1}, 8)
	check.Eq(t, []any{i, j, ok}, []any{2, 3, true})
	_, _, ok = WindowWithSum([]int{1, 2}, 7)
	check.Eq(t, ok, false)
	i, j, _ = WindowWithSum([]int{1, 2, 3}, 6)
	check.Eq(t, []int{i, j}, []int{0, 2})
}

func TestEx7_3_3(t *testing.T) {
	check.Eq(t, []bool{CanReachEnd([]int{2, 3, 1, 1, 4}), CanReachEnd([]int{3, 2, 1, 0, 4})}, []bool{true, false})
	check.Eq(t, []bool{CanReachEnd([]int{0}), CanReachEnd([]int{0, 1}), CanReachEnd([]int{1, 1, 1})}, []bool{true, false, true})
}

func TestEx7_3_4(t *testing.T) {
	check.Eq(t, RoomsNeeded([][2]int{{9, 11}, {10, 12}, {11, 13}, {9, 10}}), 2)
	check.Eq(t, []int{RoomsNeeded(nil), RoomsNeeded([][2]int{{1, 5}, {2, 6}, {3, 7}}), RoomsNeeded([][2]int{{1, 2}, {2, 3}})}, []int{0, 3, 1})
}

func TestEx7_3_5(t *testing.T) {
	check.Eq(t, MaxMeetings([][2]int{{1, 4}, {3, 5}, {0, 6}, {5, 7}, {8, 9}, {5, 9}}), 3)
	check.Eq(t, []int{MaxMeetings(nil), MaxMeetings([][2]int{{1, 10}, {2, 3}, {3, 4}, {4, 5}}), MaxMeetings([][2]int{{1, 2}, {1, 2}})}, []int{0, 3, 1})
}

func TestEx7_3_6(t *testing.T) {
	r := NewRangeSum([]int{3, 1, 4, 1, 5})
	check.Eq(t, []int{r.Total(1, 3), r.Total(0, 4), r.Total(2, 2), NewRangeSum([]int{-2, 2}).Total(0, 1)}, []int{6, 14, 4, 0})
}

func TestEx7_3_7(t *testing.T) {
	check.Eq(t, TopKFrequent([]string{"b", "a", "b", "c", "a", "b"}, 2), []string{"b", "a"})
	check.Eq(t, TopKFrequent([]string{"x", "y", "z"}, 2), []string{"x", "y"})
	check.Eq(t, TopKFrequent([]string{"q"}, 5), []string{"q"})
	check.Eq(t, TopKFrequent([]string{"m", "n", "n", "m", "o"}, 3), []string{"m", "n", "o"})
}

func TestEx7_3_8(t *testing.T) {
	check.Eq(t, []int{LoneValue([]int{4, 1, 2, 1, 2}), LoneValue([]int{7}), LoneValue([]int{0, 9, 9}), LoneValue([]int{-3, 5, 5})}, []int{4, 7, 0, -3})
}

func TestEx7_3_9(t *testing.T) {
	check.Eq(t, []int{LongestUniqueRun("abcadbe"), LongestUniqueRun("aaaa"), LongestUniqueRun(""), LongestUniqueRun("abba"), LongestUniqueRun("tmmzuxt")}, []int{5, 1, 0, 2, 5})
}

func TestEx7_3_10(t *testing.T) {
	check.Eq(t, []int{LeastTotalWaiting([]int{3, 2, 1, 2, 6}), LeastTotalWaiting(nil), LeastTotalWaiting([]int{5}), LeastTotalWaiting([]int{4, 1}), LeastTotalWaiting([]int{2, 2, 2})}, []int{17, 0, 0, 1, 6})
}

func TestEx7_3_11(t *testing.T) {
	check.Eq(t, []int{PairFinishTime([]int{5, 2, 1, 6, 4, 4}), PairFinishTime([]int{1, 2}), PairFinishTime([]int{3, 3, 3, 3}), PairFinishTime([]int{1, 9, 2, 8})}, []int{8, 3, 6, 10})
}

func TestEx7_3_12(t *testing.T) {
	check.Eq(t, []bool{CanTakePhoto([]int{5, 8, 1, 3, 4}, []int{6, 9, 2, 4, 5}), CanTakePhoto([]int{1, 2}, []int{1, 3}), CanTakePhoto([]int{3}, []int{3}), CanTakePhoto([]int{6, 5}, []int{4, 4}), CanTakePhoto([]int{1, 9}, []int{5, 5})}, []bool{true, false, false, true, false})
}

func TestEx7_3_13(t *testing.T) {
	check.Eq(t, []int{TandemTotal([]int{5, 5, 3, 9, 2}, []int{3, 6, 7, 2, 1}, true), TandemTotal([]int{5, 5, 3, 9, 2}, []int{3, 6, 7, 2, 1}, false), TandemTotal([]int{1}, []int{1}, true), TandemTotal([]int{1, 4}, []int{5, 2}, true), TandemTotal([]int{1, 4}, []int{5, 2}, false)}, []int{32, 25, 1, 9, 7})
}

func TestEx7_3_14(t *testing.T) {
	check.Eq(t, []int{LoopStart([]int{1, 5, 3, 3, 1}, []int{3, 2, 4, 1, 1}), LoopStart([]int{2, 2}, []int{3, 3}), LoopStart([]int{3}, []int{2}), LoopStart([]int{1, 2, 3}, []int{2, 3, 1})}, []int{1, -1, 0, 2})
}

func TestEx7_3_15(t *testing.T) {
	check.Eq(t, []int{FewestRewards([]int{4, 2, 6, 3, 3, 1}), FewestRewards(nil), FewestRewards([]int{5}), FewestRewards([]int{1, 2, 3}), FewestRewards([]int{3, 3, 3}), FewestRewards([]int{1, 3, 2}), FewestRewards([]int{5, 4, 3, 2, 1})}, []int{9, 0, 1, 6, 3, 4, 15})
}

func TestEx7_3_16(t *testing.T) {
	check.Eq(t, []int{BestScheduleProfit([][2]int{{2, 50}, {1, 10}, {2, 40}, {1, 30}, {3, 20}}), BestScheduleProfit(nil), BestScheduleProfit([][2]int{{1, 5}, {1, 9}}), BestScheduleProfit([][2]int{{3, 1}, {3, 2}, {3, 3}, {3, 4}})}, []int{110, 0, 9, 9})
}

func TestEx7_3_17(t *testing.T) {
	check.Eq(t, StableMatching([][]int{{0, 1, 2}, {0, 2, 1}, {1, 0, 2}}, [][]int{{1, 0, 2}, {0, 1, 2}, {0, 1, 2}}), []int{1, 0, 2})
	check.Eq(t, StableMatching([][]int{{0}}, [][]int{{0}}), []int{0})
	check.Eq(t, StableMatching([][]int{{0, 1}, {0, 1}}, [][]int{{1, 0}, {1, 0}}), []int{1, 0})
}
