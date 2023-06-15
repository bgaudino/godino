package helpers

import (
	"sort"

	"golang.org/x/exp/maps"
)

type counterElement[T comparable] struct {
	element T
	count   int
}
type counterElements[T comparable] []counterElement[T]

func (l counterElements[T]) Less(i, j int) bool {
	switch diff := l[i].count - l[j].count; {
	case diff < 0:
		return false
	case diff > 0:
		return true
	}
	return i < j
}

func (l counterElements[T]) Len() int { return len(l) }

func (l counterElements[T]) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

type Counter[T comparable] struct {
	elements map[T]int
	keys     []T
}

func NewCounter[T comparable](values []T) Counter[T] {
	c := Counter[T]{elements: make(map[T]int), keys: []T{}}
	for _, v := range values {
		c.Add(v)
	}
	return c
}

func (c *(Counter[T])) Add(value T) {
	if _, ok := c.elements[value]; ok {
		c.elements[value]++
	} else {
		c.elements[value] = 1
		c.keys = append(c.keys, value)
	}
}

func (c Counter[T]) Elements() counterElements[T] {
	elements := counterElements[T]{}
	for _, k := range c.keys {
		elements = append(elements, counterElement[T]{
			element: k,
			count:   c.elements[k],
		})
	}
	return elements
}

func (c Counter[T]) Get(value T) int {
	return c.elements[value]
}

func (c Counter[T]) MostCommon(n int) counterElements[T] {
	elements := c.Elements()
	sort.Sort(elements)
	if n < 0 {
		return elements
	}
	return elements[:n]
}

func (c *(Counter[T])) Subtract(value T) {
	if _, ok := c.elements[value]; ok {
		c.elements[value]--
	}
}

func (c Counter[T]) Total() int {
	return Sum(maps.Values(c.elements)...)
}

func (c *(Counter[T])) Update(arrs ...[]T) {
	for _, arr := range arrs {
		for _, v := range arr {
			c.Add(v)
		}
	}
}