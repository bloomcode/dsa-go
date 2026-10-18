package ch8_1

import (
	"math"
	"math/rand/v2"
	"testing"

	"dsa/internal/check"
)

func checkLRU(t *testing.T, c Cache[string, int]) {
	t.Helper()
	c.Put("a", 1)
	c.Put("b", 2)
	v, ok := c.Get("a")
	check.Eq(t, []any{v, ok}, []any{1, true})
	c.Put("c", 3)
	v, ok = c.Get("b")
	check.Eq(t, []any{v, ok}, []any{0, false})
	v, _ = c.Get("c")
	check.Eq(t, v, 3)
	c.Put("a", 10)
	v, _ = c.Get("a")
	check.Eq(t, v, 10)
}

func TestEx8_1_1(t *testing.T) { checkLRU(t, NewLRU[string, int](2)) }

func TestEx8_1_2(t *testing.T) {
	checkLRU(t, NewLinkedLRU[string, int](2))
	rng := rand.New(rand.NewPCG(1, 0))
	a, b := NewLRU[int, int](5), NewLinkedLRU[int, int](5)
	for i := 0; i < 5000; i++ {
		k := rng.IntN(12)
		if rng.IntN(2) == 0 {
			a.Put(k, i)
			b.Put(k, i)
		} else {
			x, ok1 := a.Get(k)
			y, ok2 := b.Get(k)
			check.Eq(t, []any{x, ok1}, []any{y, ok2})
		}
	}
}

func TestEx8_1_3(t *testing.T) {
	check.True(t, math.Abs(HitRate[string](NewLRU[string, bool](2), []string{"a", "b", "a", "a", "c", "b"})-1.0/3) < 1e-9, "hit rate 1/3")
}

func TestEx8_1_4(t *testing.T) {
	f := NewFIFO[string, int](2)
	f.Put("a", 1)
	f.Put("b", 2)
	f.Get("a")
	f.Put("c", 3)
	_, ok := f.Get("a")
	check.Eq(t, ok, false)
	var hot []int
	for i := 0; i < 2000; i++ {
		hot = append(hot, i%2, 100+i)
	}
	check.True(t, HitRate[int](NewLRU[int, bool](4), hot) > HitRate[int](NewFIFO[int, bool](4), hot), "LRU beats FIFO on hot keys")
}

func TestEx8_1_5(t *testing.T) {
	rc := NewRandomCache[int, int](3, rand.New(rand.NewPCG(2, 0)))
	for i := 0; i < 50; i++ {
		rc.Put(i, i)
		v, ok := rc.Get(i)
		check.True(t, rc.Len() <= 3 && ok && v == i, "random cache bounded and holds the newest")
	}
}

func TestEx8_1_6(t *testing.T) {
	cyc := CyclicRequests(10, 100)
	check.Eq(t, len(cyc), 1100)
	check.Eq(t, HitRate[int](NewLRU[int, bool](10), cyc), 0.0)
	check.True(t, HitRate[int](NewRandomCache[int, bool](10, rand.New(rand.NewPCG(3, 0))), cyc) > 0.5, "random eviction keeps hits")
}

func TestEx8_1_7(t *testing.T) {
	m := NewMemo(func(x int) int { return x * x }, 10)
	check.Eq(t, []int{m.Call(1), m.Call(2), m.Call(1)}, []int{1, 4, 1})
	check.Eq(t, []int{m.Hits, m.Misses}, []int{1, 2})
}

func TestEx8_1_8(t *testing.T) {
	g := [][]int{{1, 2}, {3, 4}}
	check.Eq(t, []int{SumRowsFirst(g), SumColumnsFirst(g)}, []int{10, 10})
	check.Eq(t, BlockedTranspose([][]int{{1, 2, 3}, {4, 5, 6}}, 2), [][]int{{1, 4}, {2, 5}, {3, 6}})
}
