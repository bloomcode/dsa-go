package ch62

import (
	"testing"

	"dsa/internal/check"
)

func TestEx62_1(t *testing.T) {
	sp := map[string][]string{"A": {"B", "C"}, "B": {"A", "D"}, "C": {"A", "D"}, "D": {"B", "C", "E"}, "E": {"D"}}
	check.Eq(t, ShortestPath(sp, "A", "E"), []string{"A", "B", "D", "E"})
	check.True(t, ShortestPath(map[string][]string{"A": {}, "B": {}}, "A", "B") == nil, "no path")
}

func TestEx62_2(t *testing.T) {
	routes := map[string]map[string]int{
		"Ames":  {"Bern": 4, "Cork": 2},
		"Bern":  {"Dover": 5},
		"Cork":  {"Bern": 1, "Dover": 8, "Eden": 10},
		"Dover": {"Eden": 2},
	}
	for i := 0; i < 20; i++ { // map order is random; the answer must not be
		cost, path, ok := Dijkstra(routes, "Ames", "Eden")
		check.Eq(t, []any{cost, path, ok}, []any{10, []string{"Ames", "Cork", "Bern", "Dover", "Eden"}, true})
	}
	cost, path, _ := Dijkstra(routes, "Ames", "Dover")
	check.Eq(t, []any{cost, path}, []any{8, []string{"Ames", "Cork", "Bern", "Dover"}})
	_, _, ok := Dijkstra(routes, "Eden", "Ames")
	check.Eq(t, ok, false)
}

func TestEx62_3(t *testing.T) {
	deps := map[string][]string{"wake": {"shower", "coffee"}, "shower": {"dress"}, "coffee": {"dress"}, "dress": {}}
	for i := 0; i < 20; i++ {
		check.Eq(t, TopologicalOrder(deps), []string{"wake", "shower", "coffee", "dress"})
	}
	check.True(t, TopologicalOrder(map[string][]string{"a": {"b"}, "b": {"a"}}) == nil, "cycle")
}
