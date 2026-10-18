package ch22

import (
	"math"
	"testing"

	"dsa/internal/check"
)

// ordered returns an OrderedArray holding sorted, filled directly so a test does not depend on Insert.
func ordered(sorted ...int) *OrderedArray { return &OrderedArray{Items: append([]int{}, sorted...)} }

func checkBinarySearch(t *testing.T, fn func([]int, int) int) {
	t.Helper()
	check.Eq(t, fn([]int{3, 17, 75, 80, 202}, 80), 3)
	check.Eq(t, []int{fn([]int{1, 3, 5, 7, 9}, 1), fn([]int{1, 3, 5, 7, 9}, 9)}, []int{0, 4})
	check.Eq(t, []int{fn([]int{1, 3, 5}, 4), fn(nil, 1)}, []int{-1, -1})
}

func TestEx22_1(t *testing.T) {
	o := ordered(3, 17, 80, 202)
	o.Insert(75)
	check.Eq(t, o.Items, []int{3, 17, 75, 80, 202})
	o.Insert(1)
	o.Insert(500)
	check.Eq(t, o.Items, []int{1, 3, 17, 75, 80, 202, 500})
	check.Eq(t, NewOrderedArray(8, 4, 6, 2).Items, []int{2, 4, 6, 8})
}

func TestEx22_2(t *testing.T) {
	check.Eq(t, LinearSearchOrdered([]int{3, 17, 75, 80, 202}, 22), -1)
	check.Eq(t, LinearSearchOrdered([]int{3, 17, 75, 80, 202}, 75), 2)
	check.Eq(t, LinearSearchOrdered(nil, 1), -1)
}

func TestEx22_3(t *testing.T) { checkBinarySearch(t, BinarySearch) }

func TestEx22_4(t *testing.T) { checkBinarySearch(t, BinarySearchRecursive) }

func TestEx22_5(t *testing.T) {
	lb := []int{10, 20, 20, 30}
	check.Eq(t, []int{LowerBound(lb, 20), LowerBound(lb, 25), LowerBound(lb, 99), LowerBound(lb, 5)}, []int{1, 3, 4, 0})
	check.Eq(t, LowerBound(nil, 5), 0)
}

func TestEx22_6(t *testing.T) {
	o := ordered(3, 17, 80, 202)
	o.InsertFast(75)
	o.InsertFast(1)
	o.InsertFast(500)
	check.Eq(t, o.Items, []int{1, 3, 17, 75, 80, 202, 500})
	f := ordered()
	for _, v := range []int{5, 1, 4, 1, 9} {
		f.InsertFast(v)
	}
	check.Eq(t, f.Items, []int{1, 1, 4, 5, 9})
}

func TestEx22_7(t *testing.T) {
	f, l := FirstAndLast([]int{1, 2, 2, 2, 5}, 2)
	check.Eq(t, []int{f, l}, []int{1, 3})
	f, l = FirstAndLast([]int{1, 2, 2, 2, 5}, 3)
	check.Eq(t, []int{f, l}, []int{-1, -1})
	f, l = FirstAndLast([]int{7, 7, 7}, 7)
	check.Eq(t, []int{f, l}, []int{0, 2})
}

func TestEx22_8(t *testing.T) {
	o := ordered(2, 4, 6, 8)
	check.True(t, o.Delete(4), "delete")
	check.Eq(t, o.Items, []int{2, 6, 8})
	check.True(t, !o.Delete(5), "absent")
	check.True(t, o.Delete(8), "delete last")
	check.Eq(t, o.Items, []int{2, 6})
}

func TestEx22_9(t *testing.T) {
	a, b := SearchSteps(100)
	check.Eq(t, []int{a, b}, []int{100, 7})
	a, b = SearchSteps(10_000)
	check.Eq(t, []int{a, b}, []int{10000, 14})
}

func TestEx22_10(t *testing.T) {
	check.Eq(t, []int{IntSqrt(27), IntSqrt(36), IntSqrt(0), IntSqrt(1), IntSqrt(math.MaxInt64)}, []int{5, 6, 0, 1, 3037000499})
}

func TestEx22_11(t *testing.T) {
	nums := []int{15, 18, 22, 3, 7, 9}
	for i, x := range nums {
		check.Eq(t, SearchRotated(nums, x), i)
	}
	check.Eq(t, []int{SearchRotated(nums, 10), SearchRotated(nums, 1), SearchRotated(nil, 4)}, []int{-1, -1, -1})
	check.Eq(t, []int{SearchRotated([]int{5}, 5), SearchRotated([]int{1, 2, 3}, 3), SearchRotated([]int{3, 1}, 1)}, []int{0, 2, 1})
}

func TestEx22_12(t *testing.T) {
	check.Eq(t, FindPeak([]int{1, 3, 2, 5, 4}), 3)
	for _, nums := range [][]int{{7}, {1, 2, 3, 4}, {4, 3, 2, 1}, {2, 9, 1, 8, 0, 7, 6}, {5, 1}} {
		i := FindPeak(nums)
		check.True(t, (i == 0 || nums[i-1] < nums[i]) && (i == len(nums)-1 || nums[i+1] < nums[i]), "index is a peak")
	}
}

func TestEx22_13(t *testing.T) {
	loads := []int{4, 9, 3, 6, 2, 8}
	got := []int{SmallestMaxLoad(loads, 3), SmallestMaxLoad(loads, 1), SmallestMaxLoad(loads, 6), SmallestMaxLoad([]int{5, 5, 5, 5}, 2), SmallestMaxLoad([]int{7}, 3)}
	check.Eq(t, got, []int{13, 32, 9, 10, 7})
}

func TestEx22_14(t *testing.T) {
	got := []int{FixedPoint([]int{-4, -1, 2, 5, 9}), FixedPoint([]int{1, 2, 3}), FixedPoint([]int{0, 1, 2}), FixedPoint([]int{-3, -1, 1, 3, 4}), FixedPoint(nil), FixedPoint([]int{-5, 0, 7})}
	check.Eq(t, got, []int{2, -1, 0, 3, -1, -1})
}

func TestEx22_15(t *testing.T) {
	got := []float64{
		MedianOfTwo([]int{1, 4, 9}, []int{2, 3}), MedianOfTwo([]int{1, 4}, []int{2, 3}), MedianOfTwo(nil, []int{5}),
		MedianOfTwo([]int{1, 2}, []int{3, 4, 5, 6}), MedianOfTwo([]int{7}, []int{1, 2, 3}), MedianOfTwo([]int{10, 20, 30}, nil),
		MedianOfTwo([]int{1, 1}, []int{1, 1}),
	}
	check.Eq(t, got, []float64{3, 2.5, 5, 3.5, 2.5, 20, 1})
}

func TestEx22_16(t *testing.T) {
	p := NewWeightedPicker([]int{3, 1, 2})
	check.Eq(t, p.Total(), 6)
	picks := []int{}
	for r := range 6 {
		picks = append(picks, p.Pick(r))
	}
	check.Eq(t, picks, []int{0, 0, 0, 1, 2, 2})
	weights := []int{5, 1, 4, 2}
	p = NewWeightedPicker(weights)
	counts := make([]int, len(weights))
	for r := range p.Total() {
		counts[p.Pick(r)]++
	}
	check.Eq(t, counts, weights)
	check.Eq(t, NewWeightedPicker([]int{9}).Pick(8), 0)
}
