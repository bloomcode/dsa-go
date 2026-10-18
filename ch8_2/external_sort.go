// Package ch8_2: External memory: sorting data bigger than RAM.
package ch8_2


// ==== EX 8.2.1 | Make sorted runs that fit in memory | Core
// PROMPT: Write `MakeRuns(data []int, M int) [][]int` that cuts the input into pieces of at most M items and sorts each piece in memory.
// EXAMPLE: `MakeRuns([]int{5, 1, 9, 3, 7, 2, 8}, 3)` -> `[[1 5 9] [2 3 7] [8]]`
// BIGO: O(N log M) CPU, one read and one write pass over the data.
func MakeRuns(data []int, M int) [][]int {
	panic("todo")
}

// ==== END

// ==== EX 8.2.2 | Two-way external mergesort | Core
// PROMPT: Write `TwoWaySort(data []int, M int) (sorted []int, passes int)` that makes runs, then merges runs in pairs, pass after pass, until one run is left. Count the run-making pass too.
// EXAMPLE: `TwoWaySort(16 down to 1, 2)` -> sorted `1..16`, passes `4`: 8 runs, then 4, 2, 1
// BIGO: 1 + ceil(log2(N / M)) passes over the data, each costing O(N / B) I/Os.
func TwoWaySort(data []int, M int) (sorted []int, passes int) {
	panic("todo")
}

func mergePasses(runs [][]int, fanIn int) ([]int, int) {
	panic("todo")
}

// ==== END

// ==== EX 8.2.3 | Merge K sorted arrays with a heap | Core
// PROMPT: Write `KWayMerge(lists [][]int) []int` with a `container/heap` of `{value, list, pos}` entries: pop the smallest, push the next item from the same list.
// EXAMPLE: `KWayMerge([][]int{{1, 4, 9}, {2, 3}, {0, 10}})` -> `[0 1 2 3 4 9 10]`
// BIGO: O(N log K) for N total items in K lists.
type entry struct{ value, list, pos int }

type entryHeap []entry

func (h entryHeap) Len() int           { panic("todo") }
func (h entryHeap) Less(i, j int) bool { panic("todo") }
func (h entryHeap) Swap(i, j int)      { panic("todo") }
func (h *entryHeap) Push(x any)        { panic("todo") }
func (h *entryHeap) Pop() any {
	panic("todo")
}

func KWayMerge(lists [][]int) []int {
	panic("todo")
}

// ==== END

// ==== EX 8.2.4 | M/B-way mergesort | Core
// PROMPT: Write `MultiwaySort(data []int, M, B int) (sorted []int, passes int)` that makes runs of M items and then merges up to `M / B - 1` runs at a time (one input block per run plus one output block).
// EXAMPLE: N = 10,000, M = 100, B = 10 -> 100 runs merged 9 at a time: passes `4` (versus 8 for two-way)
// BIGO: 1 + ceil(log_{M/B}(N / M)) passes. The wider the fan-in, the fewer passes over the disk.
func MultiwaySort(data []int, M, B int) (sorted []int, passes int) {
	panic("todo")
}

// ==== END

// ==== EX 8.2.5 | Predict the passes | Core
// PROMPT: Write `PassesTwoWay(N, M int) int` and `PassesMultiway(N, M, B int) int` from the formulas, and check them against 8.2.2 and 8.2.4.
// EXAMPLE: `PassesTwoWay(10_000, 100)` -> `8`; `PassesMultiway(10_000, 100, 10)` -> `4`
// BIGO: O(log) arithmetic. Count the halvings with integers instead of `math.Log`, which can land a hair off an exact power of two.
func PassesTwoWay(N, M int) int { panic("todo") }

func PassesMultiway(N, M, B int) int { panic("todo") }

func passes(N, M, fanIn int) int {
	panic("todo")
}

// ==== END

// ==== EX 8.2.6 | Count the I/Os of an external sort | Practice
// PROMPT: Write `ExternalSortIOs(N, M, B int) int`: every pass reads and writes the whole data set, which is N / B blocks each way.
// EXAMPLE: `ExternalSortIOs(10_000, 100, 10)` -> `8000` (4 passes x 2 x 1000 blocks)
// BIGO: O((N / B) log_{M/B}(N / M)) I/Os, the known optimum for sorting in external memory.
func ExternalSortIOs(N, M, B int) int {
	panic("todo")
}

// ==== END

// ==== EX 8.2.7 | Sort a real file that "does not fit" | Practice
// PROMPT: Write `SortFile(src, dst string, maxItems int) (runs int, err error)` that sorts a file of one integer per line using at most `maxItems` numbers in memory: write sorted runs to `os.CreateTemp` files, then stream them through a heap of one `bufio.Scanner` per run into dst.
// EXAMPLE: a file of 1,000 shuffled numbers with `maxItems = 100` -> dst holds them sorted, via 10 temporary run files
// BIGO: O(N log N) CPU, two passes over the disk. The same code sorts files far larger than RAM.
// NOTE: `defer` the cleanup right after creating each temporary file, so the files are closed and removed even when a later step returns an error.
func SortFile(src, dst string, maxItems int) (runs int, err error) {
	panic("todo")
}

// ==== END
