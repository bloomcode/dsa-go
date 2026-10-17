package ch06

import (
	"math/rand"
	"testing"

	"dsa/internal/check"
)

func TestEx6_1(t *testing.T) {
	a := []int{4, 2, 7, 1, 3}
	InsertionSort(a)
	check.Eq(t, a, []int{1, 2, 3, 4, 7})
	var empty []int
	InsertionSort(empty)
	check.Eq(t, len(empty), 0)
}

func TestEx6_2(t *testing.T) {
	up, down := make([]int, 100), make([]int, 100)
	for i := range up {
		up[i], down[i] = i, 100-i
	}
	check.Eq(t, []int{InsertionSortSteps(up), InsertionSortSteps(down)}, []int{99, 9900})
	shuffled := rand.New(rand.NewSource(2)).Perm(100)
	mid := InsertionSortSteps(shuffled)
	check.True(t, mid > 99 && mid < 9900, "random input is between best and worst")
}

func TestEx6_3(t *testing.T) {
	check.Eq(t, IntersectionNested([]int{3, 1, 4, 2}, []int{4, 5, 3, 6}), []int{3, 4})
	check.Eq(t, len(IntersectionNested([]int{1, 2}, []int{3})), 0)
}

func TestEx6_4(t *testing.T) {
	check.Eq(t, InsertSorted([]int{1, 4, 9}, 5), []int{1, 4, 5, 9})
	check.Eq(t, InsertSorted([]int{1, 4, 9}, 0), []int{0, 1, 4, 9})
	check.Eq(t, InsertSorted(nil, 5), []int{5})
}

func TestEx6_5(t *testing.T) {
	check.Eq(t, MergeSorted([]int{1, 4, 7}, []int{2, 3, 9}), []int{1, 2, 3, 4, 7, 9})
	check.Eq(t, MergeSorted(nil, []int{1}), []int{1})
}
