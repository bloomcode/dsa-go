package ch11

import (
	"slices"
	"testing"

	"dsa/internal/check"
)

func TestEx11_1(t *testing.T) {
	a := []int{1, 2, 3}
	DoubleInPlace(a)
	check.Eq(t, a, []int{2, 4, 6})
}

func TestEx11_2(t *testing.T) {
	check.Eq(t, []string{Reverse("abcde"), Reverse("")}, []string{"edcba", ""})
}

func TestEx11_3(t *testing.T) {
	check.Eq(t, []int{CountX("axbxcxd"), CountX("")}, []int{3, 0})
}

func TestEx11_4(t *testing.T) {
	check.Eq(t, TotalChars([]string{"ab", "c", "def", "ghij"}), 10)
}

func TestEx11_5(t *testing.T) {
	check.Eq(t, SelectEvens([]int{1, 2, 3, 4, 5, 6}), []int{2, 4, 6})
}

func TestEx11_6(t *testing.T) {
	check.Eq(t, []int{Triangular(7), Triangular(1)}, []int{28, 1})
}

func TestEx11_7(t *testing.T) {
	check.Eq(t, []int{IndexOfX("abcdefghijklmnopqrstuvwxyz", 0), IndexOfX("abc", 0)}, []int{23, -1})
}

func TestEx11_8(t *testing.T) {
	check.Eq(t, []int{Staircase(3), Staircase(5)}, []int{4, 13})
}

func TestEx11_9(t *testing.T) {
	an := Anagrams("abc")
	slices.Sort(an)
	check.Eq(t, an, []string{"abc", "acb", "bac", "bca", "cab", "cba"})
	check.Eq(t, len(Anagrams("abcdef")), 720)
}

func TestEx11_10(t *testing.T) {
	check.Eq(t, UniquePaths(3, 7), 28)
}

func TestEx11_11(t *testing.T) {
	check.Eq(t, Subsets([]int{1, 2}), [][]int{{}, {2}, {1}, {1, 2}})
	all := Subsets([]int{1, 2, 3, 4})
	check.Eq(t, len(all), 16)
	check.Eq(t, all[15], []int{1, 2, 3, 4})
}
