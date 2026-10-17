// Package ch09: Stacks, queues and their constraints.
package ch09


// ==== EX 9.1 | Stack | Core
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

// ==== EX 9.2 | Queue | Core
// PROMPT: Build `type Queue[T any] struct` with `Enqueue`, `Dequeue() (T, bool)`, `Read() (T, bool)` (front) and `Len()` on a slice plus a `head` index, compacting when the dead prefix grows past half. Explain why `q = q[1:]` alone keeps memory alive.
// EXAMPLE: enqueue 1, 2, 3 -> `Dequeue()` is `1`, `Read()` is `2`
// BIGO: O(1) amortized each. Re-slicing never frees the front of the backing array, and removing index 0 with a shift would be O(N) (see 1.11).
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

// ==== EX 9.3 | Bracket linter | Core
// PROMPT: Write `Lint(code string) string` using your `Stack[byte]`. Return "" if `()`, `[]`, `{}` are balanced, else a message naming the problem character.
// EXAMPLE: `"(var x = {y: [1, 2, 3]})"` -> `""`; `"(var x = 2;"` -> `"( has no closing brace"`; `"var x = 2)"` -> `") has no opening brace"`
// BIGO: O(N) time, O(N) space for the stack.
func Lint(code string) string {
	panic("todo")
}

// ==== END

// ==== EX 9.4 | Reverse a string with a stack | Core
// PROMPT: Write `ReverseWithStack(s string) string` pushing every rune and popping them back out into a `strings.Builder`.
// EXAMPLE: `"abcde"` -> `"edcba"`
// BIGO: O(N) time and space.
func ReverseWithStack(s string) string {
	panic("todo")
}

// ==== END

// ==== EX 9.5 | Print manager | Core
// PROMPT: Build `type PrintManager struct` with `QueueJob(doc string)` and `Run() []string`, which prints jobs first-come-first-served and returns them in print order.
// EXAMPLE: queue "First", "Second", "Third" -> `Run()` returns `[First Second Third]`
// BIGO: O(1) per job queued or printed.
type PrintManager struct{ queue Queue[string] }

func (p *PrintManager) QueueJob(doc string) { panic("todo") }

func (p *PrintManager) Run() []string {
	panic("todo")
}

// ==== END

// ==== EX 9.6 | Queue built from two stacks | Practice
// PROMPT: Build `type StackQueue[T any] struct` with `Enqueue` and `Dequeue() (T, bool)` using only two `Stack` values.
// EXAMPLE: enqueue 1, 2; dequeue -> `1`; enqueue 3; dequeue -> `2`, dequeue -> `3`
// BIGO: O(1) amortized: each element is moved from the inbox to the outbox at most once.
type StackQueue[T any] struct{ inbox, outbox Stack[T] }

func (q *StackQueue[T]) Enqueue(v T) { panic("todo") }

func (q *StackQueue[T]) Dequeue() (T, bool) {
	panic("todo")
}

// ==== END

// ==== EX 9.7 | Evaluate postfix (RPN) | Practice
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

// ==== EX 9.8 | Undo and redo | Practice
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

// ==== EX 9.9 | Fixed-size circular queue | Stretch
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
