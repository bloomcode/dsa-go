package ch23

import (
	"math"
	"math/rand/v2"
	"slices"
	"testing"

	"dsa/internal/check"
)

func seq(n int) []int {
	a := make([]int, n)
	for i := range a {
		a[i] = i
	}
	return a
}

func TestEx23_1(t *testing.T) {
	check.Eq(t, RandomQuicksort(seq(10), rand.New(rand.NewPCG(1, 0))), seq(10))
	rng := rand.New(rand.NewPCG(5, 0))
	for trial := 0; trial < 100; trial++ {
		d := make([]int, rng.IntN(50))
		for i := range d {
			d[i] = rng.IntN(10)
		}
		got := RandomQuicksort(d, rng)
		if len(d) == 0 {
			check.Eq(t, len(got), 0)
		} else {
			check.Eq(t, got, slices.Sorted(slices.Values(d)))
		}
	}
}

func TestEx23_2(t *testing.T) {
	s := SampleWithSeed(42, seq(100), 3)
	check.Eq(t, SampleWithSeed(42, seq(100), 3), s)
	check.Eq(t, len(s), 3)
}

func TestEx23_3(t *testing.T) {
	g := &LCG{X: 1, A: 5, C: 3, M: 16}
	check.Eq(t, []uint64{g.Next(), g.Next(), g.Next()}, []uint64{8, 11, 10})
	check.Eq(t, LCGPeriod(1, 5, 3, 16), 16)
	check.True(t, NewLCG(7).Next() < 1<<31, "LCG range")
}

func TestEx23_4(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	FisherYates(a, rand.New(rand.NewPCG(7, 0)))
	b := []int{1, 2, 3, 4, 5}
	FisherYates(b, rand.New(rand.NewPCG(7, 0)))
	check.Eq(t, a, b)
	check.Eq(t, slices.Sorted(slices.Values(a)), []int{1, 2, 3, 4, 5})
}

func TestEx23_5(t *testing.T) {
	check.Eq(t, FisherYatesDistribution(3), map[string]int{"012": 1, "021": 1, "102": 1, "120": 1, "201": 1, "210": 1})
	check.Eq(t, NaiveDistribution(3), map[string]int{"012": 4, "021": 5, "102": 5, "120": 5, "201": 4, "210": 4})
}

func TestEx23_6(t *testing.T) {
	keys := seq(200)
	check.Eq(t, BSTHeight(keys), 200)
	sh := slices.Clone(keys)
	rand.New(rand.NewPCG(3, 0)).Shuffle(len(sh), func(i, j int) { sh[i], sh[j] = sh[j], sh[i] })
	h := BSTHeight(sh)
	check.True(t, h >= 10 && h <= 30, "shuffled height logarithmic")
}

func TestEx23_7(t *testing.T) {
	r := Reservoir(func(yield func(int) bool) {
		for i := range 1_000_000 {
			if !yield(i) {
				return
			}
		}
	}, 3, rand.New(rand.NewPCG(2, 0)))
	check.Eq(t, len(r), 3)
	counts := make([]int, 10)
	rr := rand.New(rand.NewPCG(11, 0))
	for trial := 0; trial < 20000; trial++ {
		for _, x := range Reservoir(slices.Values(seq(10)), 2, rr) {
			counts[x]++
		}
	}
	for _, n := range counts {
		check.True(t, math.Abs(float64(n)-4000) < 300, "reservoir uniform")
	}
}

func TestEx23_8(t *testing.T) {
	one := MaxLoadRandom(10000, 1000, rand.New(rand.NewPCG(4, 0)))
	two := MaxLoadTwoChoices(10000, 1000, rand.New(rand.NewPCG(4, 0)))
	check.True(t, two < one && two <= 14, "two choices lower")
}

func TestEx23_9(t *testing.T) {
	pick := MakeWeightedPicker([]float64{1, 3, 6}, rand.New(rand.NewPCG(8, 0)))
	n2 := 0
	for i := 0; i < 10000; i++ {
		if pick() == 2 {
			n2++
		}
	}
	check.True(t, n2 > 5700 && n2 < 6300, "weighted about 60%")
}
