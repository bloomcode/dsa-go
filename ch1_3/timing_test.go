package ch1_3

import (
	"math"
	"math/rand/v2"
	"slices"
	"testing"

	"dsa/internal/check"
)

func descending(n int) []int {
	d := make([]int, n)
	for i := range d {
		d[i] = n - i
	}
	return d
}

func TestEx1_3_1(t *testing.T) {
	v, d := TimeOnce(func() int {
		s := 0
		for i := range 1_000_000 {
			s += i
		}
		return s
	})
	check.Eq(t, v, 499999500000)
	check.True(t, d >= 0, "duration non-negative")
}

func TestEx1_3_2(t *testing.T) {
	data := descending(1000)
	check.True(t, Bench(func() { slices.Sort(slices.Clone(data)) }, 20, 3) > 0, "bench positive")
}

func TestEx1_3_3(t *testing.T) {
	data := descending(1000)
	check.True(t, BenchSort(slices.Sort[[]int], data, 3) > 0, "benchsort positive")
	check.Eq(t, data[0], 1000) // BenchSort must not sort the caller's slice
}

func TestEx1_3_4(t *testing.T) {
	check.True(t, math.Abs(GrowthExponent(map[int]float64{1000: 0.01, 2000: 0.04, 4000: 0.16})-2) < 1e-9, "slope 2")
	check.True(t, math.Abs(GrowthExponent(map[int]float64{10: 1, 100: 10, 1000: 100})-1) < 1e-9, "slope 1")
}

func TestEx1_3_5(t *testing.T) {
	data := descending(1000)
	names, err := Race(map[string]func([]int) []int{
		"builtin": func(a []int) []int { slices.Sort(a); return a },
		"copy":    func(a []int) []int { slices.Sort(a); return a },
	}, data)
	check.Eq(t, []any{err, len(names)}, []any{nil, 2})
	_, err = Race(map[string]func([]int) []int{"broken": func(a []int) []int { return a }}, data)
	check.True(t, err != nil, "broken sort reported")
}

func TestEx1_3_6(t *testing.T) {
	ordered := make([]int, 100)
	for i := range ordered {
		ordered[i] = i
	}
	check.Eq(t, QuicksortComparisons(ordered), 4950)
	shuffled := slices.Clone(ordered)
	rand.New(rand.NewPCG(1, 0)).Shuffle(len(shuffled), func(i, j int) { shuffled[i], shuffled[j] = shuffled[j], shuffled[i] })
	check.True(t, QuicksortComparisons(shuffled) < 1500, "shuffled needs far fewer")
}

func TestEx1_3_7(t *testing.T) {
	check.Eq(t, []int{CountRuns([]int{1, 2, 3, 2, 5, 6, 1}), CountRuns([]int{0, 1, 2, 3}), CountRuns(nil)}, []int{3, 1, 0})
}
