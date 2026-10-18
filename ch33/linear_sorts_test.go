package ch33

import (
	"testing"

	"dsa/internal/check"
)

func TestEx33_1(t *testing.T) {
	check.Eq(t, CountingSort([]int{3, 0, 2, 3, 1}, 3), []int{0, 1, 2, 3, 3})
	check.Eq(t, CountingSort([]int{}, 5), []int{})
	check.Eq(t, CountingSort([]int{4, 4, 4}, 4), []int{4, 4, 4})
	check.Eq(t, CountingSort([]int{9, 0, 5, 0}, 9), []int{0, 0, 5, 9})
}

func TestEx33_2(t *testing.T) {
	check.Eq(t, SortByOrder([]int{2, 0, 1, 2, 0}, []int{2, 0, 1}), []int{2, 2, 0, 0, 1})
	check.Eq(t, SortByOrder([]int{}, []int{1, 2}), []int{})
	check.Eq(t, SortByOrder([]int{5, 5, 7}, []int{7, 5, 9}), []int{7, 5, 5})
	check.Eq(t, SortByOrder([]int{3, 1, 3}, []int{1, 3}), []int{1, 3, 3})
}

func TestEx33_3(t *testing.T) {
	check.Eq(t, RadixSort([]int{329, 57, 640, 4, 1000, 72}), []int{4, 57, 72, 329, 640, 1000})
	check.Eq(t, RadixSort([]int{}), []int{})
	check.Eq(t, RadixSort([]int{0, 0, 1}), []int{0, 0, 1})
	check.Eq(t, RadixSort([]int{21, 12, 11, 22}), []int{11, 12, 21, 22})
}

func TestEx33_4(t *testing.T) {
	check.Eq(t, BucketSort([]float64{0.42, 0.07, 0.93, 0.5, 0.31}), []float64{0.07, 0.31, 0.42, 0.5, 0.93})
	check.Eq(t, BucketSort([]float64{}), []float64{})
	check.Eq(t, BucketSort([]float64{0.25, 0.2, 0.21, 0.0}), []float64{0.0, 0.2, 0.21, 0.25})
	check.Eq(t, BucketSort([]float64{0.99}), []float64{0.99})
}
