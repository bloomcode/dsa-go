package ch61

import (
	"fmt"
	"testing"

	"dsa/internal/check"
)

// towns links Ames - Bern, Ames - Cork, Bern - Dover by hand so tests don't depend on AddAdjacent.
func towns() (ames, bern, cork, dover *Vertex) {
	ames, bern, cork, dover = &Vertex{Value: "Ames"}, &Vertex{Value: "Bern"}, &Vertex{Value: "Cork"}, &Vertex{Value: "Dover"}
	ames.Adjacent = []*Vertex{bern, cork}
	bern.Adjacent = []*Vertex{ames, dover}
	cork.Adjacent = []*Vertex{ames}
	dover.Adjacent = []*Vertex{bern}
	return
}

func TestEx61_1(t *testing.T) {
	g := map[string][]string{}
	AddEdge(g, "Ames", "Bern", false)
	check.Eq(t, fmt.Sprint(g), "map[Ames:[Bern] Bern:[Ames]]")
	AddEdge(g, "Ames", "Bern", false)
	AddEdge(g, "Bern", "Cork", true)
	check.Eq(t, []any{len(g["Ames"]), g["Bern"], g["Cork"]}, []any{1, []string{"Ames", "Cork"}, []string{}})
}

func TestEx61_2(t *testing.T) {
	ames, bern := &Vertex{Value: "Ames"}, &Vertex{Value: "Bern"}
	ames.AddAdjacent(bern)
	check.True(t, len(ames.Adjacent) == 1 && bern.Adjacent[0] == ames, "mutual")
	ames.AddAdjacent(bern)
	check.Eq(t, []int{len(ames.Adjacent), len(bern.Adjacent)}, []int{1, 1})
}

func TestEx61_3(t *testing.T) {
	ames, _, _, _ := towns()
	check.Eq(t, DFSTraverse(ames, map[*Vertex]bool{}), []string{"Ames", "Bern", "Dover", "Cork"})
}

func TestEx61_4(t *testing.T) {
	ames, _, _, dover := towns()
	check.True(t, DFSSearch(ames, "Dover", map[*Vertex]bool{}) == dover, "found")
	check.True(t, DFSSearch(ames, "Zed", map[*Vertex]bool{}) == nil, "missing")
}

func TestEx61_5(t *testing.T) {
	ames, _, _, _ := towns()
	check.Eq(t, BFSTraverse(ames), []string{"Ames", "Bern", "Cork", "Dover"})
}

func TestEx61_6(t *testing.T) {
	check.Eq(t, CountComponents(map[string][]string{"a": {"b"}, "b": {"a"}, "c": {}, "d": {"e"}, "e": {"d"}}), 3)
}
