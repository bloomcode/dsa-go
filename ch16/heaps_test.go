package ch16

import (
	"math/rand"
	"slices"
	"testing"

	"dsa/internal/check"
)

func TestEx16_1(t *testing.T) {
	check.Eq(t, []int{LeftChild(0), RightChild(0), Parent(4)}, []int{1, 2, 1})
	var h MaxHeap
	_, ok := h.Root()
	check.Eq(t, []any{ok, h.Len()}, []any{false, 0})
	h.data = []int{9, 5, 3}
	r, _ := h.Root()
	l, _ := h.Last()
	check.Eq(t, []int{r, l, h.Len()}, []int{9, 3, 3})
}

func TestEx16_2(t *testing.T) {
	var h MaxHeap
	for _, v := range []int{3, 9, 5, 12} {
		h.Insert(v)
	}
	r, _ := h.Root()
	check.Eq(t, []int{r, h.Len()}, []int{12, 4})
	for i := 1; i < len(h.data); i++ {
		check.True(t, h.data[Parent(i)] >= h.data[i], "heap order")
	}
}

func TestEx16_3(t *testing.T) {
	var h MaxHeap
	for _, v := range []int{3, 9, 5, 12} {
		h.Insert(v)
	}
	var got []int
	for h.Len() > 0 {
		v, _ := h.Delete()
		got = append(got, v)
	}
	check.Eq(t, got, []int{12, 9, 5, 3})
	_, ok := h.Delete()
	check.Eq(t, ok, false)
}

func TestEx16_4(t *testing.T) {
	check.Eq(t, HeapSort([]int{5, 1, 9, 3, 7}), []int{1, 3, 5, 7, 9})
	check.Eq(t, HeapSort([]int{4, 4, 1, 4}), []int{1, 4, 4, 4})
	rng := rand.New(rand.NewSource(1))
	data := make([]int, 200)
	for i := range data {
		data[i] = rng.Intn(201) - 100
	}
	check.Eq(t, HeapSort(slices.Clone(data)), slices.Sorted(slices.Values(data)))
}

func TestEx16_5(t *testing.T) {
	var q TriageQueue
	q.Admit("Ann", 3)
	q.Admit("Bo", 9)
	q.Admit("Cy", 9)
	a, _ := q.NextPatient()
	b, _ := q.NextPatient()
	c, _ := q.NextPatient()
	_, ok := q.NextPatient()
	check.Eq(t, []any{a, b, c, ok}, []any{"Bo", "Cy", "Ann", false})
}

func TestEx16_6(t *testing.T) {
	check.Eq(t, TopK([]int{3, 1, 9, 7, 5, 8}, 3), []int{9, 8, 7})
}

func TestEx16_7(t *testing.T) {
	check.True(t, IsMaxHeap([]int{100, 88, 25, 87, 16}) && !IsMaxHeap([]int{10, 20}) && IsMaxHeap(nil), "heap check")
}

func TestEx16_8(t *testing.T) {
	check.Eq(t, MergeKSorted([][]int{{1, 4, 7}, {2, 5}, {0, 9}}), []int{0, 1, 2, 4, 5, 7, 9})
	check.Eq(t, MergeKSorted([][]int{{}, {3}}), []int{3})
}
