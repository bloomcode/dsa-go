package ch05

import (
	"testing"

	"dsa/internal/check"
)

func TestEx5_1(t *testing.T) {
	a := []int{4, 2, 7, 1, 3}
	SelectionSort(a)
	check.Eq(t, a, []int{1, 2, 3, 4, 7})
	b := []int{3, -1, 3, 0}
	SelectionSort(b)
	check.Eq(t, b, []int{-1, 0, 3, 3})
}

func TestEx5_2(t *testing.T) {
	ten := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	c, s := BubbleCounts(ten)
	check.Eq(t, []int{c, s}, []int{45, 45})
	c, s = SelectionCounts(ten)
	check.Eq(t, []int{c, s}, []int{45, 5})
	check.Eq(t, ten[0], 10) // both sort a copy
}

func TestEx5_3(t *testing.T) {
	check.Eq(t, EveryOther([]int{10, 11, 12, 13, 14}), []int{10, 12, 14})
	check.Eq(t, EveryOther([]int{1, 2}), []int{1})
}

func TestEx5_4(t *testing.T) {
	v, ok := SecondLargest([]int{5, 1, 9, 7, 9})
	check.Eq(t, []any{v, ok}, []any{7, true})
	_, ok = SecondLargest([]int{3, 3})
	check.Eq(t, ok, false)
	v, ok = SecondLargest([]int{-5, -9})
	check.Eq(t, []any{v, ok}, []any{-9, true})
}

func TestEx5_5(t *testing.T) {
	check.Eq(t, KSmallest([]int{8, 3, 5, 1, 9}, 2), []int{1, 3})
}
