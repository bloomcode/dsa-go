package ch41

import (
	"fmt"
	"strconv"
	"testing"

	"dsa/internal/check"
)

func TestEx41_1(t *testing.T) {
	h := NewHashTable()
	h.Put("bad", 7)
	v, ok := h.Get("bad")
	check.Eq(t, []any{v, ok}, []any{7, true})
	h.Put("bad", 8)
	v, _ = h.Get("bad")
	check.Eq(t, []int{v, h.Len()}, []int{8, 1})
	_, ok = h.Get("good")
	check.Eq(t, ok, false)
	big := NewHashTable()
	for i := 0; i < 1000; i++ {
		big.Put("k"+strconv.Itoa(i), i*i)
	}
	check.True(t, big.Len() == 1000 && big.Buckets() >= 1334, "resized")
	v, _ = big.Get("k321")
	check.Eq(t, v, 103041)
	check.True(t, big.Delete("k321") && !big.Delete("k321"), "delete")
	_, ok = big.Get("k321")
	check.Eq(t, []any{ok, big.Len()}, []any{false, 999})
}

func TestEx41_2(t *testing.T) {
	check.Eq(t, []bool{CanBuild("abba", "aabbc"), CanBuild("abc", "ab")}, []bool{true, false})
	check.Eq(t, []bool{CanBuild("", "xyz"), CanBuild("aa", "a"), CanBuild("cab", "bac")}, []bool{true, false, true})
}

func TestEx41_3(t *testing.T) {
	check.Eq(t, InExactlyOne([]int{1, 2, 3}, []int{2, 3, 4}), []int{1, 4})
	check.Eq(t, InExactlyOne([]int{5, 5, 1}, []int{1}), []int{5})
	check.Eq(t, InExactlyOne(nil, []int{2, 2}), []int{2})
	check.Eq(t, InExactlyOne([]int{1}, []int{1}), []int{})
}

func TestEx41_4(t *testing.T) {
	check.Eq(t, DistinctPerDay([]int{1, 1, 1, 2}, []string{"ana", "bo", "ana", "ana"}), map[int]int{1: 2, 2: 1})
	check.Eq(t, DistinctPerDay(nil, nil), map[int]int{})
	check.Eq(t, DistinctPerDay([]int{7, 7}, []string{"x", "x"}), map[int]int{7: 1})
}

func TestEx41_5(t *testing.T) {
	check.Eq(t, []bool{SamePattern("egg", "add"), SamePattern("foo", "bar"), SamePattern("paper", "title")}, []bool{true, false, true})
	check.Eq(t, []bool{SamePattern("ab", "aa"), SamePattern("aa", "ab"), SamePattern("a", "ab"), SamePattern("", "")}, []bool{false, false, false, true})
}

func TestEx41_6(t *testing.T) {
	check.Eq(t, Positions([]string{"a", "b", "a"}), map[string][]int{"a": {0, 2}, "b": {1}})
	check.Eq(t, Positions(nil), map[string][]int{})
	check.Eq(t, Positions([]string{"x", "x", "x"}), map[string][]int{"x": {0, 1, 2}})
}

func TestEx41_7(t *testing.T) {
	check.Eq(t, fmt.Sprint(WordCounts("the cat and the hat")), "map[and:1 cat:1 hat:1 the:2]")
	check.Eq(t, MostCommonWord("the cat and the hat"), "the")
	check.Eq(t, MostCommonWord("b a"), "b")
}

func TestEx41_8(t *testing.T) {
	check.Eq(t, GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}), [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}})
}

func TestEx41_9(t *testing.T) {
	i, j, ok := TwoSumIndices([]int{2, 7, 11, 15}, 9)
	check.Eq(t, []any{i, j, ok}, []any{0, 1, true})
	_, _, ok = TwoSumIndices([]int{1}, 2)
	check.Eq(t, ok, false)
	i, j, _ = TwoSumIndices([]int{3, 3}, 6)
	check.Eq(t, []int{i, j}, []int{0, 1})
}

