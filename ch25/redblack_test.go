package ch25

import (
	"math/rand/v2"
	"slices"
	"testing"

	"dsa/internal/check"
)

func valid(t *testing.T, tr *RedBlackTree) {
	t.Helper()
	_, err := tr.Validate()
	check.Eq(t, err, nil)
}

func tree(keys ...int) *RedBlackTree {
	tr := NewRedBlackTree()
	for _, k := range keys {
		tr.Insert(k)
	}
	return tr
}

func upTo(n int) []int {
	a := make([]int, n)
	for i := range a {
		a[i] = i + 1
	}
	return a
}

func TestEx25_1(t *testing.T) {
	tr := NewRedBlackTree()
	check.True(t, tr.root == tr.NIL && !tr.NIL.Red, "empty tree is the black sentinel")
	n := func(k int, red bool) *RBNode {
		return &RBNode{Key: k, Red: red, left: tr.NIL, right: tr.NIL, parent: tr.NIL}
	}
	x, a, y, b, c := n(2, false), n(1, true), n(4, true), n(3, false), n(5, false)
	x.left, x.right, a.parent, y.parent = a, y, x, x
	y.left, y.right, b.parent, c.parent = b, c, y, y
	tr.root = x
	tr.rotateLeft(x)
	check.True(t, tr.root == y && y.left == x && x.parent == y && y.parent == tr.NIL, "left rotation")
	check.True(t, x.left == a && x.right == b && b.parent == x && y.right == c, "subtrees moved")
	tr.rotateRight(y)
	check.True(t, tr.root == x && x.right == y && y.parent == x && y.left == b && b.parent == y, "right rotation")
	check.Eq(t, tr.Rotations, 2)
}

func TestEx25_2(t *testing.T) {
	tr := NewRedBlackTree()
	for k := 1; k <= 7; k++ {
		check.Eq(t, tr.Insert(k), true)
	}
	check.Eq(t, []any{tr.root.Key, tr.root.Red, tr.Size}, []any{2, false, 7})
	check.Eq(t, tr.Insert(3), false)
}

func TestEx25_3(t *testing.T) {
	tr := NewRedBlackTree()
	for k := 1; k <= 7; k++ {
		tr.Insert(k)
		valid(t, tr)
	}
	valid(t, tree(upTo(1000)...))
	tr.root.Red = true // break the black-root rule by hand
	_, err := tr.Validate()
	check.True(t, err != nil, "invalid tree reported")
}

func TestEx25_4(t *testing.T) {
	s := tree(5, 2, 8, 1)
	check.True(t, s.Contains(8) && !s.Contains(3), "contains")
	check.Eq(t, s.Inorder(), []int{1, 2, 5, 8})
}

func TestEx25_5(t *testing.T) {
	check.Eq(t, tree(upTo(7)...).Height(), 4)
	big := tree(upTo(1000)...)
	check.True(t, big.Height() <= 19 && float64(big.Height()) <= big.HeightBound(), "height bounded")
}

func TestEx25_6(t *testing.T) {
	m := tree(20, 10, 30, 25)
	check.Eq(t, []int{m.minimum(m.root).Key, m.Maximum()}, []int{10, 30})
	v, ok := m.Successor(20)
	check.Eq(t, []any{v, ok}, []any{25, true})
	_, ok = m.Successor(30)
	check.Eq(t, ok, false)
}

func TestEx25_7(t *testing.T) {
	sorted := make([]int, 1000)
	for i := range sorted {
		sorted[i] = i
	}
	check.True(t, RotationsFor(sorted) < 1000, "rotations amortized O(1)")
}

func TestEx25_8(t *testing.T) {
	d := tree(upTo(100)...)
	var odd []int
	for k := 1; k <= 100; k++ {
		if k%2 == 0 {
			check.Eq(t, d.Delete(k), true)
			valid(t, d)
		} else {
			odd = append(odd, k)
		}
	}
	check.Eq(t, d.Inorder(), odd)
	check.Eq(t, d.Delete(2), false)
	rng := rand.New(rand.NewPCG(7, 0))
	r := NewRedBlackTree()
	ref := map[int]bool{}
	for i := 0; i < 5000; i++ {
		k := rng.IntN(300)
		if rng.IntN(3) == 0 {
			check.Eq(t, r.Delete(k), ref[k])
			delete(ref, k)
		} else {
			check.Eq(t, r.Insert(k), !ref[k])
			ref[k] = true
		}
		if i%100 == 0 {
			valid(t, r)
		}
	}
	var keys []int
	for k := range ref {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	check.Eq(t, []any{r.Inorder(), r.Size}, []any{keys, len(keys)})
}
