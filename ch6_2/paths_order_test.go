package ch6_2

import (
	"math"
	"testing"

	"dsa/internal/check"
)

func TestEx6_2_1(t *testing.T) {
	sp := map[string][]string{"A": {"B", "C"}, "B": {"A", "D"}, "C": {"A", "D"}, "D": {"B", "C", "E"}, "E": {"D"}}
	check.Eq(t, ShortestPath(sp, "A", "E"), []string{"A", "B", "D", "E"})
	check.True(t, ShortestPath(map[string][]string{"A": {}, "B": {}}, "A", "B") == nil, "no path")
}

func TestEx6_2_2(t *testing.T) {
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

func TestEx6_2_3(t *testing.T) {
	deps := map[string][]string{"wake": {"shower", "coffee"}, "shower": {"dress"}, "coffee": {"dress"}, "dress": {}}
	for i := 0; i < 20; i++ {
		check.Eq(t, TopologicalOrder(deps), []string{"wake", "shower", "coffee", "dress"})
	}
	check.True(t, TopologicalOrder(map[string][]string{"a": {"b"}, "b": {"a"}}) == nil, "cycle")
}

func TestEx6_2_4(t *testing.T) {
	inf := math.MaxInt
	dist, ok := BellmanFord(4, [][3]int{{0, 1, 4}, {0, 2, 1}, {2, 1, -2}, {1, 3, 1}}, 0)
	check.Eq(t, dist, []int{0, -1, 1, 0})
	check.True(t, ok, "no negative cycle")
	dist, _ = BellmanFord(3, [][3]int{{0, 1, 5}}, 0)
	check.Eq(t, dist, []int{0, 5, inf})
	dist, _ = BellmanFord(1, nil, 0)
	check.Eq(t, dist, []int{0})
	_, ok = BellmanFord(3, [][3]int{{0, 1, 1}, {1, 2, -1}, {2, 1, -1}}, 0)
	check.True(t, !ok, "negative cycle")
	dist, ok = BellmanFord(3, [][3]int{{1, 2, -1}, {2, 1, -1}}, 0)
	check.Eq(t, dist, []int{0, inf, inf})
	check.True(t, ok, "unreachable cycle does not count")
}

func TestEx6_2_5(t *testing.T) {
	got := []bool{
		HasProfitableLoop([][]float64{{1, 0.9, 0.8}, {1.2, 1, 0.7}, {1.3, 1.4, 1}}), HasProfitableLoop([][]float64{{1, 0.5}, {1.9, 1}}),
		HasProfitableLoop([][]float64{{1, 2}, {0.5, 1}}), HasProfitableLoop([][]float64{{1}}),
		HasProfitableLoop([][]float64{{1, 2, 0.1}, {0.4, 1, 3}, {0.2, 0.1, 1}}),
	}
	check.Eq(t, got, []bool{true, false, false, false, true})
}

func TestEx6_2_6(t *testing.T) {
	grid := []string{"....#", ".##.#", "....."}
	got := []int{
		GridPathLength(grid, [2]int{0, 0}, [2]int{2, 4}), GridPathLength(grid, [2]int{0, 0}, [2]int{0, 0}), GridPathLength(grid, [2]int{0, 3}, [2]int{2, 4}),
		GridPathLength([]string{".#.", ".#.", "..."}, [2]int{0, 0}, [2]int{0, 2}), GridPathLength([]string{".#", "#."}, [2]int{0, 0}, [2]int{1, 1}),
	}
	check.Eq(t, got, []int{6, 0, 3, 6, -1})
}
