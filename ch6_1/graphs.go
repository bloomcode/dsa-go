// Package ch6_1: Graphs: representation and traversal.
package ch6_1


// ==== EX 6.1.1 | Adjacency-list graph with a hash map | Core
// PROMPT: Write `AddEdge(g map[string][]string, a, b string, directed bool)` for a graph stored as vertex -> neighbours. Add missing vertices automatically and avoid duplicate edges.
// EXAMPLE: `g := map[string][]string{}; AddEdge(g, "Ames", "Bern", false)` -> `map[Ames:[Bern] Bern:[Ames]]`
// BIGO: O(D) for the duplicate check, where D is the vertex's degree. A `map[string]bool` of neighbours would make it O(1).
func AddEdge(g map[string][]string, a, b string, directed bool) {
	panic("todo")
}

// ==== END

// ==== EX 6.1.2 | Vertex objects with adjacency lists | Core
// PROMPT: Create `type Vertex struct{ Value string; Adjacent []*Vertex }` and `AddAdjacent(v *Vertex)` that links both ways (friendship is mutual) without recursing forever.
// EXAMPLE: `ames.AddAdjacent(bern)` -> bern is in `ames.Adjacent` and ames is in `bern.Adjacent`
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

// ==== EX 6.1.3 | Depth-first traversal | Core
// PROMPT: Write `DFSTraverse(v *Vertex, visited map[*Vertex]bool) []string` returning vertex values in visit order. Mark each vertex visited before recursing into its neighbours.
// EXAMPLE: Ames-Bern, Ames-Cork, Bern-Dover -> `[Ames Bern Dover Cork]`
// BIGO: O(V + E). Without the visited set, a cycle would loop forever.
func DFSTraverse(v *Vertex, visited map[*Vertex]bool) []string {
	panic("todo")
}

// ==== END

// ==== EX 6.1.4 | Depth-first search | Core
// PROMPT: Write `DFSSearch(v *Vertex, target string, visited map[*Vertex]bool) *Vertex` returning the matching vertex or nil.
// EXAMPLE: From Ames, search "Dover" -> the Dover vertex
// BIGO: O(V + E) worst case.
func DFSSearch(v *Vertex, target string, visited map[*Vertex]bool) *Vertex {
	panic("todo")
}

// ==== END

// ==== EX 6.1.5 | Breadth-first traversal | Core
// PROMPT: Write `BFSTraverse(start *Vertex) []string` with a slice queue: mark the start visited, then repeatedly dequeue a vertex and enqueue its unvisited neighbours.
// EXAMPLE: Same graph as 6.1.3 -> `[Ames Bern Cork Dover]`
// BIGO: O(V + E).
// NOTE: BFS spreads out in rings, so it finds close vertices first; DFS dives deep first.
func BFSTraverse(start *Vertex) []string {
	panic("todo")
}

// ==== END

// ==== EX 6.1.6 | Count connected groups | Practice
// PROMPT: Write `CountComponents(g map[string][]string) int` returning how many separate groups an undirected graph has, using an explicit stack.
// EXAMPLE: `map[a:[b] b:[a] c:[] d:[e] e:[d]]` -> `3`
// BIGO: O(V + E): each vertex is explored once across all searches. The count is the same whatever order the map is walked in.
func CountComponents(g map[string][]string) int {
	panic("todo")
}

// ==== END
