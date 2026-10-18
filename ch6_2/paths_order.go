// Package ch6_2: Shortest paths and ordering.
package ch6_2


// ==== EX 6.2.1 | Shortest path, unweighted | Core
// PROMPT: Write `ShortestPath(g map[string][]string, start, end string) []string` using BFS and a `previous` map, returning the vertices on the path or nil.
// EXAMPLE: `map[A:[B C] B:[A D] C:[A D] D:[B C E] E:[D]]`, A to E -> `[A B D E]`
// BIGO: O(V + E). The first time BFS reaches a vertex is along a fewest-edges path.
func ShortestPath(g map[string][]string, start, end string) []string {
	panic("todo")
}

// ==== END

// ==== EX 6.2.2 | Fastest route by travel time | Core
// PROMPT: Write `Dijkstra(routes map[string]map[string]int, start, end string) (cost int, path []string, ok bool)`, using `container/heap` to always visit the nearest known unvisited town next.
// EXAMPLE: Ames->Bern 4, Ames->Cork 2, Bern->Dover 5, Cork->Bern 1, Cork->Dover 8, Cork->Eden 10, Dover->Eden 2 -> Ames to Eden is `10, [Ames Cork Bern Dover Eden], true`
// BIGO: O((V + E) log V) with a heap. Scanning an array for the cheapest unvisited city instead is O(V^2).
// NOTE: Dijkstra requires non-negative weights. `container/heap` has no decrease-key either, so push a new entry and skip stale ones when they come out.
type stop struct {
	cost int
	city string
}

type stops []stop

func (s stops) Len() int           { panic("todo") }
func (s stops) Less(i, j int) bool { panic("todo") }
func (s stops) Swap(i, j int)      { panic("todo") }
func (s *stops) Push(x any)        { panic("todo") }
func (s *stops) Pop() any {
	panic("todo")
}

func Dijkstra(routes map[string]map[string]int, start, end string) (int, []string, bool) {
	panic("todo")
}

// ==== END

// ==== EX 6.2.3 | Topological order of tasks | Stretch
// PROMPT: Write `TopologicalOrder(deps map[string][]string) []string` for a directed graph task -> tasks that depend on it, using Kahn's algorithm (repeatedly take tasks with no remaining prerequisites). Seed the queue in sorted order so the answer is the same on every run. Return nil if there is a cycle.
// EXAMPLE: `map[coffee:[dress] dress:[] shower:[dress] wake:[shower coffee]]` -> `[wake shower coffee dress]`
// BIGO: O(V + E), plus O(V log V) for the deterministic seeding.
func TopologicalOrder(deps map[string][]string) []string {
	panic("todo")
}

// ==== END

// ==== EX 6.2.4 | Shortest paths with negative weights | Practice
// PROMPT: Vertices are `0` to `n - 1`, and `edges` holds directed `[u, v, w]` with any integer weight. Write `BellmanFord(n int, edges [][3]int, source int) ([]int, bool)` that returns the shortest distances from `source` (`math.MaxInt` when unreachable) and `true`, or `nil, false` when a negative cycle can be reached. Relax every edge `n - 1` times, then check once more.
// EXAMPLE: `4, [[0 1 4] [0 2 1] [2 1 -2] [1 3 1]], 0` -> `[0 -1 1 0], true`
// BIGO: O(V * E) time, O(V) space: a shortest path uses at most V - 1 edges, so V - 1 rounds of relaxing every edge settle it, and any improvement after that means a negative cycle.
// NOTE: This is the Bellman-Ford method. Dijkstra (6.2.2) is faster but wrong once a weight is negative, because a settled vertex can still get cheaper.
func BellmanFord(n int, edges [][3]int, source int) ([]int, bool) {
	panic("todo")
}

// ==== END

// ==== EX 6.2.5 | A profitable loop of exchanges | Stretch
// PROMPT: `rates[i][j]` is how many units of currency `j` one unit of `i` buys. Write `HasProfitableLoop(rates [][]float64) bool`: `true` when some cycle of exchanges ends with more than it started. Give edge `i -> j` the weight `-math.Log(rates[i][j])` and look for a negative cycle: start every distance at `0`, relax all edges `n - 1` times, then check once more (allow `1e-12` for rounding).
// EXAMPLE: `[[1 0.9 0.8] [1.2 1 0.7] [1.3 1.4 1]]` -> `true` (0 -> 1 -> 0 turns 1 into 1.08)
// BIGO: O(N^3) time for N currencies, O(N) space: logarithms turn a product above 1 into a sum below 0, so a profitable loop is a negative cycle.
func HasProfitableLoop(rates [][]float64) bool {
	panic("todo")
}

// ==== END

// ==== EX 6.2.6 | A* search on a grid | Stretch
// PROMPT: `grid` holds strings, `.` open and `#` wall; a step moves to one of the 4 neighbouring open cells. Write `GridPathLength(grid []string, start, goal [2]int) int` (cells as `[row col]`) that returns the fewest steps, or `-1`, using A*: a heap (`container/heap`) ordered by steps so far plus the Manhattan distance to `goal`, and a table of the best known steps per cell.
// EXAMPLE: `["....#" ".##.#" "....."], [0 0], [2 4]` -> `6`
// BIGO: O(R*C log(R*C)) worst case, usually much less: the Manhattan distance never overestimates, so the first time the goal leaves the queue its cost is optimal, and the estimate steers the search toward it.
type openCells [][4]int

func (q openCells) Len() int           { panic("todo") }
func (q openCells) Less(i, j int) bool { panic("todo") }
func (q openCells) Swap(i, j int)      { panic("todo") }
func (q *openCells) Push(x any)        { panic("todo") }
func (q *openCells) Pop() any {
	panic("todo")
}

func GridPathLength(grid []string, start, goal [2]int) int {
	panic("todo")
}

// ==== END
