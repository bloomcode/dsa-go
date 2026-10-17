// Package ch14: Linked lists, singly and doubly.
package ch14

import "errors"

// ==== EX 14.1 | Node and LinkedList skeleton | Core
// PROMPT: Create `type Node struct{ Data int; Next *Node }` and `type LinkedList struct{ Head *Node }`. Add `Len()` and `ToSlice()` that walk the chain, plus `FromSlice(values ...int) *LinkedList` for tests.
// EXAMPLE: `FromSlice(1, 2, 3).ToSlice()` -> `[1 2 3]`
// BIGO: `Len` and `ToSlice` are O(N): there is no length stored, so every node must be visited.
// NOTE: Nodes can live anywhere in memory; each one only knows where the next one is. That is what makes reading slow and insertion at the head fast.
type Node struct {
	Data int
	Next *Node
}

type LinkedList struct{ Head *Node }

var ErrIndex = errors.New("index out of range")

func FromSlice(values ...int) *LinkedList {
	panic("todo")
}

func (l *LinkedList) Len() int {
	panic("todo")
}

func (l *LinkedList) ToSlice() []int {
	panic("todo")
}

// ==== END

// ==== EX 14.2 | Read by index | Core
// PROMPT: Add `Read(i int) (int, bool)`: follow `Next` links i times. Return false if the list is too short.
// EXAMPLE: `[10 20 30]` -> `Read(2)` is `30, true`; `Read(5)` is `0, false`
// BIGO: O(N). An array reads in O(1); this is the linked list's biggest cost.
func (l *LinkedList) Read(i int) (int, bool) {
	panic("todo")
}

// ==== END

// ==== EX 14.3 | Search for a value | Core
// PROMPT: Add `IndexOf(value int) int` returning the position of the first match, or -1.
// EXAMPLE: `[10 20 30]` -> `IndexOf(20)` is `1`
// BIGO: O(N), the same as linear search in an array.
func (l *LinkedList) IndexOf(value int) int {
	panic("todo")
}

// ==== END

// ==== EX 14.4 | Insert at the start | Core
// PROMPT: Add `InsertAtStart(value int)`: the new node points at the old head and becomes the head.
// EXAMPLE: `[2 3]` -> `InsertAtStart(1)` -> `[1 2 3]`
// BIGO: O(1). The array's worst case is the linked list's best case.
func (l *LinkedList) InsertAtStart(value int) {
	panic("todo")
}

// ==== END

// ==== EX 14.5 | Insert at the end | Core
// PROMPT: Add `InsertAtEnd(value int)`: walk to the last node and attach the new one. Handle the empty list. Try the pointer-to-pointer version: walk `p := &l.Head` until `*p == nil`, then set `*p`.
// EXAMPLE: `[1 2]` -> `InsertAtEnd(3)` -> `[1 2 3]`
// BIGO: O(N) without a tail pointer (see 14.15 for O(1)).
func (l *LinkedList) InsertAtEnd(value int) {
	panic("todo")
}

// ==== END

// ==== EX 14.6 | Insert at an index | Core
// PROMPT: Add `InsertAt(i, value int) error`: walk to the node just before i, then rewire two links. Return `ErrIndex` if i is past the end.
// EXAMPLE: `[1 2 4]` -> `InsertAt(2, 3)` -> `[1 2 3 4]`
// BIGO: O(N) to find the spot, then O(1) to insert. No shifting, unlike an array.
// NOTE: Set the new node's `Next` before changing `prev.Next`, or the rest of the list is lost.
func (l *LinkedList) InsertAt(i, value int) error {
	panic("todo")
}

// ==== END

// ==== EX 14.7 | Delete at the start | Core
// PROMPT: Add `DeleteAtStart() (int, bool)` returning the removed value: move `Head` to `Head.Next`.
// EXAMPLE: `[1 2 3]` -> returns `1, true`, leaves `[2 3]`
// BIGO: O(1). Arrays need O(N) here.
func (l *LinkedList) DeleteAtStart() (int, bool) {
	panic("todo")
}

// ==== END

// ==== EX 14.8 | Delete at the end | Core
// PROMPT: Add `DeleteAtEnd() (int, bool)` returning the removed value: find the second-to-last node and cut its link. Handle 0 and 1 nodes.
// EXAMPLE: `[1 2 3]` -> returns `3, true`, leaves `[1 2]`
// BIGO: O(N) in a singly linked list, because you need the node before the last one.
func (l *LinkedList) DeleteAtEnd() (int, bool) {
	panic("todo")
}

// ==== END

// ==== EX 14.9 | Delete at an index | Core
// PROMPT: Add `DeleteAt(i int) (int, error)` returning the removed value: point the previous link past the doomed node. Return `ErrIndex` if out of range.
// EXAMPLE: `[1 2 9 3]` -> `DeleteAt(2)` returns `9`, leaves `[1 2 3]`
// BIGO: O(N) to reach it, O(1) to unlink.
func (l *LinkedList) DeleteAt(i int) (int, error) {
	panic("todo")
}

