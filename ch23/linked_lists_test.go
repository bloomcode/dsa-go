package ch23

import (
	"errors"
	"slices"
	"testing"

	"dsa/internal/check"
)

func TestEx23_1(t *testing.T) {
	check.Eq(t, FromSlice(1, 2, 3).ToSlice(), []int{1, 2, 3})
	check.Eq(t, []int{FromSlice(1, 2, 3).Len(), FromSlice().Len()}, []int{3, 0})
	one := FromSlice(7)
	check.True(t, one.Head.Data == 7 && one.Head.Next == nil, "one node")
}

func TestEx23_2(t *testing.T) {
	l := FromSlice(10, 20, 30)
	v, ok := l.Read(2)
	check.Eq(t, []any{v, ok}, []any{30, true})
	_, ok = l.Read(5)
	check.Eq(t, ok, false)
	_, ok = l.Read(-1)
	check.Eq(t, ok, false)
}

func TestEx23_3(t *testing.T) {
	l := FromSlice(10, 20, 30)
	check.Eq(t, []int{l.IndexOf(20), l.IndexOf(99)}, []int{1, -1})
}

func TestEx23_4(t *testing.T) {
	l := FromSlice(2, 3)
	l.InsertAtStart(1)
	check.Eq(t, l.ToSlice(), []int{1, 2, 3})
	e := FromSlice()
	e.InsertAtStart(1)
	check.Eq(t, e.ToSlice(), []int{1})
}

func TestEx23_5(t *testing.T) {
	l := FromSlice(1, 2)
	l.InsertAtEnd(3)
	check.Eq(t, l.ToSlice(), []int{1, 2, 3})
	e := FromSlice()
	e.InsertAtEnd(7)
	check.Eq(t, e.ToSlice(), []int{7})
}

func TestEx23_6(t *testing.T) {
	l := FromSlice(1, 2, 4)
	l.InsertAt(2, 3)
	l.InsertAt(4, 5)
	l.InsertAt(0, 0)
	check.Eq(t, l.ToSlice(), []int{0, 1, 2, 3, 4, 5})
	check.True(t, errors.Is(l.InsertAt(9, 1), ErrIndex), "insert past end")
}

func TestEx23_7(t *testing.T) {
	l := FromSlice(1, 2, 3)
	v, _ := l.DeleteAtStart()
	check.Eq(t, []any{v, l.ToSlice()}, []any{1, []int{2, 3}})
	_, ok := FromSlice().DeleteAtStart()
	check.Eq(t, ok, false)
}

func TestEx23_8(t *testing.T) {
	l := FromSlice(1, 2, 3)
	v, _ := l.DeleteAtEnd()
	check.Eq(t, []any{v, l.ToSlice()}, []any{3, []int{1, 2}})
	l = FromSlice(1)
	v, _ = l.DeleteAtEnd()
	check.Eq(t, []any{v, l.ToSlice()}, []any{1, []int{}})
}

func TestEx23_9(t *testing.T) {
	l := FromSlice(1, 2, 9, 3)
	v, _ := l.DeleteAt(2)
	check.Eq(t, []any{v, l.ToSlice()}, []any{9, []int{1, 2, 3}})
	_, err := l.DeleteAt(3)
	check.True(t, errors.Is(err, ErrIndex), "delete past end")
	v, _ = l.DeleteAt(0)
	check.Eq(t, []any{v, l.ToSlice()}, []any{1, []int{2, 3}})
}

func TestEx23_10(t *testing.T) {
	l := FromSlice(4, 5, 4)
	check.True(t, l.DeleteValue(4), "first")
	check.Eq(t, l.ToSlice(), []int{5, 4})
	check.True(t, l.DeleteValue(4) && !l.DeleteValue(4), "second, then absent")
	check.Eq(t, l.ToSlice(), []int{5})
}

func TestEx23_11(t *testing.T) {
	l := FromSlice(1, 2, 3, 4, 5)
	check.Eq(t, l.SplitHalf().ToSlice(), []int{4, 5})
	check.Eq(t, l.ToSlice(), []int{1, 2, 3})
	even := FromSlice(1, 2, 3, 4)
	check.Eq(t, even.SplitHalf().ToSlice(), []int{3, 4})
	check.Eq(t, even.ToSlice(), []int{1, 2})
	one := FromSlice(7)
	check.Eq(t, []int{one.SplitHalf().Len(), one.Len(), FromSlice().SplitHalf().Len()}, []int{0, 1, 0})
}

func TestEx23_12(t *testing.T) {
	l := FromSlice(1, 2, 3)
	l.Reverse()
	check.Eq(t, l.ToSlice(), []int{3, 2, 1})
	e := FromSlice()
	e.Reverse()
	check.Eq(t, e.ToSlice(), []int{})
}

