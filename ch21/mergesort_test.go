package ch21

import (
	"math/rand/v2"
	"slices"
	"testing"

	"dsa/internal/check"
)

func randomSlices(seed uint64) [][]int {
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

func TestEx21_1(t *testing.T) {
	check.Eq(t, Merge([]int{1, 4, 9}, []int{2, 3, 10, 11}), []int{1, 2, 3, 4, 9, 10, 11})
	check.Eq(t, Merge(nil, []int{1}), []int{1})
}

func TestEx21_2(t *testing.T) {
	check.Eq(t, MergeComparisons([]int{1, 2, 3}, []int{4, 5, 6}), 3)
	check.Eq(t, MergeComparisons([]int{1, 3, 5}, []int{2, 4, 6}), 5)
}

func TestEx21_3(t *testing.T) {
	check.Eq(t, Mergesort([]int{38, 27, 43, 3, 9, 82, 10}), []int{3, 9, 10, 27, 38, 43, 82})
	check.Eq(t, Mergesort([]int{}), []int{})
	for _, d := range randomSlices(1) {
		check.Eq(t, Mergesort(d), sorted(d))
	}
}

func TestEx21_4(t *testing.T) {
	a := []int{5, 2, 4, 6, 1, 3}
	MergesortInPlace(a)
	check.Eq(t, a, []int{1, 2, 3, 4, 5, 6})
	for _, d := range randomSlices(2) {
		b := slices.Clone(d)
		MergesortInPlace(b)
		check.Eq(t, b, sorted(d))
	}
}

func TestEx21_5(t *testing.T) {
	check.Eq(t, MergesortBottomUp([]int{5, 1, 4, 2, 8, 0, 2}), []int{0, 1, 2, 2, 4, 5, 8})
	for _, d := range randomSlices(3) {
		check.Eq(t, MergesortBottomUp(d), sorted(d))
	}
}

func TestEx21_6(t *testing.T) {
	check.Eq(t, []int{CountInversions([]int{2, 4, 1, 3, 5}), CountInversions([]int{5, 4, 3, 2, 1})}, []int{3, 10})
	for _, d := range randomSlices(4) {
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

func TestEx21_7(t *testing.T) {
	type rec struct {
		name string
		n    int
	}
	rs := []rec{{"ann", 3}, {"bo", 1}, {"cy", 3}, {"di", 1}}
	check.Eq(t, MergesortBy(rs, func(r rec) int { return r.n }), []rec{{"bo", 1}, {"di", 1}, {"ann", 3}, {"cy", 3}})
}

func TestEx21_8(t *testing.T) {
	check.Eq(t, HybridSort([]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, 4), []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	for _, d := range randomSlices(5) {
		check.Eq(t, HybridSort(d, 5), sorted(d))
	}
}
