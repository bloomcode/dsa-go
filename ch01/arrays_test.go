package ch01

import (
	"errors"
	"fmt"
	"testing"

	"dsa/internal/check"
)

// build returns an Array holding values, filled directly so a test needs only 1.1 and its own exercise.
func build(capacity int, values ...int) *Array {
	a := NewArray(capacity)
	copy(a.data, values)
	a.Length = len(values)
	return a
}

func TestEx1_1(t *testing.T) {
	a := NewArray(5)
	check.Eq(t, []any{a.Length, a.IsEmpty(), a.IsFull(), a.ToSlice(), a.Steps}, []any{0, true, false, []int{}, 0})
	full := build(2, 1, 2)
	check.Eq(t, []any{full.IsFull(), full.IsEmpty(), full.ToSlice()}, []any{true, false, []int{1, 2}})
	full.Steps = 7
	full.ResetSteps()
	check.Eq(t, full.Steps, 0)
}

func TestEx1_2(t *testing.T) {
	a := build(5, 10, 20, 30)
	v, err := a.Read(1)
	check.Eq(t, []any{v, err, a.Steps}, []any{20, nil, 1})
	_, err = a.Read(3)
	check.True(t, errors.Is(err, ErrIndex), "read past length")
	_, err = a.Read(-1)
	check.True(t, errors.Is(err, ErrIndex), "negative index")
}

func TestEx1_3(t *testing.T) {
	a := build(5, 4, 8, 15)
	f, ok1 := a.First()
	l, ok2 := a.Last()
	check.Eq(t, []any{f, ok1, l, ok2}, []any{4, true, 15, true})
	_, ok := NewArray(2).First()
	check.Eq(t, ok, false)
	_, ok = NewArray(2).Last()
	check.Eq(t, ok, false)
}

func TestEx1_4(t *testing.T) {
	a := build(5, 1, 2, 3)
	check.Eq(t, a.Update(0, 9), nil)
	check.Eq(t, []any{a.ToSlice(), a.Steps}, []any{[]int{9, 2, 3}, 1})
	check.True(t, errors.Is(a.Update(3, 0), ErrIndex), "update past length")
}

func TestEx1_5(t *testing.T) {
	a := build(5, 5, 3, 7)
	check.Eq(t, []int{a.Search(7), a.Steps}, []int{2, 3})
	a.ResetSteps()
	check.Eq(t, []int{a.Search(9), a.Steps}, []int{-1, 3})
	a.ResetSteps()
	check.Eq(t, []int{a.Search(5), a.Steps}, []int{0, 1})
}

func TestEx1_6(t *testing.T) {
	a := build(5, 5, 3, 7)
	check.True(t, a.Contains(3) && !a.Contains(4), "contains")
}

func TestEx1_7(t *testing.T) {
	a := build(5, 1, 2)
	check.Eq(t, a.InsertAtEnd(3), nil)
	check.Eq(t, []any{a.ToSlice(), a.Steps}, []any{[]int{1, 2, 3}, 1})
	check.True(t, errors.Is(build(1, 1).InsertAtEnd(2), ErrFull), "full")
}

func TestEx1_8(t *testing.T) {
	a := build(5, 2, 3, 4)
	a.InsertAtStart(1)
	check.Eq(t, []any{a.ToSlice(), a.Steps}, []any{[]int{1, 2, 3, 4}, 4})
}

func TestEx1_9(t *testing.T) {
	a := build(5, 1, 2, 4)
	a.InsertAt(2, 3)
	check.Eq(t, []any{a.ToSlice(), a.Steps}, []any{[]int{1, 2, 3, 4}, 2})
	a.InsertAt(4, 5)
	check.Eq(t, a.ToSlice(), []int{1, 2, 3, 4, 5})
	check.True(t, errors.Is(a.InsertAt(0, 0), ErrFull), "full")
	check.True(t, errors.Is(build(3, 1).InsertAt(5, 0), ErrIndex), "bad index")
}

func TestEx1_10(t *testing.T) {
	a := build(5, 1, 2, 3)
	v, err := a.DeleteAtEnd()
	check.Eq(t, []any{v, err, a.ToSlice(), a.Steps}, []any{3, nil, []int{1, 2}, 1})
	_, err = NewArray(2).DeleteAtEnd()
	check.True(t, errors.Is(err, ErrEmpty), "empty")
}

func TestEx1_11(t *testing.T) {
	a := build(5, 1, 2, 3, 4)
	v, _ := a.DeleteAtStart()
	check.Eq(t, []any{v, a.ToSlice(), a.Steps}, []any{1, []int{2, 3, 4}, 4})
}

func TestEx1_12(t *testing.T) {
	a := build(5, 1, 2, 9, 3)
	v, _ := a.DeleteAt(2)
	check.Eq(t, []any{v, a.ToSlice()}, []any{9, []int{1, 2, 3}})
	_, err := a.DeleteAt(3)
	check.True(t, errors.Is(err, ErrIndex), "bad index")
}

func TestEx1_13(t *testing.T) {
	a := build(5, 4, 5, 4)
	check.True(t, a.DeleteValue(4), "delete value")
	check.Eq(t, a.ToSlice(), []int{5, 4})
	check.True(t, !a.DeleteValue(7), "absent")
	check.Eq(t, a.ToSlice(), []int{5, 4})
}

func TestEx1_14(t *testing.T) {
	check.Eq(t, FindAllIndices([]int{3, 1, 3, 3}, 3), []int{0, 2, 3})
	check.Eq(t, CountOccurrences([]int{3, 1, 3, 3}, 3), 3)
	check.Eq(t, CountOccurrences(nil, 1), 0)
	check.Eq(t, len(FindAllIndices([]int{1, 2}, 9)), 0)
}

func TestEx1_15(t *testing.T) {
	lo, hi, ok := FindMinMax([]int{7, -2, 9, 4})
	check.Eq(t, []any{lo, hi, ok}, []any{-2, 9, true})
	lo, hi, ok = FindMinMax([]int{5})
	check.Eq(t, []any{lo, hi, ok}, []any{5, 5, true})
	_, _, ok = FindMinMax(nil)
	check.Eq(t, ok, false)
}

func TestEx1_16(t *testing.T) {
	r := []int{1, 2, 3, 4, 5}
	ReverseInPlace(r)
	check.Eq(t, r, []int{5, 4, 3, 2, 1})
	r = []int{1, 2}
	ReverseInPlace(r)
	check.Eq(t, r, []int{2, 1})
	ReverseInPlace(nil)
}

func TestEx1_17(t *testing.T) {
	s := NewArraySet(5)
	check.True(t, s.Add(3), "add")
	check.True(t, !s.Add(3), "dup")
	check.Eq(t, s.Length, 1)
	s.ResetSteps()
	s.Add(4)
	check.Eq(t, s.Steps, 2) // 1 search step + 1 write
	check.True(t, s.InsertAtEnd(3) != nil, "dup insert")
	check.Eq(t, s.ToSlice(), []int{3, 4})
}

func TestEx1_18(t *testing.T) {
	check.Eq(t, fmt.Sprint(StepsTable(5)), "map[deleteEnd:1 deleteStart:5 insertEnd:1 insertStart:6 read:1 searchMissing:5]")
}

func TestEx1_19(t *testing.T) {
	d := NewDynamicArray(2)
	for i := 1; i <= 3; i++ {
		check.Eq(t, d.InsertAtEnd(i), nil)
	}
	check.Eq(t, []any{d.Cap(), d.ToSlice()}, []any{4, []int{1, 2, 3}})
}
