package ch12

import (
	"testing"

	"dsa/internal/check"
)

func TestEx12_1(t *testing.T) {
	scores := []int{10, 20}
	check.Eq(t, ScaledCopy(scores, 3), []int{30, 60})
	check.Eq(t, scores, []int{10, 20})
	ScaleInPlace(scores, 3)
	check.Eq(t, scores, []int{30, 60})
	check.Eq(t, ScaledCopy([]int{}, 5), []int{})
}

func TestEx12_2(t *testing.T) {
	for _, fn := range []func([]int) int{ModeNested, ModeCounts, ModeSorted} {
		check.Eq(t, fn([]int{4, 1, 4, 2, 1, 4}), 4)
		check.Eq(t, []int{fn([]int{2, 1, 2, 1}), fn([]int{7}), fn([]int{-3, 5, -3})}, []int{1, 7, -3})
	}
}

func TestEx12_3(t *testing.T) {
	upTo := func(n int) []int {
		out := make([]int, n)
		for i := range out {
			out[i] = i
		}
		return out
	}
	check.Eq(t, SumListIterative(upTo(50_000)), 1249975000)
	check.Eq(t, SumListRecursive(upTo(1_000_000), 0), 499999500000)
	check.Eq(t, SumListRecursive(nil, 0), 0)
}

func TestEx12_4(t *testing.T) {
	orig := []int{1, 2, 3, 4, 5}
	check.Eq(t, RotateCopy(orig, 2), []int{4, 5, 1, 2, 3})
	check.Eq(t, orig, []int{1, 2, 3, 4, 5})
	RotateInPlace(orig, 2)
	check.Eq(t, orig, []int{4, 5, 1, 2, 3})
	three := []int{1, 2, 3}
	RotateInPlace(three, 7)
	check.Eq(t, three, []int{3, 1, 2})
	check.Eq(t, RotateCopy([]int{1, 2, 3}, 0), []int{1, 2, 3})
	check.Eq(t, RotateCopy([]int{}, 3), []int{})
}

func TestEx12_5(t *testing.T) {
	check.True(t, IsPalindromeClean("A man, a plan, a canal: Panama") && !IsPalindromeClean("race a car"), "pal")
}
