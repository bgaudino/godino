package helpers

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func Sum[T Number](numbers ...T) T {
	var sum T
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func Prod[T Number](numbers ...T) T {
	product := T(1)
	for _, n := range numbers {
		product *= n
	}
	return product
}

func Max[T constraints.Ordered](elements ...T) (T, bool) {
	var max T
	for i, e := range elements {
		if i == 0 || e > max {
			max = e
		}
	}
	return max, len(elements) > 0
}

func Min[T constraints.Ordered](elements ...T) (T, bool) {
	var min T
	for i, e := range elements {
		if i == 0 || e < min {
			min = e
		}
	}
	return min, len(elements) > 0
}
