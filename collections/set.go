package collections

import (
	"errors"
	"fmt"
)

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](values ...T) Set[T] {
	set := Set[T]{}
	for _, v := range values {
		set[v] = struct{}{}
	}
	return set
}

func (set Set[T]) Add(values ...T) {
	for _, v := range values {
		set[v] = struct{}{}
	}
}

func (set Set[T]) Clear() {
	for _, v := range set.Members() {
		set.Discard(v)
	}
}

func (set Set[T]) Copy() Set[T] {
	return NewSet(set.Members()...)
}

func (set Set[T]) Difference(sets ...Set[T]) Set[T] {
	difference := set.Copy()
	difference.DifferenceUpdate(sets...)
	return difference
}

func (set Set[T]) DifferenceUpdate(sets ...Set[T]) {
	for _, s := range sets {
		for value := range s {
			set.Discard(value)
		}
	}
}

func (set Set[T]) Discard(value T) {
	delete(set, value)
}

func (set Set[T]) Has(value T) bool {
	_, ok := set[value]
	return ok
}

func (set Set[T]) Intersection(sets ...Set[T]) Set[T] {
	intersection := set.Copy()
	intersection.IntersectionUpdate(sets...)
	return intersection
}

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

func (set1 Set[T]) IsDisjoint(set2 Set[T]) bool {
	for value := range set1 {
		if set2.Has(value) {
			return false
		}
	}
	return true
}

func (set1 Set[T]) IsSubset(set2 Set[T]) bool {
	for value := range set1 {
		if !set2.Has(value) {
			return false
		}
	}
	return true
}

func (set1 Set[T]) IsSuperset(set2 Set[T]) bool {
	for value := range set2 {
		if !set1.Has(value) {
			return false
		}
	}
	return true
}

func (set Set[T]) Members() []T {
	members := make([]T, len(set))
	i := 0
	for value := range set {
		members[i] = value
		i++
	}
	return members
}

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

func (set1 Set[T]) SymmetricDifference(set2 Set[T]) Set[T] {
	return set1.Difference(set2).Union(set2.Difference(set1))
}

func (set Set[T]) SymmetricDifferenceUpdate(sets ...Set[T]) {
	for _, s := range sets {
		diff := NewSet[T]()
		diff.Add(set.SymmetricDifference(s).Members()...)
		set.Clear()
		set.Add(diff.Members()...)
	}
}

func (set1 Set[T]) Union(set2 Set[T]) Set[T] {
	union := set1.Copy()
	union.Add(set2.Members()...)
	return union
}

func (set Set[T]) Update(sets ...Set[T]) {
	for _, s := range sets {
		set.Add(s.Members()...)
	}
}
