package ch07

import (
	"math"
	"testing"

	"dsa/internal/check"
)

func TestEx7_1(t *testing.T) {
	m, ok := MeanOfEvens([]int{1, 2, 3, 4, 6})
	check.Eq(t, []any{m, ok}, []any{4.0, true})
	_, ok = MeanOfEvens([]int{1, 3})
	check.Eq(t, ok, false)
}

func TestEx7_2(t *testing.T) {
	check.Eq(t, TwoLetterWords("abc"), []string{"ab", "ac", "ba", "bc", "ca", "cb"})
}

func TestEx7_3(t *testing.T) {
	check.Eq(t, ThreeLetterWords("abc"), []string{"abc", "acb", "bac", "bca", "cab", "cba"})
}

func TestEx7_4(t *testing.T) {
	check.Eq(t, Sample([]int{1, 2, 3, 4, 5}), [3]int{1, 3, 5})
}

func TestEx7_5(t *testing.T) {
	check.True(t, math.Abs(AverageCelsius([]float64{32, 212})-50) < 1e-9, "celsius")
}

func TestEx7_6(t *testing.T) {
	labels := ClothingLabels([]string{"Shirt"})
	check.Eq(t, []any{len(labels), labels[0], labels[4]}, []any{5, "Shirt Size: 1", "Shirt Size: 5"})
}

func TestEx7_7(t *testing.T) {
	check.Eq(t, CountOnes([][]int{{0, 1, 1}, {1, 0}, {1, 1, 1, 0}}), 6)
}

func TestEx7_8(t *testing.T) {
	check.True(t, IsPalindrome("racecar") && !IsPalindrome("kayaks") && IsPalindrome(""), "pal")
}

func TestEx7_9(t *testing.T) {
	check.Eq(t, PairProducts([]int{1, 2, 3, 4}), []int{2, 3, 4, 6, 8, 12})
}

func TestEx7_10(t *testing.T) {
	check.Eq(t, CrossProducts([]int{1, 2}, []int{10, 100, 1000}), []int{10, 100, 1000, 20, 200, 2000})
}

func TestEx7_11(t *testing.T) {
	check.Eq(t, AllPasswords(2, "ab"), []string{"aa", "ab", "ba", "bb"})
	check.Eq(t, len(AllPasswords(3, "abcdefghijklmnopqrstuvwxyz")), 17576)
}
