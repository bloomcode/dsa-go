package ch13

import (
	"math/rand"
	"slices"
	"testing"

	"dsa/internal/check"
)

func randomSlices(seed int64) [][]int {
	r := rand.New(rand.NewSource(seed))
	out := make([][]int, 200)
	for t := range out {
		out[t] = make([]int, r.Intn(30))
		for i := range out[t] {
			out[t][i] = r.Intn(10)
		}
	}
	return out
}

func sorted(x []int) []int {
	s := slices.Clone(x)
	slices.Sort(s)
	return s
}

func TestEx13_1(t *testing.T) {
	p := []int{0, 5, 2, 1, 6, 3}
	pi := Partition(p, 0, 5)
	check.Eq(t, []int{pi, p[3]}, []int{3, 3})
	for i := 0; i < 3; i++ {
		check.True(t, p[i] < 3, "left")
	}
	for i := 4; i < 6; i++ {
		check.True(t, p[i] > 3, "right")
	}
}

func TestEx13_2(t *testing.T) {
	a := []int{0, 5, 2, 1, 6, 3}
	Quicksort(a)
	check.Eq(t, a, []int{0, 1, 2, 3, 5, 6})
	for _, x := range randomSlices(7) {
		q := slices.Clone(x)
		Quicksort(q)
		check.Eq(t, q, sorted(x))
	}
}

func TestEx13_3(t *testing.T) {
	qs := []int{0, 50, 20, 10, 60, 30}
	check.Eq(t, Quickselect(qs, 1), 10)
	check.Eq(t, qs, []int{0, 50, 20, 10, 60, 30})
	r := rand.New(rand.NewSource(8))
	for _, x := range randomSlices(9) {
		if len(x) > 0 {
			k := r.Intn(len(x))
			check.Eq(t, Quickselect(x, k), sorted(x)[k])
		}
	}
}

func TestEx13_4(t *testing.T) {
	check.True(t, HasDuplicatesSorted([]int{5, 9, 3, 2, 4, 5, 6}) && !HasDuplicatesSorted([]int{5, 9, 3}), "dups")
}

func TestEx13_5(t *testing.T) {
	check.Eq(t, []int{MaxProductOfThree([]int{1, 10, 2, 6, 5, 3}), MaxProductOfThree([]int{-10, -10, 1, 3, 2})}, []int{300, 300})
}

func TestEx13_6(t *testing.T) {
	check.Eq(t, []int{MissingNumberSorted([]int{5, 2, 4, 1, 0}), MissingNumberSorted([]int{0, 1, 2})}, []int{3, 3})
}

func TestEx13_7(t *testing.T) {
	check.Eq(t, []float64{Median([]int{7, 1, 3}), Median([]int{4, 1, 3, 2})}, []float64{3, 2.5})
}

func TestEx13_8(t *testing.T) {
	check.Eq(t, MergeSort([]int{38, 27, 43, 3, 9, 82, 10}), []int{3, 9, 10, 27, 38, 43, 82})
	for _, x := range randomSlices(10) {
		check.Eq(t, MergeSort(x), sorted(x))
	}
}