func TestEx23_13(t *testing.T) {
	a, b := FromSlice(1, 4, 7), FromSlice(2, 3, 9)
	firstB := b.Head
	merged := MergeSortedLists(a, b)
	check.Eq(t, merged.ToSlice(), []int{1, 2, 3, 4, 7, 9})
	check.Eq(t, merged.Head.Next == firstB, true)
	check.Eq(t, MergeSortedLists(FromSlice(), FromSlice(5)).ToSlice(), []int{5})
	check.Eq(t, MergeSortedLists(FromSlice(2, 2), FromSlice(2)).ToSlice(), []int{2, 2, 2})
	check.Eq(t, MergeSortedLists(FromSlice(), FromSlice()).Len(), 0)
}

func TestEx23_14(t *testing.T) {
	l := FromSlice(1, 2, 3, 4)
	check.Eq(t, l.HasCycle(), false)
	l.Head.Next.Next.Next.Next = l.Head.Next
	check.Eq(t, l.HasCycle(), true)
	check.Eq(t, FromSlice().HasCycle(), false)
}

func TestEx23_15(t *testing.T) {
	var d DoublyLinkedList
	d.InsertAtEnd(2)
	d.InsertAtStart(1)
	d.InsertAtEnd(3)
	check.Eq(t, []any{d.ToSlice(), d.ToSliceReversed()}, []any{[]int{1, 2, 3}, []int{3, 2, 1}})
	a, _ := d.DeleteAtEnd()
	b, _ := d.DeleteAtStart()
	c, _ := d.DeleteAtStart()
	_, ok := d.DeleteAtEnd()
	check.Eq(t, []any{a, b, c, ok, d.Tail == nil}, []any{3, 1, 2, false, true})
}

func TestEx23_16(t *testing.T) {
	var q LinkedQueue
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	x, _ := q.Dequeue()
	y, _ := q.Read()
	check.Eq(t, []int{x, y}, []int{1, 2})
}

func TestEx23_17(t *testing.T) {
	l := FromSlice(1, 2, 3, 4, 5)
	check.Eq(t, RemoveKthFromEnd(l, 2), 4)
	check.Eq(t, l.ToSlice(), []int{1, 2, 3, 5})
	check.Eq(t, RemoveKthFromEnd(l, 4), 1)
	check.Eq(t, l.ToSlice(), []int{2, 3, 5})
	one := FromSlice(7)
	check.Eq(t, RemoveKthFromEnd(one, 1), 7)
	check.Eq(t, one.ToSlice(), []int{})
}

func TestEx23_18(t *testing.T) {
	check.Eq(t, AddDigitLists(FromSlice(9, 4, 1), FromSlice(3, 8)).ToSlice(), []int{2, 3, 2})
	check.Eq(t, AddDigitLists(FromSlice(9, 9), FromSlice(1)).ToSlice(), []int{0, 0, 1})
	check.Eq(t, AddDigitLists(FromSlice(), FromSlice()).ToSlice(), []int{})
	check.Eq(t, AddDigitLists(FromSlice(0), FromSlice(0)).ToSlice(), []int{0})
	check.Eq(t, AddDigitLists(FromSlice(5), FromSlice()).ToSlice(), []int{5})
}

func TestEx23_19(t *testing.T) {
	shared := FromSlice(8, 9).Head
	a, b := &LinkedList{&Node{1, &Node{2, shared}}}, &LinkedList{&Node{5, shared}}
	check.True(t, MeetingNode(a, b) == shared, "a then b")
	check.True(t, MeetingNode(b, a) == shared, "b then a")
	check.True(t, MeetingNode(FromSlice(1, 2), FromSlice(3)) == nil, "no shared node")
	check.True(t, MeetingNode(a, a) == a.Head, "same list")
	check.True(t, MeetingNode(FromSlice(), b) == nil, "empty list")
}

func TestEx23_20(t *testing.T) {
	l := FromSlice(1, 2, 3, 2, 1)
	check.True(t, IsPalindromeList(l), "odd palindrome")
	check.Eq(t, l.ToSlice(), []int{1, 2, 3, 2, 1})
	l = FromSlice(1, 2, 3, 1)
	check.True(t, !IsPalindromeList(l), "not a palindrome")
	check.Eq(t, l.ToSlice(), []int{1, 2, 3, 1})
	got := []bool{IsPalindromeList(FromSlice(1, 2, 2, 1)), IsPalindromeList(FromSlice(1, 2)), IsPalindromeList(FromSlice()), IsPalindromeList(FromSlice(4))}
	check.Eq(t, got, []bool{true, false, true, true})
}

