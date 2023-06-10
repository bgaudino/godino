package helpers

import (
	"errors"
)

type List[T any] []T

func NewList[T any](values ...T) List[T] {
	list := List[T]{}
	list.Extend(values)
	return list
}

func (list *List[T]) Append(item T) {
	*list = append(*list, item)
}

func (list List[T]) At(index int) T {
	if index < 0 {
		return list[len(list)+index]
	}
	return list[index]
}

func (list *List[T]) Clear() {
	for len(*list) > 0 {
		list.Pop()
	}
}

func (list List[T]) Copy() List[T] {
	return NewList(list...)
}

func (list *List[T]) Extend(items []T) {
	*list = append(*list, items...)
}

func (list List[T]) Every(f func(T) bool) bool {
	for _, v := range list {
		if !f(v) {
			return false
		}
	}
	return true
}

func (list *List[T]) Insert(value T, index int) {
	before, after := (*list)[:index], (*list)[index:]
	result := []T{}
	for _, v := range before {
		result = append(result, v)
	}
	result = append(result, value)
	for _, v := range after {
		result = append(result, v)
	}
	*list = result
}

func (list *List[T]) Pop() T {
	return list.Remove(len(*list) - 1)
}

func (list *List[T]) Remove(index int) T {
	value := (*list)[index]
	*list = append((*list)[:index], (*list)[index+1:]...)
	return value
}

func (list *List[T]) Reverse() {
	reversed := []T{}
	for i := len(*list) - 1; i >= 0; i-- {
		reversed = append(reversed, (*list)[i])
	}
	*list = reversed
}

func (list *List[T]) Shift() T {
	return list.Remove(0)
}

func (list List[T]) Some(f func(T) bool) bool {
	for _, v := range list {
		if f(v) {
			return true
		}
	}
	return false
}

func (list *List[T]) UnShift(item T) {
	*(list) = append([]T{item}, *list...)
}

func Contains[T comparable](list []T, value T) bool {
	return Index(list, value) != -1
}

func Count[T comparable](list []T, value T) int {
	count := 0
	for _, item := range list {
		if item == value {
			count++
		}
	}
	return count
}

func Filter[T any](list []T, condition func(T) bool) []T {
	filtered := []T{}
	for _, v := range list {
		if condition(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func Find[T any](list []T, condition func(T) bool) (value T, found bool) {
	for _, v := range list {
		if condition(v) {
			return v, true
		}
	}
	return value, false
}

func Index[T comparable](list []T, value T) int {
	for i, v := range list {
		if v == value {
			return i
		}
	}
	return -1
}

func Map[T any, V any](list []T, f func(T) V) []V {
	mapped := []V{}
	for _, v := range list {
		mapped = append(mapped, f(v))
	}
	return mapped
}

func Reduce[T any, V any](list []T, f func(V, T) V, acc V) V {
	for _, v := range list {
		acc = f(acc, v)
	}
	return acc
}

func All(conditions ...bool) bool {
	for _, c := range conditions {
		if !c {
			return false
		}
	}
	return true
}

func Any(conditions ...bool) bool {
	for _, c := range conditions {
		if c {
			return true
		}
	}
	return false
}

func Zip[T any](slices ...[]T) ([][]T, error) {
	var zipped [][]T
	if len(slices) == 0 {
		return zipped, errors.New("Zip() expected at least 1 argument but got 0")
	}
	x, y := len(slices), len(slices[0])
	for i := 1; i < y; i++ {
		if len(slices[i]) != x {
			return zipped, errors.New("Zip() received slices of different lengths")
		}
	}
	zipped = make([][]T, len(slices))
	for i := 0; i < len(slices[0]); i++ {
		z := make([]T, len(slices))
		for j := 0; j < len(slices); j++ {
			z[j] = slices[j][i]
		}
		zipped[i] = z
	}
	return zipped, nil
}
