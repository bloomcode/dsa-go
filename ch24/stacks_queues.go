// Package ch24: Stacks and queues.
package ch24


// ==== EX 24.1 | Stack | Core
// PROMPT: Build a generic `type Stack[T any] struct` on a slice, exposing only `Push`, `Pop() (T, bool)`, `Read() (T, bool)` (peek), `IsEmpty()` and `Len()`. Pop and Read return false when empty.
// EXAMPLE: push 1, 2, 3 -> `Read()` is `3`, `Pop()` is `3`, `Read()` is `2`
// BIGO: O(1) for every operation (amortized for Push). The end of a slice is the cheap end.
// NOTE: A stack is an abstract data type: the constraint (last in, first out) matters, not the slice underneath.
type Stack[T any] struct{ data []T }

func (s *Stack[T]) Push(v T) { panic("todo") }

func (s *Stack[T]) Pop() (T, bool) {
	panic("todo")
}

func (s *Stack[T]) Read() (T, bool) {
	panic("todo")
}

func (s *Stack[T]) IsEmpty() bool { panic("todo") }

func (s *Stack[T]) Len() int { panic("todo") }

// ==== END

// ==== EX 24.2 | Queue | Core
// PROMPT: Build `type Queue[T any] struct` with `Enqueue`, `Dequeue() (T, bool)`, `Read() (T, bool)` (front) and `Len()` on a slice plus a `head` index, compacting when the dead prefix grows past half. Explain why `q = q[1:]` alone keeps memory alive.
// EXAMPLE: enqueue 1, 2, 3 -> `Dequeue()` is `1`, `Read()` is `2`
// BIGO: O(1) amortized each. Re-slicing never frees the front of the backing array, and removing index 0 with a shift would be O(N) (see 21.11).
type Queue[T any] struct {
	data []T
	head int
}

func (q *Queue[T]) Enqueue(v T) { panic("todo") }

func (q *Queue[T]) Dequeue() (T, bool) {
	panic("todo")
}

func (q *Queue[T]) Read() (T, bool) {
	panic("todo")
}

func (q *Queue[T]) Len() int { panic("todo") }

// ==== END

// ==== EX 24.3 | Do the tags nest? | Core
// PROMPT: Write `TagsBalanced(markup string) bool` using your `Stack`. Tags are `<name>` and `</name>` with letters only and no attributes; text between tags is ignored. Return `true` when every tag closes, in the reverse order it opened.
// EXAMPLE: `"<p><b>hi</b></p>"` -> `true`; `"<p><b></p></b>"` -> `false`
// BIGO: O(N) time, O(N) space for the stack.
func TagsBalanced(markup string) bool {
	panic("todo")
}

// ==== END

// ==== EX 24.4 | Simplify a file path | Core
// PROMPT: Write `SimplifyPath(path string) string` for an absolute path: drop empty parts and `.`, let `..` remove the folder before it (if any), and return the canonical path. Use your `Stack` for the folders.
// EXAMPLE: `"/a/./b/../c/"` -> `"/a/c"`; `"/../"` -> `"/"`
// BIGO: O(N) time and space.
func SimplifyPath(path string) string {
	panic("todo")
}

// ==== END

// ==== EX 24.5 | Round-robin turns | Core
// PROMPT: Write `RoundRobin(names []string, times []int, slice int) []string` using your `Queue`: job `names[i]` needs `times[i]` units of work. The front job runs for up to `slice` units; if it still has work left it goes to the back. Return the names in the order the jobs finish.
// EXAMPLE: `[a b c], [3 1 2], 1` -> `[b c a]`
// BIGO: O(T / s + N) queue operations for total work T, slice s and N jobs; each one is O(1).
func RoundRobin(names []string, times []int, slice int) []string {
	panic("todo")
}

// ==== END

// ==== EX 24.6 | Queue built from two stacks | Practice
// PROMPT: Build `type StackQueue[T any] struct` with `Enqueue` and `Dequeue() (T, bool)` using only two `Stack` values.
// EXAMPLE: enqueue 1, 2; dequeue -> `1`; enqueue 3; dequeue -> `2`, dequeue -> `3`
// BIGO: O(1) amortized: each element is moved from the inbox to the outbox at most once.
type StackQueue[T any] struct{ inbox, outbox Stack[T] }

func (q *StackQueue[T]) Enqueue(v T) { panic("todo") }

func (q *StackQueue[T]) Dequeue() (T, bool) {
	panic("todo")
}

// ==== END

// ==== EX 24.7 | Evaluate postfix (RPN) | Practice
// PROMPT: Write `EvalRPN(tokens []string) (float64, error)` for `+ - * /` using a `Stack[float64]`, returning an error for a malformed expression.
// EXAMPLE: `[3 4 + 2 *]` -> `14, nil`; `[1 +]` -> an error
// BIGO: O(N).
// NOTE: Pop the right operand first: for `8 2 -` the stack gives back 2, then 8.
func EvalRPN(tokens []string) (float64, error) {
	panic("todo")
}

