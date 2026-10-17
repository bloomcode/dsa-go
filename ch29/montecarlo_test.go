package ch29

import (
	"math"
	"math/big"
	"math/rand/v2"
	"testing"

	"dsa/internal/check"
)

func TestEx29_1(t *testing.T) {
	check.True(t, math.Abs(EstimatePi(100_000, rand.New(rand.NewPCG(0, 0)))-math.Pi) < 0.02, "pi")
}

func TestEx29_2(t *testing.T) {
	rng := rand.New(rand.NewPCG(1, 0))
	data := make([]int, 1_000_000)
	for i := range data {
		data[i] = rng.IntN(1000)
	}
	exact := ExactMean(data)
	check.True(t, math.Abs(SampledMean(data, 2000, rng)-exact)/exact < 0.05, "sampled mean close")
	check.Eq(t, ExactMean([]int{1, 2, 3, 6}), 3.0)
}

func TestEx29_3(t *testing.T) {
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

func TestEx29_4(t *testing.T) {
	check.Eq(t, []bool{IsPrimeTrial(999_999_000_001), IsPrimeTrial(1), IsPrimeTrial(2), IsPrimeTrial(999_999_000_001 * 3)}, []bool{true, false, true, false})
}

func TestEx29_5(t *testing.T) {
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

func TestEx29_6(t *testing.T) {
	rng := rand.New(rand.NewPCG(6, 0))
	check.Eq(t, []bool{FermatTest(1_000_000_007, 20, rng), FermatTest(1_000_000_007*998_244_353, 20, rng)}, []bool{true, false})
}

func TestEx29_7(t *testing.T) {
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

func TestEx29_8(t *testing.T) {
	ex := BirthdayExact(23)
	check.True(t, math.Abs(ex-0.5072972) < 1e-6, "birthday exact")
	check.True(t, math.Abs(BirthdaySimulation(23, 20000, rand.New(rand.NewPCG(3, 0)))-0.5072972) < 0.02, "birthday simulated")
}
