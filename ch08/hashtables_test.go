package ch08

import (
	"fmt"
	"strconv"
	"testing"

	"dsa/internal/check"
)

func TestEx8_1(t *testing.T) {
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

func TestEx8_2(t *testing.T) {
	check.True(t, IsSubset([]string{"b", "d"}, []string{"a", "b", "c", "d"}), "subset")
	check.True(t, !IsSubset([]string{"b", "z"}, []string{"a", "b", "c", "d"}), "not subset")
}

func TestEx8_3(t *testing.T) {
	check.Eq(t, Intersection([]int{1, 2, 3, 4, 5}, []int{0, 2, 4, 6, 8}), []int{2, 4})
	check.Eq(t, Intersection([]int{2, 2, 4}, []int{2, 4}), []int{2, 4})
}

func TestEx8_4(t *testing.T) {
	w, ok := FirstDuplicate([]string{"a", "b", "c", "d", "c", "e", "f"})
	check.Eq(t, []any{w, ok}, []any{"c", true})
	_, ok = FirstDuplicate([]string{"a"})
	check.Eq(t, ok, false)
}

func TestEx8_5(t *testing.T) {
	c, _ := MissingLetter("the quick brown box jumps over a lazy dog")
	check.Eq(t, c, byte('f'))
}

func TestEx8_6(t *testing.T) {
	r, _ := FirstUniqueChar("minimum")
	check.Eq(t, r, 'n')
	_, ok := FirstUniqueChar("aa")
	check.Eq(t, ok, false)
}

func TestEx8_7(t *testing.T) {
	check.Eq(t, fmt.Sprint(WordCounts("the cat and the hat")), "map[and:1 cat:1 hat:1 the:2]")
	check.Eq(t, MostCommonWord("the cat and the hat"), "the")
	check.Eq(t, MostCommonWord("b a"), "b")
}

func TestEx8_8(t *testing.T) {
	check.Eq(t, GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}), [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}})
}

func TestEx8_9(t *testing.T) {
	i, j, ok := TwoSumIndices([]int{2, 7, 11, 15}, 9)
	check.Eq(t, []any{i, j, ok}, []any{0, 1, true})
	_, _, ok = TwoSumIndices([]int{1}, 2)
	check.Eq(t, ok, false)
	i, j, _ = TwoSumIndices([]int{3, 3}, 6)
	check.Eq(t, []int{i, j}, []int{0, 1})
}

func TestEx8_10(t *testing.T) {
	check.Eq(t, fmt.Sprint(Invert(map[string]int{"fries": 2, "soda": 2, "burger": 5})), "map[2:[fries soda] 5:[burger]]")
}
