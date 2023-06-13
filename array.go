package helpers

import (
	"errors"
)

func Append[T any](arr *([]T), value T) {
	*arr = append(*arr, value)
}

func Clear[T any](arr *([]T)) {
	for len(*arr) > 0 {
		Pop(arr)
	}
}

func Contains[T comparable](arr []T, value T) bool {
	return Index(arr, value) != -1
}

func Copy[T any](arr []T) []T {
	arr2 := make([]T, len(arr))
	copy(arr2, arr)
	return arr2
}

func Count[T comparable](arr []T, value T) int {
	count := 0
	for _, item := range arr {
		if item == value {
			count++
		}
	}
	return count
}

func Extend[T any](arr *([]T), items []T) {
	*arr = append(*arr, items...)
}

func Every[T any](arr []T, f func(T) bool) bool {
	for _, v := range arr {
		if !f(v) {
			return false
		}
	}
	return true
}

func Filter[T any](arr []T, condition func(T) bool) []T {
	filtered := []T{}
	for _, v := range arr {
		if condition(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func Find[T any](arr []T, condition func(T) bool) (value T, found bool) {
	for _, v := range arr {
		if condition(v) {
			return v, true
		}
	}
	return value, false
}

func Index[T comparable](arr []T, value T) int {
	for i, v := range arr {
		if v == value {
			return i
		}
	}
	return -1
}

func Insert[T any](arr *([]T), value T, index int) {
	*arr = append((*arr)[:index], append([]T{value}, (*arr)[index:]...)...)
}

func Map[T any, V any](arr []T, f func(T) V) []V {
	mapped := []V{}
	for _, v := range arr {
		mapped = append(mapped, f(v))
	}
	return mapped
}

func Pop[T any](arr *([]T)) T {
	return Remove(arr, len(*arr)-1)
}

func Reduce[T any, V any](arr []T, f func(V, T) V, acc V) V {
	for _, v := range arr {
		acc = f(acc, v)
	}
	return acc
}

func Remove[T any](arr *([]T), index int) T {
	value := (*arr)[index]
	*arr = append((*arr)[:index], (*arr)[index+1:]...)
	return value
}

func Reverse[T any](arr *([]T)) {
	start, end := 0, len(*arr)-1
	for start < end {
		(*arr)[start], (*arr)[end] = (*arr)[end], (*arr)[start]
		start++
		end--
	}
}

func Shift[T any](arr *([]T)) T {
	return Remove(arr, 0)
}

func Some[T any](arr []T, f func(T) bool) bool {
	for _, v := range arr {
		if f(v) {
			return true
		}
	}
	return false
}

func UnShift[T any](arr *([]T), value T) {
	*arr = append([]T{value}, *arr...)
}

func ValueAt[T any](arr []T, index int) T {
	if index < 0 {
		return arr[len(arr)+index]
	}
	return arr[index]
}

func Zip[T any](arrs ...[]T) ([][]T, error) {
	var zipped [][]T
	if len(arrs) == 0 {
		return zipped, errors.New("Zip() expected at least 1 argument but got 0")
	}
	x, y := len(arrs), len(arrs[0])
	for i := 1; i < y; i++ {
		if len(arrs[i]) != x {
			return zipped, errors.New("Zip() received slices of different lengths")
		}
	}
	zipped = make([][]T, len(arrs))
	for i := 0; i < len(arrs[0]); i++ {
		z := make([]T, len(arrs))
		for j := 0; j < len(arrs); j++ {
			z[j] = arrs[j][i]
		}
		zipped[i] = z
	}
	return zipped, nil
}
