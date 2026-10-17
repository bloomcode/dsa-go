package ch15

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

func TestEx15_1(t *testing.T) {
	root := sampleTree()
	check.Eq(t, Search(root, 33).Value, 33)
	check.True(t, Search(root, 40) == nil && Search(nil, 1) == nil, "missing")
}

func TestEx15_2(t *testing.T) {
	r := grow(50, 25, 75, 10, 33)
	check.Eq(t, []int{r.Value, r.Left.Value, r.Right.Value, r.Left.Left.Value, r.Left.Right.Value}, []int{50, 25, 75, 10, 33})
}

func TestEx15_3(t *testing.T) {
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

func TestEx15_4(t *testing.T) {
	root := sampleTree()
	check.Eq(t, Inorder(root), []int{10, 25, 33, 50, 75})
	check.Eq(t, Preorder(root), []int{50, 25, 10, 33, 75})
	check.Eq(t, Postorder(root), []int{10, 33, 25, 75, 50})
	check.Eq(t, len(Inorder(nil)), 0)
}

func TestEx15_5(t *testing.T) {
	root := sampleTree()
	lo, _ := FindMin(root)
	hi, _ := FindMax(root)
	_, ok := FindMin(nil)
	check.Eq(t, []any{lo, hi, ok}, []any{10, 75, false})
}

func TestEx15_6(t *testing.T) {
	check.Eq(t, []int{Height(grow(50, 25, 75, 10)), Height(grow(1, 2, 3, 4)), Height(nil)}, []int{3, 4, 0})
}

func TestEx15_7(t *testing.T) {
	root := BuildBST(50, 25, 75, 10, 33)
	check.Eq(t, []any{root.Value, Inorder(root)}, []any{50, []int{10, 25, 33, 50, 75}})
	check.Eq(t, TreeSort(5, 3, 8, 1, 4), []int{1, 3, 4, 5, 8})
}

func TestEx15_8(t *testing.T) {
	bad := &TreeNode{10, &TreeNode{5, nil, &TreeNode{Value: 12}}, &TreeNode{Value: 15}}
	check.True(t, !IsValidBST(bad) && IsValidBST(sampleTree()), "valid")
	check.True(t, IsValidBST(&TreeNode{0, &TreeNode{Value: math.MinInt}, &TreeNode{Value: math.MaxInt}}), "extremes")
	check.True(t, !IsValidBST(&TreeNode{5, &TreeNode{Value: 5}, nil}), "duplicate")
}

func TestEx15_9(t *testing.T) {
	check.Eq(t, LevelOrder(sampleTree()), [][]int{{50}, {25, 75}, {10, 33}})
	check.True(t, LevelOrder(nil) == nil, "empty levels")
}

func TestEx15_10(t *testing.T) {
	check.Eq(t, []int{CountNodes(sampleTree()), CountNodes(nil)}, []int{5, 0})
}
