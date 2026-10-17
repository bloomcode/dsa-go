package ch23

import (
	"errors"
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
