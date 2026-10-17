package ch74

import (
	"math"
	"math/big"
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

func TestEx74_1(t *testing.T) {
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

func TestEx74_2(t *testing.T) {
	s := SampleWithSeed(42, seq(100), 3)
	check.Eq(t, SampleWithSeed(42, seq(100), 3), s)
	check.Eq(t, len(s), 3)
}

func TestEx74_3(t *testing.T) {
	g := &LCG{X: 1, A: 5, C: 3, M: 16}
	check.Eq(t, []uint64{g.Next(), g.Next(), g.Next()}, []uint64{8, 11, 10})
	check.Eq(t, LCGPeriod(1, 5, 3, 16), 16)
	check.True(t, NewLCG(7).Next() < 1<<31, "LCG range")
}

func TestEx74_4(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	FisherYates(a, rand.New(rand.NewPCG(7, 0)))
	b := []int{1, 2, 3, 4, 5}
	FisherYates(b, rand.New(rand.NewPCG(7, 0)))
	check.Eq(t, a, b)
	check.Eq(t, slices.Sorted(slices.Values(a)), []int{1, 2, 3, 4, 5})
}

func TestEx74_5(t *testing.T) {
	check.Eq(t, FisherYatesDistribution(3), map[string]int{"012": 1, "021": 1, "102": 1, "120": 1, "201": 1, "210": 1})
	check.Eq(t, NaiveDistribution(3), map[string]int{"012": 4, "021": 5, "102": 5, "120": 5, "201": 4, "210": 4})
}

func TestEx74_6(t *testing.T) {
	keys := seq(200)
	check.Eq(t, BSTHeight(keys), 200)
	sh := slices.Clone(keys)
	rand.New(rand.NewPCG(3, 0)).Shuffle(len(sh), func(i, j int) { sh[i], sh[j] = sh[j], sh[i] })
	h := BSTHeight(sh)
	check.True(t, h >= 10 && h <= 30, "shuffled height logarithmic")
}

func TestEx74_7(t *testing.T) {
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

func TestEx74_8(t *testing.T) {
	one := MaxLoadRandom(10000, 1000, rand.New(rand.NewPCG(4, 0)))
	two := MaxLoadTwoChoices(10000, 1000, rand.New(rand.NewPCG(4, 0)))
	check.True(t, two < one && two <= 14, "two choices lower")
}

func TestEx74_9(t *testing.T) {
	pick := MakeWeightedPicker([]float64{1, 3, 6}, rand.New(rand.NewPCG(8, 0)))
	n2 := 0
	for i := 0; i < 10000; i++ {
		if pick() == 2 {
			n2++
		}
	}
	check.True(t, n2 > 5700 && n2 < 6300, "weighted about 60%")
}

func TestEx74_10(t *testing.T) {
	check.True(t, math.Abs(EstimatePi(100_000, rand.New(rand.NewPCG(0, 0)))-math.Pi) < 0.02, "pi")
}

func TestEx74_11(t *testing.T) {
	rng := rand.New(rand.NewPCG(1, 0))
	data := make([]int, 1_000_000)
	for i := range data {
		data[i] = rng.IntN(1000)
	}
	exact := ExactMean(data)
	check.True(t, math.Abs(SampledMean(data, 2000, rng)-exact)/exact < 0.05, "sampled mean close")
	check.Eq(t, ExactMean([]int{1, 2, 3, 6}), 3.0)
}

func TestEx74_12(t *testing.T) {
	rng := rand.New(rand.NewPCG(2, 0))
	half := make([]int, 100)
	for i := range half {
		half[i] = i % 2
	}
	check.Eq(t, half[LasVegasFindOne(half, rng)], 1)
	i := MonteCarloFindOne(half, rng, 20)
	check.True(t, i == -1 || half[i] == 1, "monte carlo result")
	check.Eq(t, MonteCarloFindOne(make([]int, 10), rng, 20), -1)
}

func TestEx74_13(t *testing.T) {
	check.Eq(t, []bool{IsPrimeTrial(999_999_000_001), IsPrimeTrial(1), IsPrimeTrial(2), IsPrimeTrial(999_999_000_001 * 3)}, []bool{true, false, true, false})
}

func TestEx74_14(t *testing.T) {
	check.Eq(t, []bool{FermatWitness(2, 15), FermatWitness(2, 13), FermatWitness(2, 561)}, []bool{true, false, false})
	r := rand.New(rand.NewPCG(5, 0))
	for k := 0; k < 3000; k++ { // MulMod and PowMod against math/big, including n near 2^64
		n := r.Uint64() | 1
		a, b, e := r.Uint64()%n, r.Uint64()%n, r.Uint64()
		bn := new(big.Int).SetUint64(n)
		want := new(big.Int).Mul(new(big.Int).SetUint64(a), new(big.Int).SetUint64(b))
		check.Eq(t, MulMod(a, b, n), want.Mod(want, bn).Uint64())
		pw := new(big.Int).Exp(new(big.Int).SetUint64(a), new(big.Int).SetUint64(e), bn)
		check.Eq(t, PowMod(a, e, n), pw.Uint64())
	}
}

func TestEx74_15(t *testing.T) {
	rng := rand.New(rand.NewPCG(6, 0))
	check.Eq(t, []bool{FermatTest(1_000_000_007, 20, rng), FermatTest(1_000_000_007*998_244_353, 20, rng)}, []bool{true, false})
}

func TestEx74_16(t *testing.T) {
	check.Eq(t, []bool{MillerRabin(561), MillerRabin(1<<61 - 1), MillerRabin(1)}, []bool{false, true, false})
	r := rand.New(rand.NewPCG(7, 0))
	for k := 0; k < 3000; k++ {
		n := r.Uint64() | 1
		check.Eq(t, MillerRabin(n), new(big.Int).SetUint64(n).ProbablyPrime(20))
	}
	for n := uint64(0); n < 5000; n++ {
		check.Eq(t, MillerRabin(n), new(big.Int).SetUint64(n).ProbablyPrime(20))
	}
	for _, c := range []uint64{561, 1105, 1729, 2465, 2821, 6601} {
		check.Eq(t, MillerRabin(c), false)
	}
}

func TestEx74_17(t *testing.T) {
	ex := BirthdayExact(23)
	check.True(t, math.Abs(ex-0.5072972) < 1e-6, "birthday exact")
	check.True(t, math.Abs(BirthdaySimulation(23, 20000, rand.New(rand.NewPCG(3, 0)))-0.5072972) < 0.02, "birthday simulated")
}
