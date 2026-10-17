package ch03

import (
	"testing"

	"dsa/internal/check"
)

func TestEx3_1(t *testing.T) {
	check.Eq(t, []bool{IsLeapYear(2024), IsLeapYear(1900), IsLeapYear(2000), IsLeapYear(2023)}, []bool{true, false, true, false})
}

func TestEx3_2(t *testing.T) {
	check.Eq(t, []int{ArraySum([]int{1, 2, 3, 4}), ArraySum(nil)}, []int{10, 0})
}

func TestEx3_3(t *testing.T) {
	check.Eq(t, []int{HalvingSteps(8), HalvingSteps(1024), HalvingSteps(1)}, []int{3, 10, 0})
}

func TestEx3_4(t *testing.T) {
	check.Eq(t, []int{SquaresNeeded(1), SquaresNeeded(16), SquaresNeeded(17)}, []int{1, 5, 6})
}

func TestEx3_5(t *testing.T) {
	check.Eq(t, []int{AllPairsSteps(4), UniquePairsSteps(4)}, []int{16, 6})
}
