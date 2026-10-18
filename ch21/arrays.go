// Package ch21: Arrays and dynamic arrays.
package ch21

import (
	"errors"
)

// ==== EX 21.1 | Fixed-capacity Array skeleton | Core
// PROMPT: Create `type Array struct` with a fixed-size `data []int`, a `Length` and a `Steps` counter, plus `NewArray(capacity int) *Array` and methods `IsEmpty()`, `IsFull()`, `ToSlice()` and `ResetSteps()`. Declare the sentinel errors `ErrIndex`, `ErrFull` and `ErrEmpty` for the next exercises.
// EXAMPLE: `a := NewArray(5)` -> `a.Length` is `0`; `a.IsFull()` is `false`
// BIGO: All O(1) except `ToSlice()`, which is O(N).
// NOTE: Capacity is fixed, like a real block of memory. Go slices hide this by growing on `append`; this type makes the cost of every shift visible.
var (
	ErrIndex = errors.New("index out of range")
	ErrFull  = errors.New("array is full")
	ErrEmpty = errors.New("array is empty")
)

type Array struct {
	data   []int
	Length int
	Steps  int
}

func NewArray(capacity int) *Array {
	panic("todo")
}

func (a *Array) IsEmpty() bool { panic("todo") }

func (a *Array) IsFull() bool { panic("todo") }

func (a *Array) ToSlice() []int {
	panic("todo")
}

func (a *Array) ResetSteps() { panic("todo") }

// ==== END

// ==== EX 21.2 | Read by index | Core
// PROMPT: Add `func (a *Array) Read(i int) (int, error)` that returns the element at `i` in one step, or `ErrIndex` if i is outside `0 .. Length-1`.
// EXAMPLE: `[10 20 30]` -> `Read(1)` returns `20, nil`; `Read(3)` returns `ErrIndex`
// BIGO: O(1) time. The computer jumps straight to the memory address.
// NOTE: Check against `Length`, not the capacity: slots past the length hold no real data.
func (a *Array) Read(i int) (int, error) {
	panic("todo")
}

// ==== END

// ==== EX 21.3 | Read the first and last element | Core
// PROMPT: Add `First() (int, bool)` and `Last() (int, bool)`, returning false on an empty array.
// EXAMPLE: `[4 8 15]` -> `First()` is `4, true`, `Last()` is `15, true`
// BIGO: O(1) each. Knowing `Length` makes the last slot one jump away.
// NOTE: The edges matter: insertion and deletion at the end are cheap only because the end is found in one step.
func (a *Array) First() (int, bool) {
	panic("todo")
}

func (a *Array) Last() (int, bool) {
	panic("todo")
}

// ==== END

// ==== EX 21.4 | Update by index | Core
// PROMPT: Add `Update(i, value int) error` that overwrites one slot. Same bounds rule as `Read`.
// EXAMPLE: `[1 2 3]` -> `Update(0, 9)` -> `[9 2 3]`
// BIGO: O(1) time.
func (a *Array) Update(i, value int) error {
	panic("todo")
}

// ==== END

// ==== EX 21.5 | Linear search | Core
// PROMPT: Add `Search(value int) int` that returns the index of the first match, or `-1`. Count one step per element inspected.
// EXAMPLE: `[5 3 7]` -> `Search(7)` is `2` after 3 steps; `Search(9)` is `-1` after 3 steps
// BIGO: O(N) worst case, O(1) best case (match in slot 0).
func (a *Array) Search(value int) int {
	panic("todo")
}

// ==== END

// ==== EX 21.6 | Membership test | Core
// PROMPT: Add `Contains(value int) bool` built on `Search`.
// EXAMPLE: `[5 3 7]` -> `Contains(3)` is `true`
// BIGO: O(N). This is what `slices.Contains` costs too.
func (a *Array) Contains(value int) bool {
	panic("todo")
}

// ==== END

