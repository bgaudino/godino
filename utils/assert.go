package utils

import (
	"errors"
	"fmt"
	"strings"

	"golang.org/x/exp/constraints"
)

var ErrAssertion error = errors.New("assertion failed")

func Assert(condition bool, messages ...string) {
	message := strings.Join(messages, "\n")
	err := errors.New(message)
	if len(message) == 0 {
		err = ErrAssertion
	}
	if !condition {
		panic(err)
	}
}

func AssertEqual[T comparable](x, y T) {
	Assert(x == y, fmt.Sprintf("%v does not equal %v", x, y))
}

func AssertTrue(condition bool) {
	Assert(condition)
}

func AssertFalse(condition bool) {
	Assert(!condition)
}

func AssertGreaterThan[T constraints.Ordered](x, y T) {
	Assert(x > y, fmt.Sprintf("%v is not greater than %v", x, y))
}

func AssertGreaterThanOrEqual[T constraints.Ordered](x, y T) {
	Assert(x >= y, fmt.Sprintf("%v is not greater than or equal to %v", x, y))
}

func AssertLessThan[T constraints.Ordered](x, y T) {
	Assert(x < y, fmt.Sprintf("%v is not less than %v", x, y))
}

func AssertLessThanOrEqual[T constraints.Ordered](x, y T) {
	Assert(x <= y, fmt.Sprintf("%v is not less than or equal to %v", x, y))
}
