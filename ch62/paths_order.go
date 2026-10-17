// Package ch62: Shortest paths and ordering.
package ch62


// ==== EX 62.1 | Shortest path, unweighted | Core
// PROMPT: Write `ShortestPath(g map[string][]string, start, end string) []string` using BFS and a `previous` map, returning the vertices on the path or nil.
// EXAMPLE: `map[A:[B C] B:[A D] C:[A D] D:[B C E] E:[D]]`, A to E -> `[A B D E]`
// BIGO: O(V + E). The first time BFS reaches a vertex is along a fewest-edges path.
func ShortestPath(g map[string][]string, start, end string) []string {
	panic("todo")
}

// ==== END

// ==== EX 62.2 | Fastest route by travel time | Core
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

// ==== EX 62.3 | Topological order of tasks | Stretch
// PROMPT: Write `TopologicalOrder(deps map[string][]string) []string` for a directed graph task -> tasks that depend on it, using Kahn's algorithm (repeatedly take tasks with no remaining prerequisites). Seed the queue in sorted order so the answer is the same on every run. Return nil if there is a cycle.
// EXAMPLE: `map[coffee:[dress] dress:[] shower:[dress] wake:[shower coffee]]` -> `[wake shower coffee dress]`
// BIGO: O(V + E), plus O(V log V) for the deterministic seeding.
func TopologicalOrder(deps map[string][]string) []string {
	panic("todo")
}

// ==== END
