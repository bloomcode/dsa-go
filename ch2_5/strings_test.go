package ch2_5

import (
	"testing"

	"dsa/internal/check"
)

func TestEx2_5_1(t *testing.T) {
	for _, c := range [][2]string{{"sun is up", "up is sun"}, {"solo", "solo"}, {"a b", "b a"}, {"", ""}} {
		chars := []byte(c[0])
		ReverseWords(chars)
		check.Eq(t, string(chars), c[1])
	}
}

func TestEx2_5_2(t *testing.T) {
	check.Eq(t, []bool{OneEditApart("cart", "card"), OneEditApart("cart", "art"), OneEditApart("cart", "cards")}, []bool{true, true, false})
	check.Eq(t, []bool{OneEditApart("cart", "cart"), OneEditApart("", "x"), OneEditApart("ab", "ba"), OneEditApart("art", "cart")}, []bool{false, true, false, true})
}

func TestEx2_5_3(t *testing.T) {
	check.Eq(t, []int{FirstUniqueIndex("abacabad"), FirstUniqueIndex("aabb"), FirstUniqueIndex(""), FirstUniqueIndex("z"), FirstUniqueIndex("xxyz")}, []int{3, -1, -1, 0, 2})
}

func TestEx2_5_4(t *testing.T) {
	check.Eq(t, LettersNeeded([]string{"tool", "loot", "lot"}), "loot")
	check.Eq(t, []string{LettersNeeded([]string{"a", "b"}), LettersNeeded([]string{"aa", "a"}), LettersNeeded(nil), LettersNeeded([]string{"abc", "cab", "cc"})}, []string{"ab", "aa", "", "abcc"})
}

func TestEx2_5_5(t *testing.T) {
	check.Eq(t, ReversedPairs([]string{"stop", "tops", "pots", "spot", "level"}), [][2]string{{"stop", "pots"}, {"tops", "spot"}})
	check.Eq(t, ReversedPairs(nil), [][2]string{})
	check.Eq(t, ReversedPairs([]string{"noon", "abc"}), [][2]string{})
	check.Eq(t, ReversedPairs([]string{"ab", "ba"}), [][2]string{{"ab", "ba"}})
}

func TestEx2_5_6(t *testing.T) {
	check.Eq(t, LongestPalindrome("noonabcbad"), "abcba")
	check.Eq(t, []string{LongestPalindrome(""), LongestPalindrome("x"), LongestPalindrome("ab"), LongestPalindrome("cbbd"), LongestPalindrome("aaaa")}, []string{"", "x", "a", "bb", "aaaa"})
}

func TestEx2_5_7(t *testing.T) {
	check.Eq(t, AnagramStarts("tapatpta", "pat"), []int{0, 2, 3, 5})
	check.Eq(t, AnagramStarts("aaa", "aa"), []int{0, 1})
	check.Eq(t, AnagramStarts("ab", "abc"), []int{})
	check.Eq(t, AnagramStarts("xyz", "q"), []int{})
}

func TestEx2_5_8(t *testing.T) {
	check.Eq(t, []int{LongestAfterChanges("aabcabb", 2), LongestAfterChanges("aabcabb", 0), LongestAfterChanges("", 3), LongestAfterChanges("abcd", 9), LongestAfterChanges("abab", 1)}, []int{5, 2, 0, 4, 3})
}

func TestEx2_5_9(t *testing.T) {
	check.Eq(t, UnderscoreMatches("lalala la", "la"), "_lalala_ _la_")
	check.Eq(t, []string{UnderscoreMatches("aaa", "aa"), UnderscoreMatches("abc", "x"), UnderscoreMatches("xx", "x")}, []string{"_aaa_", "abc", "_xx_"})
}

func TestEx2_5_10(t *testing.T) {
	check.Eq(t, []string{SmallestWindow("bxaxcaxb", "abc"), SmallestWindow("abxab", "aab")}, []string{"caxb", "abxa"})
	check.Eq(t, []string{SmallestWindow("abc", "d"), SmallestWindow("a", "a"), SmallestWindow("aa", "aaa"), SmallestWindow("xyz", "")}, []string{"", "a", "", ""})
}
