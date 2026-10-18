package ch3_2

import (
	"math/rand/v2"
	"slices"
	"testing"

	"dsa/internal/check"
)

func randomSlices(seed int64) [][]int {
	r := rand.New(rand.NewPCG(uint64(seed), 0))
	out := make([][]int, 200)
	for t := range out {
		out[t] = make([]int, r.IntN(30))
		for i := range out[t] {
			out[t][i] = r.IntN(10)
		}
	}
	return out
}

func TestEx3_2_1(t *testing.T) {
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

func TestEx3_2_2(t *testing.T) {
	a := []int{0, 5, 2, 1, 6, 3}
	Quicksort(a)
	check.Eq(t, a, []int{0, 1, 2, 3, 5, 6})
	for _, x := range randomSlices(7) {
		q := slices.Clone(x)
		Quicksort(q)
		check.Eq(t, q, sorted(x))
	}
}

func TestEx3_2_3(t *testing.T) {
	qs := []int{0, 50, 20, 10, 60, 30}
	check.Eq(t, Quickselect(qs, 1), 10)
	check.Eq(t, qs, []int{0, 50, 20, 10, 60, 30})
	r := rand.New(rand.NewPCG(8, 0))
	for _, x := range randomSlices(9) {
		if len(x) > 0 {
			k := r.IntN(len(x))
			check.Eq(t, Quickselect(x, k), sorted(x)[k])
		}
	}
}

func TestEx3_2_4(t *testing.T) {
	data := []int{10, 4, 15, 7}
	check.Eq(t, ClosestGapSorted(data), 3)
	check.Eq(t, data, []int{10, 4, 15, 7})
	check.Eq(t, []int{ClosestGapSorted([]int{1, 100}), ClosestGapSorted([]int{5, 5, 9}), ClosestGapSorted([]int{-3, 8, -1})}, []int{99, 0, 2})
}

func TestEx3_2_5(t *testing.T) {
	check.Eq(t, AnyOverlap([][2]int{{9, 11}, {13, 15}, {10, 12}}), true)
	check.Eq(t, AnyOverlap([][2]int{{9, 10}, {10, 11}}), false)
	check.Eq(t, []bool{AnyOverlap(nil), AnyOverlap([][2]int{{1, 5}}), AnyOverlap([][2]int{{1, 10}, {2, 3}, {11, 12}})}, []bool{false, false, true})
}

func TestEx3_2_6(t *testing.T) {
	check.Eq(t, CountDistinctSorted([]int{5, 9, 3, 5, 9, 9}), 3)
	check.Eq(t, []int{CountDistinctSorted(nil), CountDistinctSorted([]int{7, 7, 7}), CountDistinctSorted([]int{3, -3, 0})}, []int{0, 1, 3})
}

func TestEx3_2_7(t *testing.T) {
	check.Eq(t, []float64{Median([]int{7, 1, 3}), Median([]int{4, 1, 3, 2})}, []float64{3, 2.5})
}

func TestEx3_2_8(t *testing.T) {
	data := []int{3, 1, 3, 2, 3, 1}
	Quicksort3(data)
	check.Eq(t, data, []int{1, 1, 2, 3, 3, 3})
	many := make([]int, 500)
	for i := range many {
		many[i] = (i * 7919) % 5
	}
	want := sorted(many)
	Quicksort3(many)
	check.Eq(t, many, want)
	for _, x := range randomSlices(11) {
		q := slices.Clone(x)
		Quicksort3(q)
		check.Eq(t, q, sorted(x))
	}
}

func randomLongSlices(seed uint64) [][]int {
	rng := rand.New(rand.NewPCG(seed, 0))
	out := make([][]int, 200)
	for t := range out {
		out[t] = make([]int, rng.IntN(60))
		for i := range out[t] {
			out[t][i] = rng.IntN(20)
		}
	}
	return out
}

// sorted returns a sorted copy, never nil, to match the solutions' empty results.
func sorted(x []int) []int {
	s := append([]int{}, x...)
	slices.Sort(s)
	return s
}

func TestEx3_2_9(t *testing.T) {
	check.Eq(t, Merge([]int{1, 4, 9}, []int{2, 3, 10, 11}), []int{1, 2, 3, 4, 9, 10, 11})
	check.Eq(t, Merge(nil, []int{1}), []int{1})
}

func TestEx3_2_10(t *testing.T) {
	check.Eq(t, MergeComparisons([]int{1, 2, 3}, []int{4, 5, 6}), 3)
	check.Eq(t, MergeComparisons([]int{1, 3, 5}, []int{2, 4, 6}), 5)
}

func TestEx3_2_11(t *testing.T) {
	check.Eq(t, Mergesort([]int{38, 27, 43, 3, 9, 82, 10}), []int{3, 9, 10, 27, 38, 43, 82})
	check.Eq(t, Mergesort([]int{}), []int{})
	for _, d := range randomLongSlices(1) {
		check.Eq(t, Mergesort(d), sorted(d))
	}
}

func TestEx3_2_12(t *testing.T) {
	a := []int{5, 2, 4, 6, 1, 3}
	MergesortInPlace(a)
	check.Eq(t, a, []int{1, 2, 3, 4, 5, 6})
	for _, d := range randomLongSlices(2) {
		b := slices.Clone(d)
		MergesortInPlace(b)
		check.Eq(t, b, sorted(d))
	}
}

func TestEx3_2_13(t *testing.T) {
	check.Eq(t, MergesortBottomUp([]int{5, 1, 4, 2, 8, 0, 2}), []int{0, 1, 2, 2, 4, 5, 8})
	for _, d := range randomLongSlices(3) {
		check.Eq(t, MergesortBottomUp(d), sorted(d))
	}
}

func TestEx3_2_14(t *testing.T) {
	check.Eq(t, []int{CountInversions([]int{2, 4, 1, 3, 5}), CountInversions([]int{5, 4, 3, 2, 1})}, []int{3, 10})
	for _, d := range randomLongSlices(4) {
		inv := 0
		for i := range d {
			for j := i + 1; j < len(d); j++ {
				if d[i] > d[j] {
					inv++
				}
			}
		}
		check.Eq(t, CountInversions(d), inv)
	}
}

func TestEx3_2_15(t *testing.T) {
	type rec struct {
		name string
		n    int
	}
	rs := []rec{{"ann", 3}, {"bo", 1}, {"cy", 3}, {"di", 1}}
	check.Eq(t, MergesortBy(rs, func(r rec) int { return r.n }), []rec{{"bo", 1}, {"di", 1}, {"ann", 3}, {"cy", 3}})
}

func TestEx3_2_16(t *testing.T) {
	check.Eq(t, HybridSort([]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, 4), []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	for _, d := range randomLongSlices(5) {
		check.Eq(t, HybridSort(d, 5), sorted(d))
	}
}
