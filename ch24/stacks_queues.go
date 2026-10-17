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
