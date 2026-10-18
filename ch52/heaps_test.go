package ch52

import (
	"math/rand"
	"slices"
	"testing"

	"dsa/internal/check"
)

func TestEx52_1(t *testing.T) {
	check.Eq(t, []int{LeftChild(0), RightChild(0), Parent(4)}, []int{1, 2, 1})
	var h MaxHeap
	_, ok := h.Root()
	check.Eq(t, []any{ok, h.Len()}, []any{false, 0})
	h.data = []int{9, 5, 3}
	r, _ := h.Root()
	l, _ := h.Last()
	check.Eq(t, []int{r, l, h.Len()}, []int{9, 3, 3})
}

func TestEx52_2(t *testing.T) {
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

func TestEx52_3(t *testing.T) {
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

func TestEx52_4(t *testing.T) {
	check.Eq(t, HeapSort([]int{5, 1, 9, 3, 7}), []int{1, 3, 5, 7, 9})
	check.Eq(t, HeapSort([]int{4, 4, 1, 4}), []int{1, 4, 4, 4})
	rng := rand.New(rand.NewSource(1))
	data := make([]int, 200)
	for i := range data {
		data[i] = rng.Intn(201) - 100
	}
	check.Eq(t, HeapSort(slices.Clone(data)), slices.Sorted(slices.Values(data)))
}

func TestEx52_5(t *testing.T) {
	var q DeadlineQueue
	q.Add("report", 5)
	q.Add("backup", 2)
	q.Add("email", 2)
	a, _ := q.NextJob()
	b, _ := q.NextJob()
	c, _ := q.NextJob()
	_, ok := q.NextJob()
	check.Eq(t, []any{a, b, c, ok}, []any{"backup", "email", "report", false})
	q.Add("late", 9)
	q.Add("soon", 1)
	d, _ := q.NextJob()
	check.Eq(t, d, "soon")
}

func TestEx52_6(t *testing.T) {
	check.Eq(t, TopK([]int{3, 1, 9, 7, 5, 8}, 3), []int{9, 8, 7})
}

func TestEx52_7(t *testing.T) {
	check.True(t, IsMaxHeap([]int{100, 88, 25, 87, 16}) && !IsMaxHeap([]int{10, 20}) && IsMaxHeap(nil), "heap check")
}

func TestEx52_8(t *testing.T) {
	check.Eq(t, MergeKSorted([][]int{{1, 4, 7}, {2, 5}, {0, 9}}), []int{0, 1, 2, 4, 5, 7, 9})
	check.Eq(t, MergeKSorted([][]int{{}, {3}}), []int{3})
}

func TestEx52_9(t *testing.T) {
	var m RunningMedian
	got := []float64{}
	for _, x := range []int{5, 15, 1, 3} {
		m.Add(x)
		got = append(got, m.Median())
	}
	check.Eq(t, got, []float64{5, 10, 5, 4})
	stream := []int{8, -2, 7, 7, 0, 11, 3, 3, -5, 20}
	var m2 RunningMedian
	for n := 1; n <= len(stream); n++ {
		m2.Add(stream[n-1])
		s := slices.Clone(stream[:n])
		slices.Sort(s)
		want := float64(s[n/2])
		if n%2 == 0 {
			want = float64(s[n/2-1]+s[n/2]) / 2
		}
		check.Eq(t, m2.Median(), want)
	}
}

func TestEx52_10(t *testing.T) {
	check.Eq(t, SortNearlySorted([]int{3, 1, 2, 6, 4, 5, 9, 7, 8}, 2), []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	check.Eq(t, SortNearlySorted([]int{}, 3), []int{})
	check.Eq(t, SortNearlySorted([]int{2, 1, 4, 3}, 1), []int{1, 2, 3, 4})
	check.Eq(t, SortNearlySorted([]int{1, 2, 3}, 0), []int{1, 2, 3})
	check.Eq(t, SortNearlySorted([]int{5, 4, 3, 2, 1}, 4), []int{1, 2, 3, 4, 5})
	check.Eq(t, SortNearlySorted([]int{2, 2, 1}, 2), []int{1, 2, 2})
}
