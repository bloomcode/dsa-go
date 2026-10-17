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
