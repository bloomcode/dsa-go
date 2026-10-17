// Package ch22: Measuring real speed.
package ch22

import (
	"time"
)

// ==== EX 22.1 | Time one call with a monotonic clock | Core
// PROMPT: Write `TimeOnce[T any](fn func() T) (T, time.Duration)` that records `time.Now()` before the call and returns `time.Since(start)` after it.
// EXAMPLE: `TimeOnce(func() int { s := 0; for i := range 1_000_000 { s += i }; return s })` -> `499999500000, 300µs` or so
// BIGO: O(1) overhead. `time.Now` carries a monotonic clock reading, so `time.Since` is safe even if the wall clock is adjusted mid-measurement.
func TimeOnce[T any](fn func() T) (T, time.Duration) {
	panic("todo")
}

// ==== END

// ==== EX 22.2 | Repeat and keep the minimum | Core
// PROMPT: Write `Bench(task func(), number, repeat int) time.Duration` that times `repeat` batches of `number` calls and returns the best time per call. Explain why the minimum, not the average, is the usual summary.
// EXAMPLE: `Bench(func() { slices.Sort(slices.Clone(data)) }, 100, 5)` -> a few microseconds for 1,000 ints
// BIGO: O(number * repeat) calls.
// NOTE: Noise (the garbage collector, other processes) only ever adds time, so the fastest batch is closest to the true cost. For real work, write a `BenchmarkX(b *testing.B)` function and run `go test -bench .`, which picks the iteration count for you.
func Bench(task func(), number, repeat int) time.Duration {
	panic("todo")
}

// ==== END

// ==== EX 22.3 | The sort-the-same-array gotcha | Core
// PROMPT: Write `BenchSort(sortFn func([]int), data []int, repeat int) time.Duration` that sorts a fresh copy of data on every run and times only the sort, not the copy. Then show what goes wrong if every run sorts the same slice.
// EXAMPLE: `BenchSort(slices.Sort[[]int], descending10k, 5)` returns a small positive duration
// BIGO: O(repeat * cost of sortFn).
// NOTE: An in-place sort leaves the slice sorted, so runs 2..n time an already sorted input and look far faster than they are.
func BenchSort(sortFn func([]int), data []int, repeat int) time.Duration {
	panic("todo")
}

// ==== END

// ==== EX 22.4 | Read the Big O off a benchmark | Core
// PROMPT: Write `GrowthExponent(timings map[int]float64) float64` that takes n -> seconds for several sizes and returns the slope of log(time) against log(n). A slope near 1 means O(N), near 2 means O(N^2).
// EXAMPLE: `GrowthExponent(map[int]float64{1000: 0.01, 2000: 0.04, 4000: 0.16})` -> `2`
// BIGO: O(K) for K sizes (least-squares fit). The answer does not depend on the map's random iteration order.
func GrowthExponent(timings map[int]float64) float64 {
	panic("todo")
}

// ==== END

// ==== EX 22.5 | A sorting race | Practice
// PROMPT: Write `Race(sorts map[string]func([]int) []int, data []int) ([]string, error)` where each function returns a new sorted slice; return names ordered fastest first, or an error naming any sort that got the wrong answer.
// EXAMPLE: `Race(map[string]func([]int) []int{"builtin": ..., "mergesort": ...}, data)` -> `[builtin mergesort]` on most machines
// BIGO: Sum of each sort's cost.
func Race(sorts map[string]func([]int) []int, data []int) ([]string, error) {
	panic("todo")
}

// ==== END

// ==== EX 22.6 | Quicksort on a sorted array | Core
// PROMPT: Write `QuicksortComparisons(arr []int) int` counting pivot comparisons for the quicksort from Chapter 13 with the last element as pivot, using an explicit stack of ranges. Compare sorted input with shuffled input of the same size.
// EXAMPLE: `QuicksortComparisons` of `0..99` in order -> `4950`; a shuffled slice of 100 needs far fewer
// BIGO: O(N^2) on sorted input with a last-element pivot, O(N log N) on average.
func QuicksortComparisons(arr []int) int {
	panic("todo")
}

// ==== END

// ==== EX 22.7 | Why adaptive sorts love runs | Practice
// PROMPT: Write `CountRuns(arr []int) int` returning the number of maximal ascending runs. Run-adaptive sorts such as TimSort, and the pattern detection in `slices.Sort`, finish nearly sorted data in close to O(N).
// EXAMPLE: `CountRuns([]int{1, 2, 3, 2, 5, 6, 1})` -> `3`; a sorted slice -> `1`
// BIGO: O(N).
func CountRuns(arr []int) int {
	panic("todo")
}

// ==== END
