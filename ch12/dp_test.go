package ch12

import (
	"testing"

	"dsa/internal/check"
)

func TestEx12_1(t *testing.T) {
	check.Eq(t, []int{FibNaive(10), FibMemo(10, map[int]int{}), FibBottomUp(10)}, []int{55, 55, 55})
	check.Eq(t, []int{FibMemo(80, map[int]int{}), FibBottomUp(80), FibBottomUp(92)}, []int{23416728348467685, 23416728348467685, 7540113804746346429})
	check.Eq(t, []int{FibBottomUp(0), FibBottomUp(1)}, []int{0, 1})
}

func TestEx12_2(t *testing.T) {
	check.Eq(t, []int{MaxRecursive([]int{1, 2, 3, 4}), MaxRecursive([]int{9, 2, 3})}, []int{4, 9})
}

func TestEx12_3(t *testing.T) {
	check.Eq(t, []int{StaircaseMemo(5), StaircaseMemo(50)}, []int{13, 10562230626642})
}

func TestEx12_4(t *testing.T) {
	check.Eq(t, []int{UniquePathsMemo(3, 7, map[[2]int]int{}), UniquePathsMemo(18, 18, map[[2]int]int{})}, []int{28, 2333606220})
}

func TestEx12_5(t *testing.T) {
	n, m := CountFibCalls(10)
	check.Eq(t, []int{n, m}, []int{177, 19})
}

func TestEx12_6(t *testing.T) {
	check.Eq(t, []int{MinCoins([]int{1, 5, 10, 25}, 63), MinCoins([]int{2}, 3), MinCoins([]int{1, 3, 4}, 6), MinCoins([]int{1}, 0)}, []int{6, -1, 2, 0})
}