type rpnError string

func (e rpnError) Error() string { panic("todo") }

const errMalformed = rpnError("malformed expression")

// ==== END

// ==== EX 24.8 | Undo and redo | Practice
// PROMPT: Build `type Editor struct{ Text string }` with `Type(s string)`, `Undo() string` and `Redo() string`. Use one stack of past states and one of undone states; new typing clears redo.
// EXAMPLE: type "a", type "b", undo -> `"a"`; redo -> `"ab"`
// BIGO: O(1) per operation (plus the cost of building the new string).
type Editor struct {
	Text       string
	undo, redo Stack[string]
}

func (e *Editor) Type(s string) {
	panic("todo")
}

func (e *Editor) Undo() string {
	panic("todo")
}

func (e *Editor) Redo() string {
	panic("todo")
}

// ==== END

// ==== EX 24.9 | Fixed-size circular queue | Stretch
// PROMPT: Build `type CircularQueue struct` on a fixed `[]int` with `head` and `size`, so neither `Enqueue` nor `Dequeue` shifts anything. `Enqueue` returns false when full; `Dequeue` returns false when empty.
// EXAMPLE: capacity 3: enqueue 1, 2, 3 -> 4th returns `false`; dequeue -> `1`; enqueue 4 succeeds and wraps to slot 0
// BIGO: O(1) for both, with no shifting, because the index wraps with `% capacity`.
type CircularQueue struct {
	data       []int
	head, size int
}

func NewCircularQueue(capacity int) *CircularQueue {
	panic("todo")
}

func (q *CircularQueue) Enqueue(v int) bool {
	panic("todo")
}

func (q *CircularQueue) Dequeue() (int, bool) {
	panic("todo")
}

// ==== END

// ==== EX 24.10 | Min-max stack | Core
// PROMPT: Build `type MinMaxStack struct` with `Push(v int)`, `Pop() int`, `Peek() int`, `Min() int` and `Max() int`, every one O(1). Store each value together with the minimum and maximum of the stack at the moment it was pushed. Callers never pop or read an empty stack.
// EXAMPLE: push 5, 2, 8 -> `Min()` is `2`, `Max()` is `8`; `Pop()` -> `8`; `Max()` is `5`
// BIGO: O(1) per operation, O(N) space: each entry stores the minimum and maximum of itself and everything below it, so popping restores the old ones for free.
type MinMaxStack struct{ entries [][3]int }

func (s *MinMaxStack) Push(v int) {
	panic("todo")
}

func (s *MinMaxStack) Pop() int {
	panic("todo")
}

func (s *MinMaxStack) Peek() int { panic("todo") }

func (s *MinMaxStack) Min() int { panic("todo") }

func (s *MinMaxStack) Max() int { panic("todo") }

// ==== END

// ==== EX 24.11 | Next greater value | Core
// PROMPT: Write `NextGreater(nums []int) []int` that returns, for each index, the first value to its right that is larger, or `-1`. Keep a stack of indices still waiting (a slice); each new value settles every waiting index with a smaller value.
// EXAMPLE: `[2 7 3 5 4 6 8]` -> `[7 8 5 6 6 8 -1]`
// BIGO: O(N) time and space: each index is pushed and popped at most once, and the stack always holds indices still waiting for a larger value.
func NextGreater(nums []int) []int {
	panic("todo")
}

// ==== END

// ==== EX 24.12 | Sunset view | Core
// PROMPT: `heights` lists buildings from west to east, and the sun sets in the east. Write `SunsetView(heights []int) []int` that returns, in ascending order, the indices of the buildings strictly taller than every building east of them (a non-nil slice). Scan west to east with a stack, popping every building that is not taller than the new one.
// EXAMPLE: `[3 5 4 4 3 1 3 2]` -> `[1 3 6 7]`
// BIGO: O(N) time and space: a building that is not taller than a later one can never see past it, so it leaves the stack for good.
func SunsetView(heights []int) []int {
	panic("todo")
}

// ==== END

// ==== EX 24.13 | Largest rectangle under a bar chart | Stretch
// PROMPT: `heights` are the bars of a chart, each 1 wide. Write `LargestRectangle(heights []int) int` that returns the largest rectangle area that fits under the bars. Keep a stack of `[start, height]` with rising heights; a lower bar pops every taller entry, scoring `height * (i - start)`, and inherits the start of the last one popped. A final zero-height bar empties the stack.
// EXAMPLE: `[2 4 3 5 1 3]` -> `9` (height 3 across the bars `4, 3, 5`)
// BIGO: O(N) time and space: a bar is popped exactly when the first lower bar to its right appears, and the entry below it says how far left it reaches.
func LargestRectangle(heights []int) int {
	panic("todo")
}

// ==== END

