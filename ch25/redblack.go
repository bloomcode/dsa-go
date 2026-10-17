// Package ch25: Red-black trees.
// A shared black sentinel nil node stands in for every empty child, as in most textbook
// implementations; it removes many nil checks from the fix-up code.
package ch25


// ==== EX 25.1 | Nodes, the NIL sentinel and rotations | Core
// PROMPT: Create `type RBNode struct` with `Key int`, `Red bool`, and `left`, `right`, `parent` pointers, then `type RedBlackTree struct` whose empty children all point at one black `NIL` sentinel node, with `rotateLeft(x)` and `rotateRight(x)` that keep parent links and the root correct.
// EXAMPLE: after a left rotation at x, x's old right child takes x's place and x becomes its left child; in-order order is unchanged
// BIGO: O(1) per rotation: three links change.
// NOTE: A rotation never breaks binary-search order. It only changes the shape, which is all balancing needs. Compare nodes with `==` on pointers: the sentinel is one shared node.
type RBNode struct {
	Key                 int
	Red                 bool
	left, right, parent *RBNode
}

type RedBlackTree struct {
	NIL, root       *RBNode
	Size, Rotations int
}

func NewRedBlackTree() *RedBlackTree {
	panic("todo")
}

func (t *RedBlackTree) rotateLeft(x *RBNode) {
	panic("todo")
}

func (t *RedBlackTree) rotateRight(x *RBNode) {
	panic("todo")
}

// ==== END

// ==== EX 25.2 | Insert with recolouring and rotations | Core
// PROMPT: Add `(t *RedBlackTree) Insert(key int) bool`: do a normal BST insert of a red node, then fix up. While the parent is red: if the uncle is red, recolour parent, uncle and grandparent and move up; otherwise rotate (once or twice) and recolour. Finally paint the root black.
// EXAMPLE: insert 1..7 in sorted order -> root `2`, height `4` (a plain BST would be 7 tall), and `Validate()` passes after every insert
// BIGO: O(log N): the walk down is O(log N), fix-up recolours up the tree and does at most two rotations.
func (t *RedBlackTree) Insert(key int) bool {
	panic("todo")
}

func (t *RedBlackTree) insertFixup(z *RBNode) {
	panic("todo")
}

// ==== END

// ==== EX 25.3 | Check the red-black rules | Core
// PROMPT: Add `(t *RedBlackTree) Validate() (int, error)` returning the tree's black height, or an error if the root is red, a red node has a red child, BST order is broken, or two paths have different black counts.
// EXAMPLE: after any sequence of inserts and deletes, `Validate()` returns a positive number and a nil error
// BIGO: O(N). Write this before Insert: it is the test that tells you whether your fix-up works.
func (t *RedBlackTree) Validate() (int, error) {
	panic("todo")
}

func (t *RedBlackTree) check(n *RBNode, lo, hi int) (int, error) {
	panic("todo")
}

// ==== END

// ==== EX 25.4 | Search and in-order walk | Core
// PROMPT: Add `Contains(key int) bool` and `Inorder() []int` (iterative, with an explicit stack). Both ignore colour entirely: a red-black tree is still a BST.
// EXAMPLE: after inserting `5, 2, 8, 1` -> `Contains(8)` is `true`; `Inorder()` is `[1 2 5 8]`
// BIGO: Search O(log N) guaranteed; in-order O(N).
func (t *RedBlackTree) Contains(key int) bool {
	panic("todo")
}

func (t *RedBlackTree) find(key int) *RBNode {
	panic("todo")
}

func (t *RedBlackTree) Inorder() []int {
	panic("todo")
}

// ==== END

// ==== EX 25.5 | Height stays logarithmic | Core
// PROMPT: Add `Height() int` and `HeightBound() float64` = 2 * log2(N + 1), and check the height never exceeds the bound, even after inserting 1..N in sorted order (which makes a plain BST N tall).
// EXAMPLE: insert 1..1000 in order -> `Height()` is at most `19`, while a plain BST would be `1000` tall
// BIGO: O(N) to measure. The bound follows from the rules: no two reds in a row, and equal black counts on every path.
func (t *RedBlackTree) Height() int { panic("todo") }

func (t *RedBlackTree) h(n *RBNode) int {
	panic("todo")
}

func (t *RedBlackTree) HeightBound() float64 {
	panic("todo")
}

// ==== END

// ==== EX 25.6 | Minimum, maximum and successor | Practice
// PROMPT: Add `minimum(n *RBNode) *RBNode`, `Maximum() int` and `Successor(key int) (int, bool)` (the smallest key greater than key, with ok false if none), using parent links for the successor.
// EXAMPLE: keys `20, 10, 30, 25` -> `minimum(root).Key` `10`, `Maximum()` `30`, `Successor(20)` `25, true`, `Successor(30)` `0, false`
// BIGO: O(log N) each.
func (t *RedBlackTree) minimum(n *RBNode) *RBNode {
	panic("todo")
}

func (t *RedBlackTree) Maximum() int {
	panic("todo")
}

func (t *RedBlackTree) Successor(key int) (int, bool) {
	panic("todo")
}

// ==== END

// ==== EX 25.7 | Count the rotations | Practice
// PROMPT: Write `RotationsFor(keys []int) int` returning how many rotations a red-black tree performs while inserting keys in order. Compare sorted, reversed and shuffled input.
// EXAMPLE: `RotationsFor(0..999 in order)` -> fewer than 1000 rotations in total (amortized O(1) per insert)
// BIGO: At most 2 rotations per insert; amortized fewer.
func RotationsFor(keys []int) int {
	panic("todo")
}

// ==== END

// ==== EX 25.8 | Deletion | Stretch
// PROMPT: Add `Delete(key int) bool`. Splice out the node (or its successor), remember the colour removed, and if a black node was removed run the four-case "double black" fix-up with rotations and recolouring. Return `false` if the key is absent.
// EXAMPLE: insert 1..100, delete every even key -> `Validate()` passes and `Inorder()` is the odd keys
// BIGO: O(log N), with at most three rotations.
// NOTE: The fix-up may set `NIL.parent` temporarily; that is why the sentinel is a real node rather than a nil pointer.
func (t *RedBlackTree) transplant(u, v *RBNode) {
	panic("todo")
}

func (t *RedBlackTree) Delete(key int) bool {
	panic("todo")
}

func (t *RedBlackTree) deleteFixup(x *RBNode) {
	panic("todo")
}

// ==== END
