package helpers

import (
	"errors"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
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

func Max[T constraints.Ordered](elements ...T) (T, error) {
	var max T
	var err error
	if len(elements) == 0 {
		err = errors.New("Max() expected at least 1 argument, got 0")
	} else {
		for i, e := range elements {
			if i == 0 || e > max {
				max = e
			}
		}
	}
	return max, err
}

func Min[T constraints.Ordered](elements ...T) (T, error) {
	var min T
	var err error
	if len(elements) == 0 {
		err = errors.New("Min() expected at least 1 argument, got 0")
	} else {
		for i, e := range elements {
			if i == 0 || e < min {
				min = e
			}
		}
	}
	return min, err
}

func Prod[T Number](numbers ...T) T {
	product := T(1)
	for _, n := range numbers {
		product *= n
	}
	return product
}

func Sum[T Number](numbers ...T) T {
	var sum T
	for _, n := range numbers {
		sum += n
	}
	return sum
}