func TestEx23_21(t *testing.T) {
	for _, c := range []struct {
		k    int
		want []int
	}{{2, []int{4, 5, 1, 2, 3}}, {7, []int{4, 5, 1, 2, 3}}, {0, []int{1, 2, 3, 4, 5}}, {5, []int{1, 2, 3, 4, 5}}, {1, []int{5, 1, 2, 3, 4}}} {
		l := FromSlice(1, 2, 3, 4, 5)
		RotateRight(l, c.k)
		check.Eq(t, l.ToSlice(), c.want)
	}
	empty := FromSlice()
	RotateRight(empty, 3)
	check.Eq(t, empty.ToSlice(), []int{})
}

func TestEx23_22(t *testing.T) {
	for _, c := range [][2][]int{{{1, 2, 3, 4, 5}, {1, 5, 2, 4, 3}}, {{1, 2, 3, 4}, {1, 4, 2, 3}}, {{1, 2}, {1, 2}}, {{1}, {1}}, {{}, {}}} {
		l := FromSlice(c[0]...)
		InterleaveEnds(l)
		check.Eq(t, l.ToSlice(), c[1])
	}
}

func TestEx23_23(t *testing.T) {
	l := FromSlice(1, 2, 3, 4, 5)
	first := l.Head
	SwapPairs(l)
	check.Eq(t, l.ToSlice(), []int{2, 1, 4, 3, 5})
	check.True(t, l.Head.Next == first, "nodes were relinked, not copied")
	for _, c := range [][2][]int{{{}, {}}, {{1}, {1}}, {{1, 2}, {2, 1}}, {{1, 2, 3, 4}, {2, 1, 4, 3}}} {
		l := FromSlice(c[0]...)
		SwapPairs(l)
		check.Eq(t, l.ToSlice(), c[1])
	}
}

func TestEx23_24(t *testing.T) {
	l := FromSlice(3, 1, 2)
	three := l.Head
	SortList(l)
	check.Eq(t, l.ToSlice(), []int{1, 2, 3})
	check.True(t, l.Head.Next.Next == three, "nodes were relinked, not copied")
	for _, values := range [][]int{{5, 1, 4, 2, 3, 1}, {}, {2, 1}, {7}, {4, 4, 1, 9, 0, 3, 8}} {
		l := FromSlice(values...)
		SortList(l)
		want := slices.Clone(values)
		slices.Sort(want)
		check.Eq(t, l.ToSlice(), want)
	}
}

func TestEx23_25(t *testing.T) {
	for _, c := range [][2][]int{{{1, 1, 2, 3, 3, 3}, {1, 2, 3}}, {{}, {}}, {{4, 4}, {4}}, {{1, 2}, {1, 2}}, {{0, 0, 0, 5}, {0, 5}}} {
		l := FromSlice(c[0]...)
		DropSortedDuplicates(l)
		check.Eq(t, l.ToSlice(), c[1])
	}
}

func TestEx23_26(t *testing.T) {
	for _, c := range []struct {
		in   []int
		x    int
		want []int
	}{{[]int{4, 1, 6, 2, 5, 2}, 4, []int{1, 2, 2, 4, 6, 5}}, {[]int{3, 1}, 0, []int{3, 1}}, {[]int{}, 5, []int{}}, {[]int{5, 6, 1}, 9, []int{5, 6, 1}}, {[]int{9, 2, 8, 1}, 5, []int{2, 1, 9, 8}}} {
		l := FromSlice(c.in...)
		PartitionList(l, c.x)
		check.Eq(t, l.ToSlice(), c.want)
	}
}

func TestEx23_27(t *testing.T) {
	n := make([]*MultiNode, 9)
	for v := 1; v <= 8; v++ {
		n[v] = &MultiNode{Data: v}
	}
	n[1].Next, n[2].Next, n[3].Next, n[5].Next = n[2], n[3], n[4], n[6]
	n[2].Child, n[5].Child, n[4].Child = n[5], n[7], n[8]
	values := []int{}
	for node := Flatten(n[1]); node != nil; node = node.Next {
		check.True(t, node.Child == nil, "child links cleared")
		values = append(values, node.Data)
	}
	check.Eq(t, values, []int{1, 2, 5, 7, 6, 3, 4, 8})
	check.True(t, Flatten(nil) == nil, "empty list")
}

func TestEx23_28(t *testing.T) {
	digitSquares := func(x int) int {
		sum := 0
		for ; x > 0; x /= 10 {
			sum += (x % 10) * (x % 10)
		}
		return sum
	}
	check.Eq(t, CycleStartAndLength(func(x int) int { return (x*x + 1) % 11 }, 0), [2]int{4, 2})
	check.Eq(t, CycleStartAndLength(func(x int) int { return (x + 1) % 5 }, 0), [2]int{0, 5})
	check.Eq(t, CycleStartAndLength(func(x int) int { return x / 2 }, 20), [2]int{5, 1})
	check.Eq(t, CycleStartAndLength(digitSquares, 7), [2]int{5, 1})
	check.Eq(t, CycleStartAndLength(digitSquares, 4), [2]int{0, 8})
}
