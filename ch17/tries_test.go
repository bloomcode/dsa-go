package ch17

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

func TestEx17_1(t *testing.T) {
	tr := trie("cat", "cab")
	check.Eq(t, len(tr.Root.Children), 1)
	a := tr.Root.Children['c'].Children['a']
	check.Eq(t, string(sortedKeys(a)), "bt")
	check.True(t, a.Children['t'].End && !a.End, "end marks")
}

func TestEx17_2(t *testing.T) {
	tr := trie("cat", "cab")
	check.True(t, tr.Search("ca") != nil && tr.Search("co") == nil, "search")
}

func TestEx17_3(t *testing.T) {
	tr := trie("cat", "cab")
	check.True(t, tr.ContainsWord("cat") && !tr.ContainsWord("ca"), "contains")
}

func TestEx17_4(t *testing.T) {
	tr := trie("cat", "cab", "dog")
	check.Eq(t, tr.CollectAllWords(tr.Root), []string{"cab", "cat", "dog"})
	empty := NewTrie()
	check.Eq(t, len(empty.CollectAllWords(empty.Root)), 0)
}

func TestEx17_5(t *testing.T) {
	tr := trie("cat", "catnip", "cab", "dog")
	check.Eq(t, tr.Autocomplete("cat"), []string{"cat", "catnip"})
	check.True(t, tr.Autocomplete("x") == nil, "no completions")
}

func TestEx17_6(t *testing.T) {
	check.Eq(t, trie("can", "cat").AllKeys(), "cant")
}

func TestEx17_7(t *testing.T) {
	tr := trie("cat", "catnap", "catnip")
	check.Eq(t, []string{tr.Autocorrect("catnar"), tr.Autocorrect("cat"), tr.Autocorrect("zebra")}, []string{"catnap", "cat", "cat"})
}

func TestEx17_8(t *testing.T) {
	rt := NewTrie()
	rt.InsertRanked("catnip", 9)
	rt.InsertRanked("cat", 5)
	rt.InsertRanked("catalog", 7)
	check.Eq(t, rt.TopCompletions("cat", 2), []string{"catnip", "catalog"})
	check.Eq(t, rt.TopCompletions("cat", 9), []string{"catnip", "catalog", "cat"})
}

func TestEx17_9(t *testing.T) {
	tr := trie("cat", "catnip")
	check.True(t, tr.Delete("catnip") && tr.ContainsWord("cat") && !tr.ContainsWord("catnip"), "delete")
	check.Eq(t, len(tr.Search("cat").Children), 0)
	check.True(t, !tr.Delete("dog") && !tr.Delete("ca") && tr.Delete("cat"), "delete more")
	check.Eq(t, len(tr.Root.Children), 0)
}
