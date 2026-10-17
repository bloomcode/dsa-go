package ch31

import (
	"math/rand"
	"slices"
	"testing"

	"dsa/internal/check"
)

func randomData(seed int64) []int {
	rng := rand.New(rand.NewSource(seed))
	a := make([]int, 60)
	for i := range a {
		a[i] = rng.Intn(101) - 50
	}
	return a
}

func TestEx31_1(t *testing.T) {
	a := []int{4, 2, 7, 1, 3}
	BubbleSort(a)
	check.Eq(t, a, []int{1, 2, 3, 4, 7})
	data := randomData(1)
	want := slices.Sorted(slices.Values(data))
	BubbleSort(data)
	check.Eq(t, data, want)
}

func TestEx31_2(t *testing.T) {
	check.Eq(t, BubbleSortOptimized([]int{1, 2, 3, 4}), 1)
	r := []int{4, 3, 2, 1}
	check.Eq(t, BubbleSortOptimized(r), 3)
	check.Eq(t, r, []int{1, 2, 3, 4})
	data := randomData(2)
	want := slices.Sorted(slices.Values(data))
	BubbleSortOptimized(data)
	check.Eq(t, data, want)
}

func TestEx31_3(t *testing.T) {
	check.Eq(t, EqualPairsQuadratic([]int{1, 2, 1, 1, 2}), 4)
	check.Eq(t, []int{EqualPairsQuadratic([]int{1, 2, 3}), EqualPairsQuadratic(nil), EqualPairsQuadratic([]int{5, 5, 5, 5, 5})}, []int{0, 0, 10})
}

func TestEx31_4(t *testing.T) {
	check.Eq(t, EqualPairsLinear([]int{1, 2, 1, 1, 2}), 4)
	check.Eq(t, []int{EqualPairsLinear([]int{1, 2, 3}), EqualPairsLinear(nil)}, []int{0, 0})
	same := make([]int, 2000)
	check.Eq(t, EqualPairsLinear(same), 1999000)
}

func TestEx31_5(t *testing.T) {
	for _, fn := range []func([]int) int{LargestPairSumQuadratic, LargestPairSumLinear} {
		check.Eq(t, []int{fn([]int{3, 9, 1, 7}), fn([]int{5, 5}), fn([]int{-4, -1, -9}), fn([]int{9, 1, 9})}, []int{16, 10, -5, 18})
	}
}

func TestEx31_6(t *testing.T) {
	check.True(t, IsSorted([]int{1, 2, 2, 5}) && !IsSorted([]int{1, 3, 2}) && IsSorted(nil), "sorted")
}

func TestEx31_7(t *testing.T) {
	a := []int{4, 2, 7, 1, 3}
	SelectionSort(a)
	check.Eq(t, a, []int{1, 2, 3, 4, 7})
	b := []int{3, -1, 3, 0}
	SelectionSort(b)
	check.Eq(t, b, []int{-1, 0, 3, 3})
}

func TestEx31_8(t *testing.T) {
	ten := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	c, s := BubbleCounts(ten)
	check.Eq(t, []int{c, s}, []int{45, 45})
	c, s = SelectionCounts(ten)
	check.Eq(t, []int{c, s}, []int{45, 5})
	check.Eq(t, ten[0], 10) // both sort a copy
}

func TestEx31_9(t *testing.T) {
	check.Eq(t, NeighbourGaps([]int{3, 7, 4, 10}), []int{4, -3, 6})
	check.Eq(t, NeighbourGaps([]int{5}), []int{})
	check.Eq(t, NeighbourGaps(nil), []int{})
	check.Eq(t, NeighbourGaps([]int{1, 1}), []int{0})
}

func TestEx31_10(t *testing.T) {
	v, ok := SecondLargest([]int{5, 1, 9, 7, 9})
	check.Eq(t, []any{v, ok}, []any{7, true})
	_, ok = SecondLargest([]int{3, 3})
	check.Eq(t, ok, false)
	v, ok = SecondLargest([]int{-5, -9})
	check.Eq(t, []any{v, ok}, []any{-9, true})
}

func TestEx31_11(t *testing.T) {
	check.Eq(t, KSmallest([]int{8, 3, 5, 1, 9}, 2), []int{1, 3})
}

func TestEx31_12(t *testing.T) {
	a := []int{4, 2, 7, 1, 3}
	InsertionSort(a)
	check.Eq(t, a, []int{1, 2, 3, 4, 7})
	var empty []int
	InsertionSort(empty)
	check.Eq(t, len(empty), 0)
}

func TestEx31_13(t *testing.T) {
	up, down := make([]int, 100), make([]int, 100)
	for i := range up {
		up[i], down[i] = i, 100-i
	}
	check.Eq(t, []int{InsertionSortSteps(up), InsertionSortSteps(down)}, []int{99, 9900})
	shuffled := rand.New(rand.NewSource(2)).Perm(100)
	mid := InsertionSortSteps(shuffled)
	check.True(t, mid > 99 && mid < 9900, "random input is between best and worst")
}

func TestEx31_14(t *testing.T) {
	check.Eq(t, []bool{ContainsAll([]int{4, 8, 15, 16}, []int{8, 16}), ContainsAll([]int{4, 8}, []int{8, 5})}, []bool{true, false})
	check.Eq(t, []bool{ContainsAll([]int{1}, nil), ContainsAll(nil, []int{1}), ContainsAll([]int{3, 3}, []int{3, 3})}, []bool{true, false, true})
}

func TestEx31_15(t *testing.T) {
	check.Eq(t, InsertSorted([]int{1, 4, 9}, 5), []int{1, 4, 5, 9})
	check.Eq(t, InsertSorted([]int{1, 4, 9}, 0), []int{0, 1, 4, 9})
	check.Eq(t, InsertSorted(nil, 5), []int{5})
}

func TestEx31_16(t *testing.T) {
	data := []int{2, 1, 3, 5, 4}
	SortNearlySorted(data, 1)
	check.Eq(t, data, []int{1, 2, 3, 4, 5})
	more := []int{3, 1, 2, 6, 4, 5}
	SortNearlySorted(more, 2)
	check.Eq(t, more, []int{1, 2, 3, 4, 5, 6})
}
