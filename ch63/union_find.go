// Package ch63: Union-find and spanning trees.
package ch63


// ==== EX 63.1 | Union-find | Core
// PROMPT: Build `type DisjointSet struct` over elements `0` to `n - 1` with `NewDisjointSet(n int) *DisjointSet`, `Find(x int) int` (the root of x's group, pointing every node on the way straight at the root), `Union(a, b int) bool` (merge two groups, hanging the smaller under the larger; `false` if already together), `Size(x int) int` and a `Groups` count.
// EXAMPLE: `NewDisjointSet(5)`: `Union(0, 1)` -> `true`; `Union(1, 0)` -> `false`; `Size(1)` -> `2`; `Groups` is `4`
// BIGO: O(alpha(N)) amortized per operation, effectively constant: union by size keeps the trees shallow, and path compression flattens every path that `find` walks.
type DisjointSet struct {
	Parent, sizes []int
	Groups        int
}

func NewDisjointSet(n int) *DisjointSet {
	panic("todo")
}

func (d *DisjointSet) Find(x int) int {
	panic("todo")
}

func (d *DisjointSet) Union(a, b int) bool {
	panic("todo")
}

func (d *DisjointSet) Size(x int) int { panic("todo") }

// ==== END

// ==== EX 63.2 | How many groups after the merges? | Core
// PROMPT: People `0` to `n - 1` start alone; each pair in `pairs` joins two people's groups. Write `GroupCount(n int, pairs [][2]int) int` that returns how many groups remain, using your `DisjointSet` (63.1).
// EXAMPLE: `5, [[0 1] [1 2] [3 4]]` -> `2`
// BIGO: O(N + M alpha(N)) for M pairs: one union per pair, and the set keeps the count.
func GroupCount(n int, pairs [][2]int) int {
	panic("todo")
}

// ==== END

// ==== EX 63.3 | Cycle in an undirected graph | Core
// PROMPT: Write `HasCycleUndirected(n int, edges [][2]int) bool` for vertices `0` to `n - 1` and undirected `[u, v]` edges (no repeats, no self-loops): `true` when the graph has a cycle. Add edges to a `DisjointSet` one at a time; an edge inside one group closes a cycle.
// EXAMPLE: `4, [[0 1] [1 2] [2 0] [2 3]]` -> `true`; `4, [[0 1] [1 2] [2 3]]` -> `false`
// BIGO: O(N + E alpha(N)): an edge whose ends are already in one group closes a cycle.
func HasCycleUndirected(n int, edges [][2]int) bool {
	panic("todo")
}

// ==== END

// ==== EX 63.4 | Is the graph a tree? | Core
// PROMPT: Write `IsTree(n int, edges [][2]int) bool` for `n >= 1` vertices and undirected `[u, v]` edges: `true` when the graph is connected and has no cycle. Check that there are exactly `n - 1` edges and that none of them closes a cycle in a `DisjointSet`.
// EXAMPLE: `4, [[0 1] [0 2] [2 3]]` -> `true`; `4, [[0 1] [2 3]]` -> `false`
// BIGO: O(N + E alpha(N)): a tree on N vertices has exactly N - 1 edges and no cycle, and those two facts together also force it to be connected.
func IsTree(n int, edges [][2]int) bool {
	panic("todo")
}

// ==== END

// ==== EX 63.5 | Minimum spanning tree, edge by edge | Practice
// PROMPT: Write `KruskalCost(n int, edges [][3]int) int` for an undirected graph on `0` to `n - 1` with `[u, v, w]` edges: the total weight of a minimum spanning tree, or `-1` when the graph is not connected. Take edges cheapest first, keeping each one that joins two different groups of a `DisjointSet`.
// EXAMPLE: `4, [[0 1 4] [1 2 2] [0 2 5] [2 3 1] [1 3 7]]` -> `7`
// BIGO: O(E log E) for the sort, then O(E alpha(V)): the cheapest edge between two groups is always safe to take (the cut property).
// NOTE: This is Kruskal's method; 63.6 grows the tree from one vertex instead.
func KruskalCost(n int, edges [][3]int) int {
	panic("todo")
}

// ==== END

// ==== EX 63.6 | Minimum spanning tree, grown from one vertex | Practice
// PROMPT: Write `PrimCost(n int, edges [][3]int) int` with the same input and answer as 63.5, but grow one tree from vertex `0`: keep a heap (`container/heap`) of `[weight, vertex]` for edges leaving the tree, and add the cheapest one whose far end is not yet in the tree.
// EXAMPLE: `4, [[0 1 4] [1 2 2] [0 2 5] [2 3 1] [1 3 7]]` -> `7`
// BIGO: O(E log E) time, O(V + E) space: the heap always offers the cheapest edge leaving the tree built so far, and stale entries are skipped.
// NOTE: This is Prim's method. It needs no union-find, and on dense graphs an array instead of a heap makes it O(V^2).
type leavingEdges [][2]int

func (q leavingEdges) Len() int           { panic("todo") }
func (q leavingEdges) Less(i, j int) bool { panic("todo") }
func (q leavingEdges) Swap(i, j int)      { panic("todo") }
func (q *leavingEdges) Push(x any)        { panic("todo") }
func (q *leavingEdges) Pop() any {
	panic("todo")
}

func PrimCost(n int, edges [][3]int) int {
	panic("todo")
}

// ==== END

// ==== EX 63.7 | The one extra edge | Practice
// PROMPT: `edges` connect vertices `0` to `n - 1`, where `n = len(edges)`: a tree plus one extra edge, so there is exactly one cycle. Write `RedundantEdge(edges [][2]int) [2]int` that returns the edge you can remove, in order of the list, to leave a tree: the first one that joins two vertices already in one group.
// EXAMPLE: `[[0 1] [1 2] [2 3] [1 3] [3 4]]` -> `[1 3]`
// BIGO: O(N alpha(N)): add edges in order; the first one whose ends are already joined is the extra one.
func RedundantEdge(edges [][2]int) [2]int {
	panic("todo")
}

// ==== END

// ==== EX 63.8 | When is everyone connected? | Practice
// PROMPT: `logs` holds `[time, a, b]` entries, in any order, each saying people `a` and `b` became friends at `time`; friendship joins groups. Write `EarliestConnected(n int, logs [][3]int) int` for `n >= 2` people that returns the earliest time everyone is in one group, or `-1`. Sort by time and merge until `Groups` reaches one.
// EXAMPLE: `4, [[7 2 3] [2 0 1] [5 1 2] [9 0 3]]` -> `7`
// BIGO: O(M log M) for sorting M log entries, then O(M alpha(N)): the group count falls to one at exactly the answer.
func EarliestConnected(n int, logs [][3]int) int {
	panic("todo")
}

// ==== END
