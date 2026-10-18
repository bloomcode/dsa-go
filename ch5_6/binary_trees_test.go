package ch5_6

import (
	"math"
	"testing"

	"dsa/internal/check"
)

// X marks a missing child in tree's level-order values.
const X = math.MinInt

// tree builds a tree from level-order values, X for a missing child (children of a missing node are not listed).
func tree(values ...int) *TreeNode {
	if len(values) == 0 || values[0] == X {
		return nil
	}
	root := &TreeNode{Value: values[0]}
	queue, i := []*TreeNode{root}, 1
	for q := 0; q < len(queue) && i < len(values); q++ {
		for _, child := range []**TreeNode{&queue[q].Left, &queue[q].Right} {
			if i < len(values) && values[i] != X {
				*child = &TreeNode{Value: values[i]}
				queue = append(queue, *child)
			}
			i++
		}
	}
	return root
}

// sample is 1(2(4, 5(7)), 3(-, 6)), used by most tests.
func sample() *TreeNode { return tree(1, 2, 3, 4, 5, X, 6, X, X, 7) }

func inorder(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}
	return append(append(inorder(node.Left), node.Value), inorder(node.Right)...)
}

func preorder(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}
	return append(append([]int{node.Value}, preorder(node.Left)...), preorder(node.Right)...)
}

func TestEx5_6_1(t *testing.T) {
	root := sample()
	got := []bool{}
	for _, target := range []int{10, 15, 7, 8, 3} {
		got = append(got, HasPathSum(root, target))
	}
	check.Eq(t, got, []bool{true, true, true, false, false})
	check.True(t, !HasPathSum(nil, 0) && HasPathSum(&TreeNode{Value: 5}, 5), "empty tree and single leaf")
}

func TestEx5_6_2(t *testing.T) {
	got := []int{TotalDepth(sample()), TotalDepth(nil), TotalDepth(&TreeNode{Value: 1}), TotalDepth(tree(1, 2, X, 3)), TotalDepth(tree(1, 2, 3, 4, 5, 6, 7))}
	check.Eq(t, got, []int{11, 0, 0, 3, 10})
}

func TestEx5_6_3(t *testing.T) {
	root := sample()
	check.True(t, Mirror(root) == root, "returns the root")
	check.Eq(t, inorder(root), []int{6, 3, 1, 5, 7, 2, 4})
	check.Eq(t, root.Left.Value, 3)
	check.True(t, Mirror(nil) == nil, "empty tree")
}

func TestEx5_6_4(t *testing.T) {
	got := []bool{IsSymmetric(tree(1, 2, 2, 3, 4, 4, 3)), IsSymmetric(tree(1, 2, 2, X, 3, X, 3)), IsSymmetric(nil), IsSymmetric(&TreeNode{Value: 1}), IsSymmetric(tree(1, 2, 3)), IsSymmetric(sample())}
	check.Eq(t, got, []bool{true, false, true, true, false, false})
}

func TestEx5_6_5(t *testing.T) {
	check.Eq(t, RightView(sample()), []int{1, 3, 6, 7})
	check.Eq(t, RightView(tree(1, 2, 3, 4)), []int{1, 3, 4})
	check.Eq(t, RightView(nil), []int{})
}

func TestEx5_6_6(t *testing.T) {
	check.Eq(t, LevelAverages(sample()), []float64{1, 2.5, 5, 7})
	check.Eq(t, LevelAverages(nil), []float64{})
	check.Eq(t, LevelAverages(tree(3, -1, 2)), []float64{3, 0.5})
}

func TestEx5_6_7(t *testing.T) {
	check.Eq(t, LeafPaths(sample()), [][]int{{1, 2, 4}, {1, 2, 5, 7}, {1, 3, 6}})
	check.Eq(t, LeafPaths(nil), [][]int{})
	check.Eq(t, LeafPaths(&TreeNode{Value: 9}), [][]int{{9}})
}

func TestEx5_6_8(t *testing.T) {
	check.Eq(t, InorderIterative(sample()), []int{4, 2, 7, 5, 1, 3, 6})
	check.Eq(t, InorderIterative(nil), []int{})
	check.Eq(t, InorderIterative(tree(1, 2, X, 3, X, 4)), []int{4, 3, 2, 1})
}

