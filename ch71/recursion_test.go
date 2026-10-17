package ch71

import (
	"testing"

	"dsa/internal/check"
)

func TestEx71_1(t *testing.T) {
	check.Eq(t, []int{DigitCount(12345), DigitCount(0), DigitCount(9), DigitCount(10)}, []int{5, 1, 1, 2})
	check.Eq(t, DigitCount(1<<62), 19)
}

func TestEx71_2(t *testing.T) {
	check.Eq(t, []int{Factorial(5), Factorial(0), Factorial(20)}, []int{120, 1, 2432902008176640000})
}

func TestEx71_3(t *testing.T) {
	check.Eq(t, []int{SumRecursive([]int{1, 2, 3, 4}), SumRecursive(nil)}, []int{10, 0})
}

func TestEx71_4(t *testing.T) {
	check.Eq(t, MaxDepth([]any{1, []any{2, []any{3}}, 4}), 3)
	check.Eq(t, []int{MaxDepth(nil), MaxDepth([]any{1, 2}), MaxDepth([]any{[]any{[]any{[]any{}}}, []any{1}})}, []int{1, 1, 4})
}

func TestEx71_5(t *testing.T) {
	check.Eq(t, []int{Power(2, 10), Power(3, 0), Power(5, 3)}, []int{1024, 1, 125})
}

func TestEx71_6(t *testing.T) {
	ana := Person{"ana", []Person{{"bo", nil}, {"cy", []Person{{"di", nil}}}}}
	check.Eq(t, []int{Headcount(ana), Headcount(Person{Name: "solo"})}, []int{4, 1})
}

func TestEx71_7(t *testing.T) {
	check.Eq(t, []string{ToBinary(5), ToBinary(0), ToBinary(8), ToBinary(255)}, []string{"101", "0", "1000", "11111111"})
}

func TestEx71_8(t *testing.T) {
	data := []int{1, 2, 3, 4}
	RunningTotals(data, 1)
	check.Eq(t, data, []int{1, 3, 6, 10})
	mixed := []int{2, -2, 2}
	RunningTotals(mixed, 1)
	check.Eq(t, mixed, []int{2, 0, 2})
}

func TestEx71_9(t *testing.T) {
	for _, s := range []string{"level", "", "x", "abba"} {
		check.Eq(t, IsPalindromeRec(s, 0, len(s)-1), true)
	}
	for _, s := range []string{"levels", "abca"} {
		check.Eq(t, IsPalindromeRec(s, 0, len(s)-1), false)
	}
}

func TestEx71_10(t *testing.T) {
	check.Eq(t, []int{DigitSum(9045), DigitSum(0), DigitSum(7), DigitSum(99999)}, []int{18, 0, 7, 45})
}

func TestEx71_11(t *testing.T) {
	check.Eq(t, LongestWord([]string{"go", "java", "py", "rust"}, 0), "java")
	check.Eq(t, []string{LongestWord(nil, 0), LongestWord([]string{"a", "bb", "cc"}, 0)}, []string{"", "bb"})
}

func TestEx71_12(t *testing.T) {
	check.Eq(t, CollapseRuns([]int{1, 1, 2, 2, 2, 3, 1}, 0), []int{1, 2, 3, 1})
	check.Eq(t, CollapseRuns(nil, 0), []int{})
	check.Eq(t, CollapseRuns([]int{4, 4, 4}, 0), []int{4})
	check.Eq(t, CollapseRuns([]int{1, 2, 1}, 0), []int{1, 2, 1})
}

func TestEx71_13(t *testing.T) {
	check.Eq(t, []int{CollatzSteps(6), CollatzSteps(1), CollatzSteps(2), CollatzSteps(27)}, []int{8, 0, 1, 111})
}

func TestEx71_14(t *testing.T) {
	check.Eq(t, []int{GCD(48, 18), GCD(7, 0), GCD(0, 5), GCD(17, 5), GCD(270, 192)}, []int{6, 7, 5, 1, 6})
}

func TestEx71_15(t *testing.T) {
	var first []int
	for n := 0; n < 8; n++ {
		first = append(first, Tilings(n))
	}
	check.Eq(t, first, []int{1, 1, 1, 2, 3, 4, 6, 9})
	check.Eq(t, Tilings(10), 28)
}

func TestEx71_16(t *testing.T) {
	check.Eq(t, ParenStrings(2), []string{"(())", "()()"})
	check.Eq(t, ParenStrings(0), []string{""})
	check.Eq(t, ParenStrings(3), []string{"((()))", "(()())", "(())()", "()(())", "()()()"})
	check.Eq(t, len(ParenStrings(5)), 42)
}

func TestEx71_17(t *testing.T) {
	check.Eq(t, []int{DiceWays(2, 7), DiceWays(1, 7), DiceWays(0, 0), DiceWays(3, 10), DiceWays(2, 13)}, []int{6, 0, 1, 27, 0})
}

func TestEx71_18(t *testing.T) {
	check.Eq(t, Subsets([]int{1, 2}), [][]int{{}, {2}, {1}, {1, 2}})
	all := Subsets([]int{1, 2, 3, 4})
	check.Eq(t, len(all), 16)
	check.Eq(t, all[15], []int{1, 2, 3, 4})
}
