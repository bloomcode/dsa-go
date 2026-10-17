package ch72

import (
	"testing"

	"dsa/internal/check"
)

func TestEx72_1(t *testing.T) {
	check.Eq(t, []int{FibNaive(10), FibMemo(10, map[int]int{}), FibBottomUp(10)}, []int{55, 55, 55})
	check.Eq(t, []int{FibMemo(80, map[int]int{}), FibBottomUp(80), FibBottomUp(92)}, []int{23416728348467685, 23416728348467685, 7540113804746346429})
	check.Eq(t, []int{FibBottomUp(0), FibBottomUp(1)}, []int{0, 1})
}

func TestEx72_2(t *testing.T) {
	lo, hi := MinMax([]int{4, -2, 9, 0}, 0)
	check.Eq(t, []int{lo, hi}, []int{-2, 9})
	lo, hi = MinMax([]int{7}, 0)
	check.Eq(t, []int{lo, hi}, []int{7, 7})
	big := make([]int, 60)
	for i := range big {
		big[i] = i
	}
	lo, hi = MinMax(big, 0)
	check.Eq(t, []int{lo, hi}, []int{0, 59})
}

func TestEx72_3(t *testing.T) {
	check.Eq(t, TilingsMemo(10, map[int]int{}), 28)
	check.Eq(t, TilingsMemo(60, map[int]int{}), 5582216355)
	check.Eq(t, TilingsMemo(90, map[int]int{}), 533352837917926)
}

func TestEx72_4(t *testing.T) {
	m := func() map[[2]int]int { return map[[2]int]int{} }
	check.Eq(t, []int{DiceWaysMemo(2, 7, m()), DiceWaysMemo(3, 10, m()), DiceWaysMemo(10, 35, m())}, []int{6, 27, 4395456})
	check.Eq(t, DiceWaysMemo(20, 70, m()), 189456975899496)
}

func TestEx72_5(t *testing.T) {
	n, m := CountFibCalls(10)
	check.Eq(t, []int{n, m}, []int{177, 19})
}

func TestEx72_6(t *testing.T) {
	check.Eq(t, []int{MinCoins([]int{1, 5, 10, 25}, 63), MinCoins([]int{2}, 3), MinCoins([]int{1, 3, 4}, 6), MinCoins([]int{1}, 0)}, []int{6, -1, 2, 0})
}