func TestEx5_6_9(t *testing.T) {
	leaf := func(token string) *ExprNode { return &ExprNode{Token: token} }
	op := func(token string, l, r *ExprNode) *ExprNode { return &ExprNode{token, l, r} }
	got := []int{
		Evaluate(op("*", op("+", leaf("3"), leaf("4")), op("-", leaf("5"), leaf("2")))), Evaluate(leaf("7")),
		Evaluate(op("-", leaf("2"), leaf("-9"))), Evaluate(op("-", op("-", leaf("10"), leaf("4")), leaf("3"))),
	}
	check.Eq(t, got, []int{21, 7, 11, 3})
}

func TestEx5_6_10(t *testing.T) {
	got := []int{Diameter(sample()), Diameter(nil), Diameter(&TreeNode{Value: 1}), Diameter(tree(1, 2, 3)), Diameter(tree(1, 2, X, 3, 4, 5, X, X, 6))}
	check.Eq(t, got, []int{5, 0, 0, 2, 4})
}

func TestEx5_6_11(t *testing.T) {
	got := []bool{IsBalanced(sample()), IsBalanced(tree(1, 2, X, 3)), IsBalanced(nil), IsBalanced(tree(1, 2, 3, 4, X, X, X, 5)), IsBalanced(tree(1, 2, 3, 4))}
	check.Eq(t, got, []bool{true, false, true, false, true})
}

func TestEx5_6_12(t *testing.T) {
	got := []int{WidestLevel(sample()), WidestLevel(tree(1, 2, 3, 4, X, X, 5)), WidestLevel(nil), WidestLevel(tree(1, 2, X, 3)), WidestLevel(tree(1, 2, 3))}
	check.Eq(t, got, []int{4, 4, 0, 1, 2})
}

func TestEx5_6_13(t *testing.T) {
	root := sample()
	got := []int{}
	for _, p := range [][2]int{{4, 7}, {7, 6}, {5, 7}, {3, 6}, {4, 4}} {
		got = append(got, LowestCommonAncestor(root, p[0], p[1]).Value)
	}
	check.Eq(t, got, []int{2, 1, 5, 3, 4})
}

func TestEx5_6_14(t *testing.T) {
	text := "1,2,4,#,#,5,7,#,#,#,3,#,6,#,#"
	check.Eq(t, Serialize(sample()), text)
	check.Eq(t, Serialize(Parse(text)), text)
	check.Eq(t, inorder(Parse(text)), inorder(sample()))
	check.Eq(t, Serialize(nil), "#")
	check.True(t, Parse("#") == nil, "empty tree")
	check.Eq(t, Serialize(Parse(Serialize(tree(-1, 20)))), "-1,20,#,#,#")
}

func TestEx5_6_15(t *testing.T) {
	check.Eq(t, VerticalColumns(sample()), [][]int{{4}, {2, 7}, {1, 5}, {3}, {6}})
	check.Eq(t, VerticalColumns(nil), [][]int{})
	check.Eq(t, VerticalColumns(tree(1, 2, 3, X, 4, 5)), [][]int{{2}, {1, 4, 5}, {3}})
}

func TestEx5_6_16(t *testing.T) {
	root := sample()
	check.Eq(t, NodesAtDistance(root, 5, 2), []int{1, 4})
	check.Eq(t, NodesAtDistance(root, 5, 0), []int{5})
	check.Eq(t, NodesAtDistance(root, 4, 3), []int{3, 7})
	check.Eq(t, NodesAtDistance(root, 1, 2), []int{4, 5, 6})
	check.Eq(t, NodesAtDistance(root, 7, 5), []int{6})
	check.Eq(t, NodesAtDistance(root, 7, 9), []int{})
}

func TestEx5_6_17(t *testing.T) {
	got := []int{MaxPathSum(tree(2, -1, 3, 4, -6, X, 5)), MaxPathSum(tree(-5, -2, -7)), MaxPathSum(&TreeNode{Value: 7}), MaxPathSum(sample()), MaxPathSum(tree(1, -9, -9))}
	check.Eq(t, got, []int{13, -2, 7, 24, 1})
}

func TestEx5_6_18(t *testing.T) {
	root := Rebuild([]int{1, 2, 4, 5, 7, 3, 6}, []int{4, 2, 7, 5, 1, 3, 6})
	check.Eq(t, []int{root.Left.Right.Left.Value, root.Right.Right.Value}, []int{7, 6})
	check.True(t, Rebuild(nil, nil) == nil, "empty input")
	for _, tr := range []*TreeNode{sample(), tree(1, 2, X, 3), tree(5, X, 8, 6), {Value: 4}} {
		again := Rebuild(preorder(tr), inorder(tr))
		check.Eq(t, preorder(again), preorder(tr))
		check.Eq(t, inorder(again), inorder(tr))
	}
}
