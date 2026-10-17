package ch53

import (
	"testing"

	"dsa/internal/check"
)

func trie(words ...string) *Trie {
	t := NewTrie()
	for _, w := range words {
		t.Insert(w)
	}
	return t
}

func TestEx53_1(t *testing.T) {
	tr := trie("pin", "pit")
	check.Eq(t, len(tr.Root.Children), 1)
	a := tr.Root.Children['p'].Children['i']
	check.Eq(t, string(sortedKeys(a)), "nt")
	check.True(t, a.Children['n'].End && !a.End, "end marks")
}

func TestEx53_2(t *testing.T) {
	tr := trie("pin", "pit")
	check.True(t, tr.Search("pi") != nil && tr.Search("po") == nil, "search")
}

func TestEx53_3(t *testing.T) {
	tr := trie("pin", "pit")
	check.True(t, tr.ContainsWord("pin") && !tr.ContainsWord("pi"), "contains")
}

func TestEx53_4(t *testing.T) {
	tr := trie("pin", "pit", "sun")
	check.Eq(t, tr.CollectAllWords(tr.Root), []string{"pin", "pit", "sun"})
	empty := NewTrie()
	check.Eq(t, len(empty.CollectAllWords(empty.Root)), 0)
}

func TestEx53_5(t *testing.T) {
	tr := trie("pin", "pinch", "pit", "sun")
	check.Eq(t, tr.Autocomplete("pin"), []string{"pin", "pinch"})
	check.True(t, tr.Autocomplete("x") == nil, "no completions")
}

func TestEx53_6(t *testing.T) {
	check.Eq(t, trie("pie", "pin").AllKeys(), "pien")
}

func TestEx53_7(t *testing.T) {
	check.Eq(t, trie("pine", "pint", "pinto").LongestCommonPrefix(), "pin")
	check.Eq(t, []string{trie("pin", "pinch").LongestCommonPrefix(), trie("sun").LongestCommonPrefix()}, []string{"pin", "sun"})
	check.Eq(t, []string{trie("pin", "sun").LongestCommonPrefix(), NewTrie().LongestCommonPrefix()}, []string{"", ""})
}

func TestEx53_8(t *testing.T) {
	rt := NewTrie()
	rt.InsertRanked("pinch", 9)
	rt.InsertRanked("pin", 5)
	rt.InsertRanked("pinball", 7)
	check.Eq(t, rt.TopCompletions("pin", 2), []string{"pinch", "pinball"})
	check.Eq(t, rt.TopCompletions("pin", 9), []string{"pinch", "pinball", "pin"})
}

func TestEx53_9(t *testing.T) {
	tr := trie("pin", "pinch")
	check.True(t, tr.Delete("pinch") && tr.ContainsWord("pin") && !tr.ContainsWord("pinch"), "delete")
	check.Eq(t, len(tr.Search("pin").Children), 0)
	check.True(t, !tr.Delete("sun") && !tr.Delete("pi") && tr.Delete("pin"), "delete more")
	check.Eq(t, len(tr.Root.Children), 0)
}
