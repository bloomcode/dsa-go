package ch09

import (
	"testing"

	"dsa/internal/check"
)

func TestEx9_1(t *testing.T) {
	var s Stack[int]
	s.Push(1)
	s.Push(2)
	s.Push(3)
	r, _ := s.Read()
	p, _ := s.Pop()
	r2, _ := s.Read()
	check.Eq(t, []int{r, p, r2, s.Len()}, []int{3, 3, 2, 2})
	var e Stack[int]
	_, ok := e.Pop()
	check.Eq(t, ok, false)
}

func TestEx9_2(t *testing.T) {
	var q Queue[int]
	for i := 1; i <= 3; i++ {
		q.Enqueue(i)
	}
	d, _ := q.Dequeue()
	f, _ := q.Read()
	check.Eq(t, []int{d, f}, []int{1, 2})
	var big Queue[int]
	for i := 0; i < 1000; i++ {
		big.Enqueue(i)
		if i%3 == 0 {
			big.Dequeue()
		}
	}
	first, _ := big.Read()
	check.Eq(t, []int{big.Len(), first}, []int{666, 334})
}

func TestEx9_3(t *testing.T) {
	check.Eq(t, Lint("(var x = {y: [1, 2, 3]})"), "")
	check.Eq(t, Lint("(var x = 2;"), "( has no closing brace")
	check.Eq(t, Lint("var x = 2)"), ") has no opening brace")
	check.Eq(t, Lint("(var x = [1, 2)"), ") has mismatched opening brace")
}

func TestEx9_4(t *testing.T) {
	check.Eq(t, []string{ReverseWithStack("abcde"), ReverseWithStack("")}, []string{"edcba", ""})
}

func TestEx9_5(t *testing.T) {
	var pm PrintManager
	pm.QueueJob("First")
	pm.QueueJob("Second")
	pm.QueueJob("Third")
	check.Eq(t, pm.Run(), []string{"First", "Second", "Third"})
}

func TestEx9_6(t *testing.T) {
	var sq StackQueue[int]
	sq.Enqueue(1)
	sq.Enqueue(2)
	a, _ := sq.Dequeue()
	sq.Enqueue(3)
	b, _ := sq.Dequeue()
	c, _ := sq.Dequeue()
	_, ok := sq.Dequeue()
	check.Eq(t, []any{a, b, c, ok}, []any{1, 2, 3, false})
}

func TestEx9_7(t *testing.T) {
	v, err := EvalRPN([]string{"3", "4", "+", "2", "*"})
	check.Eq(t, []any{v, err}, []any{14.0, nil})
	v, _ = EvalRPN([]string{"8", "2", "-"})
	check.Eq(t, v, 6.0)
	v, _ = EvalRPN([]string{"9", "2", "/"})
	check.Eq(t, v, 4.5)
	_, err = EvalRPN([]string{"1", "+"})
	check.True(t, err != nil, "malformed")
	_, err = EvalRPN([]string{"1", "x"})
	check.True(t, err != nil, "bad token")
}

func TestEx9_8(t *testing.T) {
	var ed Editor
	ed.Type("a")
	ed.Type("b")
	check.Eq(t, ed.Undo(), "a")
	check.Eq(t, ed.Redo(), "ab")
	ed.Undo()
	ed.Type("c")
	check.Eq(t, ed.Redo(), "ac")
}

func TestEx9_9(t *testing.T) {
	cq := NewCircularQueue(3)
	check.True(t, cq.Enqueue(1) && cq.Enqueue(2) && cq.Enqueue(3) && !cq.Enqueue(4), "full")
	x, _ := cq.Dequeue()
	check.Eq(t, x, 1)
	check.True(t, cq.Enqueue(4), "wrap")
	check.Eq(t, cq.data[0], 4)
	var got []int
	for {
		v, ok := cq.Dequeue()
		if !ok {
			break
		}
		got = append(got, v)
	}
	check.Eq(t, got, []int{2, 3, 4})
}