func TestEx41_10(t *testing.T) {
	check.Eq(t, fmt.Sprint(Invert(map[string]int{"fries": 2, "soda": 2, "burger": 5})), "map[2:[fries soda] 5:[burger]]")
}

func TestEx41_11(t *testing.T) {
	got := []int{LongestConsecutive([]int{10, 4, 20, 1, 3, 2, 11}), LongestConsecutive(nil), LongestConsecutive([]int{5, 5, 5}), LongestConsecutive([]int{0, -1, 1, 9, 8})}
	check.Eq(t, got, []int{4, 0, 1, 3})
}

func TestEx41_12(t *testing.T) {
	got := []bool{HasZeroSumRun([]int{4, 2, -3, 1, 6}), HasZeroSumRun([]int{4, 2, 1}), HasZeroSumRun([]int{0}), HasZeroSumRun(nil), HasZeroSumRun([]int{1, -1}), HasZeroSumRun([]int{3, 5, -8, 2}), HasZeroSumRun([]int{-2, 5})}
	check.Eq(t, got, []bool{true, false, true, false, true, true, false})
}

func TestEx41_13(t *testing.T) {
	got := []int{CountSumRuns([]int{1, 2, 3, -2, 2}, 3), CountSumRuns([]int{1, 1, 1}, 2), CountSumRuns(nil, 0), CountSumRuns([]int{0, 0}, 0), CountSumRuns([]int{5}, -5), CountSumRuns([]int{2, -2, 2}, 2)}
	check.Eq(t, got, []int{4, 2, 0, 3, 0, 3})
}

func TestEx41_14(t *testing.T) {
	check.Eq(t, SymmetricPairs([][2]int{{1, 2}, {3, 4}, {5, 6}, {4, 3}, {2, 1}}), [][2]int{{3, 4}, {1, 2}})
	check.Eq(t, SymmetricPairs(nil), [][2]int{})
	check.Eq(t, SymmetricPairs([][2]int{{1, 2}, {2, 3}}), [][2]int{})
	check.Eq(t, SymmetricPairs([][2]int{{7, 8}, {8, 7}, {9, 7}, {7, 9}}), [][2]int{{7, 8}, {9, 7}})
}

func TestEx41_15(t *testing.T) {
	check.Eq(t, TraceJourney([][2]string{{"MAA", "BOM"}, {"DEL", "MAA"}, {"BOM", "GOI"}}), []string{"DEL", "MAA", "BOM", "GOI"})
	check.Eq(t, TraceJourney(nil), []string{})
	check.Eq(t, TraceJourney([][2]string{{"A", "B"}}), []string{"A", "B"})
	check.Eq(t, TraceJourney([][2]string{{"C", "D"}, {"B", "C"}, {"D", "E"}, {"A", "B"}}), []string{"A", "B", "C", "D", "E"})
}

func TestEx41_16(t *testing.T) {
	got := []bool{Disjoint([]int{1, 5, 9}, []int{2, 6}), Disjoint([]int{1, 5}, []int{5}), Disjoint(nil, []int{1}), Disjoint(nil, nil), Disjoint([]int{3, 3}, []int{4, 3})}
	check.Eq(t, got, []bool{true, false, true, true, false})
}

func TestEx41_17(t *testing.T) {
	got := []int{
		MostOnALine([][2]int{{0, 0}, {1, 1}, {2, 2}, {3, 1}, {0, 3}}), MostOnALine(nil), MostOnALine([][2]int{{1, 1}}),
		MostOnALine([][2]int{{1, 1}, {1, 1}, {2, 3}}), MostOnALine([][2]int{{2, 0}, {2, 5}, {2, -1}, {3, 3}}),
		MostOnALine([][2]int{{0, 0}, {2, 1}, {4, 2}, {6, 3}, {1, 0}}), MostOnALine([][2]int{{0, 0}, {-2, 1}, {2, -1}}),
	}
	check.Eq(t, got, []int{3, 0, 1, 3, 3, 4, 3})
}
