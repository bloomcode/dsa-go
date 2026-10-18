package ch24

import (
	"strconv"
	"testing"

	"dsa/internal/check"
)

func TestEx24_1(t *testing.T) {
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

func TestEx24_2(t *testing.T) {
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

func TestEx24_3(t *testing.T) {
	check.Eq(t, []bool{TagsBalanced("<p><b>hi</b></p>"), TagsBalanced("<p><b></p></b>")}, []bool{true, false})
	check.Eq(t, []bool{TagsBalanced("<p>"), TagsBalanced("</p>"), TagsBalanced("plain text")}, []bool{false, false, true})
	check.Eq(t, []bool{TagsBalanced("<a><b></b><c>x</c></a>"), TagsBalanced("<a></a></a>")}, []bool{true, false})
}

func TestEx24_4(t *testing.T) {
	check.Eq(t, SimplifyPath("/a/./b/../c/"), "/a/c")
	check.Eq(t, SimplifyPath("/../"), "/")
	check.Eq(t, SimplifyPath("/home//user/"), "/home/user")
	check.Eq(t, SimplifyPath("/a/b/../../.."), "/")
	check.Eq(t, SimplifyPath("/x/y/./z"), "/x/y/z")
}

func TestEx24_5(t *testing.T) {
	check.Eq(t, RoundRobin([]string{"a", "b", "c"}, []int{3, 1, 2}, 1), []string{"b", "c", "a"})
	check.Eq(t, RoundRobin([]string{"a", "b", "c"}, []int{3, 1, 2}, 5), []string{"a", "b", "c"})
	check.Eq(t, RoundRobin([]string{"x", "y"}, []int{4, 4}, 2), []string{"x", "y"})
	check.Eq(t, RoundRobin(nil, nil, 3), []string{})
}

func TestEx24_6(t *testing.T) {
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

func TestEx24_7(t *testing.T) {
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

func TestEx24_8(t *testing.T) {
	var ed Editor
	ed.Type("a")
	ed.Type("b")
	check.Eq(t, ed.Undo(), "a")
	check.Eq(t, ed.Redo(), "ab")
	ed.Undo()
	ed.Type("c")
	check.Eq(t, ed.Redo(), "ac")
}

func TestEx24_9(t *testing.T) {
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

func TestEx24_10(t *testing.T) {
	var s MinMaxStack
	for _, v := range []int{5, 2, 8} {
		s.Push(v)
	}
	check.Eq(t, []int{s.Min(), s.Max(), s.Peek()}, []int{2, 8, 8})
	check.Eq(t, s.Pop(), 8)
	check.Eq(t, []int{s.Min(), s.Max()}, []int{2, 5})
	check.Eq(t, s.Pop(), 2)
	check.Eq(t, []int{s.Min(), s.Max()}, []int{5, 5})
	s.Push(5)
	s.Push(-1)
	check.Eq(t, []int{s.Min(), s.Max(), s.Peek()}, []int{-1, 5, -1})
}

func TestEx24_11(t *testing.T) {
	check.Eq(t, NextGreater([]int{2, 7, 3, 5, 4, 6, 8}), []int{7, 8, 5, 6, 6, 8, -1})
	check.Eq(t, NextGreater([]int{}), []int{})
	check.Eq(t, NextGreater([]int{5, 4, 3}), []int{-1, -1, -1})
	check.Eq(t, NextGreater([]int{1, 1, 2}), []int{2, 2, -1})
}

func TestEx24_12(t *testing.T) {
	check.Eq(t, SunsetView([]int{3, 5, 4, 4, 3, 1, 3, 2}), []int{1, 3, 6, 7})
	check.Eq(t, SunsetView(nil), []int{})
	check.Eq(t, SunsetView([]int{1, 2, 3}), []int{2})
	check.Eq(t, SunsetView([]int{3, 2, 1}), []int{0, 1, 2})
	check.Eq(t, SunsetView([]int{2, 2}), []int{1})
}

func TestEx24_13(t *testing.T) {
	got := []int{LargestRectangle([]int{2, 4, 3, 5, 1, 3}), LargestRectangle([]int{2, 1, 2}), LargestRectangle(nil), LargestRectangle([]int{5}), LargestRectangle([]int{1, 1, 1, 1}), LargestRectangle([]int{2, 3, 4, 5})}
	check.Eq(t, got, []int{9, 3, 0, 5, 4, 9})
}

func TestEx24_14(t *testing.T) {
	check.Eq(t, SortStack([]int{3, 1, 4, 2}), []int{1, 2, 3, 4})
	check.Eq(t, SortStack(nil), []int{})
	check.Eq(t, SortStack([]int{2, 2, 1}), []int{1, 2, 2})
	check.Eq(t, SortStack([]int{1, 2, 3}), []int{1, 2, 3})
}

func TestEx24_15(t *testing.T) {
	check.Eq(t, AfterCollisions([]int{-3, 4, 9, -6, 2, -2, 7}), []int{-3, 4, 9, 7})
	check.Eq(t, AfterCollisions([]int{4, 9, -6, 2, -2, -10}), []int{-10})
	check.Eq(t, AfterCollisions([]int{5, -5}), []int{})
	check.Eq(t, AfterCollisions([]int{-1, 1}), []int{-1, 1})
	check.Eq(t, AfterCollisions(nil), []int{})
}

func TestEx24_16(t *testing.T) {
	check.Eq(t, WindowMax([]int{4, 2, 12, 3, 8, 1, 7}, 3), []int{12, 12, 12, 8, 8})
	check.Eq(t, WindowMax([]int{3, 1, 2}, 1), []int{3, 1, 2})
	check.Eq(t, WindowMax([]int{3, 1, 2}, 3), []int{3})
	check.Eq(t, WindowMax([]int{5, 4, 3, 2, 1}, 2), []int{5, 4, 3, 2})
	check.Eq(t, WindowMax([]int{1, 3, 3, 1}, 2), []int{3, 3, 3})
}

func TestEx24_17(t *testing.T) {
	got := []int{}
	for _, s := range []string{"(()())())(", "", ")(", "()(()", "(()", "()()", "))()(())"} {
		got = append(got, LongestBalanced(s))
	}
	check.Eq(t, got, []int{8, 0, 0, 2, 2, 4, 6})
}

func TestEx24_18(t *testing.T) {
	var s QueueStack
	check.True(t, s.IsEmpty(), "new stack is empty")
	for _, v := range []int{1, 2, 3} {
		s.Push(v)
	}
	v, ok := s.Pop()
	check.Eq(t, []any{v, ok}, []any{3, true})
	v, ok = s.Peek()
	check.Eq(t, []any{v, ok}, []any{2, true})
	s.Push(4)
	got := []int{}
	for !s.IsEmpty() {
		v, _ := s.Pop()
		got = append(got, v)
	}
	check.Eq(t, got, []int{4, 2, 1})
	_, ok = s.Pop()
	check.True(t, !ok, "pop on empty")
}

func TestEx24_19(t *testing.T) {
	check.Eq(t, ReverseFirstK([]int{1, 2, 3, 4, 5}, 3), []int{3, 2, 1, 4, 5})
	check.Eq(t, ReverseFirstK([]int{1, 2, 3}, 0), []int{1, 2, 3})
	check.Eq(t, ReverseFirstK([]int{1, 2, 3}, 3), []int{3, 2, 1})
	check.Eq(t, ReverseFirstK([]int{}, 0), []int{})
}

func TestEx24_20(t *testing.T) {
	check.Eq(t, BinaryNumbers(5), []string{"1", "10", "11", "100", "101"})
	check.Eq(t, BinaryNumbers(0), []string{})
	check.Eq(t, BinaryNumbers(1), []string{"1"})
	sixteen := BinaryNumbers(16)
	for i := 1; i <= 16; i++ {
		check.Eq(t, sixteen[i-1], strconv.FormatInt(int64(i), 2))
	}
}
