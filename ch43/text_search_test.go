package ch43

import (
	"math/rand/v2"
	"strings"
	"testing"

	"dsa/internal/check"
)

var text = func() string {
	rng := rand.New(rand.NewPCG(1, 0))
	b := make([]byte, 10000)
	for i := range b {
		b[i] = byte('a' + rng.IntN(26))
	}
	return string(b)
}()

var needles = []string{"ab", "abc", "zz", text[500:506]}

// matches finds every match with strings.HasPrefix, independent of the exercises.
func matches(s, needle string) []int {
	out := []int{}
	for i := 0; i+len(needle) <= len(s); i++ {
		if strings.HasPrefix(s[i:], needle) {
			out = append(out, i)
		}
	}
	return out
}

func TestEx43_1(t *testing.T) {
	check.Eq(t, BruteForceSearch("abracadabra", "abra"), []int{0, 7})
	for _, nd := range needles {
		check.Eq(t, len(BruteForceSearch(text, nd)), len(matches(text, nd)))
	}
}

func TestEx43_2(t *testing.T) {
	check.Eq(t, WindowSums([]int{1, 2, 3, 4, 5}, 3), []int{6, 9, 12})
}

func TestEx43_3(t *testing.T) {
	check.Eq(t, RollingHashes("abcd", 2), []int{1, 28, 55})
}

func TestEx43_4(t *testing.T) {
	check.Eq(t, RabinKarp("abracadabra", "abra", Prime), []int{0, 7})
	check.Eq(t, len(RabinKarp("ab", "abc", Prime)), 0)
	for _, nd := range needles { // verified search is exact, even with a tiny prime
		want := matches(text, nd)
		check.Eq(t, len(RabinKarp(text, nd, Prime)), len(want))
		check.Eq(t, len(RabinKarp(text, nd, 7)), len(want))
	}
}

func TestEx43_5(t *testing.T) {
	check.True(t, SpuriousHits(text, "abcde", 7) > 1000, "tiny prime spurious")
	check.Eq(t, SpuriousHits(text, "abcde", Prime), 0)
}

func TestEx43_6(t *testing.T) {
	rng := rand.New(rand.NewPCG(2, 0))
	check.Eq(t, RabinKarpMonteCarlo("abracadabra", "abra", rng), []int{0, 7})
	for _, nd := range needles {
		check.Eq(t, len(RabinKarpMonteCarlo(text, nd, rng)), len(matches(text, nd)))
	}
}

func TestEx43_7(t *testing.T) {
	check.Eq(t, FindAny("thecatsatonthemat", []string{"cat", "mat", "dog"}), []Hit{{3, "cat"}, {14, "mat"}})
}

func TestEx43_8(t *testing.T) {
	check.Eq(t, []string{LongestRepeat("banana"), LongestRepeat("abcd"), LongestRepeat("abcxyzabcq")}, []string{"ana", "", "abc"})
}

func TestEx43_9(t *testing.T) {
	check.Eq(t, PrefixTable("abacabab"), []int{0, 0, 1, 0, 1, 2, 3, 2})
	check.Eq(t, PrefixTable("aaaa"), []int{0, 1, 2, 3})
	check.Eq(t, PrefixTable(""), []int{})
	check.Eq(t, PrefixTable("abcd"), []int{0, 0, 0, 0})
	check.Eq(t, PrefixTable("aabaaab"), []int{0, 1, 0, 1, 2, 2, 3})
}

func TestEx43_10(t *testing.T) {
	check.Eq(t, PrefixSearch("abababca", "abab"), []int{0, 2})
	check.Eq(t, PrefixSearch("aaaa", "aa"), []int{0, 1, 2})
	check.Eq(t, PrefixSearch("abc", "d"), []int{})
	check.Eq(t, PrefixSearch("ab", "abc"), []int{})
	check.Eq(t, PrefixSearch("xabcabcx", "abcx"), []int{4})
	check.Eq(t, PrefixSearch("", "a"), []int{})
}
