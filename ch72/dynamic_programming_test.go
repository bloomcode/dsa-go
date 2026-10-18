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

func TestEx72_7(t *testing.T) {
	check.Eq(t, []int{BestWithoutNeighbours([]int{5, 1, 2, 9, 3}), BestWithoutNeighbours(nil), BestWithoutNeighbours([]int{4}), BestWithoutNeighbours([]int{2, 7, 9, 3, 1}), BestWithoutNeighbours([]int{3, 1, 1, 3})}, []int{14, 0, 4, 12, 6})
}

func TestEx72_8(t *testing.T) {
	check.Eq(t, []int{ChangeWays([]int{2, 3, 5}, 10), ChangeWays([]int{2, 3, 5}, 0), ChangeWays([]int{3}, 7), ChangeWays([]int{1}, 4), ChangeWays([]int{1, 2}, 4)}, []int{4, 1, 0, 1, 3})
}

func TestEx72_9(t *testing.T) {
	check.Eq(t, []int{DecodeWays("1213"), DecodeWays(""), DecodeWays("06"), DecodeWays("10"), DecodeWays("27"), DecodeWays("101"), DecodeWays("0"), DecodeWays("2626")}, []int{5, 1, 0, 1, 1, 1, 0, 4})
}

func TestEx72_10(t *testing.T) {
	check.Eq(t, []int{BestStreak([]int{3, -4, 5, -1, 2, -6, 4}), BestStreak([]int{-3, -1, -2}), BestStreak([]int{1}), BestStreak([]int{2, 3, -1, 4})}, []int{6, -1, 1, 8})
}

func TestEx72_11(t *testing.T) {
	check.Eq(t, []int{GridPaths([]string{"...", ".#.", "..."}), GridPaths([]string{"#"}), GridPaths([]string{"."}), GridPaths([]string{"..", ".."}), GridPaths([]string{"...", "...", "..."}), GridPaths([]string{".#", "#."}), GridPaths([]string{"..#", "..."})}, []int{2, 0, 1, 2, 6, 0, 2})
}

func TestEx72_12(t *testing.T) {
	check.Eq(t, []int{EditDistance("garden", "pardon"), EditDistance("", "abc"), EditDistance("abc", "abc"), EditDistance("abc", "yabd"), EditDistance("flour", "floor"), EditDistance("abc", "")}, []int{2, 3, 0, 2, 1, 3})
}

func TestEx72_13(t *testing.T) {
	check.Eq(t, []int{CommonSubsequenceLength("stone", "longest"), CommonSubsequenceLength("abc", "def"), CommonSubsequenceLength("", "a"), CommonSubsequenceLength("abcde", "ace"), CommonSubsequenceLength("aaaa", "aa")}, []int{3, 0, 0, 3, 2})
}

func TestEx72_14(t *testing.T) {
	check.Eq(t, []int{LongestIncreasingLength([]int{4, 1, 3, 2, 8, 5, 6}), LongestIncreasingLength(nil), LongestIncreasingLength([]int{5, 5, 5}), LongestIncreasingLength([]int{1, 2, 3}), LongestIncreasingLength([]int{3, 2, 1}), LongestIncreasingLength([]int{2, 9, 3, 10, 4, 5})}, []int{4, 0, 1, 3, 1, 4})
}

func TestEx72_15(t *testing.T) {
	check.Eq(t, []int{Knapsack([]int{3, 4, 2, 6}, []int{30, 50, 15, 60}, 8), Knapsack([]int{3, 4, 2, 6}, []int{30, 50, 15, 60}, 0), Knapsack([]int{3, 4, 2, 6}, []int{30, 50, 15, 60}, 10), Knapsack([]int{3, 4, 2, 6}, []int{30, 50, 15, 60}, 2), Knapsack([]int{3, 4, 2, 6}, []int{30, 50, 15, 60}, 1), Knapsack(nil, nil, 5)}, []int{80, 0, 110, 15, 0, 0})
}

func TestEx72_16(t *testing.T) {
	check.Eq(t, []int{FewestJumps([]int{3, 1, 2, 1, 1, 4}), FewestJumps([]int{0}), FewestJumps([]int{1, 0, 2}), FewestJumps([]int{5, 1, 1}), FewestJumps([]int{1, 1, 1}), FewestJumps([]int{2, 0, 0})}, []int{3, 0, -1, 1, 2, 1})
}

func TestEx72_17(t *testing.T) {
	check.Eq(t, []int{LargestSquare([]string{"0111", "1111", "0111"}), LargestSquare([]string{"0"}), LargestSquare([]string{"1"}), LargestSquare([]string{"11", "10"}), LargestSquare([]string{"110", "111", "011"}), LargestSquare(nil)}, []int{9, 0, 1, 1, 4, 0})
}

func TestEx72_18(t *testing.T) {
	words := []string{"sun", "flower", "field", "sunflow", "er"}
	got := []bool{CanSplit("sunflowerfield", words), CanSplit("sunfield", []string{"sun", "flower"}), CanSplit("", nil), CanSplit("erersun", words), CanSplit("sunflowe", words)}
	check.Eq(t, got, []bool{true, false, true, true, false})
}

func TestEx72_19(t *testing.T) {
	check.Eq(t, []int{MinPalindromeCuts("noonabbad"), MinPalindromeCuts("racecarx"), MinPalindromeCuts("abcb"), MinPalindromeCuts("abc"), MinPalindromeCuts(""), MinPalindromeCuts("aaaa"), MinPalindromeCuts("ab")}, []int{2, 1, 1, 2, 0, 0, 1})
}

func TestEx72_20(t *testing.T) {
	check.Eq(t, []int{MaxProfitK([]int{3, 8, 2, 6, 1, 9}, 2), MaxProfitK([]int{3, 8, 2, 6, 1, 9}, 1), MaxProfitK([]int{3, 8, 2, 6, 1, 9}, 3), MaxProfitK([]int{3, 8, 2, 6, 1, 9}, 0), MaxProfitK(nil, 2), MaxProfitK([]int{5, 4, 3}, 2)}, []int{13, 8, 17, 0, 0, 0})
}
