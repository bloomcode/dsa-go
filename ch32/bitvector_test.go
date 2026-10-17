package ch32

import (
	"testing"

	"dsa/internal/check"
)

func bv(size int, xs ...int) *BitVector {
	b := NewBitVector(size)
	for _, x := range xs {
		b.Add(x)
	}
	return b
}

func TestEx32_1(t *testing.T) {
	s := NewBoolArraySet(10)
	s.Add(3)
	s.Add(7)
	s.Add(7)
	s.Remove(3)
	check.Eq(t, []any{s.Contains(7), s.Contains(3), s.Len()}, []any{true, false, 1})
}

func TestEx32_2(t *testing.T) {
	b := bv(100, 42)
	check.Eq(t, []any{b.Contains(42), b.Contains(43), b.Words[0]}, []any{true, false, uint64(4398046511104)})
	b.Add(99)
	b.Remove(42)
	check.Eq(t, []bool{b.Contains(42), b.Contains(99)}, []bool{false, true})
}

func TestEx32_3(t *testing.T) {
	b := bv(100, 1, 5, 99)
	check.Eq(t, b.Count(), 3)
	b.Add(5)
	check.Eq(t, []int{b.Count(), NewBitVector(10).Count()}, []int{3, 0})
}

func TestEx32_4(t *testing.T) {
	x, y := bv(10, 1, 2, 3), bv(10, 2, 3, 4)
	u, i, d := x.Union(y), x.Intersection(y), x.Difference(y)
	for k := 0; k < 10; k++ {
		check.Eq(t, []bool{u.Contains(k), i.Contains(k), d.Contains(k)}, []bool{k >= 1 && k <= 4, k == 2 || k == 3, k == 1})
	}
}

func TestEx32_5(t *testing.T) {
	check.Eq(t, bv(200, 9, 2, 63, 64, 199).Members(), []int{2, 9, 63, 64, 199})
	check.Eq(t, bv(70, 9, 2, 63).Members(), []int{2, 9, 63})
}

func TestEx32_6(t *testing.T) {
	r := SpaceReport(1_000_000)
	check.Eq(t, []int{r["boolSlice"], r["bitVector"], r["map"]}, []int{1000024, 125024, 9 * (1 << 21)})
}

func TestEx32_7(t *testing.T) {
	v, ok := FirstDuplicate([]int{4, 9, 1, 9, 4}, 10)
	check.Eq(t, []any{v, ok}, []any{9, true})
	_, ok = FirstDuplicate([]int{1, 2, 3}, 10)
	check.Eq(t, ok, false)
}

func TestEx32_8(t *testing.T) {
	check.Eq(t, UnionList([]int{1, 5}, []int{5, 7}), []int{1, 5, 7})
	check.Eq(t, []int{CommonCount([]int{1, 5}, []int{5, 7}), CommonCount([]int{1, 300, 900}, []int{300, 900, 2})}, []int{1, 2})
}

func TestEx32_9(t *testing.T) {
	check.Eq(t, Missing([]int{0, 3, 1, 6}, 7), []int{2, 4, 5})
	check.Eq(t, len(Missing([]int{0, 1, 2}, 3)), 0)
}
