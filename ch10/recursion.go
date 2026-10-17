// Package ch10: Recursion: a base case and a smaller problem.
package ch10

// ==== EX 10.1 | Countdown | Core
// PROMPT: Write `Countdown(n int) []int` recursively, returning n down to 0. Identify the base case first.
// EXAMPLE: `Countdown(3)` -> `[3 2 1 0]`
// BIGO: O(N) calls and stack depth. Prepending with `append([]int{n}, rest...)` copies, so this version is O(N^2) in total.
func Countdown(n int) []int {
	panic("todo")
}

// ==== END

// ==== EX 10.2 | Factorial | Core
// PROMPT: Write `Factorial(n int) int` recursively.
// EXAMPLE: `Factorial(5)` -> `120`; `Factorial(0)` -> `1`; `Factorial(20)` -> `2432902008176640000`
// BIGO: O(N).
// NOTE: 20! is the largest factorial that fits in a 64-bit int; use `math/big` beyond that.
func Factorial(n int) int {
	panic("todo")
}

// ==== END

// ==== EX 10.3 | Sum an array recursively | Core
// PROMPT: Write `SumRecursive(arr []int) int` that adds `arr[0]` to the sum of `arr[1:]`. Re-slicing is free in Go: it makes a new header, not a copy.
// EXAMPLE: `[1 2 3 4]` -> `10`
// BIGO: O(N) time and stack depth. Unlike Python's slicing, `arr[1:]` costs O(1).
func SumRecursive(arr []int) int {
	panic("todo")
}

// ==== END

// ==== EX 10.4 | Flatten nested lists | Core
// PROMPT: Write `Flatten(nested []any) []int` returning every int in a slice that may contain `[]any` to any depth, in order. Use a type switch.
// EXAMPLE: `[]any{1, []any{2, []any{3, 4}}, 5}` -> `[1 2 3 4 5]`
// BIGO: O(N) in the total number of items. Recursion fits because the depth is unknown in advance.
func Flatten(nested []any) []int {
	panic("todo")
}

// ==== END

// ==== EX 10.5 | Fast power | Practice
// PROMPT: Write `Power(x, n int) int` for n >= 0 by squaring: x^n = (x^(n/2))^2, times x if n is odd.
// EXAMPLE: `Power(2, 10)` -> `1024`; `Power(3, 0)` -> `1`
// BIGO: O(log N): each call halves n.
func Power(x, n int) int {
	panic("todo")
}

// ==== END

// ==== EX 10.6 | Filesystem traversal | Core
// PROMPT: Model a filesystem as `type Dir struct{ Name string; Dirs []Dir; Files []string }`. Write `AllDirectories(d Dir, path string) []string` returning every directory path below d, depth first.
// EXAMPLE: root with dirs `docs` (containing `old` and file `a.txt`) and `src` -> `[/docs /docs/old /src]`
// BIGO: O(N) in the number of entries.
// NOTE: The real version uses `os.ReadDir(path)` and `entry.IsDir()`, or `filepath.WalkDir`: same shape, same recursion.
type Dir struct {
	Name  string
	Dirs  []Dir
	Files []string
}

func AllDirectories(d Dir, path string) []string {
	panic("todo")
}

// ==== END

// ==== EX 10.7 | Every number from low to high | Core
// PROMPT: Write `NumbersBetween(low, high int) []int` recursively returning `[low ... high]`, appending `high` after the recursive call on `high - 1` so nothing is prepended.
// EXAMPLE: `NumbersBetween(3, 6)` -> `[3 4 5 6]`
// BIGO: O(N).
func NumbersBetween(low, high int) []int {
	panic("todo")
}

// ==== END
