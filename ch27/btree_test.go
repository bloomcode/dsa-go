package ch27

import (
	"math/rand/v2"
	"slices"
	"testing"

	"dsa/internal/check"
)

func valid(t *testing.T, b *BTree) int {
	t.Helper()
	h, err := b.Validate()
	check.Eq(t, err, nil)
	return h
}

func btree(deg, lo, hi int) *BTree {
	b := NewBTree(deg)
	for k := lo; k <= hi; k++ {
		b.Insert(k)
	}
	return b
}

func TestEx27_1(t *testing.T) {
	check.Eq(t, []int{ScanIOs(1000, 64), RandomReadsIOs(1000)}, []int{16, 1000})
}

func TestEx27_2(t *testing.T) {
	data := make([]int, 10000)
	for i := range data {
		data[i] = 2 * i
	}
	ext := NewExternalArray(data, 100)
	check.Eq(t, ExternalBinarySearch(ext, 2*4321), 4321)
	check.True(t, ext.IOs <= 9, "about 8 I/Os")
	check.Eq(t, []int{ExternalBinarySearch(ext, 7), ExternalBinarySearch(ext, 0), ExternalBinarySearch(ext, 19998)}, []int{-1, 0, 9999})
}

func TestEx27_3(t *testing.T) {
	check.Eq(t, []int{BSTIOs(1_000_000_000), BTreeIOs(1_000_000_000, 1000)}, []int{30, 3})
}

func TestEx27_4(t *testing.T) {
	b := NewBTree(2)
	check.Eq(t, b.Search(1), false)
	b.root = &BTreeNode{keys: []int{3}, children: []*BTreeNode{{leaf: true, keys: []int{1, 2}}, {leaf: true, keys: []int{4, 5}}}}
	for k := 1; k <= 5; k++ {
		check.Eq(t, b.Search(k), true)
	}
	check.True(t, !b.Search(0) && !b.Search(6), "missing")
}

func TestEx27_5(t *testing.T) {
	b := NewBTree(2)
	for k := 1; k <= 10; k++ {
		check.Eq(t, b.Insert(k), true)
	}
	check.Eq(t, b.Insert(5), false)
	for k := 1; k <= 10; k++ {
		check.Eq(t, b.Search(k), true)
	}
	check.True(t, len(b.root.keys) <= 3, "root within 2t - 1 keys")
	h := btree(3, 1, 100)
	check.True(t, h.Search(42) && !h.Search(101), "search after inserts")
}

func TestEx27_6(t *testing.T) {
	b := NewBTree(2)
	for k := 1; k <= 10; k++ {
		b.Insert(k)
		valid(t, b)
	}
	check.Eq(t, b.Inorder(), []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
}

func TestEx27_7(t *testing.T) {
	check.Eq(t, []int{BTreeHeightBound(1_000_000, 50), BTreeHeightBound(10000, 3)}, []int{4, 8})
	check.True(t, valid(t, btree(3, 1, 10000)) <= 8, "within height bound")
}

func TestEx27_8(t *testing.T) {
	h := btree(3, 1, 100)
	check.Eq(t, h.RangeQuery(40, 45), []int{40, 41, 42, 43, 44, 45})
	check.Eq(t, h.RangeQuery(99, 500), []int{99, 100})
	check.Eq(t, len(h.RangeQuery(200, 300)), 0)
}

func TestEx27_9(t *testing.T) {
	d := btree(3, 1, 200)
	var odds, evens []int
	for k := 1; k <= 200; k++ {
		if k%2 == 1 {
			odds = append(odds, k)
		} else {
			evens = append(evens, k)
		}
	}
	rng := rand.New(rand.NewPCG(4, 0))
	rng.Shuffle(len(odds), func(i, j int) { odds[i], odds[j] = odds[j], odds[i] })
	for _, k := range odds {
		check.Eq(t, d.Delete(k), true)
		valid(t, d)
	}
	check.Eq(t, d.Inorder(), evens)
	check.Eq(t, d.Delete(1), false)
	for _, deg := range []int{2, 3, 5} {
		r := NewBTree(deg)
		ref := map[int]bool{}
		for i := 0; i < 4000; i++ {
			k := rng.IntN(500)
			if rng.IntN(3) == 0 {
				check.Eq(t, r.Delete(k), ref[k])
				delete(ref, k)
			} else {
				check.Eq(t, r.Insert(k), !ref[k])
				ref[k] = true
			}
			if i%50 == 0 {
				valid(t, r)
			}
		}
		var keys []int
		for k := range ref {
			keys = append(keys, k)
		}
		slices.Sort(keys)
		check.Eq(t, r.Inorder(), keys)
	}
}
