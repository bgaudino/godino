package functools

import (
	"errors"
)

func Append[L ~[]T, T any](arr *L, value T) {
	*arr = append(*arr, value)
}

func Clear[L ~[]T, T any](arr *L) {
	for len(*arr) > 0 {
		Pop(arr)
	}
}

func Contains[L ~[]T, T comparable](arr L, value T) bool {
	return Index(arr, value) != -1
}

func Copy[L ~[]T, T any](arr L) L {
	arr2 := make([]T, len(arr))
	copy(arr2, arr)
	return arr2
}

func Count[L ~[]T, T comparable](arr L, value T) int {
	count := 0
	for _, item := range arr {
		if item == value {
			count++
		}
	}
	return count
}

func Extend[L ~[]T, T any](arr *L, items []T) {
	*arr = append(*arr, items...)
}

func Index[L ~[]T, T comparable](arr L, value T) int {
	for i, v := range arr {
		if v == value {
			return i
		}
	}
	return -1
}

func Insert[L ~[]T, T any](arr *L, value T, index int) {
	*arr = append((*arr)[:index], append([]T{value}, (*arr)[index:]...)...)
}

func Pop[L ~[]T, T any](arr *L) T {
	return Remove(arr, len(*arr)-1)
}

func Remove[L ~[]T, T any](arr *L, index int) T {
	value := (*arr)[index]
	*arr = append((*arr)[:index], (*arr)[index+1:]...)
	return value
}

func Reverse[L ~[]T, T any](arr *L) {
	start, end := 0, len(*arr)-1
	for start < end {
		(*arr)[start], (*arr)[end] = (*arr)[end], (*arr)[start]
		start++
		end--
	}
}

func Shift[L ~[]T, T any](arr *L) T {
	return Remove(arr, 0)
}

func UnShift[L ~[]T, T any](arr *L, value T) {
	*arr = append([]T{value}, *arr...)
}

func ValueAt[L ~[]T, T any](arr L, index int) T {
	if index < 0 {
		return arr[len(arr)+index]
	}
	return arr[index]
}

func Zip[L ~[]T, T any](arrs ...[]T) ([][]T, error) {
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
