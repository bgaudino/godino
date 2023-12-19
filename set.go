package godino

import (
	"errors"
	"fmt"
)

// A data structure for storing unique values without any particular order
type Set[T comparable] map[T]struct{}

// Returns a new set containing the values provided
func NewSet[T comparable](values ...T) Set[T] {
	set := Set[T]{}
	for _, v := range values {
		set[v] = struct{}{}
	}
	return set
}

// Adds element(s) to the set
func (set Set[T]) Add(values ...T) {
	for _, v := range values {
		set[v] = struct{}{}
	}
}

// Removes all the elements from the set
func (set Set[T]) Clear() {
	for _, v := range set.Members() {
		set.Discard(v)
	}
}

// Returns a copy of the set
func (set Set[T]) Copy() Set[T] {
	return NewSet(set.Members()...)
}

// Returns a set containing the difference between two or more sets
func (set Set[T]) Difference(sets ...Set[T]) Set[T] {
	difference := set.Copy()
	difference.DifferenceUpdate(sets...)
	return difference
}

// Removes the items in this set that are also included in one or more other sets
func (set Set[T]) DifferenceUpdate(sets ...Set[T]) {
	for _, s := range sets {
		for value := range s {
			set.Discard(value)
		}
	}
}

// Remove the specified item. If the item is not present this is a noop
func (set Set[T]) Discard(value T) {
	delete(set, value)
}

// Returns true if the sets contain the same members regardless of order
func (set Set[T]) Equals(set2 Set[T]) bool {
	return len(set.SymmetricDifference(set2)) == 0
}

// Returns true if the set contins the specified element
func (set Set[T]) Has(value T) bool {
	_, ok := set[value]
	return ok
}

// Returns a set, that is the intersection of two or more sets
func (set Set[T]) Intersection(sets ...Set[T]) Set[T] {
	intersection := set.Copy()
	intersection.IntersectionUpdate(sets...)
	return intersection
}

// Removes the items in this set that are not present in other, specified set(s)
func (set Set[T]) IntersectionUpdate(sets ...Set[T]) {
	for value := range set {
		for _, s := range sets {
			if !s.Has(value) {
				set.Discard(value)
				break
			}
		}
	}
}

// Returns true if the sets have no common elements
func (set1 Set[T]) IsDisjoint(set2 Set[T]) bool {
	for value := range set1 {
		if set2.Has(value) {
			return false
		}
	}
	return true
}

// Returns true if the set is contained by the given set
func (set1 Set[T]) IsSubset(set2 Set[T]) bool {
	for value := range set1 {
		if !set2.Has(value) {
			return false
		}
	}
	return true
}

// Returns true is the set contains the given set
func (set1 Set[T]) IsSuperset(set2 Set[T]) bool {
	for value := range set2 {
		if !set1.Has(value) {
			return false
		}
	}
	return true
}

// Returns the elements contained by the set. Elements are unordered.
func (set Set[T]) Members() []T {
	members := make([]T, len(set))
	i := 0
	for value := range set {
		members[i] = value
		i++
	}
	return members
}

// Returns an element from the set and an error if the set is empty
func (set Set[T]) Pop() (T, error) {
	var popped T
	if len(set) == 0 {
		return popped, errors.New("cannot pop item from an empty set")
	}
	for value := range set {
		popped = value
		set.Discard(value)
		break
	}
	return popped, nil
}

// Removes the specified element from the set. Returns true if the element was in the set
func (set Set[T]) Remove(value T) bool {
	_, ok := set[value]
	if ok {
		set.Discard(value)
	}
	return ok
}

func (set Set[T]) String() string {
	return fmt.Sprintf("%v", set.Members())
}

// Returns a set that contains all items from both sets, except items that are present in both sets
func (set1 Set[T]) SymmetricDifference(set2 Set[T]) Set[T] {
	return set1.Difference(set2).Union(set2.Difference(set1))
}

// Removes the items that are present in both sets, and inserts the items that are not present in both sets
func (set Set[T]) SymmetricDifferenceUpdate(sets ...Set[T]) {
	for _, s := range sets {
		diff := NewSet[T]()
		diff.Add(set.SymmetricDifference(s).Members()...)
		set.Clear()
		set.Add(diff.Members()...)
	}
}

// Return a set that contains all items from both sets
func (set1 Set[T]) Union(set2 Set[T]) Set[T] {
	union := set1.Copy()
	union.Add(set2.Members()...)
	return union
}

// Adds all the items from the given sets
func (set Set[T]) Update(sets ...Set[T]) {
	for _, s := range sets {
		set.Add(s.Members()...)
	}
}
