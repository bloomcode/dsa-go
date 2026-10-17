package ch20

import (
	"testing"

	"dsa/internal/check"
)

func TestEx20_1(t *testing.T) {
	check.True(t, HasTwoSum([]int{2, 0, 4, 1, 7, 9}, 10) && !HasTwoSum([]int{2, 0, 4, 1, 7, 9}, 100) && !HasTwoSum([]int{5}, 10), "two sum")
}

func TestEx20_2(t *testing.T) {
	i, j, ok := SumSwap([]int{5, 3, 2, 9, 1}, []int{1, 12, 5})
	check.Eq(t, []any{i, j, ok}, []any{2, 0, true})
	_, _, ok = SumSwap([]int{1}, []int{2})
	check.Eq(t, ok, false)
}

func TestEx20_3(t *testing.T) {
	check.Eq(t, []int{MaxSubarraySum([]int{3, -4, 4, -3, 5, -9}), MaxSubarraySum([]int{-3, -1, -2})}, []int{6, -1})
}

func TestEx20_4(t *testing.T) {
	check.True(t, HasIncreasingTriplet([]float64{22, 25, 21, 18, 19.6, 17, 16, 20.5}) && !HasIncreasingTriplet([]float64{5, 4, 3}), "triplet")
}

func TestEx20_5(t *testing.T) {
	check.Eq(t, []int{MaxProfit([]int{10, 7, 5, 8, 11, 2, 6}), MaxProfit([]int{5, 4, 3})}, []int{6, 0})
}

func TestEx20_6(t *testing.T) {
	check.True(t, AreAnagrams("listen", "silent") && !AreAnagrams("listen", "silenc") && !AreAnagrams("aab", "abb"), "anagrams")
}

func TestEx20_7(t *testing.T) {
	check.Eq(t, SortTemperatures([]float64{98.6, 98.0, 97.1, 99.0, 98.9, 97.8, 98.5, 98.2, 98.0, 97.1}),
		[]float64{97.1, 97.1, 97.8, 98.0, 98.0, 98.2, 98.5, 98.6, 98.9, 99.0})
}

func TestEx20_8(t *testing.T) {
	check.Eq(t, MissingNumber([]int{5, 2, 4, 1, 0}), 3)
}

func TestEx20_9(t *testing.T) {
	check.Eq(t, []int{LongestConsecutive([]int{10, 5, 12, 3, 55, 30, 4, 11, 2}), LongestConsecutive(nil)}, []int{4, 0})
}
