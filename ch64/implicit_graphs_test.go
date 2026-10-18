package ch64

import (
	"slices"
	"testing"

	"dsa/internal/check"
)

func TestEx64_1(t *testing.T) {
	got := []int{CountIslands([]string{"##..#", "#...#", "..#..", "....#"}), CountIslands(nil), CountIslands([]string{"..."}), CountIslands([]string{"###", "#.#", "###"}), CountIslands([]string{"#.#", ".#.", "#.#"})}
	check.Eq(t, got, []int{4, 0, 0, 1, 5})
}

func TestEx64_2(t *testing.T) {
	got := []bool{
		TwoColourable(4, [][2]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}}), TwoColourable(3, [][2]int{{0, 1}, {1, 2}, {2, 0}}), TwoColourable(5, nil),
		TwoColourable(6, [][2]int{{0, 1}, {3, 4}, {4, 5}, {5, 3}}), TwoColourable(4, [][2]int{{0, 1}, {2, 3}}),
	}
	check.Eq(t, got, []bool{true, false, true, false, true})
}

func TestEx64_3(t *testing.T) {
	ring := make([]*GraphNode, 4)
	for i := range ring {
		ring[i] = &GraphNode{Value: i + 1}
	}
	for i, node := range ring {
		node.Neighbours = []*GraphNode{ring[(i+1)%4], ring[(i+3)%4]}
	}
	cp := CloneGraph(ring[0])
	seen, stack, values := map[*GraphNode]bool{cp: true}, []*GraphNode{cp}, []int{}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		check.True(t, !slices.Contains(ring, node), "copy shares no node")
		values = append(values, node.Value)
		for _, other := range node.Neighbours {
			if !seen[other] {
				seen[other] = true
				stack = append(stack, other)
			}
		}
	}
	slices.Sort(values)
	check.Eq(t, values, []int{1, 2, 3, 4})
	check.Eq(t, []int{cp.Neighbours[0].Value, cp.Neighbours[1].Value}, []int{2, 4})
	check.True(t, cp.Neighbours[0].Neighbours[1] == cp, "links point at copies")
	check.True(t, CloneGraph(nil) == nil, "nil graph")
	alone := &GraphNode{Value: 7}
	alone.Neighbours = []*GraphNode{alone}
	twin := CloneGraph(alone)
	check.True(t, twin != alone && len(twin.Neighbours) == 1 && twin.Neighbours[0] == twin, "self-loop")
}

func TestEx64_4(t *testing.T) {
	got := []int{
		LargestIsland([]string{"##..#", "#...#", "..#..", "....#"}), LargestIsland(nil), LargestIsland([]string{"..."}), LargestIsland([]string{"###", "#.#", "###"}),
		LargestIsland([]string{"#.#", ".#.", "#.#"}), LargestIsland([]string{"##.", "...", "###"}),
	}
	check.Eq(t, got, []int{3, 0, 0, 8, 1, 3})
}

func TestEx64_5(t *testing.T) {
	got := []int{
		RoundsToSpread([]string{"S..#", ".#..", "...S"}), RoundsToSpread([]string{"S#."}), RoundsToSpread([]string{"S#"}),
		RoundsToSpread([]string{"..."}), RoundsToSpread([]string{"S...."}), RoundsToSpread([]string{"S..", "#..", "..S"}),
	}
	check.Eq(t, got, []int{2, -1, 0, -1, 4, 2})
}

func TestEx64_6(t *testing.T) {
	got := []bool{
		HasCycleDirected(4, [][2]int{{0, 1}, {1, 2}, {2, 0}, {2, 3}}), HasCycleDirected(3, [][2]int{{0, 1}, {0, 2}, {1, 2}}), HasCycleDirected(3, nil),
		HasCycleDirected(2, [][2]int{{1, 1}}), HasCycleDirected(5, [][2]int{{0, 1}, {3, 4}, {4, 2}, {2, 3}}), HasCycleDirected(4, [][2]int{{0, 1}, {2, 1}, {3, 2}, {3, 0}}),
	}
	check.Eq(t, got, []bool{true, false, false, true, true, false})
}

func TestEx64_7(t *testing.T) {
	words := []string{"cord", "card", "ward", "warm", "word", "worm"}
	got := []int{
		LadderLength("cold", "warm", words), LadderLength("cold", "cord", words), LadderLength("cold", "warn", words),
		LadderLength("ab", "cd", []string{"ad", "cd"}), LadderLength("ab", "cd", []string{"xy", "cd"}),
	}
	check.Eq(t, got, []int{5, 2, 0, 3, 0})
}

func TestEx64_8(t *testing.T) {
	got := []int{
		KnightDistance(8, [2]int{0, 0}, [2]int{7, 7}), KnightDistance(8, [2]int{0, 0}, [2]int{1, 2}), KnightDistance(8, [2]int{3, 3}, [2]int{3, 3}),
		KnightDistance(3, [2]int{0, 0}, [2]int{1, 1}), KnightDistance(8, [2]int{0, 0}, [2]int{1, 1}), KnightDistance(1, [2]int{0, 0}, [2]int{0, 0}),
	}
	check.Eq(t, got, []int{6, 1, 0, -1, 4, 0})
}

func TestEx64_9(t *testing.T) {
	got := []int{
		LongestRisingPath([][]int{{1, 2, 9}, {5, 3, 8}, {4, 6, 7}}), LongestRisingPath([][]int{{7}}), LongestRisingPath([][]int{{1, 1}, {1, 1}}),
		LongestRisingPath([][]int{{1, 2, 3, 4}}), LongestRisingPath([][]int{{5, 4}, {2, 3}}),
	}
	check.Eq(t, got, []int{7, 1, 1, 4, 4})
}
