package ch51

import (
	"math"
	"math/rand"
	"slices"
	"testing"

	"dsa/internal/check"
)

// sampleTree is 50 -> (25 -> 10, 33), 75, built by hand so tests don't depend on Insert.
func sampleTree() *TreeNode {
	return &TreeNode{50, &TreeNode{25, &TreeNode{Value: 10}, &TreeNode{Value: 33}}, &TreeNode{Value: 75}}
}

func grow(values ...int) *TreeNode {
	var root *TreeNode
	for _, v := range values {
		root = Insert(root, v)
	}
	return root
}

func TestEx51_1(t *testing.T) {
	root := sampleTree()
	check.Eq(t, Search(root, 33).Value, 33)
	check.True(t, Search(root, 40) == nil && Search(nil, 1) == nil, "missing")
}

func TestEx51_2(t *testing.T) {
	r := grow(50, 25, 75, 10, 33)
	check.Eq(t, []int{r.Value, r.Left.Value, r.Right.Value, r.Left.Left.Value, r.Left.Right.Value}, []int{50, 25, 75, 10, 33})
}

func TestEx51_3(t *testing.T) {
	r := grow(50, 25, 75, 61, 89, 52)
	r = Delete(r, 50)
	check.Eq(t, r.Value, 52)
	check.Eq(t, Inorder(r), []int{25, 52, 61, 75, 89})
	r = Delete(Delete(Delete(r, 25), 89), 1000)
	check.Eq(t, Inorder(r), []int{52, 61, 75})
	rng := rand.New(rand.NewSource(1))
	for trial := 0; trial < 100; trial++ {
		vals := rng.Perm(200)[:30]
		tree := grow(vals...)
		for _, v := range vals[:15] {
			tree = Delete(tree, v)
		}
		check.Eq(t, Inorder(tree), slices.Sorted(slices.Values(vals[15:])))
	}
}

func TestEx51_4(t *testing.T) {
	root := sampleTree()
	check.Eq(t, Inorder(root), []int{10, 25, 33, 50, 75})
	check.Eq(t, Preorder(root), []int{50, 25, 10, 33, 75})
	check.Eq(t, Postorder(root), []int{10, 33, 25, 75, 50})
	check.Eq(t, len(Inorder(nil)), 0)
}

func TestEx51_5(t *testing.T) {
	root := sampleTree()
	lo, _ := FindMin(root)
	hi, _ := FindMax(root)
	_, ok := FindMin(nil)
	check.Eq(t, []any{lo, hi, ok}, []any{10, 75, false})
}

func TestEx51_6(t *testing.T) {
	check.Eq(t, []int{Height(grow(50, 25, 75, 10)), Height(grow(1, 2, 3, 4)), Height(nil)}, []int{3, 4, 0})
}

func TestEx51_7(t *testing.T) {
	root := BuildBST(50, 25, 75, 10, 33)
	check.Eq(t, []any{root.Value, Inorder(root)}, []any{50, []int{10, 25, 33, 50, 75}})
	check.Eq(t, TreeSort(5, 3, 8, 1, 4), []int{1, 3, 4, 5, 8})
}

func TestEx51_8(t *testing.T) {
	bad := &TreeNode{10, &TreeNode{5, nil, &TreeNode{Value: 12}}, &TreeNode{Value: 15}}
	check.True(t, !IsValidBST(bad) && IsValidBST(sampleTree()), "valid")
	check.True(t, IsValidBST(&TreeNode{0, &TreeNode{Value: math.MinInt}, &TreeNode{Value: math.MaxInt}}), "extremes")
	check.True(t, !IsValidBST(&TreeNode{5, &TreeNode{Value: 5}, nil}), "duplicate")
}

func TestEx51_9(t *testing.T) {
	check.Eq(t, LevelOrder(sampleTree()), [][]int{{50}, {25, 75}, {10, 33}})
	check.True(t, LevelOrder(nil) == nil, "empty levels")
}

func TestEx51_10(t *testing.T) {
	check.Eq(t, []int{CountNodes(sampleTree()), CountNodes(nil)}, []int{5, 0})
}

func TestEx51_11(t *testing.T) {
	root := grow(50, 25, 75, 10, 33)
	got := []int{}
	for _, target := range []int{30, 60, 29, 100, 0, 50} {
		got = append(got, ClosestValue(root, target))
	}
	check.Eq(t, got, []int{33, 50, 25, 75, 10, 50})
}

func TestEx51_12(t *testing.T) {
	root := BalancedBST([]int{1, 2, 3, 4, 5, 6, 7})
	check.Eq(t, []int{root.Value, root.Left.Value, root.Right.Value}, []int{4, 2, 6})
	check.Eq(t, Height(root), 3)
	ten := BalancedBST([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	check.Eq(t, Inorder(ten), []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	check.Eq(t, Height(ten), 4)
	check.True(t, BalancedBST(nil) == nil, "empty input")
}

func TestEx51_13(t *testing.T) {
	root := grow(50, 25, 75, 10, 33)
	got := []int{}
	for k := 1; k <= 5; k++ {
		got = append(got, KthLargest(root, k))
	}
	check.Eq(t, got, []int{75, 50, 33, 25, 10})
	check.Eq(t, KthLargest(&TreeNode{Value: 7}, 1), 7)
}

func TestEx51_14(t *testing.T) {
	root := BSTFromPreorder([]int{40, 20, 10, 30, 60, 70})
	check.Eq(t, []int{root.Value, root.Left.Value, root.Left.Left.Value, root.Left.Right.Value}, []int{40, 20, 10, 30})
	check.True(t, root.Right.Value == 60 && root.Right.Left == nil && root.Right.Right.Value == 70, "right side")
	for _, values := range [][]int{{40, 20, 10, 30, 60, 70}, {5, 4, 3, 2, 1}, {1, 2, 3}, {8, 3, 1, 6, 4, 7, 10, 14, 13}} {
		root := BSTFromPreorder(values)
		check.Eq(t, Preorder(root), values)
		sorted := slices.Clone(values)
		slices.Sort(sorted)
		check.Eq(t, Inorder(root), sorted)
	}
	check.True(t, BSTFromPreorder(nil) == nil, "empty input")
}

func TestEx51_15(t *testing.T) {
	root := grow(50, 25, 75, 10, 33)
	got := []int{}
	for _, v := range []int{10, 25, 33, 50, 40, 0} {
		s, ok := Successor(root, v)
		check.True(t, ok, "has a successor")
		got = append(got, s)
	}
	check.Eq(t, got, []int{25, 33, 50, 75, 50, 10})
	_, ok := Successor(root, 75)
	check.True(t, !ok, "largest value")
	_, ok = Successor(nil, 3)
	check.True(t, !ok, "empty tree")
}

func TestEx51_16(t *testing.T) {
	root := &TreeNode{50, &TreeNode{75, &TreeNode{Value: 10}, &TreeNode{Value: 33}}, &TreeNode{Value: 25}}
	RepairSwapped(root)
	check.Eq(t, []int{root.Left.Value, root.Right.Value}, []int{25, 75})
	for _, s := range [][2]int{{10, 33}, {50, 52}, {10, 89}, {61, 75}} {
		root := grow(50, 25, 75, 61, 89, 52, 10, 33)
		x, y := Search(root, s[0]), Search(root, s[1])
		x.Value, y.Value = s[1], s[0]
		RepairSwapped(root)
		check.Eq(t, Inorder(root), []int{10, 25, 33, 50, 52, 61, 75, 89})
		check.True(t, IsValidBST(root), "valid after repair")
	}
}
