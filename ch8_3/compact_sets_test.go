package ch8_3

import (
	"math"
	"strconv"
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

func TestEx8_3_1(t *testing.T) {
	s := NewBoolArraySet(10)
	s.Add(3)
	s.Add(7)
	s.Add(7)
	s.Remove(3)
	check.Eq(t, []any{s.Contains(7), s.Contains(3), s.Len()}, []any{true, false, 1})
}

func TestEx8_3_2(t *testing.T) {
	b := bv(100, 42)
	check.Eq(t, []any{b.Contains(42), b.Contains(43), b.Words[0]}, []any{true, false, uint64(4398046511104)})
	b.Add(99)
	b.Remove(42)
	check.Eq(t, []bool{b.Contains(42), b.Contains(99)}, []bool{false, true})
}

func TestEx8_3_3(t *testing.T) {
	b := bv(100, 1, 5, 99)
	check.Eq(t, b.Count(), 3)
	b.Add(5)
	check.Eq(t, []int{b.Count(), NewBitVector(10).Count()}, []int{3, 0})
}

func TestEx8_3_4(t *testing.T) {
	x, y := bv(10, 1, 2, 3), bv(10, 2, 3, 4)
	u, i, d := x.Union(y), x.Intersection(y), x.Difference(y)
	for k := 0; k < 10; k++ {
		check.Eq(t, []bool{u.Contains(k), i.Contains(k), d.Contains(k)}, []bool{k >= 1 && k <= 4, k == 2 || k == 3, k == 1})
	}
}

func TestEx8_3_5(t *testing.T) {
	check.Eq(t, bv(200, 9, 2, 63, 64, 199).Members(), []int{2, 9, 63, 64, 199})
	check.Eq(t, bv(70, 9, 2, 63).Members(), []int{2, 9, 63})
}

func TestEx8_3_6(t *testing.T) {
	r := SpaceReport(1_000_000)
	check.Eq(t, []int{r["boolSlice"], r["bitVector"], r["map"]}, []int{1000024, 125024, 9 * (1 << 21)})
}

func TestEx8_3_7(t *testing.T) {
	v, ok := FirstDuplicate([]int{4, 9, 1, 9, 4}, 10)
	check.Eq(t, []any{v, ok}, []any{9, true})
	_, ok = FirstDuplicate([]int{1, 2, 3}, 10)
	check.Eq(t, ok, false)
}

func TestEx8_3_8(t *testing.T) {
	check.Eq(t, UnionList([]int{1, 5}, []int{5, 7}), []int{1, 5, 7})
	check.Eq(t, []int{CommonCount([]int{1, 5}, []int{5, 7}), CommonCount([]int{1, 300, 900}, []int{300, 900, 2})}, []int{1, 2})
}

func TestEx8_3_9(t *testing.T) {
	check.Eq(t, Missing([]int{0, 3, 1, 6}, 7), []int{2, 4, 5})
	check.Eq(t, len(Missing([]int{0, 1, 2}, 3)), 0)
}

func TestEx8_3_10(t *testing.T) {
	p := BitPositions("cat", 3, 32)
	check.Eq(t, len(p), 3)
	check.Eq(t, BitPositions("cat", 3, 32), p)
	for _, i := range p {
		check.True(t, i >= 0 && i < 32, "in range")
	}
}

func TestEx8_3_11(t *testing.T) {
	bf := NewBloomFilter(1000, 5)
	bf.Add("apple")
	bf.Add("pear")
	check.Eq(t, []bool{bf.MightContain("apple"), bf.MightContain("pear"), bf.MightContain("grape")}, []bool{true, true, false})
	big := NewBloomFilter(5000, 4)
	for i := 0; i < 300; i++ {
		big.Add("w" + strconv.Itoa(i))
	}
	for i := 0; i < 300; i++ {
		check.True(t, big.MightContain("w"+strconv.Itoa(i)), "no false negatives")
	}
}

func TestEx8_3_12(t *testing.T) {
	check.True(t, math.Abs(FalsePositiveRate(1000, 10_000, 7)-0.0082) < 0.0002, "predicted rate")
}

func TestEx8_3_13(t *testing.T) {
	m, k := OptimalParams(1_000_000, 0.01)
	check.Eq(t, []int{m, k}, []int{9585059, 7})
}

func TestEx8_3_14(t *testing.T) {
	check.True(t, math.Abs(MeasuredFPRate(1000, 10_000, 7, 20_000)-0.0082) < 0.005, "measured close to predicted")
}

func TestEx8_3_15(t *testing.T) {
	var stream []string
	for i := 0; i < 5000; i++ {
		stream = append(stream, "u-"+strconv.Itoa(i))
	}
	for i := 0; i < 100; i++ {
		stream = append(stream, "u-"+strconv.Itoa(i*37))
	}
	rep := StreamDuplicates(stream, 64_000, 6)
	alarms := FalseAlarms(stream, 64_000, 6)
	check.Eq(t, len(rep), 100+alarms)
	check.True(t, alarms < 20, "few false alarms")
}

func TestEx8_3_16(t *testing.T) {
	a, b := NewBloomFilter(512, 4), NewBloomFilter(512, 4)
	a.Add("x")
	b.Add("y")
	u, err := a.Union(b)
	check.Eq(t, err, nil)
	check.Eq(t, []bool{u.MightContain("x"), u.MightContain("y")}, []bool{true, true})
	check.Eq(t, u.Count, 2)
	_, err = a.Union(NewBloomFilter(256, 4))
	check.True(t, err != nil, "mismatched union errors")
}

func TestEx8_3_17(t *testing.T) {
	c := NewCountingBloomFilter(1000, 4)
	c.Add("a")
	c.Add("b")
	check.Eq(t, c.Remove("a"), true)
	check.Eq(t, []bool{c.MightContain("a"), c.MightContain("b")}, []bool{false, true})
	check.Eq(t, c.Remove("zzz"), false)
}
