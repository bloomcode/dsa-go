// Package ch18: Graphs: traversal and cheapest paths.
package ch18


// ==== EX 18.1 | Adjacency-list graph with a hash map | Core
// PROMPT: Write `AddEdge(g map[string][]string, a, b string, directed bool)` for a graph stored as vertex -> neighbours. Add missing vertices automatically and avoid duplicate edges.
// EXAMPLE: `g := map[string][]string{}; AddEdge(g, "Alice", "Bob", false)` -> `map[Alice:[Bob] Bob:[Alice]]`
// BIGO: O(D) for the duplicate check, where D is the vertex's degree. A `map[string]bool` of neighbours would make it O(1).
func AddEdge(g map[string][]string, a, b string, directed bool) {
	panic("todo")
}

// ==== END

// ==== EX 18.2 | Object-oriented vertex | Core
// PROMPT: Create `type Vertex struct{ Value string; Adjacent []*Vertex }` and `AddAdjacent(v *Vertex)` that links both ways (friendship is mutual) without recursing forever.
// EXAMPLE: `alice.AddAdjacent(bob)` -> bob is in `alice.Adjacent` and alice is in `bob.Adjacent`
// BIGO: O(D) per link because of the membership check.
// NOTE: `slices.Contains` on `[]*Vertex` compares pointers, which is exactly identity: two vertices with the same name are still different vertices.
type Vertex struct {
	Value    string
	Adjacent []*Vertex
}

func (v *Vertex) AddAdjacent(w *Vertex) {
	panic("todo")
}

// ==== END

// ==== EX 18.3 | Depth-first traversal | Core
// PROMPT: Write `DFSTraverse(v *Vertex, visited map[*Vertex]bool) []string` returning vertex values in visit order. Mark each vertex visited before recursing into its neighbours.
// EXAMPLE: Alice-Bob, Alice-Candy, Bob-Derek -> `[Alice Bob Derek Candy]`
// BIGO: O(V + E). Without the visited set, a cycle would loop forever.
func DFSTraverse(v *Vertex, visited map[*Vertex]bool) []string {
	panic("todo")
}

// ==== END

// ==== EX 18.4 | Depth-first search | Core
// PROMPT: Write `DFSSearch(v *Vertex, target string, visited map[*Vertex]bool) *Vertex` returning the matching vertex or nil.
// EXAMPLE: From Alice, search "Derek" -> the Derek vertex
// BIGO: O(V + E) worst case.
func DFSSearch(v *Vertex, target string, visited map[*Vertex]bool) *Vertex {
	panic("todo")
}

// ==== END

// ==== EX 18.5 | Breadth-first traversal | Core
// PROMPT: Write `BFSTraverse(start *Vertex) []string` with a slice queue: mark the start visited, then repeatedly dequeue a vertex and enqueue its unvisited neighbours.
// EXAMPLE: Same graph as 18.3 -> `[Alice Bob Candy Derek]`
// BIGO: O(V + E).
// NOTE: BFS spreads out in rings, so it finds close vertices first; DFS dives deep first.
func BFSTraverse(start *Vertex) []string {
	panic("todo")
}

// ==== END

// ==== EX 18.6 | Shortest path, unweighted | Core
// PROMPT: Write `ShortestPath(g map[string][]string, start, end string) []string` using BFS and a `previous` map, returning the vertices on the path or nil.
// EXAMPLE: `map[A:[B C] B:[A D] C:[A D] D:[B C E] E:[D]]`, A to E -> `[A B D E]`
// BIGO: O(V + E). The first time BFS reaches a vertex is along a fewest-edges path.
func ShortestPath(g map[string][]string, start, end string) []string {
	panic("todo")
}

// ==== END

// ==== EX 18.7 | Count connected groups | Practice
// PROMPT: Write `CountComponents(g map[string][]string) int` returning how many separate groups an undirected graph has, using an explicit stack.
// EXAMPLE: `map[a:[b] b:[a] c:[] d:[e] e:[d]]` -> `3`
// BIGO: O(V + E): each vertex is explored once across all searches. The count is the same whatever order the map is walked in.
func CountComponents(g map[string][]string) int {
	panic("todo")
}

// ==== END

// ==== EX 18.8 | Dijkstra's algorithm: cheapest flight | Core
// PROMPT: Write `Dijkstra(routes map[string]map[string]int, start, end string) (cost int, path []string, ok bool)`, using `container/heap` to always visit the cheapest known unvisited city next.
// EXAMPLE: Atlanta->Boston 100, Atlanta->Denver 160, Boston->Chicago 120, Boston->Denver 180, Chicago->El Paso 80, Denver->Chicago 40, Denver->El Paso 140 -> Atlanta to El Paso is `280, [Atlanta Denver Chicago El Paso], true`
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

// ==== EX 18.9 | Topological order of tasks | Stretch
// PROMPT: Write `TopologicalOrder(deps map[string][]string) []string` for a directed graph task -> tasks that depend on it, using Kahn's algorithm (repeatedly take tasks with no remaining prerequisites). Seed the queue in sorted order so the answer is the same on every run. Return nil if there is a cycle.
// EXAMPLE: `map[coffee:[dress] dress:[] shower:[dress] wake:[shower coffee]]` -> `[wake shower coffee dress]`
// BIGO: O(V + E), plus O(V log V) for the deterministic seeding.
func TopologicalOrder(deps map[string][]string) []string {
	panic("todo")
}

// ==== END
