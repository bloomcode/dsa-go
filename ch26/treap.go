// Package ch26: Treaps: balance from random priorities.
// A treap is a BST by key and a heap by a random priority (here a max-heap: a parent's
// priority is at least its children's). Functions take a node and return the new subtree root.
package ch26

import "math/rand/v2"

// ==== EX 26.1 | Treap nodes and rotations | Core
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

// ==== EX 26.2 | Insert, then rotate up | Core
// PROMPT: Write `Insert(n *TreapNode, key int, rng *rand.Rand) *TreapNode` that gives the new key `rng.Float64()` as its priority, inserts it as a BST leaf, and on the way back up rotates it above any parent with a smaller priority.
// EXAMPLE: `rng := rand.New(rand.NewPCG(1, 0)); var root *TreapNode; for k := range 10 { root = Insert(root, k, rng) }` -> a valid treap whose in-order walk is 0..9
// BIGO: O(log N) expected, whatever order the keys arrive in.
func Insert(n *TreapNode, key int, rng *rand.Rand) *TreapNode {
	panic("todo")
}

// ==== END

// ==== EX 26.3 | Search | Core
// PROMPT: Write `Contains(n *TreapNode, key int) bool`. Priorities play no part in search.
// EXAMPLE: after inserting `5, 2, 8` -> `Contains(root, 8)` is `true`, `Contains(root, 3)` is `false`
// BIGO: O(log N) expected.
func Contains(n *TreapNode, key int) bool {
	panic("todo")
}

// ==== END

// ==== EX 26.4 | Delete by rotating down | Core
// PROMPT: Write `Delete(n *TreapNode, key int) *TreapNode`: find the node, then rotate it down (always lifting the child with the larger priority) until it has at most one child, and splice it out.
// EXAMPLE: insert 0..99, delete the multiples of 3 -> in-order gives the other 66 keys and the treap is still valid
// BIGO: O(log N) expected.
func Delete(n *TreapNode, key int) *TreapNode {
	panic("todo")
}

// ==== END

// ==== EX 26.5 | Validate both orders at once | Core
// PROMPT: Write `IsValid(n *TreapNode, lo, hi int) bool` checking BST order on keys and heap order on priorities in one recursive pass, and `Inorder(n *TreapNode) []int`.
// EXAMPLE: any treap built by 26.2 with bounds `math.MinInt, math.MaxInt` -> `true`; swap two priorities by hand -> `false`
// BIGO: O(N).
func IsValid(n *TreapNode, lo, hi int) bool {
	panic("todo")
}

func Inorder(n *TreapNode) []int {
	panic("todo")
}

// ==== END

// ==== EX 26.6 | Random priorities do the balancing | Core
// PROMPT: Write `Height(keys []int, rng *rand.Rand) int` building a treap from keys inserted in the given order. Insert 1..1000 in sorted order and compare with a plain BST (23.6).
// EXAMPLE: sorted 1..1000 with `rand.New(rand.NewPCG(0, 0))` -> height around 20 to 30, versus 1000 for a plain BST
// BIGO: Expected height about 3 x log2 N, the same as a BST built from shuffled keys, so roughly 25 to 30 for N = 1000.
func Height(keys []int, rng *rand.Rand) int {
	panic("todo")
}

// ==== END

// ==== EX 26.7 | Split and merge | Stretch
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
