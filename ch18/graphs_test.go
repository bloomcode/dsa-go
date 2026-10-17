package ch18

import (
	"fmt"
	"testing"

	"dsa/internal/check"
)

// people links Alice - Bob, Alice - Candy, Bob - Derek by hand so tests don't depend on AddAdjacent.
func people() (alice, bob, candy, derek *Vertex) {
	alice, bob, candy, derek = &Vertex{Value: "Alice"}, &Vertex{Value: "Bob"}, &Vertex{Value: "Candy"}, &Vertex{Value: "Derek"}
	alice.Adjacent = []*Vertex{bob, candy}
	bob.Adjacent = []*Vertex{alice, derek}
	candy.Adjacent = []*Vertex{alice}
	derek.Adjacent = []*Vertex{bob}
	return
}

func TestEx18_1(t *testing.T) {
	g := map[string][]string{}
	AddEdge(g, "Alice", "Bob", false)
	check.Eq(t, fmt.Sprint(g), "map[Alice:[Bob] Bob:[Alice]]")
	AddEdge(g, "Alice", "Bob", false)
	AddEdge(g, "Bob", "Cy", true)
	check.Eq(t, []any{len(g["Alice"]), g["Bob"], g["Cy"]}, []any{1, []string{"Alice", "Cy"}, []string{}})
}

func TestEx18_2(t *testing.T) {
	alice, bob := &Vertex{Value: "Alice"}, &Vertex{Value: "Bob"}
	alice.AddAdjacent(bob)
	check.True(t, len(alice.Adjacent) == 1 && bob.Adjacent[0] == alice, "mutual")
	alice.AddAdjacent(bob)
	check.Eq(t, []int{len(alice.Adjacent), len(bob.Adjacent)}, []int{1, 1})
}

func TestEx18_3(t *testing.T) {
	alice, _, _, _ := people()
	check.Eq(t, DFSTraverse(alice, map[*Vertex]bool{}), []string{"Alice", "Bob", "Derek", "Candy"})
}

func TestEx18_4(t *testing.T) {
	alice, _, _, derek := people()
	check.True(t, DFSSearch(alice, "Derek", map[*Vertex]bool{}) == derek, "found")
	check.True(t, DFSSearch(alice, "Zed", map[*Vertex]bool{}) == nil, "missing")
}

func TestEx18_5(t *testing.T) {
	alice, _, _, _ := people()
	check.Eq(t, BFSTraverse(alice), []string{"Alice", "Bob", "Candy", "Derek"})
}

func TestEx18_6(t *testing.T) {
	sp := map[string][]string{"A": {"B", "C"}, "B": {"A", "D"}, "C": {"A", "D"}, "D": {"B", "C", "E"}, "E": {"D"}}
	check.Eq(t, ShortestPath(sp, "A", "E"), []string{"A", "B", "D", "E"})
	check.True(t, ShortestPath(map[string][]string{"A": {}, "B": {}}, "A", "B") == nil, "no path")
}

func TestEx18_7(t *testing.T) {
	check.Eq(t, CountComponents(map[string][]string{"a": {"b"}, "b": {"a"}, "c": {}, "d": {"e"}, "e": {"d"}}), 3)
}

func TestEx18_8(t *testing.T) {
	routes := map[string]map[string]int{
		"Atlanta": {"Boston": 100, "Denver": 160},
		"Boston":  {"Chicago": 120, "Denver": 180},
		"Chicago": {"El Paso": 80},
		"Denver":  {"Chicago": 40, "El Paso": 140},
	}
	for i := 0; i < 20; i++ { // map order is random; the answer must not be
		cost, path, ok := Dijkstra(routes, "Atlanta", "El Paso")
		check.Eq(t, []any{cost, path, ok}, []any{280, []string{"Atlanta", "Denver", "Chicago", "El Paso"}, true})
	}
	_, _, ok := Dijkstra(routes, "El Paso", "Atlanta")
	check.Eq(t, ok, false)
}

func TestEx18_9(t *testing.T) {
	deps := map[string][]string{"wake": {"shower", "coffee"}, "shower": {"dress"}, "coffee": {"dress"}, "dress": {}}
	for i := 0; i < 20; i++ {
		check.Eq(t, TopologicalOrder(deps), []string{"wake", "shower", "coffee", "dress"})
	}
	check.True(t, TopologicalOrder(map[string][]string{"a": {"b"}, "b": {"a"}}) == nil, "cycle")
}
