package ch1_1

import (
	"testing"

	"dsa/internal/check"
)

func TestEx1_1_1(t *testing.T) {
	check.Eq(t, []bool{IsLeapYear(2024), IsLeapYear(1900), IsLeapYear(2000), IsLeapYear(2023)}, []bool{true, false, true, false})
}

func TestEx1_1_2(t *testing.T) {
	check.Eq(t, []int{ArraySum([]int{1, 2, 3, 4}), ArraySum(nil)}, []int{10, 0})
}

func TestEx1_1_3(t *testing.T) {
	check.Eq(t, []int{HalvingSteps(8), HalvingSteps(1024), HalvingSteps(1)}, []int{3, 10, 0})
}

func TestEx1_1_4(t *testing.T) {
	check.Eq(t, []int{RoundsToReach(1), RoundsToReach(2), RoundsToReach(3)}, []int{0, 1, 2})
	check.Eq(t, []int{RoundsToReach(1000), RoundsToReach(1024), RoundsToReach(1025)}, []int{10, 10, 11})
}

func TestEx1_1_5(t *testing.T) {
	check.Eq(t, []int{AllPairsSteps(4), UniquePairsSteps(4)}, []int{16, 6})
}

func TestEx1_1_6(t *testing.T) {
	check.Eq(t, FinalBalance([]int{50, -20, -40, 10}), 40)
	check.Eq(t, []int{FinalBalance(nil), FinalBalance([]int{-5}), FinalBalance([]int{5, -5, -1})}, []int{0, 0, 0})
}

func TestEx1_1_7(t *testing.T) {
	check.Eq(t, Pairings([]string{"ana", "bo", "cy"}), []string{"ana-bo", "ana-cy", "bo-cy"})
	check.Eq(t, Pairings([]string{"solo"}), []string{})
	check.Eq(t, len(Pairings(make([]string, 10))), 45)
}

func TestEx1_1_8(t *testing.T) {
	check.Eq(t, Committees([]string{"a", "b", "c", "d"}), []string{"a,b,c", "a,b,d", "a,c,d", "b,c,d"})
	check.Eq(t, Committees([]string{"a", "b"}), []string{})
	check.Eq(t, len(Committees(make([]string, 10))), 120)
}

func TestEx1_1_9(t *testing.T) {
	check.Eq(t, Newest([]string{"a", "b", "c", "d"}, 2), []string{"c", "d"})
	check.Eq(t, Newest([]string{"a"}, 5), []string{"a"})
	check.Eq(t, Newest([]string{"a", "b"}, 0), []string{})
}

func TestEx1_1_10(t *testing.T) {
	check.Eq(t, BusiestHour([]string{"09:15", "10:02", "09:40", "11:00"}), 9)
	check.Eq(t, BusiestHour([]string{"10:00", "09:00"}), 9)
	check.Eq(t, BusiestHour([]string{"23:59", "23:00", "00:10"}), 23)
}

func TestEx1_1_11(t *testing.T) {
	check.Eq(t, WeeklySlots([]string{"ana"}), []string{"ana Mon", "ana Tue", "ana Wed", "ana Thu", "ana Fri", "ana Sat", "ana Sun"})
	slots := WeeklySlots([]string{"ana", "bo"})
	check.Eq(t, []any{len(slots), slots[7]}, []any{14, "bo Mon"})
	check.Eq(t, WeeklySlots(nil), []string{})
}

func TestEx1_1_12(t *testing.T) {
	check.Eq(t, CountVowels([]string{"Hello", "World"}), 3)
	check.Eq(t, []int{CountVowels([]string{"AEIOU xyz", ""}), CountVowels(nil)}, []int{5, 0})
}

func TestEx1_1_13(t *testing.T) {
	check.Eq(t, IsSymmetric([][]int{{1, 2}, {2, 1}}), true)
	check.Eq(t, IsSymmetric([][]int{{1, 2}, {3, 1}}), false)
	check.Eq(t, []bool{IsSymmetric(nil), IsSymmetric([][]int{{7}})}, []bool{true, true})
	check.Eq(t, IsSymmetric([][]int{{1, 2, 3}, {2, 5, 6}, {3, 6, 9}}), true)
	check.Eq(t, IsSymmetric([][]int{{1, 2, 3}, {2, 5, 6}, {3, 0, 9}}), false)
}

func TestEx1_1_14(t *testing.T) {
	check.Eq(t, ClosestGap([]int{10, 4, 15, 7}), 3)
	check.Eq(t, []int{ClosestGap([]int{1, 100}), ClosestGap([]int{5, 5, 9}), ClosestGap([]int{-3, 8, -1})}, []int{99, 0, 2})
}

func TestEx1_1_15(t *testing.T) {
	check.Eq(t, SharedHours([]int{9, 10, 13, 15}, []int{10, 11, 15}), []int{10, 15})
	check.Eq(t, SharedHours(nil, []int{1}), []int{})
	check.Eq(t, SharedHours([]int{8}, nil), []int{})
	check.Eq(t, SharedHours([]int{7, 8}, []int{8, 7}), []int{7, 8})
}

func TestEx1_1_16(t *testing.T) {
	check.Eq(t, SwitchSettings(2), []string{"00", "01", "10", "11"})
	check.Eq(t, SwitchSettings(0), []string{""})
	check.Eq(t, SwitchSettings(1), []string{"0", "1"})
	three := SwitchSettings(3)
	check.Eq(t, []any{len(three), three[0], three[5], three[7]}, []any{8, "000", "101", "111"})
}
