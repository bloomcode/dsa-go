package ch19

import (
	"testing"

	"dsa/internal/check"
)

func TestEx19_1(t *testing.T) {
	orig := []string{"amy", "bob"}
	check.Eq(t, UppercaseCopy(orig), []string{"AMY", "BOB"})
	check.Eq(t, orig, []string{"amy", "bob"})
	UppercaseInPlace(orig)
	check.Eq(t, orig, []string{"AMY", "BOB"})
}

func TestEx19_2(t *testing.T) {
	check.True(t, DupNested([]int{3, 1, 4, 1}) && DupMap([]int{3, 1, 4, 1}) && DupSortInPlace([]int{3, 1, 4, 1}), "dups")
	check.True(t, !DupNested([]int{3, 1, 4}) && !DupMap([]int{3, 1, 4}) && !DupSortInPlace([]int{3, 1, 4}), "no dups")
}

func TestEx19_3(t *testing.T) {
	check.Eq(t, SumToIterative(50_000), 1250025000)
	check.Eq(t, SumToRecursive(1_000_000), 500000500000)
}

func TestEx19_4(t *testing.T) {
	orig := []int{1, 2, 3}
	check.Eq(t, ReverseCopy(orig), []int{3, 2, 1})
	check.Eq(t, orig, []int{1, 2, 3})
}

func TestEx19_5(t *testing.T) {
	check.True(t, IsPalindromeClean("A man, a plan, a canal: Panama") && !IsPalindromeClean("race a car"), "pal")
}
