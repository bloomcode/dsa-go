// Package ch56: Binary trees.
package ch56


// TreeNode is a binary tree node; unlike a BST (Chapter 51), the values follow no order.
type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

// ==== EX 56.1 | Is there a path with this sum? | Core
// PROMPT: Write `HasPathSum(root *TreeNode, target int) bool`: `true` when some root-to-leaf path has values summing to `target`. Recurse with `target - root.Value`; only a leaf can complete a path (an empty tree has none).
// EXAMPLE: `1(2(4, 5(7)), 3(-, 6))`, `10` -> `true` (1 + 3 + 6); `8` -> `false`
// BIGO: O(N) time, O(H) stack: each call passes down what is left of the target, and only a leaf can finish a path.
// NOTE: Examples write a tree as `value(left, right)`, with `-` for a missing child: `5(7)` is a 5 whose left child is 7.
func HasPathSum(root *TreeNode, target int) bool {
	panic("todo")
}

// ==== END

// ==== EX 56.2 | Total depth | Core
// PROMPT: Write `TotalDepth(root *TreeNode) int` that returns the sum of every node's depth (the root has depth 0). Pass the depth down as you recurse.
// EXAMPLE: `1(2(4, 5(7)), 3(-, 6))` -> `11`
// BIGO: O(N) time, O(H) stack: each node adds the depth it is called with.
func TotalDepth(root *TreeNode) int {
	panic("todo")
}

// ==== END

// ==== EX 56.3 | Mirror a tree | Core
// PROMPT: Write `Mirror(root *TreeNode) *TreeNode` that swaps the left and right children of every node, in place, and returns `root`.
// EXAMPLE: `1(2(4, 5(7)), 3(-, 6))` -> `1(3(6, -), 2(5(-, 7), 4))`
// BIGO: O(N) time, O(H) stack: every node swaps its two children once.
func Mirror(root *TreeNode) *TreeNode {
	panic("todo")
}

// ==== END

// ==== EX 56.4 | Is the tree symmetric? | Core
// PROMPT: Write `IsSymmetric(root *TreeNode) bool`: `true` when the tree is its own mirror image (an empty tree is). Compare two subtrees at once: their values must match, and the left of one must mirror the right of the other.
// EXAMPLE: `1(2(3, 4), 2(4, 3))` -> `true`; `1(2(-, 3), 2(-, 3))` -> `false`
// BIGO: O(N) time, O(H) stack: compare the two halves in step, each outer child against the opposite outer child.
func IsSymmetric(root *TreeNode) bool {
	panic("todo")
}

// ==== END

// ==== EX 56.5 | The view from the right | Core
// PROMPT: Write `RightView(root *TreeNode) []int` that returns the values seen looking at the tree from the right (a non-nil slice): the last node of each level, top to bottom. Walk level by level, building the next level from the children of the current one.
// EXAMPLE: `1(2(4, 5(7)), 3(-, 6))` -> `[1 3 6 7]`
// BIGO: O(N) time, O(W) space for the widest level W: a level-by-level walk, keeping the last node of each level.
func RightView(root *TreeNode) []int {
	panic("todo")
}

// ==== END

// ==== EX 56.6 | Average of each level | Core
// PROMPT: Write `LevelAverages(root *TreeNode) []float64` that returns the average value of each level, top to bottom (a non-nil slice), walking the tree level by level.
// EXAMPLE: `1(2(4, 5(7)), 3(-, 6))` -> `[1.0 2.5 5.0 7.0]`
// BIGO: O(N) time, O(W) space for the widest level W: one level-by-level walk.
func LevelAverages(root *TreeNode) []float64 {
	panic("todo")
}

// ==== END

// ==== EX 56.7 | Every root-to-leaf path | Core
// PROMPT: Write `LeafPaths(root *TreeNode) [][]int` that returns every root-to-leaf path as a slice of values, leftmost path first (a non-nil slice). Keep one path slice: append on the way down, record a copy at a leaf, and shrink it on the way back up.
// EXAMPLE: `1(2(4, 5(7)), 3(-, 6))` -> `[[1 2 4] [1 2 5 7] [1 3 6]]`
// BIGO: O(N * H) time and space for the copies, O(H) for the path being built: one path list grows and shrinks as the walk goes down and back up.
func LeafPaths(root *TreeNode) [][]int {
	panic("todo")
}

