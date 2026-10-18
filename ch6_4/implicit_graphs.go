// Package ch6_4: Grid and implicit graphs.
package ch6_4

// ==== EX 6.4.1 | Count the islands | Core
// PROMPT: `grid` holds strings, `#` land and `.` water. Write `CountIslands(grid []string) int` that returns how many islands there are, an island being land cells joined up, down, left or right. Scan the cells; at each unvisited land cell, count one island and visit all of it with a stack.
// EXAMPLE: `["##..#" "#...#" "..#.." "....#"]` -> `4`
// BIGO: O(R*C) time and space: each land cell is pushed once, when the island it belongs to is first found.
func CountIslands(grid []string) int {
	panic("todo")
}

// ==== END

// ==== EX 6.4.2 | Can the graph be two-coloured? | Core
// PROMPT: Write `TwoColourable(n int, edges [][2]int) bool` for vertices `0` to `n - 1` and undirected `[u, v]` edges: `true` when every vertex can be coloured one of two colours with no edge joining two of the same colour. Breadth-first search from each uncoloured vertex, giving neighbours the other colour.
// EXAMPLE: `4, [[0 1] [1 2] [2 3] [3 0]]` -> `true`; `3, [[0 1] [1 2] [2 0]]` -> `false`
// BIGO: O(V + E): breadth-first search colours each vertex once, and an edge with both ends the same colour proves an odd cycle.
func TwoColourable(n int, edges [][2]int) bool {
	panic("todo")
}

// ==== END

// ==== EX 6.4.3 | Copy a graph | Core
// PROMPT: Create `type GraphNode struct{ Value int; Neighbours []*GraphNode }`. Write `CloneGraph(node *GraphNode) *GraphNode` that returns a deep copy of the connected graph reachable from `node` (`nil` for `nil`): new nodes with the same values and the same links, sharing nothing with the original. Keep a map from original to copy and walk the graph.
// EXAMPLE: a ring `1 - 2 - 3 - 4 - 1` -> a new ring `1 - 2 - 3 - 4 - 1` that shares no node
// BIGO: O(V + E) time and space: a map from each original node to its copy makes sure every node is copied once, even around cycles.
type GraphNode struct {
	Value      int
	Neighbours []*GraphNode
}

func CloneGraph(node *GraphNode) *GraphNode {
	panic("todo")
}

// ==== END

// ==== EX 6.4.4 | Largest island | Practice
// PROMPT: With the grid of 6.4.1, write `LargestIsland(grid []string) int` that returns the number of cells in the biggest island (`0` when there is no land). Fill each island with a depth-first walk and count its cells.
// EXAMPLE: `["##..#" "#...#" "..#.." "....#"]` -> `3`
// BIGO: O(R*C) time and space: each cell is visited once, and each island's area is counted while it is filled.
func LargestIsland(grid []string) int {
	panic("todo")
}

// ==== END

// ==== EX 6.4.5 | Rounds until everything is reached | Practice
// PROMPT: `grid` holds `S` (a source), `.` (a cell to reach) and `#` (a wall). Each round, every reached cell reaches its 4 neighbours that are not walls. Write `RoundsToSpread(grid []string) int` that returns how many rounds it takes to reach every `.` cell, `0` if there is none, or `-1` if some can never be reached. Start one breadth-first search from all sources together and go one level per round.
// EXAMPLE: `["S..#" ".#.." "...S"]` -> `2`
// BIGO: O(R*C) time and space: a breadth-first search started from every source at once reaches each cell in the fewest rounds.
func RoundsToSpread(grid []string) int {
	panic("todo")
}

// ==== END

// ==== EX 6.4.6 | Cycle in a directed graph | Practice
// PROMPT: Write `HasCycleDirected(n int, edges [][2]int) bool` for vertices `0` to `n - 1` and directed `[u, v]` edges: `true` when the graph has a directed cycle. Depth-first search with three states per vertex; reaching a vertex that is still on the current path means a cycle (a finished one does not).
// EXAMPLE: `4, [[0 1] [1 2] [2 0] [2 3]]` -> `true`; `3, [[0 1] [0 2] [1 2]]` -> `false`
// BIGO: O(V + E): with three colours (unvisited, on the current path, finished), an edge back to a vertex on the current path is exactly a cycle.
func HasCycleDirected(n int, edges [][2]int) bool {
	panic("todo")
}

// ==== END

// ==== EX 6.4.7 | Word ladder | Practice
// PROMPT: Write `LadderLength(start, end string, words []string) int` that returns how many words are in the shortest chain from `start` to `end` (both counted) where each step changes exactly one letter and every word after `start` is in `words`; `0` when there is no chain. Breadth-first search, generating neighbours by trying each letter `a` to `z` at each position.
// EXAMPLE: `"cold", "warm", ["cord" "card" "ward" "warm" "word" "worm"]` -> `5` (cold, cord, card, ward, warm)
// BIGO: O(N * L * 26) for N words of length L: breadth-first search where each word's neighbours are found by trying every letter at every position against a set.
func LadderLength(start, end string, words []string) int {
	panic("todo")
}

// ==== END

// ==== EX 6.4.8 | Knight moves | Practice
// PROMPT: Write `KnightDistance(n int, start, goal [2]int) int` that returns the fewest knight moves from `start` to `goal` on an `n` by `n` board (squares as `[row col]`), or `-1` when the goal cannot be reached. Breadth-first search over squares.
// EXAMPLE: `8, [0 0], [7 7]` -> `6`
// BIGO: O(N^2) time and space on an N by N board: breadth-first search over squares, eight moves from each.
func KnightDistance(n int, start, goal [2]int) int {
	panic("todo")
}

// ==== END

// ==== EX 6.4.9 | Longest rising path in a grid | Stretch
// PROMPT: `grid` is a non-empty rectangle of integers. Write `LongestRisingPath(grid [][]int) int` that returns the number of cells on the longest path that moves up, down, left or right to a strictly larger value each step. A memoized depth-first search gives each cell `1 +` the best of its larger neighbours.
// EXAMPLE: `[[1 2 9] [5 3 8] [4 6 7]]` -> `7` (1, 2, 3, 6, 7, 8, 9)
// BIGO: O(R*C) time and space: rising steps can never loop, so the graph is acyclic, and memoizing each cell's best path means each is solved once.
func LongestRisingPath(grid [][]int) int {
	panic("todo")
}

// ==== END
