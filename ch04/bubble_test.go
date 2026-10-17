package ch04

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

func TestEx4_1(t *testing.T) {
	a := []int{4, 2, 7, 1, 3}
	BubbleSort(a)
	check.Eq(t, a, []int{1, 2, 3, 4, 7})
	data := randomData(1)
	want := slices.Sorted(slices.Values(data))
	BubbleSort(data)
	check.Eq(t, data, want)
}

func TestEx4_2(t *testing.T) {
	check.Eq(t, BubbleSortOptimized([]int{1, 2, 3, 4}), 1)
	r := []int{4, 3, 2, 1}
	check.Eq(t, BubbleSortOptimized(r), 3)
	check.Eq(t, r, []int{1, 2, 3, 4})
	data := randomData(2)
	want := slices.Sorted(slices.Values(data))
	BubbleSortOptimized(data)
	check.Eq(t, data, want)
}

func TestEx4_3(t *testing.T) {
	check.True(t, HasDuplicateQuadratic([]int{1, 5, 3, 9, 1, 4}) && !HasDuplicateQuadratic([]int{1, 2, 3}), "quad")
}

func TestEx4_4(t *testing.T) {
	check.True(t, HasDuplicateLinear([]int{1, 5, 3, 9, 1, 4}) && !HasDuplicateLinear([]int{1, 2, 3}), "lin")
}

func TestEx4_5(t *testing.T) {
	g, _ := GreatestQuadratic([]int{3, 9, 2, 7})
	h, _ := GreatestLinear([]int{3, 9, 2, 7})
	check.Eq(t, []int{g, h}, []int{9, 9})
	_, ok := GreatestLinear(nil)
	check.Eq(t, ok, false)
	_, ok = GreatestQuadratic(nil)
	check.Eq(t, ok, false)
}

func TestEx4_6(t *testing.T) {
	check.True(t, IsSorted([]int{1, 2, 2, 5}) && !IsSorted([]int{1, 3, 2}) && IsSorted(nil), "sorted")
}