// ==== END

// ==== EX 14.10 | Delete by value | Core
// PROMPT: Add `DeleteValue(value int) bool` that removes the first matching node and returns true, or false.
// EXAMPLE: `[4 5 4]` -> `DeleteValue(4)` -> `[5 4]`
// BIGO: O(N). With a pointer to the link (`p := &l.Head`), deleting the head needs no special case.
func (l *LinkedList) DeleteValue(value int) bool {
	panic("todo")
}

// ==== END

// ==== EX 14.11 | Last value and middle value | Practice
// PROMPT: Add `Last() (int, bool)` and `Middle() (int, bool)`. For Middle, move a slow pointer one step and a fast pointer two steps; when fast reaches the end, slow is in the middle (second middle for an even count).
// EXAMPLE: `[1 2 3 4 5]` -> `Last()` is `5`, `Middle()` is `3`; `[1 2 3 4]` -> `Middle()` is `3`
// BIGO: O(N) each, one pass, O(1) space.
func (l *LinkedList) Last() (int, bool) {
	panic("todo")
}

func (l *LinkedList) Middle() (int, bool) {
	panic("todo")
}

// ==== END

// ==== EX 14.12 | Reverse in place | Core
// PROMPT: Add `Reverse()` that flips every `Next` link using `prev`, `cur` and `next` variables.
// EXAMPLE: `[1 2 3]` -> `[3 2 1]`
// BIGO: O(N) time, O(1) space.
func (l *LinkedList) Reverse() {
	panic("todo")
}

// ==== END

// ==== EX 14.13 | Delete a node given only that node | Practice
// PROMPT: Write `DeleteNode(node *Node)` that removes node from its list when you have no access to the head. (It is never the last node.)
// EXAMPLE: In `a -> b -> c -> d`, `DeleteNode(b)` leaves `a -> c -> d`
// BIGO: O(1). Copy the next node's data in, then skip the next node.
func DeleteNode(node *Node) {
	panic("todo")
}

// ==== END

// ==== EX 14.14 | Detect a cycle | Stretch
// PROMPT: Add `HasCycle() bool` using Floyd's slow/fast pointers: if they ever meet, the list loops.
// EXAMPLE: A list whose last node points back to its second node -> `true`
// BIGO: O(N) time, O(1) space (a map of visited nodes would be O(N) space).
// NOTE: Compare pointers with `==` (same node), not by `Data`: two different nodes may hold equal values.
func (l *LinkedList) HasCycle() bool {
	panic("todo")
}

// ==== END

// ==== EX 14.15 | Doubly linked list | Core
// PROMPT: Build `type DoublyLinkedList struct{ Head, Tail *DNode }` with nodes holding `Prev` and `Next`. Implement `InsertAtStart`, `InsertAtEnd`, `DeleteAtStart() (int, bool)`, `DeleteAtEnd() (int, bool)`, `ToSlice()` and `ToSliceReversed()` (walk back from the tail).
// EXAMPLE: InsertAtEnd 2, InsertAtStart 1, InsertAtEnd 3 -> `[1 2 3]`; reversed `[3 2 1]`
// BIGO: O(1) at both ends, because the tail pointer and `Prev` links remove the walk. `container/list` is the standard library's doubly linked list.
type DNode struct {
	Data       int
	Prev, Next *DNode
}

type DoublyLinkedList struct{ Head, Tail *DNode }

func (d *DoublyLinkedList) InsertAtStart(v int) {
	panic("todo")
}

func (d *DoublyLinkedList) InsertAtEnd(v int) {
	panic("todo")
}

func (d *DoublyLinkedList) DeleteAtStart() (int, bool) {
	panic("todo")
}

func (d *DoublyLinkedList) DeleteAtEnd() (int, bool) {
	panic("todo")
}

func (d *DoublyLinkedList) ToSlice() []int {
	panic("todo")
}

func (d *DoublyLinkedList) ToSliceReversed() []int {
	panic("todo")
}

// ==== END

// ==== EX 14.16 | Queue on a doubly linked list | Core
// PROMPT: Build `type LinkedQueue struct` with `Enqueue`, `Dequeue() (int, bool)` and `Read() (int, bool)` on top of your `DoublyLinkedList`.
// EXAMPLE: enqueue 1, 2, 3 -> dequeue `1`, read `2`
// BIGO: O(1) for every operation: insert at the tail, remove from the head.
type LinkedQueue struct{ items DoublyLinkedList }

func (q *LinkedQueue) Enqueue(v int) { panic("todo") }

func (q *LinkedQueue) Dequeue() (int, bool) { panic("todo") }

func (q *LinkedQueue) Read() (int, bool) {
	panic("todo")
}

// ==== END
