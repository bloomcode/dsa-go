package ch6_3

import (
	"testing"

	"dsa/internal/check"
)

func TestEx6_3_1(t *testing.T) {
	ds := NewDisjointSet(5)
	check.True(t, ds.Union(0, 1) && !ds.Union(1, 0), "first union merges, second does not")
	check.Eq(t, []int{ds.Size(1), ds.Groups}, []int{2, 4})
	ds = NewDisjointSet(6)
	check.True(t, ds.Union(0, 1) && ds.Union(2, 1) && !ds.Union(0, 2), "unions")
	check.Eq(t, []int{ds.Find(2), ds.Size(2)}, []int{ds.Find(0), 3})
	root := ds.Find(0)
	check.True(t, ds.Union(3, 0), "merge a single into the big group")
	check.Eq(t, []int{ds.Find(3), ds.Size(3), ds.Groups}, []int{root, 4, 3})
	chain := NewDisjointSet(4)
	chain.Parent = []int{0, 0, 1, 2}
	check.Eq(t, chain.Find(3), 0)
	check.Eq(t, chain.Parent, []int{0, 0, 0, 0})
}

func TestEx6_3_2(t *testing.T) {
	got := []int{GroupCount(5, [][2]int{{0, 1}, {1, 2}, {3, 4}}), GroupCount(3, nil), GroupCount(4, [][2]int{{0, 1}, {1, 0}, {2, 3}, {3, 1}})}
	check.Eq(t, got, []int{2, 3, 1})
}

func TestEx6_3_3(t *testing.T) {
	got := []bool{
		HasCycleUndirected(4, [][2]int{{0, 1}, {1, 2}, {2, 0}, {2, 3}}), HasCycleUndirected(4, [][2]int{{0, 1}, {1, 2}, {2, 3}}),
		HasCycleUndirected(3, nil), HasCycleUndirected(5, [][2]int{{3, 4}, {0, 1}, {1, 3}, {4, 0}}),
	}
	check.Eq(t, got, []bool{true, false, false, true})
}

func TestEx6_3_4(t *testing.T) {
	got := []bool{
		IsTree(4, [][2]int{{0, 1}, {0, 2}, {2, 3}}), IsTree(4, [][2]int{{0, 1}, {2, 3}}), IsTree(1, nil),
		IsTree(4, [][2]int{{0, 1}, {1, 2}, {2, 0}}), IsTree(3, [][2]int{{0, 1}, {1, 2}, {2, 0}}),
	}
	check.Eq(t, got, []bool{true, false, true, false, false})
}

func TestEx6_3_5(t *testing.T) {
	got := []int{
		KruskalCost(4, [][3]int{{0, 1, 4}, {1, 2, 2}, {0, 2, 5}, {2, 3, 1}, {1, 3, 7}}), KruskalCost(1, nil),
		KruskalCost(3, [][3]int{{0, 1, 3}}), KruskalCost(3, [][3]int{{0, 1, -2}, {1, 2, 6}, {0, 2, 6}}),
	}
	check.Eq(t, got, []int{7, 0, -1, 4})
}

func TestEx6_3_6(t *testing.T) {
	got := []int{
		PrimCost(4, [][3]int{{0, 1, 4}, {1, 2, 2}, {0, 2, 5}, {2, 3, 1}, {1, 3, 7}}), PrimCost(1, nil), PrimCost(3, [][3]int{{0, 1, 3}}),
		PrimCost(3, [][3]int{{0, 1, -2}, {1, 2, 6}, {0, 2, 6}}), PrimCost(5, [][3]int{{0, 1, 9}, {0, 2, 3}, {1, 2, 1}, {1, 3, 8}, {2, 3, 7}, {3, 4, 2}, {2, 4, 6}}),
	}
	check.Eq(t, got, []int{7, 0, -1, 4, 12})
}

func TestEx6_3_7(t *testing.T) {
	check.Eq(t, RedundantEdge([][2]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {3, 4}}), [2]int{1, 3})
	check.Eq(t, RedundantEdge([][2]int{{0, 1}, {1, 2}, {0, 2}}), [2]int{0, 2})
	check.Eq(t, RedundantEdge([][2]int{{2, 3}, {0, 1}, {1, 3}, {0, 2}}), [2]int{0, 2})
}

func TestEx6_3_8(t *testing.T) {
	got := []int{
		EarliestConnected(4, [][3]int{{7, 2, 3}, {2, 0, 1}, {5, 1, 2}, {9, 0, 3}}), EarliestConnected(3, [][3]int{{1, 0, 1}}),
		EarliestConnected(2, [][3]int{{4, 1, 0}, {3, 0, 1}}), EarliestConnected(3, [][3]int{{5, 0, 1}, {6, 1, 0}, {8, 2, 1}}),
	}
	check.Eq(t, got, []int{7, -1, 3, 8})
}
