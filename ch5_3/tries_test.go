package ch5_3

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

func TestEx5_3_1(t *testing.T) {
	tr := trie("pin", "pit")
	check.Eq(t, len(tr.Root.Children), 1)
	a := tr.Root.Children['p'].Children['i']
	check.Eq(t, string(sortedKeys(a)), "nt")
	check.True(t, a.Children['n'].End && !a.End, "end marks")
}

func TestEx5_3_2(t *testing.T) {
	tr := trie("pin", "pit")
	check.True(t, tr.Search("pi") != nil && tr.Search("po") == nil, "search")
}

func TestEx5_3_3(t *testing.T) {
	tr := trie("pin", "pit")
	check.True(t, tr.ContainsWord("pin") && !tr.ContainsWord("pi"), "contains")
}

func TestEx5_3_4(t *testing.T) {
	tr := trie("pin", "pit", "sun")
	check.Eq(t, tr.CollectAllWords(tr.Root), []string{"pin", "pit", "sun"})
	empty := NewTrie()
	check.Eq(t, len(empty.CollectAllWords(empty.Root)), 0)
}

func TestEx5_3_5(t *testing.T) {
	tr := trie("pin", "pinch", "pit", "sun")
	check.Eq(t, tr.Autocomplete("pin"), []string{"pin", "pinch"})
	check.True(t, tr.Autocomplete("x") == nil, "no completions")
}

func TestEx5_3_6(t *testing.T) {
	check.Eq(t, trie("pie", "pin").AllKeys(), "pien")
}

func TestEx5_3_7(t *testing.T) {
	check.Eq(t, trie("pine", "pint", "pinto").LongestCommonPrefix(), "pin")
	check.Eq(t, []string{trie("pin", "pinch").LongestCommonPrefix(), trie("sun").LongestCommonPrefix()}, []string{"pin", "sun"})
	check.Eq(t, []string{trie("pin", "sun").LongestCommonPrefix(), NewTrie().LongestCommonPrefix()}, []string{"", ""})
}

func TestEx5_3_8(t *testing.T) {
	rt := NewTrie()
	rt.InsertRanked("pinch", 9)
	rt.InsertRanked("pin", 5)
	rt.InsertRanked("pinball", 7)
	check.Eq(t, rt.TopCompletions("pin", 2), []string{"pinch", "pinball"})
	check.Eq(t, rt.TopCompletions("pin", 9), []string{"pinch", "pinball", "pin"})
}

func TestEx5_3_9(t *testing.T) {
	tr := trie("pin", "pinch")
	check.True(t, tr.Delete("pinch") && tr.ContainsWord("pin") && !tr.ContainsWord("pinch"), "delete")
	check.Eq(t, len(tr.Search("pin").Children), 0)
	check.True(t, !tr.Delete("sun") && !tr.Delete("pi") && tr.Delete("pin"), "delete more")
	check.Eq(t, len(tr.Root.Children), 0)
}

func TestEx5_3_10(t *testing.T) {
	check.Eq(t, PrefixCounts([]string{"car", "card", "care", "dog"}, []string{"car", "ca", "d", "x"}), []int{3, 3, 1, 0})
	check.Eq(t, PrefixCounts([]string{"a", "a", "ab"}, []string{"a", "ab", "abc", ""}), []int{3, 1, 0, 3})
	check.Eq(t, PrefixCounts(nil, []string{"a"}), []int{0})
}

func TestEx5_3_11(t *testing.T) {
	check.Eq(t, TrieSort([]string{"pear", "apple", "pea", "app"}), []string{"app", "apple", "pea", "pear"})
	check.Eq(t, TrieSort(nil), []string{})
	check.Eq(t, TrieSort([]string{"b", "a", "b"}), []string{"a", "b", "b"})
	check.Eq(t, TrieSort([]string{"ba", "b", "abc", "ab"}), []string{"ab", "abc", "b", "ba"})
}

func TestEx5_3_12(t *testing.T) {
	words := []string{"bad", "bed", "bid", "bud", "bead", "bad"}
	check.Eq(t, WildcardMatches(words, "b.d"), []string{"bad", "bed", "bid", "bud"})
	check.Eq(t, WildcardMatches(words, ".."), []string{})
	check.Eq(t, WildcardMatches(words, "b..d"), []string{"bead"})
	check.Eq(t, WildcardMatches(words, "bad"), []string{"bad"})
	check.Eq(t, WildcardMatches(words, ".e."), []string{"bed"})
}

func TestEx5_3_13(t *testing.T) {
	check.Eq(t, SubstringChecks("banana", []string{"nan", "ana", "nab", ""}), []bool{true, true, false, true})
	check.Eq(t, SubstringChecks("banana", []string{"banana", "bananas", "a"}), []bool{true, false, true})
	check.Eq(t, SubstringChecks("", []string{""}), []bool{true})
	check.Eq(t, SubstringChecks("", []string{"a"}), []bool{false})
}

func TestEx5_3_14(t *testing.T) {
	check.Eq(t, ShortestUniquePrefixes([]string{"zebra", "dog", "duck", "dove"}), []string{"z", "dog", "du", "dov"})
	check.Eq(t, ShortestUniquePrefixes([]string{"a", "b"}), []string{"a", "b"})
	check.Eq(t, ShortestUniquePrefixes([]string{"abc"}), []string{"a"})
	check.Eq(t, ShortestUniquePrefixes([]string{"ab", "ac", "bd"}), []string{"ab", "ac", "b"})
	check.Eq(t, ShortestUniquePrefixes(nil), []string{})
}

func TestEx5_3_15(t *testing.T) {
	words := []string{"good", "home", "gone", "hood", "hoof", "in"}
	check.Eq(t, KeypadWords("4663", words), []string{"gone", "good", "home", "hood", "hoof"})
	check.Eq(t, KeypadWords("46", words), []string{"in"})
	check.Eq(t, KeypadWords("2", []string{"a", "b", "ab"}), []string{"a", "b"})
	check.Eq(t, KeypadWords("", []string{"x"}), []string{})
	check.Eq(t, KeypadWords("9", words), []string{})
}

func TestEx5_3_16(t *testing.T) {
	words := []string{"car", "cat", "core", "tea", "rat", "bin", "rob", "cab", "tin", "ore", "rare"}
	check.Eq(t, BoardWords([]string{"cat", "ore", "bin"}, words), []string{"bin", "car", "cat", "core", "ore", "rat", "rob", "tea"})
	check.Eq(t, BoardWords([]string{"aa"}, []string{"aaa", "aa", "a"}), []string{"a", "aa"})
	check.Eq(t, len(BoardWords(nil, []string{"a"})), 0)
	check.Eq(t, BoardWords([]string{"ab", "cd"}, []string{"abdc", "acdb", "adcb", "bad", "abca"}), []string{"abdc", "acdb", "adcb", "bad"})
}