// ==== EX 21.7 | Insert at the end | Core
// PROMPT: Add `InsertAtEnd(value int) error`, returning `ErrFull` when there is no room.
// EXAMPLE: `[1 2]` -> `InsertAtEnd(3)` -> `[1 2 3]` in 1 step
// BIGO: O(1). Nothing has to move.
func (a *Array) InsertAtEnd(value int) error {
	panic("todo")
}

// ==== END

// ==== EX 21.8 | Insert at the start | Core
// PROMPT: Add `InsertAtStart(value int) error`. Shift every element one slot right, starting from the end, then write the new value into slot 0.
// EXAMPLE: `[2 3 4]` -> `InsertAtStart(1)` -> `[1 2 3 4]` in 4 steps (3 shifts + 1 write)
// BIGO: O(N). Worst case for array insertion: N shifts + 1 write = N + 1 steps.
// NOTE: Shift from the right end backwards. Shifting from the left would overwrite values before they move.
func (a *Array) InsertAtStart(value int) error {
	panic("todo")
}

// ==== END

// ==== EX 21.9 | Insert at an index | Core
// PROMPT: Add `InsertAt(i, value int) error` for any i from `0` to `Length` (inserting at `Length` means append).
// EXAMPLE: `[1 2 4]` -> `InsertAt(2, 3)` -> `[1 2 3 4]` in 2 steps
// BIGO: O(N - i): N + 1 steps at the start, 1 step at the end.
func (a *Array) InsertAt(i, value int) error {
	panic("todo")
}

// ==== END

// ==== EX 21.10 | Delete at the end | Core
// PROMPT: Add `DeleteAtEnd() (int, error)` that removes and returns the last element, or `ErrEmpty`.
// EXAMPLE: `[1 2 3]` -> returns `3`, leaves `[1 2]`, 1 step
// BIGO: O(1).
func (a *Array) DeleteAtEnd() (int, error) {
	panic("todo")
}

// ==== END

// ==== EX 21.11 | Delete at the start | Core
// PROMPT: Add `DeleteAtStart() (int, error)` that removes and returns element 0, then shifts every remaining element one slot left to close the gap.
// EXAMPLE: `[1 2 3 4]` -> returns `1`, leaves `[2 3 4]` in 4 steps (1 delete + 3 shifts)
// BIGO: O(N). Worst case for deletion: 1 delete + (N - 1) shifts = N steps.
func (a *Array) DeleteAtStart() (int, error) {
	panic("todo")
}

// ==== END

// ==== EX 21.12 | Delete at an index | Core
// PROMPT: Add `DeleteAt(i int) (int, error)` that removes and returns the element at i and closes the gap.
// EXAMPLE: `[1 2 9 3]` -> `DeleteAt(2)` returns `9`, leaves `[1 2 3]`
// BIGO: O(N - i).
func (a *Array) DeleteAt(i int) (int, error) {
	panic("todo")
}

// ==== END

// ==== EX 21.13 | Delete by value | Core
// PROMPT: Add `DeleteValue(value int) bool` that removes the first occurrence and returns `true`, or returns `false` if absent.
// EXAMPLE: `[4 5 4]` -> `DeleteValue(4)` -> `[5 4]`, returns `true`
// BIGO: O(N) in every case: the search and the shifts together touch every element once.
// NOTE: If the match is near the front, search is short but shifting is long, and the reverse near the back.
func (a *Array) DeleteValue(value int) bool {
	panic("todo")
}

// ==== END

// ==== EX 21.14 | Count and locate every match | Core
// PROMPT: Write `CountOccurrences(arr []int, value int) int` and `FindAllIndices(arr []int, value int) []int`, each with one loop.
// EXAMPLE: `FindAllIndices([]int{3, 1, 3, 3}, 3)` -> `[0 2 3]`; count is `3`
// BIGO: O(N) time; O(1) space for the count, O(K) for the K indices returned.
func CountOccurrences(arr []int, value int) int {
	panic("todo")
}

