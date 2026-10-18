package ch4_2

import (
	"math/big"
	"math/rand/v2"
	"testing"

	"dsa/internal/check"
)

func words(n int) []string {
	rng := rand.New(rand.NewPCG(1, 0))
	seen := map[string]bool{}
	var out []string
	for len(out) < n {
		b := make([]byte, 4+rng.IntN(5))
		for i := range b {
			b[i] = byte('a' + rng.IntN(26))
		}
		if w := string(b); !seen[w] {
			seen[w] = true
			out = append(out, w)
		}
	}
	return out
}

func eights() []int {
	var out []int
	for k := 0; k < 8000; k += 8 {
		out = append(out, k)
	}
	return out
}

func TestEx4_2_1(t *testing.T) {
	check.Eq(t, SumHash("stop", 101) == SumHash("pots", 101) && SumHash("pots", 101) == SumHash("tops", 101), true)
	h := SumHash("anything", 7)
	check.True(t, h >= 0 && h < 7, "in range")
}

func TestEx4_2_2(t *testing.T) {
	check.Eq(t, PolyHash("stop", 1_000_000_007) != PolyHash("pots", 1_000_000_007), true)
}

func TestEx4_2_3(t *testing.T) {
	check.Eq(t, []int{BucketsUsed(eights(), 64, DivisionHash), BucketsUsed(eights(), 61, DivisionHash)}, []int{8, 61})
	check.Eq(t, DivisionHash(-3, 10), 7)
}

func TestEx4_2_4(t *testing.T) {
	ws := words(10000)
	poly, sum := MaxBucket(ws, 1000, PolyHash), MaxBucket(ws, 1000, SumHash)
	check.True(t, poly <= 30, "poly spreads well")
	check.True(t, sum > 2*poly, "sum hash crowds")
}

func TestEx4_2_5(t *testing.T) {
	r := rand.New(rand.NewPCG(2, 0))
	pBig := new(big.Int).SetUint64(P)
	for i := 0; i < 5000; i++ {
		a, b := r.Uint64N(P), r.Uint64N(P)
		want := new(big.Int).Mul(new(big.Int).SetUint64(a), new(big.Int).SetUint64(b))
		check.Eq(t, MulModP(a, b), want.Mod(want, pBig).Uint64())
	}
	check.Eq(t, MulModP(P-1, P-1), uint64(1))
	u1 := NewUniversalHash(1000, rand.New(rand.NewPCG(1, 0)))
	u2 := NewUniversalHash(1000, rand.New(rand.NewPCG(2, 0)))
	diff := 0
	for k := 0; k < 100; k++ {
		h := u1.Hash(k)
		check.True(t, h >= 0 && h < 1000, "in range")
		if h != u2.Hash(k) {
			diff++
		}
	}
	check.True(t, diff > 90, "different seeds differ")
	check.True(t, u1.Hash(-5) >= 0, "negative key in range")
}

func TestEx4_2_6(t *testing.T) {
	adv := AdversarialKeys(1000, 5000)
	check.Eq(t, len(adv), 5000)
	check.Eq(t, Fullest(adv, func(k int) int { return k % 1000 }), 5000)
	check.True(t, Fullest(adv, NewUniversalHash(1000, rand.New(rand.NewPCG(1, 0))).Hash) < 25, "universal spreads adversarial keys")
}

func TestEx4_2_7(t *testing.T) {
	tbl := NewRandomizedHashTable(rand.New(rand.NewPCG(3, 0)))
	adv := AdversarialKeys(1024, 10000)
	for _, k := range adv {
		tbl.Put(k, "v")
	}
	tbl.Put(adv[5], "changed")
	check.Eq(t, tbl.Len(), 10000)
	for _, k := range adv[:100] {
		_, ok := tbl.Get(k)
		check.Eq(t, ok, true)
	}
	v, _ := tbl.Get(adv[5])
	check.Eq(t, v, "changed")
	_, ok := tbl.Get(7)
	check.Eq(t, ok, false)
	check.True(t, tbl.LongestChain() < 15, "chains stay short")
}

func TestEx4_2_8(t *testing.T) {
	check.Eq(t, BucketsUsed(eights(), 64, MultHash), 64)
	for _, k := range eights() {
		h := MultHash(k, 64)
		check.True(t, h >= 0 && h < 64, "in range")
	}
}
