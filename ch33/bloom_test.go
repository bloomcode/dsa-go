package ch33

import (
	"math"
	"strconv"
	"testing"

	"dsa/internal/check"
)

func TestEx33_1(t *testing.T) {
	p := BitPositions("cat", 3, 32)
	check.Eq(t, len(p), 3)
	check.Eq(t, BitPositions("cat", 3, 32), p)
	for _, i := range p {
		check.True(t, i >= 0 && i < 32, "in range")
	}
}

func TestEx33_2(t *testing.T) {
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

func TestEx33_3(t *testing.T) {
	check.True(t, math.Abs(FalsePositiveRate(1000, 10_000, 7)-0.0082) < 0.0002, "predicted rate")
}

func TestEx33_4(t *testing.T) {
	m, k := OptimalParams(1_000_000, 0.01)
	check.Eq(t, []int{m, k}, []int{9585059, 7})
}

func TestEx33_5(t *testing.T) {
	check.True(t, math.Abs(MeasuredFPRate(1000, 10_000, 7, 20_000)-0.0082) < 0.005, "measured close to predicted")
}

func TestEx33_6(t *testing.T) {
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

func TestEx33_7(t *testing.T) {
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

func TestEx33_8(t *testing.T) {
	c := NewCountingBloomFilter(1000, 4)
	c.Add("a")
	c.Add("b")
	check.Eq(t, c.Remove("a"), true)
	check.Eq(t, []bool{c.MightContain("a"), c.MightContain("b")}, []bool{false, true})
	check.Eq(t, c.Remove("zzz"), false)
}
