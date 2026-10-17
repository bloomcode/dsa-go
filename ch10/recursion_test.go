package ch10

import (
	"testing"

	"dsa/internal/check"
)

func TestEx10_1(t *testing.T) {
	check.Eq(t, Countdown(3), []int{3, 2, 1, 0})
	check.Eq(t, Countdown(0), []int{0})
}

func TestEx10_2(t *testing.T) {
	check.Eq(t, []int{Factorial(5), Factorial(0), Factorial(20)}, []int{120, 1, 2432902008176640000})
}

func TestEx10_3(t *testing.T) {
	check.Eq(t, []int{SumRecursive([]int{1, 2, 3, 4}), SumRecursive(nil)}, []int{10, 0})
}

func TestEx10_4(t *testing.T) {
	check.Eq(t, Flatten([]any{1, []any{2, []any{3, 4}}, 5}), []int{1, 2, 3, 4, 5})
}

func TestEx10_5(t *testing.T) {
	check.Eq(t, []int{Power(2, 10), Power(3, 0), Power(5, 3)}, []int{1024, 1, 125})
}

func TestEx10_6(t *testing.T) {
	root := Dir{Dirs: []Dir{{Name: "docs", Dirs: []Dir{{Name: "old"}}, Files: []string{"a.txt"}}, {Name: "src"}}}
	check.Eq(t, AllDirectories(root, ""), []string{"/docs", "/docs/old", "/src"})
}

func TestEx10_7(t *testing.T) {
	check.Eq(t, NumbersBetween(3, 6), []int{3, 4, 5, 6})
}