// ==== END

// ==== EX 56.8 | In-order walk without recursion | Core
// PROMPT: Write `InorderIterative(root *TreeNode) []int` that returns the in-order values (left, node, right) as a non-nil slice, using an explicit stack instead of recursion: push the whole left spine, pop and visit a node, then continue from its right child.
// EXAMPLE: `1(2(4, 5(7)), 3(-, 6))` -> `[4 2 7 5 1 3 6]`
// BIGO: O(N) time, O(H) space: the stack holds the left spine still waiting to be visited, which is what the call stack held in the recursive version.
func InorderIterative(root *TreeNode) []int {
	panic("todo")
}

// ==== END

// ==== EX 56.9 | Evaluate an expression tree | Core
// PROMPT: Create `type ExprNode struct{ Token string; Left, Right *ExprNode }`: a leaf holds an integer as text (`"-9"`), any other node holds `"+"`, `"-"` or `"*"` and has two children. Write `Evaluate(node *ExprNode) int` that returns the value of the expression.
// EXAMPLE: `*(+(3, 4), -(5, 2))` -> `21`
// BIGO: O(N) time, O(H) stack: a post-order walk, because an operator needs both operands first.
type ExprNode struct {
	Token       string
	Left, Right *ExprNode
}

func Evaluate(node *ExprNode) int {
	panic("todo")
}

// ==== END

// ==== EX 56.10 | Diameter | Practice
// PROMPT: Write `Diameter(root *TreeNode) int` that returns the number of edges on the longest path between any two nodes. In one walk, return each subtree's height and update the best `left height + right height` seen at any node.
// EXAMPLE: `1(2(4, 5(7)), 3(-, 6))` -> `5` (7 to 6)
// BIGO: O(N) time, O(H) stack: one post-order walk returns each subtree's height and, on the side, tries the path that bends at that node.
func Diameter(root *TreeNode) int {
	panic("todo")
}

// ==== END

// ==== EX 56.11 | Is the tree height-balanced? | Practice
// PROMPT: Write `IsBalanced(root *TreeNode) bool`: `true` when at every node the heights of the two subtrees differ by at most one. Do it in O(N): a helper returns a subtree's height, or `-1` once any part of it is unbalanced.
// EXAMPLE: `1(2(4, 5(7)), 3(-, 6))` -> `true`; `1(2(3), -)` -> `false`
// BIGO: O(N) time, O(H) stack: each call returns its height, or -1 as soon as any subtree is out of balance, so no height is computed twice.
func IsBalanced(root *TreeNode) bool {
	panic("todo")
}

// ==== END

// ==== EX 56.12 | Widest level | Practice
// PROMPT: The width of a level is the distance from its leftmost to its rightmost node, counting the missing places between them. Write `WidestLevel(root *TreeNode) int` that returns the largest width (`0` for an empty tree). Walk level by level with heap-style positions, renumbering each level from 0 so the numbers stay small.
// EXAMPLE: `1(2(4, 5(7)), 3(-, 6))` -> `4` (the level `4, 5, -, 6`)
// BIGO: O(N) time, O(W) space: number positions like a heap (children of `p` at `2p` and `2p + 1`), and a level's width is its last position minus its first, plus one.
func WidestLevel(root *TreeNode) int {
	panic("todo")
}

// ==== END

// ==== EX 56.13 | Lowest common ancestor | Practice
// PROMPT: Values are distinct, and `a` and `b` are both in the tree. Write `LowestCommonAncestor(root *TreeNode, a, b int) *TreeNode` that returns the deepest node having both values in its subtree (a node counts as its own descendant). Return the node holding `a` or `b` when you reach it; a node that gets non-nil answers from both sides is the ancestor.
// EXAMPLE: `1(2(4, 5(7)), 3(-, 6))`, `4, 7` -> the node `2`; `7, 6` -> the node `1`
// BIGO: O(N) time, O(H) stack: a node whose two subtrees each report a match is the answer, and it is the only one that sees both.
func LowestCommonAncestor(root *TreeNode, a, b int) *TreeNode {
	panic("todo")
}

