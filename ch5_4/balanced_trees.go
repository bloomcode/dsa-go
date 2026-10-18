// Package ch5_4: Self-balancing trees: red-black trees and treaps.
package ch5_4

import (
	"math/rand/v2"
)

// ==== EX 5.4.1 | Nodes, the NIL sentinel and rotations | Core
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

// ==== EX 5.4.2 | Insert with recolouring and rotations | Core
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

// ==== EX 5.4.3 | Check the red-black rules | Core
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

// ==== EX 5.4.4 | Search and in-order walk | Core
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

// ==== EX 5.4.5 | Height stays logarithmic | Core
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

// ==== EX 5.4.6 | Minimum, maximum and successor | Practice
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

// ==== EX 5.4.7 | Count the rotations | Practice
// PROMPT: Write `RotationsFor(keys []int) int` returning how many rotations a red-black tree performs while inserting keys in order. Compare sorted, reversed and shuffled input.
// EXAMPLE: `RotationsFor(0..999 in order)` -> fewer than 1000 rotations in total (amortized O(1) per insert)
// BIGO: At most 2 rotations per insert; amortized fewer.
func RotationsFor(keys []int) int {
	panic("todo")
}

// ==== END

// ==== EX 5.4.8 | Deletion | Stretch
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

// ==== EX 5.4.9 | Treap nodes and rotations | Core
// PROMPT: Create `type TreapNode struct` with `Key int`, `Priority float64`, `Left` and `Right`, and write `RotateRight(n *TreapNode) *TreapNode` and `RotateLeft` that return the new subtree root.
// EXAMPLE: rotating right at a node with left child L makes L the subtree root and the old node L's right child
// BIGO: O(1).
type TreapNode struct {
	Key         int
	Priority    float64
	Left, Right *TreapNode
}

func RotateRight(n *TreapNode) *TreapNode {
	panic("todo")
}

func RotateLeft(n *TreapNode) *TreapNode {
	panic("todo")
}

// ==== END

// ==== EX 5.4.10 | Insert, then rotate up | Core
// PROMPT: Write `Insert(n *TreapNode, key int, rng *rand.Rand) *TreapNode` that gives the new key `rng.Float64()` as its priority, inserts it as a BST leaf, and on the way back up rotates it above any parent with a smaller priority.
// EXAMPLE: `rng := rand.New(rand.NewPCG(1, 0)); var root *TreapNode; for k := range 10 { root = Insert(root, k, rng) }` -> a valid treap whose in-order walk is 0..9
// BIGO: O(log N) expected, whatever order the keys arrive in.
func Insert(n *TreapNode, key int, rng *rand.Rand) *TreapNode {
	panic("todo")
}

// ==== END

// ==== EX 5.4.11 | Search | Core
// PROMPT: Write `Contains(n *TreapNode, key int) bool`. Priorities play no part in search.
// EXAMPLE: after inserting `5, 2, 8` -> `Contains(root, 8)` is `true`, `Contains(root, 3)` is `false`
// BIGO: O(log N) expected.
func Contains(n *TreapNode, key int) bool {
	panic("todo")
}

// ==== END

// ==== EX 5.4.12 | Delete by rotating down | Core
// PROMPT: Write `Delete(n *TreapNode, key int) *TreapNode`: find the node, then rotate it down (always lifting the child with the larger priority) until it has at most one child, and splice it out.
// EXAMPLE: insert 0..99, delete the multiples of 3 -> in-order gives the other 66 keys and the treap is still valid
// BIGO: O(log N) expected.
func Delete(n *TreapNode, key int) *TreapNode {
	panic("todo")
}

// ==== END

// ==== EX 5.4.13 | Validate both orders at once | Core
// PROMPT: Write `IsValid(n *TreapNode, lo, hi int) bool` checking BST order on keys and heap order on priorities in one recursive pass, and `Inorder(n *TreapNode) []int`.
// EXAMPLE: any treap built by 5.4.10 with bounds `math.MinInt, math.MaxInt` -> `true`; swap two priorities by hand -> `false`
// BIGO: O(N).
func IsValid(n *TreapNode, lo, hi int) bool {
	panic("todo")
}

func Inorder(n *TreapNode) []int {
	panic("todo")
}

// ==== END

// ==== EX 5.4.14 | Random priorities do the balancing | Core
// PROMPT: Write `Height(keys []int, rng *rand.Rand) int` building a treap from keys inserted in the given order. Insert 1..1000 in sorted order and compare with a plain BST (7.4.6).
// EXAMPLE: sorted 1..1000 with `rand.New(rand.NewPCG(0, 0))` -> height around 20 to 30, versus 1000 for a plain BST
// BIGO: Expected height about 3 x log2 N, the same as a BST built from shuffled keys, so roughly 25 to 30 for N = 1000.
func Height(keys []int, rng *rand.Rand) int {
	panic("todo")
}

// ==== END

// ==== EX 5.4.15 | Split and merge | Stretch
// PROMPT: Write `Split(n *TreapNode, key int) (left, right *TreapNode)` returning treaps with keys < key and >= key, and `Merge(a, b *TreapNode) *TreapNode` joining treaps where every key in a is smaller than every key in b. Use them to delete a whole key range.
// EXAMPLE: split a treap of 0..9 at 4 -> in-orders `[0 1 2 3]` and `[4 5 6 7 8 9]`; merging them restores 0..9
// BIGO: O(log N) expected each. Many contest "implicit treap" tricks are built from these two operations.
func Split(n *TreapNode, key int) (left, right *TreapNode) {
	panic("todo")
}

func Merge(a, b *TreapNode) *TreapNode {
	panic("todo")
}

// ==== END