// ==== EX 24.14 | Sort a stack with one more stack | Practice
// PROMPT: A stack is a slice with its top at the end. Write `SortStack(stack []int) []int` that empties `stack` into a new stack whose values rise from bottom to top (largest on top) and returns it, using only push, pop and a look at the top of the two stacks: pop a value, move every larger value from the new stack back onto the old one, then push it.
// EXAMPLE: push 3, 1, 4, 2 -> popping the result gives `4, 3, 2, 1`
// BIGO: O(N^2) time, O(N) space for the second stack: each value may have to move every larger value aside before it settles.
func SortStack(stack []int) []int {
	panic("todo")
}

// ==== END

// ==== EX 24.15 | Collisions on a track | Practice
// PROMPT: `carts` lists carts on a track from left to right: the size is the absolute value, and the sign is the direction (positive moves right). When two meet, the smaller one is destroyed; equal sizes destroy each other. Write `AfterCollisions(carts []int) []int` that returns the carts left, in order, as a non-nil slice. A stack of survivors works: only a left-mover arriving at a right-mover on top collides.
// EXAMPLE: `[-3 4 9 -6 2 -2 7]` -> `[-3 4 9 7]`
// BIGO: O(N) time and space: every cart is pushed once and destroyed at most once.
func AfterCollisions(carts []int) []int {
	panic("todo")
}

// ==== END

// ==== EX 24.16 | Window maximum with a deque | Practice
// PROMPT: Write `WindowMax(nums []int, k int) []int` that returns the maximum of every window of `k` consecutive values (`1 <= k <= len(nums)`), in O(N). Keep a slice used as a deque of indices whose values fall from front to back: drop smaller values from the back before adding, and drop the front once it leaves the window.
// EXAMPLE: `[4 2 12 3 8 1 7], 3` -> `[12 12 12 8 8]`
// BIGO: O(N) time, O(K) space: the deque keeps indices of falling values, so its front is the window's maximum, and each index enters and leaves once.
func WindowMax(nums []int, k int) []int {
	panic("todo")
}

// ==== END

// ==== EX 24.17 | Longest balanced bracket run | Practice
// PROMPT: `s` holds only `(` and `)`. Write `LongestBalanced(s string) int` that returns the length of the longest substring of correctly matched brackets. Keep a stack of indices starting with `-1`: push each `(`; on `)` pop, then either measure from the new top or, if the stack is empty, push this index as the new boundary.
// EXAMPLE: `"(()())())("` -> `8`
// BIGO: O(N) time and space: the stack holds open-bracket indices above the index of the last unmatched close, which marks where a balanced run can start.
func LongestBalanced(s string) int {
	panic("todo")
}

// ==== END

// ==== EX 24.18 | Stack from a queue | Core
// PROMPT: Build `type QueueStack struct` with `Push(v int)`, `Pop() (int, bool)`, `Peek() (int, bool)` and `IsEmpty() bool` on one slice used only as a queue (append at the back, take from the front). After appending, rotate the older values from the front to the back so the newest is in front. `Pop` and `Peek` return `false` on an empty stack.
// EXAMPLE: push 1, 2, 3; `Pop()` -> `3`; `Peek()` -> `2`
// BIGO: O(N) push, O(1) pop and peek: after each push the queue is rotated so the newest value is at the front.
type QueueStack struct{ queue []int }

func (s *QueueStack) Push(v int) {
	panic("todo")
}

func (s *QueueStack) Pop() (int, bool) {
	panic("todo")
}

func (s *QueueStack) Peek() (int, bool) {
	panic("todo")
}

func (s *QueueStack) IsEmpty() bool { panic("todo") }

// ==== END

// ==== EX 24.19 | Reverse the first k of a queue | Core
// PROMPT: Write `ReverseFirstK(queue []int, k int) []int` that returns the queue with its first `k` values reversed (`0 <= k <= len(queue)`), using only queue steps on it (append at the back, take from the front) and a second slice as a stack: move `k` values onto the stack, append them back, then rotate the remaining `len - k` values from the front to the back.
// EXAMPLE: `[1 2 3 4 5], 3` -> `[3 2 1 4 5]`
// BIGO: O(N) time, O(K) space for the stack: k values go through the stack, then the other N - k rotate to the back once.
func ReverseFirstK(queue []int, k int) []int {
	panic("todo")
}

// ==== END

// ==== EX 24.20 | Binary numbers 1 to n with a queue | Core
// PROMPT: Write `BinaryNumbers(n int) []string` that returns the binary forms of `1` to `n` as strings, in order (a non-nil slice), without converting numbers: start a queue (a slice) with `"1"`, and each time you take a string from the front, record it and add it with `"0"` and `"1"` appended to the back.
// EXAMPLE: `5` -> `["1" "10" "11" "100" "101"]`
// BIGO: O(N log N) time and space: N strings of about log N digits each, every one produced by appending a digit to an earlier one.
func BinaryNumbers(n int) []string {
	panic("todo")
}

// ==== END
