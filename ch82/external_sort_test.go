package ch82

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"testing"

	"dsa/internal/check"
)

func randomTenK() []int {
	rng := rand.New(rand.NewPCG(3, 0))
	data := make([]int, 10000)
	for i := range data {
		data[i] = rng.IntN(1_000_000)
	}
	return data
}

func TestEx82_1(t *testing.T) {
	check.Eq(t, MakeRuns([]int{5, 1, 9, 3, 7, 2, 8}, 3), [][]int{{1, 5, 9}, {2, 3, 7}, {8}})
	check.Eq(t, len(MakeRuns(nil, 3)), 0)
}

func TestEx82_2(t *testing.T) {
	desc := make([]int, 16)
	for i := range desc {
		desc[i] = 16 - i
	}
	sorted, passes := TwoWaySort(desc, 2)
	check.Eq(t, []any{sorted, passes}, []any{slices.Sorted(slices.Values(desc)), 4})
	data := randomTenK()
	s2, p2 := TwoWaySort(data, 100)
	check.Eq(t, []any{s2, p2}, []any{slices.Sorted(slices.Values(data)), 8})
	e, ep := TwoWaySort(nil, 4)
	check.Eq(t, []any{len(e), ep}, []any{0, 1})
}

func TestEx82_3(t *testing.T) {
	check.Eq(t, KWayMerge([][]int{{1, 4, 9}, {2, 3}, {0, 10}}), []int{0, 1, 2, 3, 4, 9, 10})
	check.Eq(t, len(KWayMerge(nil)), 0)
}

func TestEx82_4(t *testing.T) {
	data := randomTenK()
	sm, pm := MultiwaySort(data, 100, 10)
	check.Eq(t, []any{sm, pm}, []any{slices.Sorted(slices.Values(data)), 4})
}

func TestEx82_5(t *testing.T) {
	check.Eq(t, []int{PassesTwoWay(10_000, 100), PassesMultiway(10_000, 100, 10), PassesTwoWay(16, 2)}, []int{8, 4, 4})
}

func TestEx82_6(t *testing.T) {
	check.Eq(t, ExternalSortIOs(10_000, 100, 10), 8000)
}

func TestEx82_7(t *testing.T) {
	dir := t.TempDir()
	src, dst := filepath.Join(dir, "in.txt"), filepath.Join(dir, "out.txt")
	f, _ := os.Create(src)
	w := bufio.NewWriter(f)
	for _, x := range rand.New(rand.NewPCG(6, 0)).Perm(1000) {
		fmt.Fprintln(w, x)
	}
	w.Flush()
	f.Close()
	runs, err := SortFile(src, dst, 100)
	check.Eq(t, []any{runs, err}, []any{10, nil})
	out, _ := os.Open(dst)
	defer out.Close()
	sc := bufio.NewScanner(out)
	var got []int
	for sc.Scan() {
		v, _ := strconv.Atoi(sc.Text())
		got = append(got, v)
	}
	want := make([]int, 1000)
	for i := range want {
		want[i] = i
	}
	check.Eq(t, got, want)
}
