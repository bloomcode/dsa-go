// Package ch27: B-trees and counting I/Os.
// The B-tree follows the textbook "minimum degree t" layout: every node except the root
// holds between t - 1 and 2t - 1 keys, and an internal node with k keys has k + 1 children.
package ch27


// ==== EX 27.1 | Count I/Os, not steps | Core
// PROMPT: Memory moves in blocks of B items. Write `ScanIOs(n, B int) int` for reading n items in order, and `RandomReadsIOs(k int) int` for k reads that each land in a different block.
// EXAMPLE: `ScanIOs(1000, 64)` -> `16`; `RandomReadsIOs(1000)` -> `1000`
// BIGO: A scan costs O(N / B) I/Os; scattered access costs one I/O per item. On disk, I/Os dominate the running time.
func ScanIOs(n, B int) int { panic("todo") }

func RandomReadsIOs(k int) int { panic("todo") }

// ==== END

// ==== EX 27.2 | Binary search over blocks | Core
// PROMPT: Write `type ExternalArray struct` whose `ReadBlock(i int) []int` returns a copy of block i and counts one I/O, and `ExternalBinarySearch(ext *ExternalArray, target int) int` that binary-searches the blocks by their first item, then searches inside one block in memory. Return the index or -1.
// EXAMPLE: 10,000 sorted items with B = 100 -> a lookup costs 7 or 8 I/Os, versus about 14 for plain binary search touching one block per probe
// BIGO: O(log(N / B)) I/Os, a small saving; the B-tree (27.4) does much better with O(log_B N).
type ExternalArray struct {
	data      []int
	B, Blocks int
	IOs       int
}

func NewExternalArray(data []int, B int) *ExternalArray {
	panic("todo")
}

func (e *ExternalArray) ReadBlock(i int) []int {
	panic("todo")
}

func ExternalBinarySearch(ext *ExternalArray, target int) int {
	panic("todo")
}

// ==== END

// ==== EX 27.3 | Lookups: BST versus B-tree | Core
// PROMPT: Write `BSTIOs(n int) int` (one I/O per level of a balanced binary tree, log2 N) and `BTreeIOs(n, B int) int` (log base B of N) and compare them for a billion keys.
// EXAMPLE: `BSTIOs(1_000_000_000)` -> `30`; `BTreeIOs(1_000_000_000, 1000)` -> `3`
// BIGO: O(log_2 N) versus O(log_B N) I/Os: a factor of log2 B fewer.
// NOTE: Go has `math.Log2` but no log-to-any-base: use `math.Log(n) / math.Log(B)`, and subtract a tiny epsilon before `math.Ceil` so an exact power such as 1000^3 does not round up to 4.
func BSTIOs(n int) int { panic("todo") }

func BTreeIOs(n, B int) int {
	panic("todo")
}

// ==== END

// ==== EX 27.4 | B-tree nodes and search | Core
// PROMPT: Create `type BTreeNode struct` with `leaf bool`, sorted `keys []int` and `children []*BTreeNode`, and `type BTree struct` with minimum degree `t` and a `Search(key int) bool` that binary-searches each node's keys with `slices.BinarySearch` and descends into the right child.
// EXAMPLE: in a B-tree holding 1..100 -> `Search(42)` is `true`, `Search(101)` is `false`
// BIGO: O(log_t N) nodes visited, O(log t) comparisons in each.
// NOTE: `slices.BinarySearch` returns the first index whose key is >= key, plus whether it is an exact match: exactly the bisect-left a B-tree needs.
type BTreeNode struct {
	leaf     bool
	keys     []int
	children []*BTreeNode
}

type BTree struct {
	t    int
	root *BTreeNode
}

func NewBTree(t int) *BTree { panic("todo") }

func (b *BTree) Search(key int) bool {
	panic("todo")
}

// ==== END

// ==== EX 27.5 | Insert with proactive splits | Core
// PROMPT: Add `(b *BTree) Insert(key int) bool`: if the root is full (2t - 1 keys), split it and grow a new root. Walk down, splitting any full child before entering it, so a leaf always has room.
// EXAMPLE: t = 2, insert 1..10 -> `Validate()` passes and every leaf sits at the same depth
// BIGO: O(t log_t N) CPU, O(log_t N) node reads. The tree grows only at the root, which keeps it perfectly balanced.
// NOTE: `full.keys[t:]` and `full.keys[:t-1]` share one backing array, so a later append to the left half would overwrite the right half's keys. Clone the halves when splitting.
func (b *BTree) splitChild(parent *BTreeNode, i int) {
	panic("todo")
}

func (b *BTree) Insert(key int) bool {
	panic("todo")
}

// ==== END

// ==== EX 27.6 | Validate and walk in order | Core
// PROMPT: Add `Inorder() []int` returning all keys sorted, and `Validate() (int, error)` that checks key counts per node, sorted keys, child counts, and that all leaves share one depth. Return the height.
// EXAMPLE: after random inserts and deletes -> `Inorder()` equals the sorted set of keys and `Validate()` returns the height with a nil error
// BIGO: O(N).
func (b *BTree) Inorder() []int {
	panic("todo")
}

func (b *BTree) Validate() (int, error) {
	panic("todo")
}

// ==== END

// ==== EX 27.7 | How tall can a B-tree get? | Core
// PROMPT: Write `BTreeHeightBound(n, t int) int` returning the most levels a B-tree of minimum degree t with n keys can have: 1 + floor(log_t((n + 1) / 2)). Check it against real trees.
// EXAMPLE: `BTreeHeightBound(1_000_000, 50)` -> `4`; a real tree of 1..10,000 with t = 3 stays within its bound of 8
// BIGO: O(1). With t in the hundreds, a billion keys fit in 4 or 5 levels.
func BTreeHeightBound(n, t int) int {
	panic("todo")
}

// ==== END

// ==== EX 27.8 | Range queries, like a database index | Practice
// PROMPT: Add `RangeQuery(lo, hi int) []int` returning the keys in [lo, hi] in order, skipping whole subtrees that cannot contain them.
// EXAMPLE: B-tree of 1..100 -> `RangeQuery(40, 45)` is `[40 41 42 43 44 45]`
// BIGO: O(log_t N + K / t) node visits for K results. This is why databases index with B-trees.
func (b *BTree) RangeQuery(lo, hi int) []int {
	panic("todo")
}

// ==== END

// ==== EX 27.9 | Deletion | Stretch
// PROMPT: Add `Delete(key int) bool`. Before descending into a child with only t - 1 keys, top it up by borrowing from a sibling or merging with one. Delete from a leaf directly; for an internal key, replace it with its predecessor or successor, or merge. Shrink the root when it empties.
// EXAMPLE: t = 3, insert 1..200, delete the odd keys in shuffled order -> `Validate()` passes and `Inorder()` is the even keys
// BIGO: O(log_t N) node visits.
// NOTE: `slices.Delete(s, i, i+1)` removes one item and returns the shortened slice; forgetting to assign the result back is the usual bug.
func (b *BTree) Delete(key int) bool {
	panic("todo")
}

func (b *BTree) del(n *BTreeNode, key int) {
	panic("todo")
}

func merge(n *BTreeNode, i int) {
	panic("todo")
}

// ==== END
