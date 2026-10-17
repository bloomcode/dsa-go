package ch54

import (
	"math"
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

func TestEx54_1(t *testing.T) {
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

func TestEx54_2(t *testing.T) {
	tr := NewRedBlackTree()
	for k := 1; k <= 7; k++ {
		check.Eq(t, tr.Insert(k), true)
	}
	check.Eq(t, []any{tr.root.Key, tr.root.Red, tr.Size}, []any{2, false, 7})
	check.Eq(t, tr.Insert(3), false)
}

func TestEx54_3(t *testing.T) {
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

func TestEx54_4(t *testing.T) {
	s := tree(5, 2, 8, 1)
	check.True(t, s.Contains(8) && !s.Contains(3), "contains")
	check.Eq(t, s.Inorder(), []int{1, 2, 5, 8})
}

func TestEx54_5(t *testing.T) {
	check.Eq(t, tree(upTo(7)...).Height(), 4)
	big := tree(upTo(1000)...)
	check.True(t, big.Height() <= 19 && float64(big.Height()) <= big.HeightBound(), "height bounded")
}

func TestEx54_6(t *testing.T) {
	m := tree(20, 10, 30, 25)
	check.Eq(t, []int{m.minimum(m.root).Key, m.Maximum()}, []int{10, 30})
	v, ok := m.Successor(20)
	check.Eq(t, []any{v, ok}, []any{25, true})
	_, ok = m.Successor(30)
	check.Eq(t, ok, false)
}

func TestEx54_7(t *testing.T) {
	sorted := make([]int, 1000)
	for i := range sorted {
		sorted[i] = i
	}
	check.True(t, RotationsFor(sorted) < 1000, "rotations amortized O(1)")
}

func TestEx54_8(t *testing.T) {
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

func seq(lo, hi int) []int {
	var a []int
	for i := lo; i < hi; i++ {
		a = append(a, i)
	}
	return a
}

func treap(n int, rng *rand.Rand) *TreapNode {
	var root *TreapNode
	for k := range n {
		root = Insert(root, k, rng)
	}
	return root
}

// keys walks in order, written here so tests before 54.13 don't depend on its Inorder.
func keys(n *TreapNode) []int {
	if n == nil {
		return nil
	}
	return append(append(keys(n.Left), n.Key), keys(n.Right)...)
}

func TestEx54_9(t *testing.T) {
	a, b, c := &TreapNode{Key: 1, Priority: 0.1}, &TreapNode{Key: 4, Priority: 0.2}, &TreapNode{Key: 7, Priority: 0.3}
	left := &TreapNode{Key: 3, Priority: 0.9, Left: a, Right: b}
	top := &TreapNode{Key: 5, Priority: 0.5, Left: left, Right: c}
	root := RotateRight(top)
	check.True(t, root == left && root.Right == top && top.Left == b, "right rotation")
	check.Eq(t, keys(root), []int{1, 3, 4, 5, 7})
	root = RotateLeft(root)
	check.True(t, root == top && top.Left == left && left.Right == b, "left rotation")
	check.Eq(t, keys(root), []int{1, 3, 4, 5, 7})
}

func TestEx54_10(t *testing.T) {
	root := treap(10, rand.New(rand.NewPCG(1, 0)))
	check.Eq(t, keys(root), seq(0, 10))
	stack := []*TreapNode{root}
	for len(stack) > 0 {
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for _, child := range []*TreapNode{n.Left, n.Right} {
			if child != nil {
				check.True(t, child.Priority <= n.Priority, "heap order")
				stack = append(stack, child)
			}
		}
	}
}

func TestEx54_11(t *testing.T) {
	rng := rand.New(rand.NewPCG(2, 0))
	var s *TreapNode
	for _, k := range []int{5, 2, 8} {
		s = Insert(s, k, rng)
	}
	check.True(t, Contains(s, 8) && !Contains(s, 3) && !Contains(nil, 1), "contains")
}

func TestEx54_12(t *testing.T) {
	d := treap(100, rand.New(rand.NewPCG(3, 0)))
	var rest []int
	for k := range 100 {
		if k%3 == 0 {
			d = Delete(d, k)
		} else {
			rest = append(rest, k)
		}
	}
	check.Eq(t, keys(d), rest)
	check.Eq(t, keys(Delete(d, 1000)), rest)
}

func TestEx54_13(t *testing.T) {
	root := treap(100, rand.New(rand.NewPCG(1, 0)))
	check.Eq(t, IsValid(root, math.MinInt, math.MaxInt), true)
	check.Eq(t, Inorder(root), seq(0, 100))
	if root.Left != nil { // break heap order by hand
		root.Left.Priority, root.Priority = root.Priority, root.Left.Priority
		check.Eq(t, IsValid(root, math.MinInt, math.MaxInt), false)
	}
	unordered := &TreapNode{Key: 5, Priority: 0.9, Left: &TreapNode{Key: 7, Priority: 0.1}}
	check.Eq(t, IsValid(unordered, math.MinInt, math.MaxInt), false)
}

func TestEx54_14(t *testing.T) {
	h := Height(seq(1, 1001), rand.New(rand.NewPCG(0, 0)))
	check.True(t, h >= 15 && h <= 40, "treap height logarithmic")
}

func TestEx54_15(t *testing.T) {
	l, r := Split(treap(10, rand.New(rand.NewPCG(5, 0))), 4)
	check.Eq(t, keys(l), []int{0, 1, 2, 3})
	check.Eq(t, keys(r), []int{4, 5, 6, 7, 8, 9})
	check.True(t, IsValid(l, math.MinInt, math.MaxInt) && IsValid(r, math.MinInt, math.MaxInt), "halves valid")
	m := Merge(l, r)
	check.Eq(t, keys(m), seq(0, 10))
	check.Eq(t, IsValid(m, math.MinInt, math.MaxInt), true)
}
