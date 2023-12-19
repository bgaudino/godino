package godino

type List[T any] []T

type ComparableList[T comparable] struct {
	List[T]
}

func (list *List[T]) Append(value T) {
	Append(list, value)
}

func (list List[T]) At(index int) T {
	return ValueAt(list, index)
}

func (list *List[T]) Clear() {
	Clear(list)
}

func (list ComparableList[T]) Contains(value T) bool {
	return Contains(list.List, value)
}

func (list List[T]) Copy() List[T] {
	return Copy(list)
}

func (list ComparableList[T]) Count(value T) int {
	return Count(list.List, value)
}

func (list *List[T]) Extend(items []T) {
	Extend(list, items)
}

func (list ComparableList[T]) Every(f func(T) bool) bool {
	return Every(list.List, f)
}

func (list List[T]) ForEach(f func(int, T)) {
	ForEach(list, f)
}

func (list List[T]) ForEachRef(f func(int, *T)) {
	ForEachRef(list, f)
}

func (list ComparableList[T]) Filter(f func(T) bool) List[T] {
	return Filter(list.List, f)
}

func (list List[T]) Find(f func(T) bool) (value T, found bool) {
	return Find(list, f)
}

func (list ComparableList[T]) Index(value T) int {
	return Index(list.List, value)
}

func (list *List[T]) Insert(value T, index int) {
	Insert(list, value, index)
}

func (list List[T]) Map(f func(T) any) []any {
	return Map(list, f)
}

func (list *List[T]) Pop() T {
	return Pop(list)
}

func (list List[T]) Reduce(f func(any, T) any, acc any) any {
	return Reduce(list, f, acc)
}

func (list *List[T]) Remove(index int) T {
	return Remove(list, index)
}

func (list *List[T]) Reverse() {
	Reverse(list)
}

func (list *List[T]) Shift() T {
	return Shift(list)
}

func (list List[T]) Some(f func(T) bool) bool {
	return Some(list, f)
}

func (list *List[T]) UnShift(value T) {
	UnShift(list, value)
}
