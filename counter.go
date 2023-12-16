package helpers

import (
	"fmt"
	"sort"

	"golang.org/x/exp/maps"
)

type counterElement[T comparable] struct {
	Element T
	Count   int
}
type counterElements[T comparable] []counterElement[T]

func (l counterElements[T]) Less(i, j int) bool {
	switch diff := l[i].Count - l[j].Count; {
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

// A datastructure for counting comparable elements. Elements and their counts are stored in a map as key value pairs
type Counter[T comparable] struct {
	counts     map[T]int
	keys       []T
	elements   counterElements[T]
	mostCommon counterElements[T]
}

// Returns a new counter containing the counts of the specified elements
func NewCounter[T comparable](values []T) Counter[T] {
	c := Counter[T]{counts: make(map[T]int), keys: []T{}}
	for _, v := range values {
		c.Add(v)
	}
	return c
}

// Increments the count for the specified element
func (c *Counter[T]) Add(value T) {
	c.clearCache()
	if _, ok := c.counts[value]; ok {
		c.counts[value]++
	} else {
		c.counts[value] = 1
		c.keys = append(c.keys, value)
	}
}

func (c *Counter[T]) clearCache() {
	c.elements = nil
	c.mostCommon = nil
}

// Returns the elements and their counts in the order they were added
func (c *Counter[T]) Elements() counterElements[T] {
	if c.elements == nil {
		elements := counterElements[T]{}
		for _, k := range c.keys {
			elements = append(elements, counterElement[T]{
				Element: k,
				Count:   c.counts[k],
			})
		}
		c.elements = elements
	}
	return c.elements
}

// Returns the count of the specified element
func (c Counter[T]) Get(value T) int {
	return c.counts[value]
}

// Returns the n most common elements and their counts. If n is less than 0, returns
// all elements sorted by most common. Elements with they same count are returned in the
// order in which they were added.
func (c *Counter[T]) MostCommon(n int) counterElements[T] {
	if c.mostCommon == nil {
		elements := c.Elements()
		sort.Sort(elements)
		c.mostCommon = elements
	}
	if n < 0 {
		return c.mostCommon
	}
	return c.mostCommon[:n]
}

func (c Counter[T]) String() string {
	return fmt.Sprintf("%v", c.Elements())
}

// Decrements the count of the specified element
func (c *Counter[T]) Subtract(value T) {
	c.clearCache()
	c.counts[value]--
}

// Returns the sum of the counts all of all elements
func (c Counter[T]) Total() int {
	return Sum(maps.Values(c.counts)...)
}

// Adds the counts of the elements in the provided arrays to the counter
func (c *Counter[T]) Update(arrs ...[]T) {
	c.clearCache()
	for _, arr := range arrs {
		for _, v := range arr {
			c.Add(v)
		}
	}
}