// ==== END

// ==== EX 56.14 | Turn a tree into text and back | Practice
// PROMPT: Write `Serialize(root *TreeNode) string`, which lists the values in pre-order separated by commas, with `#` for each missing child, and `Parse(text string) *TreeNode`, which rebuilds the tree from that text. Parse by reading tokens in order: `#` is an empty subtree; anything else is a node whose left and then right subtrees follow.
// EXAMPLE: `1(2(4, 5(7)), 3(-, 6))` -> `"1,2,4,#,#,5,7,#,#,#,3,#,6,#,#"`
// BIGO: O(N) time and space both ways: pre-order with a marker for every missing child records the shape exactly, so parsing never has to guess.
func Serialize(root *TreeNode) string {
	panic("todo")
}

func Parse(text string) *TreeNode {
	panic("todo")
}

// ==== END

// ==== EX 56.15 | Vertical columns | Practice
// PROMPT: Give the root column 0, a left child its parent's column minus 1 and a right child plus 1. Write `VerticalColumns(root *TreeNode) [][]int` that returns the values column by column from left to right (a non-nil slice), each column top to bottom, and left to right within a level. Walk breadth-first, carrying each node's column.
// EXAMPLE: `1(2(4, 5(7)), 3(-, 6))` -> `[[4] [2 7] [1 5] [3] [6]]`
// BIGO: O(N log N) time for sorting the column numbers (O(N) with a running minimum), O(N) space: a breadth-first walk lists each column top to bottom.
func VerticalColumns(root *TreeNode) [][]int {
	panic("todo")
}

// ==== END

// ==== EX 56.16 | Nodes at distance k | Stretch
// PROMPT: Values are distinct and `target` is in the tree. Write `NodesAtDistance(root *TreeNode, target, k int) []int` that returns, in ascending order, the values of every node exactly `k` edges from the node holding `target`, whether up, down or across (a non-nil slice). Map each node to its parent, then walk outward breadth-first from the target, `k` rounds, never revisiting a node.
// EXAMPLE: `1(2(4, 5(7)), 3(-, 6))`, `5, 2` -> `[1 4]`
// BIGO: O(N) time and space: record every node's parent, and the tree becomes an undirected graph that a breadth-first search can walk k steps out.
func NodesAtDistance(root *TreeNode, target, k int) []int {
	panic("todo")
}

// ==== END

// ==== EX 56.17 | Largest path sum | Stretch
// PROMPT: Values may be negative. A path is any chain of connected nodes, at least one, each used once. Write `MaxPathSum(root *TreeNode) int` that returns the largest path sum in a non-empty tree. A helper returns the best sum of a path going down from a node (dropping a negative branch), and records `node + left + right` as a candidate.
// EXAMPLE: `2(-1(4, -6), 3(-, 5))` -> `13` (4, -1, 2, 3, 5)
// BIGO: O(N) time, O(H) stack: each call returns the best downward path from its node, and the best bent path through it is tried on the side.
func MaxPathSum(root *TreeNode) int {
	panic("todo")
}

// ==== END

// ==== EX 56.18 | Rebuild from pre-order and in-order | Stretch
// PROMPT: Write `Rebuild(preorder, inorder []int) *TreeNode` that returns the tree with these two traversals (distinct values; `nil` when empty). Take roots from `preorder` in order; each root's position in `inorder`, found through a map, splits the remaining range into its left and right subtrees.
// EXAMPLE: `[1 2 4 5 7 3 6]`, `[4 2 7 5 1 3 6]` -> `1(2(4, 5(7)), 3(-, 6))`
// BIGO: O(N) time and space: the next pre-order value is the root of the current range, and a map from value to in-order index splits the range at once.
func Rebuild(preorder, inorder []int) *TreeNode {
	panic("todo")
}

// ==== END
