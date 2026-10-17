// Package ch51: Binary search trees.
package ch51

// ==== EX 51.1 | Tree node and search | Core
// PROMPT: Create `type TreeNode struct{ Value int; Left, Right *TreeNode }`. Write `Search(node *TreeNode, value int) *TreeNode` returning the node holding value, or nil: go left if smaller, right if larger.
// EXAMPLE: Tree from `[50 25 75 10 33]` -> `Search(root, 33).Value` is `33`; `Search(root, 40)` is `nil`
// BIGO: O(log N) in a balanced tree; O(N) in a degenerate one (sorted inserts make a linked list).
type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

func Search(node *TreeNode, value int) *TreeNode {
	panic("todo")
}

// ==== END

// ==== EX 51.2 | Insert | Core
// PROMPT: Write `Insert(node *TreeNode, value int) *TreeNode` that returns the (possibly new) root. Walk down as in search; attach a new leaf at the first empty child. Ignore duplicates.
// EXAMPLE: `var root *TreeNode; for _, v := range []int{50, 25, 75} { root = Insert(root, v) }` -> root 50, left 25, right 75
// BIGO: O(log N) balanced, O(N) worst case.
func Insert(node *TreeNode, value int) *TreeNode {
	panic("todo")
}

// ==== END

// ==== EX 51.3 | Delete (all three cases) | Core
// PROMPT: Write `Delete(node *TreeNode, value int) *TreeNode` returning the new subtree root. Handle: no children (remove it), one child (the child takes its place), two children (replace the value with its successor, the smallest value in the right subtree, then delete the successor from there).
// EXAMPLE: Deleting 50 from `[50 25 75 61 89 52]` makes 52 the new root value
// BIGO: O(log N) balanced.
func Delete(node *TreeNode, value int) *TreeNode {
	panic("todo")
}

// ==== END

// ==== EX 51.4 | Traversals | Core
// PROMPT: Write `Inorder`, `Preorder` and `Postorder`, each taking a root and returning a slice of values, by passing a visit closure to one recursive walker.
// EXAMPLE: Tree from `[50 25 75 10 33]` -> inorder `[10 25 33 50 75]`, preorder `[50 25 10 33 75]`, postorder `[10 33 25 75 50]`
// BIGO: O(N): every node is visited once.
// NOTE: In-order traversal of a BST always yields sorted values.
func walk(node *TreeNode, order int, visit func(int)) {
	panic("todo")
}

func collect(node *TreeNode, order int) []int {
	panic("todo")
}

func Preorder(node *TreeNode) []int  { panic("todo") }
func Inorder(node *TreeNode) []int   { panic("todo") }
func Postorder(node *TreeNode) []int { panic("todo") }

// ==== END

// ==== EX 51.5 | Minimum and maximum | Core
// PROMPT: Write `FindMin(node *TreeNode) (int, bool)` and `FindMax` by walking only left or only right.
// EXAMPLE: Tree from `[50 25 75 10 33]` -> `10` and `75`
// BIGO: O(log N) balanced (the height of the tree).
func FindMin(node *TreeNode) (int, bool) {
	panic("todo")
}

func FindMax(node *TreeNode) (int, bool) {
	panic("todo")
}

// ==== END

// ==== EX 51.6 | Height | Core
// PROMPT: Write `Height(node *TreeNode) int`: the number of nodes on the longest root-to-leaf path (empty tree = 0).
// EXAMPLE: Tree from `[50 25 75 10]` -> `3`; from sorted `[1 2 3 4]` -> `4`
// BIGO: O(N). Compare log2(N) to the height to see how balanced a tree is.
func Height(node *TreeNode) int {
	panic("todo")
}

// ==== END

// ==== EX 51.7 | Build a tree and tree-sort | Practice
// PROMPT: Write `BuildBST(values ...int) *TreeNode` and `TreeSort(values ...int) []int` = build, then in-order. Try sorted input and look at `Height`.
// EXAMPLE: `TreeSort(5, 3, 8, 1, 4)` -> `[1 3 4 5 8]`
// BIGO: O(N log N) on random input, O(N^2) on sorted input (the tree degenerates).
func BuildBST(values ...int) *TreeNode {
	panic("todo")
}

func TreeSort(values ...int) []int { panic("todo") }

// ==== END

// ==== EX 51.8 | Validate a BST | Practice
// PROMPT: Write `IsValidBST(node *TreeNode) bool`. Checking only each node's direct children is not enough: pass the allowed range down to a helper as two `*int` bounds, where nil means "no bound".
// EXAMPLE: `&TreeNode{10, &TreeNode{5, nil, &TreeNode{Value: 12}}, &TreeNode{Value: 15}}` -> `false` (12 sits left of 10)
// BIGO: O(N).
// NOTE: Pointer bounds avoid picking sentinel numbers, so a tree may hold `math.MinInt` and `math.MaxInt` themselves.
func IsValidBST(node *TreeNode) bool {
	panic("todo")
}

func valid(node *TreeNode, lo, hi *int) bool {
	panic("todo")
}

// ==== END

// ==== EX 51.9 | Level-order traversal | Practice
// PROMPT: Write `LevelOrder(root *TreeNode) [][]int` returning a slice of levels, processing one level at a time with a `current` and a `next` slice of nodes.
// EXAMPLE: Tree from `[50 25 75 10 33]` -> `[[50] [25 75] [10 33]]`
// BIGO: O(N) time, O(W) space for the widest level.
func LevelOrder(root *TreeNode) [][]int {
	panic("todo")
}

// ==== END

// ==== EX 51.10 | Count the nodes | Core
// PROMPT: Write `CountNodes(node *TreeNode) int` recursively.
// EXAMPLE: Tree from `[50 25 75 10 33]` -> `5`
// BIGO: O(N).
func CountNodes(node *TreeNode) int {
	panic("todo")
}

// ==== END