func FindAllIndices(arr []int, value int) []int {
	panic("todo")
}

// ==== END

// ==== EX 21.15 | Minimum and maximum in one pass | Core
// PROMPT: Write `FindMinMax(arr []int) (lo, hi int, ok bool)` in a single loop without the built-in `min` and `max`. Return ok = false for an empty slice.
// EXAMPLE: `[7 -2 9 4]` -> `-2, 9, true`
// BIGO: O(N) time, O(1) space.
func FindMinMax(arr []int) (lo, hi int, ok bool) {
	panic("todo")
}

// ==== END

// ==== EX 21.16 | Reverse in place | Core
// PROMPT: Write `ReverseInPlace(arr []int)` with two indices that walk toward each other, swapping with Go's parallel assignment.
// EXAMPLE: `[1 2 3 4 5]` -> `[5 4 3 2 1]`
// BIGO: O(N) time (N/2 swaps), O(1) extra space.
func ReverseInPlace(arr []int) {
	panic("todo")
}

// ==== END

// ==== EX 21.17 | Array-based set | Core
// PROMPT: Write `type ArraySet struct{ *Array }` (embedding) with `Add(value int) bool`, and give it its own `InsertAtEnd` that refuses duplicates. Return `true` if added, `false` if already present.
// EXAMPLE: `s.Add(3); s.Add(3)` -> second call returns `false`, `Length` stays 1
// BIGO: O(N): the duplicate check is a full search before every insert. Best-case insert is N + 1 steps, not 1.
// NOTE: Embedding is not inheritance. `s.InsertAtEnd` calls the set's version, but methods defined on `Array` that call `a.InsertAtEnd` still call Array's own: Go has no virtual methods.
type ArraySet struct{ *Array }

func NewArraySet(capacity int) *ArraySet { panic("todo") }

func (s *ArraySet) Add(value int) bool {
	panic("todo")
}

func (s *ArraySet) InsertAtEnd(value int) error {
	panic("todo")
}

// ==== END

// ==== EX 21.18 | Predict the steps | Practice
// PROMPT: Write `StepsTable(n int) map[string]int` that builds an `Array` of N elements and returns the step counts for: read, search for a missing value, insert at end, insert at start, delete at end, delete at start. Predict the numbers on paper first.
// EXAMPLE: `StepsTable(5)` -> `map[deleteEnd:1 deleteStart:5 insertEnd:1 insertStart:6 read:1 searchMissing:5]`
// BIGO: Confirms read O(1), search O(N), insert/delete O(1) at the end and O(N) at the start.
func StepsTable(n int) map[string]int {
	panic("todo")
}

// ==== END

// ==== EX 21.19 | Grow when full (dynamic array) | Stretch
// PROMPT: Write `type DynamicArray struct{ *Array }` whose `InsertAtEnd` doubles the capacity (copying every element to a new block, one step each) when full, instead of failing. Add `Cap() int`.
// EXAMPLE: capacity 2, append 1, 2, 3 -> `Cap()` becomes `4`, contents `[1 2 3]`
// BIGO: A single append is O(N) when it triggers a copy, but O(1) amortized across many appends.
// NOTE: This is what the built-in `append` does when a slice runs out of capacity.
type DynamicArray struct{ *Array }

func NewDynamicArray(capacity int) *DynamicArray { panic("todo") }

func (d *DynamicArray) Cap() int { panic("todo") }

func (d *DynamicArray) InsertAtEnd(value int) error {
	panic("todo")
}

// ==== END

// ==== EX 21.20 | Majority value by voting | Practice
// PROMPT: A value fills more than half of `nums`. Write `Majority(nums []int) int` in one pass with O(1) extra space: keep a candidate and a count; a matching value adds a vote, a different one removes a vote, and an empty count picks the next value as the candidate.
// EXAMPLE: `[3 1 3 3 2 3]` -> `3`
// BIGO: O(N) time, O(1) space: each vote cancels against a different value, and the majority outlasts every cancellation.
func Majority(nums []int) int {
	panic("todo")
}

