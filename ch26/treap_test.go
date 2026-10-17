package ch26

import (
	"math"
	"math/rand/v2"
	"testing"

	"dsa/internal/check"
)

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

// keys walks in order, written here so tests before 26.5 don't depend on its Inorder.
func keys(n *TreapNode) []int {
	if n == nil {
		return nil
	}
	return append(append(keys(n.Left), n.Key), keys(n.Right)...)
}

func TestEx26_1(t *testing.T) {
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

func TestEx26_2(t *testing.T) {
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

func TestEx26_3(t *testing.T) {
	rng := rand.New(rand.NewPCG(2, 0))
	var s *TreapNode
	for _, k := range []int{5, 2, 8} {
		s = Insert(s, k, rng)
	}
	check.True(t, Contains(s, 8) && !Contains(s, 3) && !Contains(nil, 1), "contains")
}

func TestEx26_4(t *testing.T) {
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

func TestEx26_5(t *testing.T) {
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

func TestEx26_6(t *testing.T) {
	h := Height(seq(1, 1001), rand.New(rand.NewPCG(0, 0)))
	check.True(t, h >= 15 && h <= 40, "treap height logarithmic")
}

func TestEx26_7(t *testing.T) {
	l, r := Split(treap(10, rand.New(rand.NewPCG(5, 0))), 4)
	check.Eq(t, keys(l), []int{0, 1, 2, 3})
	check.Eq(t, keys(r), []int{4, 5, 6, 7, 8, 9})
	check.True(t, IsValid(l, math.MinInt, math.MaxInt) && IsValid(r, math.MinInt, math.MaxInt), "halves valid")
	m := Merge(l, r)
	check.Eq(t, keys(m), seq(0, 10))
	check.Eq(t, IsValid(m, math.MinInt, math.MaxInt), true)
}
