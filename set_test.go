package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/maps"
)

func TestNewSet(t *testing.T) {
	t.Run("creating a new set of ints", func(t *testing.T) {
		set := NewSet(1, 2, 3)
		expected := map[int]struct{}{1: {}, 2: {}, 3: {}}
		assert.Equal(t, maps.Equal(set, expected), true)
	})
	t.Run("creating a new set of strings", func(t *testing.T) {
		set := NewSet("foo", "bar", "baz")
		expected := map[string]struct{}{"foo": {}, "bar": {}, "baz": {}}
		assert.Equal(t, maps.Equal(set, expected), true)
	})
}

func TestAdd(t *testing.T) {
	t.Run("adding a single value", func(t *testing.T) {
		set := NewSet[int]()
		val := 1
		set.Add(val)
		_, ok := set[val]
		assert.Equal(t, ok, true)
	})
	t.Run("adding many values", func(t *testing.T) {
		values := []string{"foo", "bar", "baz"}
		set := NewSet[string](values...)
		assert.ElementsMatch(t, set.Members(), values)
	})
}

func TestClear(t *testing.T) {
	set := NewSet(1, 2, 3)
	set.Clear()
	assert.Equal(t, len(set), 0)
}

func TestDifference(t *testing.T) {
	set1 := NewSet(1, 2, 3)
	set2 := NewSet(3, 4, 5)
	diff := set1.Difference(set2)
	assert.ElementsMatch(t, diff.Members(), []int{1, 2})
}

func TestDifferenceUpdate(t *testing.T) {
	set1 := NewSet(1, 2, 3)
	set2 := NewSet(3, 4, 5)
	set1.DifferenceUpdate(set2)
	assert.ElementsMatch(t, set1.Members(), []int{1, 2})
}

func TestHas(t *testing.T) {
	set := NewSet(1, 2, 3)
	t.Run("has value", func(t *testing.T) {
		assert.Equal(t, set.Has(1), true)
	})
	t.Run("does not have value", func(t *testing.T) {
		assert.Equal(t, set.Has(4), false)
	})
}

func TestIntersection(t *testing.T) {
	set1 := NewSet(1, 2, 3)
	set2 := NewSet(2, 3, 4)
	intersection := set1.Intersection(set2)
	assert.ElementsMatch(t, intersection.Members(), []int{2, 3})
}

func TestIntersectionUpdate(t *testing.T) {
	set1 := NewSet(1, 2, 3)
	set2 := NewSet(2, 3, 4)
	set1.IntersectionUpdate(set2)
	assert.ElementsMatch(t, set1.Members(), []int{2, 3})
}

func TestIsDisjoint(t *testing.T) {
	set1 := NewSet(1, 2, 3)
	t.Run("is disjoint", func(t *testing.T) {
		set2 := NewSet(4, 5, 6)
		assert.Equal(t, set1.IsDisjoint(set2), true)
	})
	t.Run("is not disjoint", func(t *testing.T) {
		set2 := NewSet(3, 4, 5)
		assert.Equal(t, set1.IsDisjoint(set2), false)
	})
}

func TestIsSubset(t *testing.T) {
	set1 := NewSet(1, 2, 3)
	t.Run("is subset", func(t *testing.T) {
		set2 := NewSet(1, 2, 3, 4)
		assert.Equal(t, set1.IsSubset(set2), true)
	})
	t.Run("is not subset", func(t *testing.T) {
		set2 := NewSet(2, 3, 4, 5)
		assert.Equal(t, set1.IsSubset(set2), false)
	})
}

func TestIsSuperset(t *testing.T) {
	set1 := NewSet(1, 2, 3, 4)
	t.Run("is superset", func(t *testing.T) {
		set2 := NewSet(1, 2, 3)
		assert.Equal(t, set1.IsSuperset(set2), true)
	})
	t.Run("is not superset", func(t *testing.T) {
		set2 := NewSet(3, 4, 5)
		assert.Equal(t, set1.IsSuperset(set2), false)
	})
}

func TestMembers(t *testing.T) {
	set := NewSet(1, 2, 3)
	assert.ElementsMatch(t, set.Members(), []int{1, 2, 3})
}

func TestRemove(t *testing.T) {
	set := NewSet(1, 2, 3)
	set.Remove(1)
	_, ok := set[1]
	assert.Equal(t, ok, false)
}

func TestSymmetricDifference(t *testing.T) {
	set1 := NewSet(1, 2, 3)
	set2 := NewSet(2, 3, 4)
	diff := set1.SymmetricDifference(set2)
	assert.ElementsMatch(t, diff.Members(), []int{1, 4})
}

func TestSymmetricDifferenceUpdate(t *testing.T) {
	set1 := NewSet(1, 2, 3)
	set2 := NewSet(2, 3, 4)
	set1.SymmetricDifferenceUpdate(set2)
	assert.ElementsMatch(t, set1.Members(), []int{1, 4})
}

func TestUnion(t *testing.T) {
	set1 := NewSet(1, 2, 3)
	set2 := NewSet(3, 4, 5)
	union := set1.Union(set2)
	assert.ElementsMatch(t, union.Members(), []int{1, 2, 3, 4, 5})
}

func TestUnionUpdate(t *testing.T) {
	set1 := NewSet(1, 2, 3)
	set2 := NewSet(3, 4, 5)
	set1.Update(set2)
	assert.ElementsMatch(t, set1.Members(), []int{1, 2, 3, 4, 5})
}