// ==== END

// ==== EX 21.21 | Products of everything but self | Core
// PROMPT: Write `ProductsExceptSelf(nums []int) []int` returning, for each index, the product of every other value, without using division (so zeros need no special case).
// EXAMPLE: `[2 3 4 5]` -> `[60 40 30 24]`
// BIGO: O(N) time: one pass left to right for the product before each index, one pass right to left for the product after it.
func ProductsExceptSelf(nums []int) []int {
	panic("todo")
}

// ==== END

// ==== EX 21.22 | Longest mountain run | Core
// PROMPT: A mountain is a run that strictly rises and then strictly falls, with at least one step each way. Write `LongestMountain(nums []int) int` returning the length of the longest mountain, or 0 if there is none. Find each tip, then walk down both sides.
// EXAMPLE: `[1 3 5 4 2 2 6 1]` -> `5`
// BIGO: O(N): each tip is found once, and the climbs and descents around different tips never overlap by more than one value.
func LongestMountain(nums []int) int {
	panic("todo")
}

// ==== END

// ==== EX 21.23 | Smallest amount the coins cannot make | Practice
// PROMPT: Each coin in `coins` (positive values, repeats allowed) can be used at most once. Write `SmallestUnmakeable(coins []int) int` returning the smallest positive amount that no selection of coins adds up to. Sort a copy, then track the largest amount `reach` that every smaller amount can be made up to.
// EXAMPLE: `[2 1 6]` -> `4`; `[1 1 3]` -> `6`
// BIGO: O(N log N) for the sort, then one pass: if every amount up to `reach` can be paid and the next coin is at most reach + 1, every amount up to reach + coin can be paid.
func SmallestUnmakeable(coins []int) int {
	panic("todo")
}

// ==== END

// ==== EX 21.24 | Lowest missing positive, in place | Stretch
// PROMPT: Write `LowestMissingPositive(nums []int) int` returning the smallest positive integer not in `nums`, in O(N) time with no extra slice. Rearrange in place so that value `v` (when 1 <= v <= N) sits at index `v - 1`; then the first index whose value is wrong gives the answer.
// EXAMPLE: `[4 -2 1 9 2]` -> `3`; `[2 3]` -> `1`
// BIGO: O(N) time and O(1) extra space: each swap puts one value in its home slot for good, so there are at most N swaps.
func LowestMissingPositive(nums []int) int {
	panic("todo")
}

// ==== END

// ==== EX 21.25 | Missing numbers in 1..n | Core
// PROMPT: `nums` holds distinct values from 1 to `n`. Write `MissingFromRange(nums []int, n int) []int` returning the values in 1..n that are absent, in increasing order, using a slice of `n + 1` marks rather than sorting.
// EXAMPLE: `[4 1 6], 6` -> `[2 3 5]`
// BIGO: O(N) time and O(N) space: one pass marks what is present, one pass collects what is not.
func MissingFromRange(nums []int, n int) []int {
	panic("todo")
}

// ==== END

// ==== EX 21.26 | First repeat, using the array as its own table | Practice
// PROMPT: Every value in `nums` is between 1 and `len(nums)`. Write `FirstRepeatInPlace(nums []int) int` returning the value whose second appearance comes first, or -1, using no extra structure: when you meet `v`, flip `nums[v-1]` negative; finding it already negative means `v` is a repeat. Restore the signs before returning.
// EXAMPLE: `[3 1 4 4 1]` -> `4`; `[1 2 3]` -> `-1`
// BIGO: O(N) time, O(1) extra space: the sign of `nums[v - 1]` records whether `v` has been seen.
func FirstRepeatInPlace(nums []int) int {
	panic("todo")
}

// ==== END
