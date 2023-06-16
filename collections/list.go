package collections

import "github.com/bgaudino/go-helpers/functools"

type List[T any] []T

type ComparableList[T comparable] List[T]

func (list *List[T]) Append(value T) {
	functools.Append(list, value)
}

func (list List[T]) At(index int) T {
	return functools.ValueAt(list, index)
}

func (list *List[T]) Clear() {
	functools.Clear(list)
}

func (list ComparableList[T]) Contains(value T) bool {
	return functools.Contains(list, value)
}

func (list List[T]) Copy() List[T] {
	return functools.Copy(list)
}

func (list ComparableList[T]) Count(value T) int {
	return functools.Count(list, value)
}

func (list *List[T]) Extend(items []T) {
	functools.Extend(list, items)
}

func (list ComparableList[T]) Every(f func(T) bool) bool {
	return functools.Every(list, f)
}

func (list List[T]) ForEach(f func(int, T)) {
	functools.ForEach(list, f)
}

func (list List[T]) ForEachRef(f func(int, *T)) {
	functools.ForEachRef(list, f)
}

func (list ComparableList[T]) Filter(f func(T) bool) ComparableList[T] {
	return functools.Filter(list, f)
}

func (list List[T]) Find(f func(T) bool) (value T, found bool) {
	return functools.Find(list, f)
}

func (list ComparableList[T]) Index(value T) int {
	return functools.Index(list, value)
}

func (list *List[T]) Insert(value T, index int) {
	functools.Insert(list, value, index)
}

func (list List[T]) Map(f func(T) any) []any {
	return functools.Map(list, f)
}

func (list *List[T]) Pop() T {
	return functools.Pop(list)
}

func (list List[T]) Reduce(f func(any, T) any, acc any) any {
	return functools.Reduce(list, f, acc)
}

func (list *List[T]) Remove(index int) T {
	return functools.Remove(list, index)
}

func (list *List[T]) Reverse() {
	functools.Reverse(list)
}

func (list *List[T]) Shift() T {
	return functools.Shift(list)
}

func (list List[T]) Some(f func(T) bool) bool {
	return functools.Some(list, f)
}

func (list *List[T]) UnShift(value T) {
	functools.UnShift(list, value)
}
